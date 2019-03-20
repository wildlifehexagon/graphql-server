package order

import (
	"context"
	"fmt"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
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
	return obj.Data.Id, nil
}
func (r *OrderResolver) CreatedAt(ctx context.Context, obj *pb.Order) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *OrderResolver) UpdatedAt(ctx context.Context, obj *pb.Order) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *OrderResolver) Courier(ctx context.Context, obj *pb.Order) (*string, error) {
	return &obj.Data.Attributes.Courier, nil
}
func (r *OrderResolver) CourierAccount(ctx context.Context, obj *pb.Order) (*string, error) {
	return &obj.Data.Attributes.CourierAccount, nil
}
func (r *OrderResolver) Comments(ctx context.Context, obj *pb.Order) (*string, error) {
	return &obj.Data.Attributes.Comments, nil
}
func (r *OrderResolver) Payment(ctx context.Context, obj *pb.Order) (*string, error) {
	return &obj.Data.Attributes.Payment, nil
}
func (r *OrderResolver) PurchaseOrderNum(ctx context.Context, obj *pb.Order) (*string, error) {
	return &obj.Data.Attributes.PurchaseOrderNum, nil
}
func (r *OrderResolver) Status(ctx context.Context, obj *pb.Order) (*string, error) {
	panic("not implemented")
}
func (r *OrderResolver) Consumer(ctx context.Context, obj *pb.Order) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.Consumer
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return &user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *OrderResolver) Payer(ctx context.Context, obj *pb.Order) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.Payer
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return &user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *OrderResolver) Purchaser(ctx context.Context, obj *pb.Order) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.Purchaser
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return &user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *OrderResolver) Items(ctx context.Context, obj *pb.Order) ([]models.Stock, error) {
	panic("not implemented")
}
