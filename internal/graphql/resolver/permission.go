package resolver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
)

func (m *MutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*pb.Permission, error) {
	n, err := m.GetPermissionClient(registry.PERMISSION).CreatePermission(context.Background(), &pb.CreatePermissionRequest{
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
		return nil, fmt.Errorf("error creating new permission %s", err)
	}
	m.Logger.Debugf("successfully created new permission with ID %d", n.Data.Id)
	return n, nil
}
func (m *MutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*pb.Permission, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	n, err := m.GetPermissionClient(registry.PERMISSION).UpdatePermission(context.Background(), &pb.UpdatePermissionRequest{
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
		return nil, fmt.Errorf("error updating permission %d: %s", n.Data.Id, err)
	}
	o, err := m.GetPermissionClient(registry.PERMISSION).GetPermission(context.Background(), &jsonapi.GetRequestWithFields{Id: i})
	if err != nil {
		return nil, fmt.Errorf("error fetching recently updated permission: %s", err)
	}
	m.Logger.Debugf("successfully updated permission with ID %d", n.Data.Id)
	return o, nil
}
func (m *MutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	if _, err := m.GetPermissionClient(registry.PERMISSION).DeletePermission(context.Background(), &jsonapi.DeleteRequest{Id: i}); err != nil {
		return &models.DeleteItem{
			Success: false,
		}, fmt.Errorf("error deleting permission with ID %s: %s", id, err)
	}
	m.Logger.Debugf("successfully deleted permission with ID %s", id)
	return &models.DeleteItem{
		Success: true,
	}, nil
}

func (q *QueryResolver) Permission(ctx context.Context, id string) (*pb.Permission, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	g, err := q.GetPermissionClient(registry.PERMISSION).GetPermission(ctx, &jsonapi.GetRequestWithFields{Id: i})
	if err != nil {
		return nil, fmt.Errorf("error in getting permission by ID %d: %s", i, err)
	}
	q.Logger.Debugf("successfully found permission with ID %s", id)
	return g, nil
}
func (q *QueryResolver) ListPermissions(ctx context.Context) ([]pb.Permission, error) {
	permissions := []pb.Permission{}
	l, err := q.GetPermissionClient(registry.PERMISSION).ListPermissions(ctx, &jsonapi.SimpleListRequest{})
	if err != nil {
		return nil, fmt.Errorf("error in listing permissions %s", err)
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
	q.Logger.Debugf("successfully provided list of %d permissions", len(permissions))
	return permissions, nil
}
