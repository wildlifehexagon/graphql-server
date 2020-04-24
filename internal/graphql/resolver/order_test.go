package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/order"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	o, err := ord.Order(context.Background(), "999")
	assert.NoError(err, "expect no error from getting order information")
	assert.Exactly(o.Data.Id, "999", "should match id")
	assert.Exactly(o.Data.Attributes.Courier, "USPS", "should match courier")
	assert.Exactly(o.Data.Attributes.CourierAccount, "123456", "should match courier account")
	assert.Exactly(o.Data.Attributes.Comments, "first order", "should match comments")
	assert.Exactly(o.Data.Attributes.Payment, "credit", "should match payment")
	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, "987654", "should match purchase order number")
	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_In_preparation, "should match status")
	assert.Exactly(o.Data.Attributes.Consumer, "art@vandelayindustries.com", "should match consumer")
	assert.Exactly(o.Data.Attributes.Payer, "george@costanza.com", "should match payer")
	assert.Exactly(o.Data.Attributes.Purchaser, "thatsgold@jerry.org", "should match purchaser")
	assert.ElementsMatch(o.Data.Attributes.Items, []string{"DBS123456"}, "should match items")
}

func TestListOrders(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	filter := "type===strain"
	o, err := ord.ListOrders(context.Background(), &models.ListOrderInput{
		Cursor: &cursor,
		Limit:  &limit,
		Filter: &filter,
	})
	assert.NoError(err, "expect no error from getting list of orders")
	assert.Exactly(o.Limit, &limit, "should match limit")
	assert.Exactly(o.PreviousCursor, 0, "should match previous cursor")
	assert.Exactly(o.NextCursor, 10000, "should match next cursor")
	assert.Exactly(o.TotalCount, 3, "should match total count (length) of items")
	assert.Len(o.Orders, 3, "should have three orders")
}

func TestCreateOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	comments := "first order"
	pon := "987654"
	id := "DBS123456"
	o, err := ord.CreateOrder(context.Background(), &models.CreateOrderInput{
		Courier:          "USPS",
		CourierAccount:   "123456",
		Comments:         &comments,
		Payment:          "credit",
		PurchaseOrderNum: &pon,
		Status:           models.StatusEnumInPreparation,
		Consumer:         "art@vandelayindustries.com",
		Payer:            "george@costanza.com",
		Purchaser:        "thatsgold@jerry.org",
		Items:            []*string{&id},
	})
	assert.NoError(err, "expect no error from creating an order")
	assert.Exactly(o.Data.Attributes.Courier, "USPS", "should match courier")
	assert.Exactly(o.Data.Attributes.CourierAccount, "123456", "should match courier account")
	assert.Exactly(o.Data.Attributes.Comments, "first order", "should match comments")
	assert.Exactly(o.Data.Attributes.Payment, "credit", "should match payment")
	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, "987654", "should match purchase order number")
	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_In_preparation, "should match status")
	assert.Exactly(o.Data.Attributes.Consumer, "art@vandelayindustries.com", "should match consumer")
	assert.Exactly(o.Data.Attributes.Payer, "george@costanza.com", "should match payer")
	assert.Exactly(o.Data.Attributes.Purchaser, "thatsgold@jerry.org", "should match purchaser")
	assert.ElementsMatch(o.Data.Attributes.Items, []string{"DBS123456"}, "should match items")
}

func TestUpdateOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	courier := "FedEx"
	courierAccount := "444444"
	comments := "Please send ASAP"
	status := models.StatusEnumGrowing
	o, err := ord.UpdateOrder(
		context.Background(),
		"999",
		&models.UpdateOrderInput{
			Courier:        &courier,
			CourierAccount: &courierAccount,
			Comments:       &comments,
			Status:         &status,
		},
	)
	assert.NoError(err, "expect no error from updating an order")
	assert.Exactly(o.Data.Attributes.Courier, courier, "should match updated courier")
	assert.Exactly(o.Data.Attributes.CourierAccount, courierAccount, "should match updated courier account")
	assert.Exactly(o.Data.Attributes.Comments, comments, "should match updated comments")
	assert.Exactly(o.Data.Attributes.Payment, "credit", "should match payment")
	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, "987654", "should match purchase order number")
	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_Growing, "should match updated status")
	assert.Exactly(o.Data.Attributes.Consumer, "art@vandelayindustries.com", "should match consumer")
	assert.Exactly(o.Data.Attributes.Payer, "george@costanza.com", "should match payer")
	assert.Exactly(o.Data.Attributes.Purchaser, "thatsgold@jerry.org", "should match purchaser")
	assert.ElementsMatch(o.Data.Attributes.Items, []string{"DBS123456"}, "should match items")
}
