package resolver

import (
	"context"
	"time"

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
	panic("not implemented")
}
func (r *queryResolver) ListPermissions(ctx context.Context) ([]user.Permission, error) {
	panic("not implemented")
}

func (r *permissionResolver) ID(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}
func (r *permissionResolver) Permission(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}
func (r *permissionResolver) Description(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}
func (r *permissionResolver) CreatedAt(ctx context.Context, obj *user.Permission) (time.Time, error) {
	panic("not implemented")
}
func (r *permissionResolver) UpdatedAt(ctx context.Context, obj *user.Permission) (time.Time, error) {
	panic("not implemented")
}
func (r *permissionResolver) Resource(ctx context.Context, obj *user.Permission) (*string, error) {
	panic("not implemented")
}
