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
	"github.com/sirupsen/logrus"
)

type PlasmidResolver struct {
	Client     pb.StockServiceClient
	UserClient user.UserServiceClient
	Logger     *logrus.Entry
}

func (r *PlasmidResolver) ID(ctx context.Context, obj *pb.Stock) (string, error) {
	return obj.Data.Id, nil
}
func (r *PlasmidResolver) CreatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *PlasmidResolver) UpdatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *PlasmidResolver) CreatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return *g, nil
}
func (r *PlasmidResolver) UpdatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return *g, nil
}
func (r *PlasmidResolver) Summary(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.Summary, nil
}
func (r *PlasmidResolver) EditableSummary(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.EditableSummary, nil
}
func (r *PlasmidResolver) Depositor(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.Depositor, nil
}
func (r *PlasmidResolver) Genes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	genes := []*string{}
	for _, n := range obj.Data.Attributes.Genes {
		genes = append(genes, &n)
	}
	return genes, nil
}
func (r *PlasmidResolver) Dbxrefs(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	dbxrefs := []*string{}
	for _, n := range obj.Data.Attributes.Dbxrefs {
		dbxrefs = append(dbxrefs, &n)
	}
	return dbxrefs, nil
}
func (r *PlasmidResolver) Publications(ctx context.Context, obj *pb.Stock) ([]*publication.Publication, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) ImageMap(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.PlasmidProperties.ImageMap, nil
}
func (r *PlasmidResolver) Sequence(ctx context.Context, obj *pb.Stock) (*string, error) {
	return &obj.Data.Attributes.PlasmidProperties.Sequence, nil
}

/*
* Note: none of the below have been implemented yet.
 */
func (r *PlasmidResolver) InStock(ctx context.Context, obj *pb.Stock) (bool, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Keywords(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
func (r *PlasmidResolver) GenbankAccession(ctx context.Context, obj *pb.Stock) (*string, error) {
	s := ""
	return &s, nil
}
