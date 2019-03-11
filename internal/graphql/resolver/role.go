package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *Resolver) Role() generated.RoleResolver {
	return &roleResolver{r}
}

type roleResolver struct{ *Resolver }

func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*pb.Role, error) {
	n, err := r.RoleClient.CreateRole(context.Background(), &pb.CreateRoleRequest{
		Data: &pb.CreateRoleRequest_Data{
			Type: "role",
			Attributes: &pb.RoleAttributes{
				Role:        input.Role,
				Description: input.Description,
			},
		},
	})
	if err != nil {
		r.Logger.Errorf("error creating new role %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully created new role with ID %d", n.Data.Id)
	return n, nil
}
func (r *mutationResolver) CreateRolePermissionRelationship(ctx context.Context, roleId string, permissionId string) (*pb.Role, error) {
	rid, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", roleId, err)
		return nil, err
	}
	pid, err := strconv.ParseInt(permissionId, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", permissionId, err)
		return nil, err
	}
	rr, err := r.RoleClient.CreatePermissionRelationship(ctx, &jsonapi.DataCollection{
		Id: rid,
		Data: []*jsonapi.Data{
			{
				Type: "permission",
				Id:   pid,
			},
		}})
	if err != nil {
		r.Logger.Errorf("error in creating permission relationship with role %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully created role ID %d relationship permission with ID %d %s", rid, pid, rr)
	g, err := r.RoleClient.GetRole(ctx, &jsonapi.GetRequest{Id: rid})
	if err != nil {
		r.Logger.Errorf("error in getting role by ID %d: %s", rid, err)
		return nil, err
	}
	return g, nil
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*pb.Role, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	n, err := r.RoleClient.UpdateRole(context.Background(), &pb.UpdateRoleRequest{
		Id: i,
		Data: &pb.UpdateRoleRequest_Data{
			Id:   i,
			Type: "role",
			Attributes: &pb.RoleAttributes{
				Role:        input.Role,
				Description: input.Description,
				UpdatedAt:   aphgrpc.TimestampProto(time.Now()),
			},
		},
	})
	if err != nil {
		r.Logger.Errorf("error updating role %d: %s", n.Data.Id, err)
		return nil, err
	}
	o, err := r.RoleClient.GetRole(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("error fetching recently updated role: %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully updated role with ID %d", n.Data.Id)
	return o, nil
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteItem, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	if _, err := r.RoleClient.DeleteRole(context.Background(), &jsonapi.DeleteRequest{Id: i}); err != nil {
		r.Logger.Errorf("error deleting role with ID %s: %s", id, err)
		return &models.DeleteItem{
			Success: false,
		}, err
	}
	r.Logger.Infof("successfully deleted role with ID %s", id)
	return &models.DeleteItem{
		Success: true,
	}, nil
}
func (r *queryResolver) Role(ctx context.Context, id string) (*pb.Role, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	g, err := r.RoleClient.GetRole(ctx, &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("error in getting role by ID %d: %s", i, err)
		return nil, err
	}
	r.Logger.Infof("successfully found role with ID %s", id)
	return g, nil
}
func (r *queryResolver) ListRoles(ctx context.Context) ([]pb.Role, error) {
	roles := []pb.Role{}
	l, err := r.RoleClient.ListRoles(ctx, &jsonapi.SimpleListRequest{})
	if err != nil {
		r.Logger.Errorf("error in listing roles %s", err)
		return nil, err
	}
	for _, n := range l.Data {
		item := pb.Role{
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
	r.Logger.Infof("successfully provided list of %d roles", len(roles))
	return roles, nil
}

func (r *roleResolver) ID(ctx context.Context, obj *pb.Role) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *roleResolver) Role(ctx context.Context, obj *pb.Role) (string, error) {
	return obj.Data.Attributes.Role, nil
}
func (r *roleResolver) Description(ctx context.Context, obj *pb.Role) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *roleResolver) CreatedAt(ctx context.Context, obj *pb.Role) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *roleResolver) UpdatedAt(ctx context.Context, obj *pb.Role) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *roleResolver) Permissions(ctx context.Context, obj *pb.Role) ([]pb.Permission, error) {
	permissions := []pb.Permission{}
	rp, err := r.RoleClient.GetRelatedPermissions(ctx, &jsonapi.RelationshipRequest{Id: obj.Data.Id})
	if err != nil {
		r.Logger.Errorf("error getting list of related permissions for role ID %d: %s", obj.Data.Id, err)
		return permissions, err
	}
	for _, n := range rp.Data {
		item := pb.Permission{
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
