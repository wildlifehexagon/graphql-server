package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*models.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*models.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *queryResolver) Role(ctx context.Context, id string) (*models.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) ListRoles(ctx context.Context) ([]models.Role, error) {
	panic("not implemented")
}
