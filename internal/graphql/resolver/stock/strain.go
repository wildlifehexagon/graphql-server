package stock

import (
	"context"
	"fmt"
	"regexp"

	"github.com/dictyBase/aphgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/cache"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/fetch"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type StrainResolver struct {
	Client           pb.StockServiceClient
	UserClient       user.UserServiceClient
	AnnotationClient annotation.TaggedAnnotationServiceClient
	Registry         registry.Registry
	Logger           *logrus.Entry
}

func (r *StrainResolver) CreatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	return g, nil
}

func (r *StrainResolver) UpdatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	return g, nil
}

func (r *StrainResolver) Depositor(ctx context.Context, obj *models.Strain) (*user.User, error) {
	email := *obj.Depositor
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user.User{
			Data: &user.UserData{
				Attributes: &user.UserAttributes{
					FirstName: "",
					LastName:  "",
				},
			},
		}, nil
	}
	return g, nil
}

func (r *StrainResolver) Genes(ctx context.Context, obj *models.Strain) ([]*models.Gene, error) {
	g := []*models.Gene{}
	redis := r.Registry.GetRedisRepository(cache.RedisKey)
	for _, v := range obj.Genes {
		if *v == "" {
			continue
		}
		gene, err := cache.GetGeneFromCache(ctx, redis, *v)
		if err != nil {
			r.Logger.Error(err)
			continue
		}
		g = append(g, gene)
	}
	return g, nil
}

func (r *StrainResolver) Publications(ctx context.Context, obj *models.Strain) ([]*publication.Publication, error) {
	pubs := []*publication.Publication{}
	for _, id := range obj.Publications {
		if len(*id) < 1 {
			continue
		}
		// GWDI IDs come back as doi:10.1101/582072
		doi := regexp.MustCompile(`^doi:10.\d{4,9}/[-._;()/:A-Z0-9]+$`)
		if doi.MatchString(*id) {
			url := fmt.Sprintf("https://doi.org/%s", *id)
			r.Logger.Debugf("fetching doi with address %s", url)
			p, err := fetch.FetchDOI(ctx, url)
			if err != nil {
				errorutils.AddGQLError(ctx, err)
				r.Logger.Error(err)
				return pubs, err
			}
			pubs = append(pubs, p)
		} else {
			endpoint := r.Registry.GetAPIEndpoint(registry.PUBLICATION)
			p, err := fetch.FetchPublication(ctx, endpoint, *id)
			if err != nil {
				errorutils.AddGQLError(ctx, err)
				r.Logger.Error(err)
				return pubs, err
			}
			pubs = append(pubs, p)
		}
	}
	return pubs, nil
}

func (r *StrainResolver) Parent(ctx context.Context, obj *models.Strain) (*models.Strain, error) {
	parent := obj.Parent
	if parent != nil {
		n, err := r.Client.GetStrain(ctx, &pb.StockId{Id: *parent})
		if err != nil {
			r.Logger.Debugf("could not find parent strain with ID %s", *parent)
			return nil, nil
		}
		r.Logger.Debugf("successfully found parent strain with ID %s", *parent)
		return ConvertToStrainModel(*parent, n.Data.Attributes), nil
	}
	return &models.Strain{}, nil
}

func (r *StrainResolver) Names(ctx context.Context, obj *models.Strain) ([]*string, error) {
	names := []*string{}
	if len(obj.Names) > 0 {
		for _, v := range obj.Names {
			names = append(names, v)
		}
	}
	n, err := r.AnnotationClient.ListAnnotations(
		ctx,
		&annotation.ListParameters{
			Limit: 20,
			Filter: fmt.Sprintf(
				"entry_id===%s;tag===%s;ontology===%s",
				obj.ID, registry.SynTag, registry.DictyAnnoOntology,
			)})
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return names, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return names, err
	}
	for _, syn := range n.Data {
		names = append(names, &syn.Attributes.Value)
	}
	return names, nil
}

func (r *StrainResolver) Phenotypes(ctx context.Context, obj *models.Strain) ([]*models.Phenotype, error) {
	p := []*models.Phenotype{}
	strainId := obj.ID
	gc, err := r.AnnotationClient.ListAnnotationGroups(
		ctx,
		&annotation.ListGroupParameters{
			Filter: fmt.Sprintf(
				"entry_id==%s;ontology==%s",
				strainId,
				registry.PhenoOntology,
			),
			Limit: 30,
		})
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return p, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return p, err
	}
	p = getPhenotypes(ctx, r, gc.Data)
	return p, nil
}

func (r *StrainResolver) GeneticModification(ctx context.Context, obj *models.Strain) (*string, error) {
	var gm string
	gc, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      registry.MuttypeTag,
			Ontology: registry.DictyAnnoOntology,
			EntryId:  obj.ID,
		},
	)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return &gm, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &gm, err
	}
	gm = gc.Data.Attributes.Value
	return &gm, nil
}

func (r *StrainResolver) MutagenesisMethod(ctx context.Context, obj *models.Strain) (*string, error) {
	var m string
	gc, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      registry.MutmethodTag,
			Ontology: registry.MutagenesisOntology,
			EntryId:  obj.ID,
		},
	)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return &m, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &m, err
	}
	m = gc.Data.Attributes.Value
	return &m, nil
}

func (r *StrainResolver) SystematicName(ctx context.Context, obj *models.Strain) (string, error) {
	sn, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      registry.SysnameTag,
			Ontology: registry.DictyAnnoOntology,
			EntryId:  obj.ID,
		},
	)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return "", nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return "", err
	}
	return sn.Data.Attributes.Value, nil
}

func (r *StrainResolver) Characteristics(ctx context.Context, obj *models.Strain) ([]*string, error) {
	c := []*string{}
	cg, err := r.AnnotationClient.ListAnnotations(
		ctx, &annotation.ListParameters{Filter: fmt.Sprintf(
			"entry_id===%s;ontology===%s",
			obj.ID, registry.StrainCharOnto,
		)},
	)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return c, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return c, err
	}
	for _, item := range cg.Data {
		c = append(c, &item.Attributes.Tag)
	}
	return c, nil
}

func (r *StrainResolver) Genotypes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	g := []*string{}
	gl, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			EntryId:  obj.ID,
			Ontology: registry.DictyAnnoOntology,
			Tag:      registry.GenoTag,
		})
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return g, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return g, err
	}
	g = append(g, &gl.Data.Attributes.Value)
	return g, nil
}

func (r *StrainResolver) InStock(ctx context.Context, obj *models.Strain) (bool, error) {
	id := obj.ID
	_, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      registry.StrainInvTag,
			Ontology: registry.StrainInvOnto,
			EntryId:  id,
		},
	)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return false, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Errorf("error getting %s from inventory: %v", id, err)
		return false, err
	}
	return true, nil
}

func getPhenotypes(ctx context.Context, r *StrainResolver, data []*annotation.TaggedAnnotationGroupCollection_Data) []*models.Phenotype {
	p := []*models.Phenotype{}
	for _, item := range data {
		m := &models.Phenotype{}
		for _, g := range item.Group.Data {
			switch g.Attributes.Ontology {
			case registry.PhenoOntology:
				m.Phenotype = g.Attributes.Tag
			case registry.EnvOntology:
				m.Environment = &g.Attributes.Tag
			case registry.AssayOntology:
				m.Assay = &g.Attributes.Tag
			case registry.DictyAnnoOntology:
				if g.Attributes.Tag == registry.LiteratureTag {
					endpoint := r.Registry.GetAPIEndpoint(registry.PUBLICATION)
					pub, err := fetch.FetchPublication(ctx, endpoint, g.Attributes.Value)
					if err != nil {
						r.Logger.Error(err)
						errorutils.AddGQLError(ctx, err)
					}
					m.Publication = pub
				}
				if g.Attributes.Tag == registry.NoteTag {
					m.Note = &g.Attributes.Value
				}
			}
		}
		p = append(p, m)
	}
	return p
}

func ConvertToStrainModel(id string, attr *pb.StrainAttributes) *models.Strain {
	return &models.Strain{
		ID:              id,
		CreatedAt:       aphgrpc.ProtoTimeStamp(attr.CreatedAt),
		UpdatedAt:       aphgrpc.ProtoTimeStamp(attr.UpdatedAt),
		CreatedBy:       attr.CreatedBy,
		UpdatedBy:       attr.UpdatedBy,
		Summary:         &attr.Summary,
		EditableSummary: &attr.EditableSummary,
		Depositor:       &attr.Depositor,
		Genes:           sliceConverter(attr.Genes),
		Dbxrefs:         sliceConverter(attr.Dbxrefs),
		Publications:    sliceConverter(attr.Publications),
		Label:           attr.Label,
		Species:         attr.Species,
		Plasmid:         &attr.Plasmid,
		Names:           sliceConverter(attr.Names),
	}
}

func sliceConverter(s []string) []*string {
	c := []*string{}
	// need to use for loop here, not range
	// https://github.com/golang/go/issues/22791#issuecomment-345391395
	for i := 0; i < len(s); i++ {
		c = append(c, &s[i])
	}
	return c
}
