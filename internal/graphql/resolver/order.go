package resolver

import (
	"context"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/order"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (m *mutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*pb.Order, error) {
	panic("not implemented")
}
func (m *mutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*pb.Order, error) {
	panic("not implemented")
}

func (q *queryResolver) Order(ctx context.Context, id string) (*pb.Order, error) {
	panic("not implemented")
}
func (q *queryResolver) ListOrders(ctx context.Context, input *models.ListOrderInput) (*models.OrderListWithCursor, error) {
	panic("not implemented")
}
