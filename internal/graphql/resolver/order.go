package resolver

import (
	"context"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// CreateOrder creates a new stock order.
func (m *MutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*pb.Order, error) {
	attr := &pb.NewOrderAttributes{}
	if input.Comments != nil {
		attr.Comments = *input.Comments
	}
	attr.Consumer = input.Consumer
	attr.Courier = input.Courier
	attr.CourierAccount = input.CourierAccount
	attr.Items = convertPtrToStr(input.Items)
	attr.Payer = input.Payer
	attr.Payment = input.Payment
	attr.PurchaseOrderNum = *input.PurchaseOrderNum
	attr.Purchaser = input.Purchaser
	attr.Status = statusConverter(input.Status)
	o, err := m.GetOrderClient(registry.ORDER).CreateOrder(ctx, &pb.NewOrder{
		Data: &pb.NewOrder_Data{
			Type:       "order",
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully created new order with ID %s", o.Data.Id)
	return o, nil
}

// statusConverter converts the enum status string to protocol buffer int32 value
func statusConverter(e models.StatusEnum) pb.OrderStatus {
	var status pb.OrderStatus
	switch e {
	case "IN_PREPARATION":
		status = pb.OrderStatus_In_preparation
	case "GROWING":
		status = pb.OrderStatus_Growing
	case "CANCELLED":
		status = pb.OrderStatus_Cancelled
	case "SHIPPED":
		status = pb.OrderStatus_Shipped
	}
	return status
}

// convertPtrToStr converts a slice of string pointers to a slice of strings
func convertPtrToStr(items []*string) []string {
	var sl []string
	for _, n := range items {
		sl = append(sl, *n)
	}
	return sl
}

// UpdateOrder updates an existing stock order.
func (m *MutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*pb.Order, error) {
	g, err := m.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	attr := &pb.OrderUpdateAttributes{}
	norm := normalizeUpdateOrderAttr(input)
	err = mapstructure.Decode(norm, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
	if input.Status != nil {
		attr.Status = statusConverter(*input.Status)
	} else {
		attr.Status = g.Data.Attributes.Status
	}
	o, err := m.GetOrderClient(registry.ORDER).UpdateOrder(ctx, &pb.OrderUpdate{
		Data: &pb.OrderUpdate_Data{
			Type:       "order",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	u, err := m.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: o.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully updated order with ID %s", u.Data.Id)
	return u, nil
}

func normalizeUpdateOrderAttr(attr *models.UpdateOrderInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if k.Name() == "Status" {
			newAttr["Status"] = statusConverter(*attr.Status)
		} else if !k.IsZero() {
			newAttr[k.Name()] = k.Value()
		}
	}
	return newAttr
}

// Order retrieves an individual order by ID.
func (q *QueryResolver) Order(ctx context.Context, id string) (*pb.Order, error) {
	g, err := q.GetOrderClient(registry.ORDER).GetOrder(ctx, &pb.OrderId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
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
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	orders := []*pb.Order{}
	for _, n := range list.Data {
		item := &pb.Order{
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
