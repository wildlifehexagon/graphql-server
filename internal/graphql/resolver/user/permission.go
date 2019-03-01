package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*models.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*models.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}

func (r *queryResolver) Permission(ctx context.Context, id string) (*models.Permission, error) {
	panic("not implemented")
}
func (r *queryResolver) ListPermissions(ctx context.Context) ([]models.Permission, error) {
	panic("not implemented")
}
