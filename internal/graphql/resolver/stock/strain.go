package stock

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type StrainResolver struct {
	Client pb.StockServiceClient
	Logger *logrus.Entry
}

func (r *StrainResolver) ID(ctx context.Context, obj *pb.Stock) (string, error) {
	panic("not implemented")
}
func (r *StrainResolver) CreatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	panic("not implemented")
}
func (r *StrainResolver) UpdatedAt(ctx context.Context, obj *pb.Stock) (time.Time, error) {
	panic("not implemented")
}
func (r *StrainResolver) CreatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	panic("not implemented")
}
func (r *StrainResolver) UpdatedBy(ctx context.Context, obj *pb.Stock) (user.User, error) {
	panic("not implemented")
}
func (r *StrainResolver) Summary(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) EditableSummary(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Depositor(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Genes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Dbxrefs(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Publications(ctx context.Context, obj *pb.Stock) ([]*publication.Publication, error) {
	panic("not implemented")
}
func (r *StrainResolver) SystematicName(ctx context.Context, obj *pb.Stock) (string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Descriptor(ctx context.Context, obj *pb.Stock) (string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Species(ctx context.Context, obj *pb.Stock) (string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Plasmid(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Parent(ctx context.Context, obj *pb.Stock) (*pb.Stock, error) {
	panic("not implemented")
}
func (r *StrainResolver) Names(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) InStock(ctx context.Context, obj *pb.Stock) (bool, error) {
	panic("not implemented")
}
func (r *StrainResolver) Phenotypes(ctx context.Context, obj *pb.Stock) ([]*models.Phenotype, error) {
	panic("not implemented")
}
func (r *StrainResolver) GeneticModification(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) MutagenesisMethod(ctx context.Context, obj *pb.Stock) (*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Characteristics(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
func (r *StrainResolver) Genotypes(ctx context.Context, obj *pb.Stock) ([]*string, error) {
	panic("not implemented")
}
