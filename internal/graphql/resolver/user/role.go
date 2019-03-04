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

func (r *Resolver) Role() generated.RoleResolver {
	return &roleResolver{r}
}

type roleResolver struct{ *Resolver }

func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*user.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*user.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *queryResolver) Role(ctx context.Context, id string) (*user.Role, error) {
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
func (r *queryResolver) ListRoles(ctx context.Context) ([]user.Role, error) {
	panic("not implemented")
}

func (r *roleResolver) ID(ctx context.Context, obj *user.Role) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *roleResolver) Role(ctx context.Context, obj *user.Role) (string, error) {
	return obj.Data.Attributes.Role, nil
}
func (r *roleResolver) Description(ctx context.Context, obj *user.Role) (string, error) {
	return obj.Data.Attributes.Description, nil
}
func (r *roleResolver) CreatedAt(ctx context.Context, obj *user.Role) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *roleResolver) UpdatedAt(ctx context.Context, obj *user.Role) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *roleResolver) Permissions(ctx context.Context, obj *user.Role) ([]user.Permission, error) {
	panic("not implemented")
}
