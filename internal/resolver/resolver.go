package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/generated"
	"github.com/dictyBase/graphql-server/internal/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAnnotation(ctx context.Context, input *models.CreateAnnotationInput) (*models.Annotation, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAnnotation(ctx context.Context, id string, input *models.UpdateAnnotationInput) (*models.Annotation, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAnnotation(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*models.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*models.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePublication(ctx context.Context, input *models.CreatePublicationInput) (*models.Publication, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePublication(ctx context.Context, id string, input *models.UpdatePublicationInput) (*models.Publication, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePublication(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*models.Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePlasmid(ctx context.Context, input *models.CreatePlasmidInput) (*models.Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateStock(ctx context.Context, id string, input *models.UpdateStockInput) (*models.Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteStock(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*models.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*models.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*models.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*models.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Annotation(ctx context.Context, id string) (*models.Annotation, error) {
	panic("not implemented")
}
func (r *queryResolver) AnnotationByEntry(ctx context.Context, tag string, entry_id string, ontology string, rank *int, is_obsolete *bool) (*models.Annotation, error) {
	panic("not implemented")
}
func (r *queryResolver) ListAnnotations(ctx context.Context, cursor *string, limit *int, filter *string) (*models.AnnotationListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	panic("not implemented")
}
func (r *queryResolver) ListOrders(ctx context.Context, cursor *string, limit *int, filter *string) (*models.OrderListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Publication(ctx context.Context, id string) (*models.Publication, error) {
	panic("not implemented")
}
func (r *queryResolver) ListPublications(ctx context.Context, cursor *string, limit *int, filter *string) (*models.PublicationListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Stock(ctx context.Context, id string) (*models.Stock, error) {
	panic("not implemented")
}
func (r *queryResolver) ListStocks(ctx context.Context, cursor *string, limit *int, filter *string) (*models.StockListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) ListUsers(ctx context.Context, cursor *string, limit *int, filter *string) (*models.UserListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Role(ctx context.Context, id string) (*models.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) ListRoles(ctx context.Context) ([]models.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) Permission(ctx context.Context, id string) (*models.Permission, error) {
	panic("not implemented")
}
func (r *queryResolver) ListPermissions(ctx context.Context) ([]models.Permission, error) {
	panic("not implemented")
}
