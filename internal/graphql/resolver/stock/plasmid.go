package stock

import (
	"context"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/fetch"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type PlasmidResolver struct {
	Client           pb.StockServiceClient
	UserClient       user.UserServiceClient
	AnnotationClient annotation.TaggedAnnotationServiceClient
	Registry         registry.Registry
	Logger           *logrus.Entry
}

func (r *PlasmidResolver) CreatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	user := user.User{}
	email := obj.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}

func (r *PlasmidResolver) UpdatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	user := user.User{}
	email := obj.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}

func (r *PlasmidResolver) Genes(ctx context.Context, obj *models.Plasmid) ([]*models.Gene, error) {
	g := []*models.Gene{}
	return g, nil
}

func (r *PlasmidResolver) Publications(ctx context.Context, obj *models.Plasmid) ([]*publication.Publication, error) {
	pubs := []*publication.Publication{}
	for _, id := range obj.Publications {
		if len(*id) < 1 {
			continue
		}
		endpoint := r.Registry.GetAPIEndpoint(registry.PUBLICATION)
		p, err := fetch.FetchPublication(ctx, endpoint, *id)
		if err != nil {
			errorutils.AddGQLError(ctx, err)
			r.Logger.Error(err)
			return pubs, err
		}
		pubs = append(pubs, p)
	}
	return pubs, nil
}

func (r *PlasmidResolver) InStock(ctx context.Context, obj *models.Plasmid) (bool, error) {
	id := obj.ID
	_, err := r.AnnotationClient.GetEntryAnnotation(
		ctx,
		&annotation.EntryAnnotationRequest{
			Tag:      registry.PlasmidInvTag,
			Ontology: registry.PlasmidInvOnto,
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

/*
* Note: none of the below have been implemented yet.
 */
func (r *PlasmidResolver) Keywords(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
func (r *PlasmidResolver) GenbankAccession(ctx context.Context, obj *models.Plasmid) (*string, error) {
	s := ""
	return &s, nil
}

func ConvertToPlasmidModel(id string, attr *pb.PlasmidAttributes) *models.Plasmid {
	return &models.Plasmid{
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
		ImageMap:        &attr.ImageMap,
		Sequence:        &attr.Sequence,
		Name:            attr.Name,
	}
}
