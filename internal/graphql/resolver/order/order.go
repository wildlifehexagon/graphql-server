package order

import (
	"context"
	"time"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type OrderResolver struct {
	Client      pb.OrderServiceClient
	StockClient stock.StockServiceClient
	UserClient  user.UserServiceClient
	Logger      *logrus.Entry
}

func (r *OrderResolver) ID(ctx context.Context, obj *pb.Order) (string, error) {
	panic("not implemented")
}
func (r *OrderResolver) CreatedAt(ctx context.Context, obj *pb.Order) (time.Time, error) {
	panic("not implemented")
}
func (r *OrderResolver) UpdatedAt(ctx context.Context, obj *pb.Order) (time.Time, error) {
	panic("not implemented")
}
func (r *OrderResolver) Courier(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) CourerAccount(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) Comments(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) Payment(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) PurchaseOrderNum(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) Status(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) Consumer(ctx context.Context, obj *pb.Order) (*user.User, error) {
	panic("not implemented")
}
func (r *OrderResolver) Payer(ctx context.Context, obj *pb.Order) (*user.User, error) {
	panic("not implemented")
}
func (r *OrderResolver) Purchaser(ctx context.Context, obj *pb.Order) (*user.User, error) {
	panic("not implemented")
}
func (r *OrderResolver) Items(ctx context.Context, obj *pb.Order) ([]*models.Stock, error) {
	panic("not implemented")
}
