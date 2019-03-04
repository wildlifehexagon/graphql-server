package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *Resolver) Permission() generated.PermissionResolver {
	return &permissionResolver{r}
}

type permissionResolver struct{ *Resolver }

func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*user.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*user.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}

func (r *queryResolver) Permission(ctx context.Context, id string) (*user.Permission, error) {
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
func (r *queryResolver) ListPermissions(ctx context.Context) ([]user.Permission, error) {
	panic("not implemented")
}

func (r *permissionResolver) ID(ctx context.Context, obj *user.Permission) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *permissionResolver) Permission(ctx context.Context, obj *user.Permission) (string, error) {
	return obj.Data.Attributes.Permission, nil
}
func (r *permissionResolver) Description(ctx context.Context, obj *user.Permission) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *permissionResolver) CreatedAt(ctx context.Context, obj *user.Permission) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *permissionResolver) UpdatedAt(ctx context.Context, obj *user.Permission) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *permissionResolver) Resource(ctx context.Context, obj *user.Permission) (*string, error) {
	return &obj.Data.Attributes.Resource, nil
}
