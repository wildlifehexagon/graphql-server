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

func (r *Resolver) Permission() generated.PermissionResolver {
	return &permissionResolver{r}
}

type permissionResolver struct{ *Resolver }

func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*pb.Permission, error) {
	n, err := r.PermissionClient.CreatePermission(context.Background(), &pb.CreatePermissionRequest{
		Data: &pb.CreatePermissionRequest_Data{
			Type: "permission",
			Attributes: &pb.PermissionAttributes{
				Permission:  input.Permission,
				Description: input.Description,
				Resource:    input.Resource,
			},
		},
	})
	if err != nil {
		r.Logger.Errorf("error creating new permission %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully created new permission with ID %d", n.Data.Id)
	return n, nil
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*pb.Permission, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	n, err := r.PermissionClient.UpdatePermission(context.Background(), &pb.UpdatePermissionRequest{
		Id: i,
		Data: &pb.UpdatePermissionRequest_Data{
			Id:   i,
			Type: "permission",
			Attributes: &pb.PermissionAttributes{
				Permission:  input.Permission,
				Description: input.Description,
				Resource:    input.Resource,
				UpdatedAt:   aphgrpc.TimestampProto(time.Now()),
			},
		},
	})
	if err != nil {
		r.Logger.Errorf("error updating permission %d: %s", n.Data.Id, err)
		return nil, err
	}
	o, err := r.PermissionClient.GetPermission(context.Background(), &jsonapi.GetRequestWithFields{Id: i})
	if err != nil {
		r.Logger.Errorf("error fetching recently updated permission: %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully updated permission with ID %d", n.Data.Id)
	return o, nil
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	if _, err := r.PermissionClient.DeletePermission(context.Background(), &jsonapi.DeleteRequest{Id: i}); err != nil {
		r.Logger.Errorf("error deleting permission with ID %s: %s", id, err)
		return &models.DeleteItem{
			Success: false,
		}, err
	}
	r.Logger.Infof("successfully deleted permission with ID %s", id)
	return &models.DeleteItem{
		Success: true,
	}, nil
}

func (r *queryResolver) Permission(ctx context.Context, id string) (*pb.Permission, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	g, err := r.PermissionClient.GetPermission(ctx, &jsonapi.GetRequestWithFields{Id: i})
	if err != nil {
		r.Logger.Errorf("error in getting permission by ID %d: %s", i, err)
		return nil, err
	}
	r.Logger.Infof("successfully found permission with ID %s", id)
	return g, nil
}
func (r *queryResolver) ListPermissions(ctx context.Context) ([]pb.Permission, error) {
	permissions := []pb.Permission{}
	l, err := r.PermissionClient.ListPermissions(ctx, &jsonapi.SimpleListRequest{})
	if err != nil {
		r.Logger.Errorf("error in listing permissions %s", err)
		return nil, err
	}
	for _, n := range l.Data {
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
	r.Logger.Infof("successfully provided list of %d permissions", len(permissions))
	return permissions, nil
}

func (r *permissionResolver) ID(ctx context.Context, obj *pb.Permission) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *permissionResolver) Permission(ctx context.Context, obj *pb.Permission) (string, error) {
	return obj.Data.Attributes.Permission, nil
}
func (r *permissionResolver) Description(ctx context.Context, obj *pb.Permission) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *permissionResolver) CreatedAt(ctx context.Context, obj *pb.Permission) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *permissionResolver) UpdatedAt(ctx context.Context, obj *pb.Permission) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *permissionResolver) Resource(ctx context.Context, obj *pb.Permission) (*string, error) {
	return &obj.Data.Attributes.Resource, nil
}
