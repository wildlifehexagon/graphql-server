package stock

import (
	"context"
	"fmt"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type StrainResolver struct {
	Client     pb.StockServiceClient
	UserClient user.UserServiceClient
	Logger     *logrus.Entry
}

func (r *StrainResolver) ID(ctx context.Context, obj *pb.Stock) (string, error) {
	return obj.Data.Id, nil
}
func (r *StrainResolver) CreatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *StrainResolver) UpdatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *StrainResolver) CreatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return *g, nil
}
func (r *StrainResolver) UpdatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return *g, nil
}
func (r *StrainResolver) Summary(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.Summary, nil
}
func (r *StrainResolver) EditableSummary(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.EditableSummary, nil
}
func (r *StrainResolver) Depositor(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.Depositor, nil
}
func (r *StrainResolver) Genes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	genes := []*string{}
	for _, n := range obj.Data.Attributes.Genes {
		genes = append(genes, &n)
	}
	return genes, nil
}
func (r *StrainResolver) Dbxrefs(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	dbxrefs := []*string{}
	for _, n := range obj.Data.Attributes.Dbxrefs {
		dbxrefs = append(dbxrefs, &n)
	}
	return dbxrefs, nil
}
func (r *StrainResolver) Publications(ctx context.Context, obj *pb.Stock) ([]*publication.Publication, error) {
	pub := []*publication.Publication{}
	// for _, n := range obj.Data.Attributes.Publications {

	// }
	return pub, nil
}
func (r *StrainResolver) SystematicName(ctx context.Context, obj *pb.Stock) (string, error) {
	return obj.Data.Attributes.StrainProperties.SystematicName, nil
}
func (r *StrainResolver) Descriptor(ctx context.Context, obj *pb.Stock) (string, error) {
	return obj.Data.Attributes.StrainProperties.Label, nil
}
func (r *StrainResolver) Species(ctx context.Context, obj *pb.Stock) (string, error) {
	return obj.Data.Attributes.StrainProperties.Species, nil
}
func (r *StrainResolver) Plasmid(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.StrainProperties.Plasmid, nil
}
func (r *StrainResolver) Parent(ctx context.Context, obj *pb.Stock) (*pb.Stock, error) {
	parent := obj.Data.Attributes.StrainProperties.Parent
	strain, err := r.Client.GetStock(ctx, &pb.StockId{Id: parent})
	if err != nil {
		return nil, fmt.Errorf("error in getting parent strain with ID %d: %s", parent, err)
	}
	r.Logger.Debugf("successfully found parent strain with ID %s", parent)
	return strain, nil
}
func (r *StrainResolver) Names(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	names := []*string{}
	for _, n := range obj.Data.Attributes.StrainProperties.Names {
		names = append(names, &n)
	}
	return names, nil
}

/*
* Note: none of the below have been implemented yet.
 */
func (r *StrainResolver) InStock(ctx context.Context, obj *pb.Stock) (bool, error) {
	panic("not implemented")
}
func (r *StrainResolver) Phenotypes(ctx context.Context, obj *pb.Stock) ([]*models.Phenotype, error) {
	return []*models.Phenotype{}, nil
}
func (r *StrainResolver) GeneticModification(ctx context.Context, obj *pb.Stock) (*string, error) {
	s := ""
	return &s, nil
}
func (r *StrainResolver) MutagenesisMethod(ctx context.Context, obj *pb.Stock) (*string, error) {
	s := ""
	return &s, nil
}
func (r *StrainResolver) Characteristics(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
func (r *StrainResolver) Genotypes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
