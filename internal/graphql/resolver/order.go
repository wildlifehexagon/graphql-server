package resolver

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *Resolver) Order() generated.OrderResolver {
	return &orderResolver{r}
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*order.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*order.Order, error) {
	panic("not implemented")
}

type orderResolver struct{ *Resolver }

func (r *orderResolver) ID(ctx context.Context, obj *order.Order) (string, error) {
	panic("not implemented")
}
func (r *orderResolver) CreatedAt(ctx context.Context, obj *order.Order) (time.Time, error) {
	panic("not implemented")
}
func (r *orderResolver) UpdatedAt(ctx context.Context, obj *order.Order) (time.Time, error) {
	panic("not implemented")
}
func (r *orderResolver) Courier(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) CourerAccount(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Comments(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Payment(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) PurchaseOrderNum(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Status(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Consumer(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Payer(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Purchaser(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}
func (r *orderResolver) Items(ctx context.Context, obj *order.Order) ([]*string, error) {
	panic("not implemented")
}

func (r *queryResolver) Order(ctx context.Context, id string) (*order.Order, error) {
	panic("not implemented")
}
func (r *queryResolver) ListOrders(ctx context.Context, cursor *string, limit *int, filter *string) (*models.OrderListWithCursor, error) {
	panic("not implemented")
}
