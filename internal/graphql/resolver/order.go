package resolver

import (
	"context"
	"fmt"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/order"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
)

// CreateOrder creates a new stock order.
func (m *MutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*pb.Order, error) {
	panic("not implemented")
}

// UpdateOrder updates an existing stock order.
func (m *MutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*pb.Order, error) {
	panic("not implemented")
}

// Order retrieves an individual order by ID.
func (q *QueryResolver) Order(ctx context.Context, id string) (*pb.Order, error) {
	g, err := q.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting order with id %s: %s", id, err)
	}
	q.Logger.Debugf("successfully found order with id %s", id)
	return g, nil
}

// ListOrders retrieves all orders in the database.
func (q *QueryResolver) ListOrders(ctx context.Context, input *models.ListOrderInput) (*models.OrderListWithCursor, error) {
	var cursor, limit int64
	var filter string
	if input.Cursor != nil {
		cursor = int64(*input.Cursor)
	} else {
		cursor = 0
	}
	if input.Limit != nil {
		limit = int64(*input.Limit)
	} else {
		limit = 10
	}
	if input.Filter != nil {
		filter = *input.Filter
	} else {
		filter = ""
	}
	list, err := q.GetOrderClient(registry.ORDER).ListOrders(ctx, &pb.ListParameters{Cursor: cursor, Limit: limit, Filter: filter})
	if err != nil {
		return nil, fmt.Errorf("error in getting list of orders %s", err)
	}
	orders := []pb.Order{}
	for _, n := range list.Data {
		item := pb.Order{
			Data: &pb.Order_Data{
				Type:       n.Type,
				Id:         n.Id,
				Attributes: n.Attributes,
			},
		}
		orders = append(orders, item)
	}
	l := int(limit)
	return &models.OrderListWithCursor{
		Orders:         orders,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(cursor),
		Limit:          &l,
		TotalCount:     len(orders),
	}, nil
}
