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
	attr := &pb.NewOrderAttributes{}
	if input.Comments != nil {
		attr.Comments = input.Comments
	}
	attr.Consumer = input.Consumer
	attr.Courier = input.Courier
	attr.CourierAccount = input.CourierAccount
	attr.Items = input.Items
	attr.Payer = input.Payer
	attr.Payment = input.Payment
	attr.PurchaseOrderNum = input.PurchaseOrderNum
	attr.Purchaser = input.Purchaser
	attr.Status = input.Status
	o, err := m.GetOrderClient(registry.ORDER).CreateOrder(ctx, &pb.NewOrder{
		Data: &pb.NewOrder_Data{
			Type:       "order",
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating new order %s", err)
	}
	m.Logger.Debugf("successfully created new order with ID %s", o.Data.Id)
	return o, nil
}

// UpdateOrder updates an existing stock order.
func (m *MutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*pb.Order, error) {
	_, err := q.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting order with id %s: %s", id, err)
	}
	attr := &pb.OrderUpdateAttributes{}
	norm := normalizeUpdateOrderAttr(input)
	mapstructure.Decode(norm, attr)
	o, err := m.GetOrderClient(registry.ORDER).UpdateOrder(ctx, &pb.OrderUpdate{
		Data: &pb.OrderUpdate_Data{
			Type:       "order",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error updating order %s: %s", o.Data.Id, err)
	}
	u, err := q.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting order with id %s: %s", id, err)
	}
	m.Logger.Debugf("successfully updated order with ID %s", u.Data.Id)
	return u, nil
}

func normalizeUpdateOrderAttr(attr *models.UpdateOrderInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
				newAttr[k.Name()] = k.Value()
		}
	}
	return newAttr
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
