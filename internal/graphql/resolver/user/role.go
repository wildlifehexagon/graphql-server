package user

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/sirupsen/logrus"
)

type RoleResolver struct {
	Client pb.RoleServiceClient
	Logger *logrus.Entry
}

func (r *RoleResolver) ID(ctx context.Context, obj *pb.Role) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *RoleResolver) Role(ctx context.Context, obj *pb.Role) (string, error) {
	return obj.Data.Attributes.Role, nil
}
func (r *RoleResolver) Description(ctx context.Context, obj *pb.Role) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *RoleResolver) CreatedAt(ctx context.Context, obj *pb.Role) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *RoleResolver) UpdatedAt(ctx context.Context, obj *pb.Role) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *RoleResolver) Permissions(ctx context.Context, obj *pb.Role) ([]*pb.Permission, error) {
	permissions := []*pb.Permission{}
	rp, err := r.Client.GetRelatedPermissions(ctx, &jsonapi.RelationshipRequest{Id: obj.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return permissions, err
	}
	for _, n := range rp.Data {
		item := &pb.Permission{
			Data: &pb.PermissionData{
				Type: "permission",
				Id:   n.Id,
				Attributes: &pb.PermissionAttributes{
					Permission:  n.Attributes.Permission,
					Description: n.Attributes.Description,
					CreatedAt:   n.Attributes.CreatedAt,
					UpdatedAt:   n.Attributes.UpdatedAt,
					Resource:    n.Attributes.Resource,
				},
			},
		}
		permissions = append(permissions, item)
	}
	r.Logger.Infof("successfully retrieved list of %d permissions for role ID %d", len(permissions), obj.Data.Id)
	return permissions, nil
}
