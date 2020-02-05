package content

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/sirupsen/logrus"
)

type ContentResolver struct {
	Client     pb.ContentServiceClient
	UserClient user.UserServiceClient
	Logger     *logrus.Entry
}

func (r *ContentResolver) ID(ctx context.Context, obj *pb.Content) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *ContentResolver) Name(ctx context.Context, obj *pb.Content) (string, error) {
	return obj.Data.Attributes.Name, nil
}
func (r *ContentResolver) Slug(ctx context.Context, obj *pb.Content) (string, error) {
	return obj.Data.Attributes.Slug, nil
}
func (r *ContentResolver) CreatedBy(ctx context.Context, obj *pb.Content) (*user.User, error) {
	user := user.User{}
	id := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUser(ctx, &jsonapi.GetRequest{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with id %d", id)
	return g, nil
}
func (r *ContentResolver) UpdatedBy(ctx context.Context, obj *pb.Content) (*user.User, error) {
	user := user.User{}
	id := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUser(ctx, &jsonapi.GetRequest{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with id %d", id)
	return g, nil
}
func (r *ContentResolver) CreatedAt(ctx context.Context, obj *pb.Content) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *ContentResolver) UpdatedAt(ctx context.Context, obj *pb.Content) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *ContentResolver) Content(ctx context.Context, obj *pb.Content) (string, error) {
	return obj.Data.Attributes.Content, nil
}
func (r *ContentResolver) Namespace(ctx context.Context, obj *pb.Content) (string, error) {
	return obj.Data.Attributes.Namespace, nil
}
