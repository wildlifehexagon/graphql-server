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

type UserResolver struct {
	Client pb.UserServiceClient
	Logger *logrus.Entry
}

func (r *UserResolver) ID(ctx context.Context, obj *pb.User) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *UserResolver) FirstName(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.FirstName, nil
}
func (r *UserResolver) LastName(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.LastName, nil
}
func (r *UserResolver) Email(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.Email, nil
}
func (r *UserResolver) Organization(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Organization, nil
}
func (r *UserResolver) GroupName(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.GroupName, nil
}
func (r *UserResolver) FirstAddress(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.FirstAddress, nil
}
func (r *UserResolver) SecondAddress(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.SecondAddress, nil
}
func (r *UserResolver) City(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.City, nil
}
func (r *UserResolver) State(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.State, nil
}
func (r *UserResolver) Zipcode(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Zipcode, nil
}
func (r *UserResolver) Country(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Country, nil
}
func (r *UserResolver) Phone(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Phone, nil
}
func (r *UserResolver) IsActive(ctx context.Context, obj *pb.User) (bool, error) {
	return obj.Data.Attributes.IsActive, nil
}
func (r *UserResolver) CreatedAt(ctx context.Context, obj *pb.User) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *UserResolver) UpdatedAt(ctx context.Context, obj *pb.User) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *UserResolver) Roles(ctx context.Context, obj *pb.User) ([]*pb.Role, error) {
	roles := []*pb.Role{}
	rr, err := r.Client.GetRelatedRoles(ctx, &jsonapi.RelationshipRequest{Id: obj.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return roles, err
	}
	for _, n := range rr.Data {
		item := &pb.Role{
			Data: &pb.RoleData{
				Type: "role",
				Id:   n.Id,
				Attributes: &pb.RoleAttributes{
					Role:        n.Attributes.Role,
					Description: n.Attributes.Description,
					CreatedAt:   n.Attributes.CreatedAt,
					UpdatedAt:   n.Attributes.UpdatedAt,
				},
			},
		}
		roles = append(roles, item)
	}
	r.Logger.Infof("successfully retrieved list of %d roles for user ID %d", len(roles), obj.Data.Id)
	return roles, nil
}
