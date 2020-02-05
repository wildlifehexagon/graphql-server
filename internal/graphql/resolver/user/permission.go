package user

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/aphgrpc"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/sirupsen/logrus"
)

type PermissionResolver struct {
	Client pb.PermissionServiceClient
	Logger *logrus.Entry
}

func (r *PermissionResolver) ID(ctx context.Context, obj *pb.Permission) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *PermissionResolver) Permission(ctx context.Context, obj *pb.Permission) (string, error) {
	return obj.Data.Attributes.Permission, nil
}
func (r *PermissionResolver) Description(ctx context.Context, obj *pb.Permission) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *PermissionResolver) CreatedAt(ctx context.Context, obj *pb.Permission) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *PermissionResolver) UpdatedAt(ctx context.Context, obj *pb.Permission) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *PermissionResolver) Resource(ctx context.Context, obj *pb.Permission) (*string, error) {
	return &obj.Data.Attributes.Resource, nil
}
