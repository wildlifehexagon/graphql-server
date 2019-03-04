package resolver

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *Resolver) Role() generated.RoleResolver {
	return &roleResolver{r}
}

type roleResolver struct{ *Resolver }

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

func (r *roleResolver) ID(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}
func (r *roleResolver) Role(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}
func (r *roleResolver) Description(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}
func (r *roleResolver) CreatedAt(ctx context.Context, obj *user.Role) (time.Time, error) {
	panic("not implemented")
}
func (r *roleResolver) UpdatedAt(ctx context.Context, obj *user.Role) (time.Time, error) {
	panic("not implemented")
}
func (r *roleResolver) Permissions(ctx context.Context, obj *user.Role) ([]user.Permission, error) {
	panic("not implemented")
}
