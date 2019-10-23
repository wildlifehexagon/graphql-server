package stock

import (
	"context"
	"fmt"
	"time"

	"github.com/dictyBase/graphql-server/internal/graphql/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

const (
	phenoOntology       = "Dicty Phenotypes"
	envOntology         = "Dicty Environment"
	assayOntology       = "Dictyostellium Assay"
	mutagenesisOntology = "Dd Mutagenesis Method"
	geneticModOntology  = "genetic modification"
	dictyAnnoOntology   = "dicty_annotation"
	strainCharOnto      = "strain_characteristics"
	literatureTag       = "literature_tag"
	noteTag             = "public note"
	sysnameTag          = "systematic name"
	mutmethodTag        = "mutagenesis method"
	muttypeTag          = "mutant type"
	genoTag             = "genotype"
	synTag              = "synonym"
)

type StrainResolver struct {
	Client           pb.StockServiceClient
	UserClient       user.UserServiceClient
	AnnotationClient annotation.TaggedAnnotationServiceClient
	Registry         registry.Registry
	Logger           *logrus.Entry
}

func (r *StrainResolver) ID(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Id, nil
}
func (r *StrainResolver) CreatedAt(ctx context.Context, obj *models.Strain) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *StrainResolver) UpdatedAt(ctx context.Context, obj *models.Strain) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *StrainResolver) CreatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *StrainResolver) UpdatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *StrainResolver) Summary(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.Summary, nil
}
func (r *StrainResolver) EditableSummary(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.EditableSummary, nil
}
func (r *StrainResolver) Depositor(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.Depositor, nil
}
func (r *StrainResolver) Genes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	g := obj.Data.Attributes.Genes
	pg := []*string{}
	// need to use for loop here, not range
	// https://github.com/golang/go/issues/22791#issuecomment-345391395
	for i := 0; i < len(g); i++ {
		pg = append(pg, &g[i])
	}
	return pg, nil
}
func (r *StrainResolver) Dbxrefs(ctx context.Context, obj *models.Strain) ([]*string, error) {
	d := obj.Data.Attributes.Dbxrefs
	pd := []*string{}
	for i := 0; i < len(d); i++ {
		pd = append(pd, &d[i])
	}
	return pd, nil
}
func (r *StrainResolver) Publications(ctx context.Context, obj *models.Strain) ([]*publication.Publication, error) {
	pubs := []*publication.Publication{}
	for _, id := range obj.Data.Attributes.Publications {
		if len(id) < 1 {
			continue
		}
		p, err := utils.FetchPublication(ctx, r.Registry, id)
		if err != nil {
			errorutils.AddGQLError(ctx, err)
			r.Logger.Error(err)
			return pubs, err
		}
		pubs = append(pubs, p)
	}
	return pubs, nil
}
func (r *StrainResolver) Label(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.Label, nil
}
func (r *StrainResolver) Species(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.Species, nil
}
func (r *StrainResolver) Plasmid(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.Plasmid, nil
}
func (r *StrainResolver) Parent(ctx context.Context, obj *models.Strain) (*models.Strain, error) {
	parent := obj.Data.Attributes.Parent
	strain, err := r.Client.GetStrain(ctx, &pb.StockId{Id: parent})
	if err != nil {
		r.Logger.Debugf("could not find parent strain with ID %s", parent)
		return nil, nil
	}
	r.Logger.Debugf("successfully found parent strain with ID %s", parent)
	return &models.Strain{
		Data: strain.Data,
	}, nil
}
func (r *StrainResolver) Names(ctx context.Context, obj *models.Strain) ([]*string, error) {
	names := []*string{}
	n, err := r.AnnotationClient.ListAnnotations(
		ctx,
		&annotation.ListParameters{
			Limit: 20,
			Filter: fmt.Sprintf(
				"entry_id===%s;tag===%s;ontology===%s",
				obj.Data.Id, synTag, dictyAnnoOntology,
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
	strainId := obj.Data.Id
	gc, err := r.AnnotationClient.ListAnnotationGroups(
		ctx,
		&annotation.ListGroupParameters{
			Filter: fmt.Sprintf(
				"entry_id==%s;ontology==%s",
				strainId,
				phenoOntology,
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
	for _, item := range gc.Data {
		m := &models.Phenotype{}
		for _, g := range item.Group.Data {
			switch g.Attributes.Ontology {
			case phenoOntology:
				m.Phenotype = g.Attributes.Tag
			case envOntology:
				m.Environment = &g.Attributes.Tag
			case assayOntology:
				m.Assay = &g.Attributes.Tag
			case literatureTag:
				pub, err := utils.FetchPublication(ctx, r.Registry, g.Attributes.Value)
				if err != nil {
					errorutils.AddGQLError(ctx, err)
					r.Logger.Error(err)
				}
				m.Publication = pub
			case noteTag:
				m.Note = &g.Attributes.Value
			}
		}
		p = append(p, m)
	}
	return p, nil
}

func (r *StrainResolver) GeneticModification(ctx context.Context, obj *models.Strain) (*string, error) {
	var gm string
	gc, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      muttypeTag,
			Ontology: dictyAnnoOntology,
			EntryId:  obj.Data.Id,
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
			Tag:      mutmethodTag,
			Ontology: mutagenesisOntology,
			EntryId:  obj.Data.Id,
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
			Tag:      sysnameTag,
			Ontology: dictyAnnoOntology,
			EntryId:  obj.Data.Id,
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
	cg, err := r.AnnotationClient.ListAnnotationGroups(
		ctx,
		&annotation.ListGroupParameters{
			Filter: fmt.Sprintf(
				"entry_id==%s;ontology==%s",
				obj.Data.Id,
				strainCharOnto,
			),
			Limit: 30,
		})
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return c, nil
		}
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return c, err
	}
	for _, item := range cg.Data {
		for _, t := range item.Group.Data {
			c = append(c, &t.Attributes.Tag)
		}
	}
	return c, nil
}
func (r *StrainResolver) Genotypes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	g := []*string{}
	gl, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			EntryId:  obj.Data.Id,
			Ontology: dictyAnnoOntology,
			Tag:      genoTag,
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

// still needs to be implemented properly
func (r *StrainResolver) InStock(ctx context.Context, obj *models.Strain) (bool, error) {
	return true, nil
}
