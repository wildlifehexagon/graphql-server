package stock

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/sirupsen/logrus"
)

type PlasmidResolver struct {
	Client pb.StockServiceClient
	Logger *logrus.Entry
}

func (r *PlasmidResolver) ID(ctx context.Context, obj *pb.Stock) (string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) CreatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) UpdatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) CreatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) UpdatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Summary(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) EditableSummary(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Depositor(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Genes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Dbxrefs(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Publications(ctx context.Context, obj *pb.Stock) ([]*publication.Publication, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) ImageMap(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Sequence(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) InStock(ctx context.Context, obj *pb.Stock) (bool, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) Keywords(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *PlasmidResolver) GenbankAccession(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
