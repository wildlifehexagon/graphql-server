package graph

import (
	"context"
)

type mutationResolver struct {
	client *Client
}

func (r *mutationResolver) CreateAnnotation(ctx context.Context, input *CreateAnnotationInput) (*Annotation, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAnnotation(ctx context.Context, id string, input *UpdateAnnotationInput) (*Annotation, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAnnotation(ctx context.Context, id string) (*DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateOrder(ctx context.Context, input *CreateOrderInput) (*Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateOrder(ctx context.Context, id string, input *UpdateOrderInput) (*Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePublication(ctx context.Context, input *CreatePublicationInput) (*Publication, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePublication(ctx context.Context, id string, input *UpdatePublicationInput) (*Publication, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePublication(ctx context.Context, id string) (*DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateStrain(ctx context.Context, input *CreateStrainInput) (*Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePlasmid(ctx context.Context, input *CreatePlasmidInput) (*Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateStock(ctx context.Context, id string, input *UpdateStockInput) (*Stock, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteStock(ctx context.Context, id string) (*DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input *CreateUserInput) (*User, error) {
	user := &User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Email:         input.Email,
		Organization:  input.Organization,
		GroupName:     input.GroupName,
		FirstAddress:  input.FirstAddress,
		SecondAddress: input.SecondAddress,
		City:          input.City,
		State:         input.State,
		Zipcode:       input.Zipcode,
		Country:       input.Country,
		Phone:         input.Phone,
		IsActive:      input.IsActive,
	}
	// need to save to database
	return user, nil
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *UpdateUserInput) (*User, error) {
	// Steps:
	// 1) Get user by ID, check if it exists
	// 2) If it exists, then update with given input
	// 3) Save to database
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*DeleteItem, error) {
	// need to remove user by given id
	panic("not implemented")
	// return true
}
func (r *mutationResolver) CreateRole(ctx context.Context, input *CreateRoleInput) (*Role, error) {
	role := &Role{
		Role:        input.Role,
		Description: input.Description,
	}
	// need to save to database
	return role, nil
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *UpdateRoleInput) (*Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePermission(ctx context.Context, input *CreatePermissionInput) (*Permission, error) {
	resource := input.Resource
	permission := &Permission{
		Permission:  input.Permission,
		Description: input.Description,
		Resource:    &resource,
	}
	// need to save to database
	return permission, nil
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *UpdatePermissionInput) (*Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*DeleteItem, error) {
	panic("not implemented")
}

type queryResolver struct {
	client *Client
}

func (r *queryResolver) Annotation(ctx context.Context, id string) (*Annotation, error) {
	panic("not implemented")
}
func (r *queryResolver) AnnotationByEntry(ctx context.Context, tag string, entry_id string, ontology string, rank *int, is_obsolete *bool) (*Annotation, error) {
	panic("not implemented")
}
func (r *queryResolver) ListAnnotations(ctx context.Context, cursor *string, limit *int, filter *string) (*AnnotationListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Order(ctx context.Context, id string) (*Order, error) {
	panic("not implemented")
}
func (r *queryResolver) ListOrders(ctx context.Context, cursor *string, limit *int, filter *string) (*OrderListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Publication(ctx context.Context, id string) (*Publication, error) {
	panic("not implemented")
}
func (r *queryResolver) ListPublications(ctx context.Context, cursor *string, limit *int, filter *string) (*PublicationListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Stock(ctx context.Context, id string) (*Stock, error) {
	panic("not implemented")
}
func (r *queryResolver) ListStocks(ctx context.Context, cursor *string, limit *int, filter *string) (*StockListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	panic("not implemented")
}
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*User, error) {
	panic("not implemented")
}
func (r *queryResolver) ListUsers(ctx context.Context, cursor *string, limit *int, filter *string) (*UserListWithCursor, error) {
	panic("not implemented")
}
func (r *queryResolver) Role(ctx context.Context, id string) (*Role, error) {
	panic("not implemented")
}
func (r *queryResolver) ListRoles(ctx context.Context) ([]Role, error) {
	panic("not implemented")
}
func (r *queryResolver) Permission(ctx context.Context, id string) (*Permission, error) {
	panic("not implemented")
}
func (r *queryResolver) ListPermissions(ctx context.Context) ([]Permission, error) {
	panic("not implemented")
}
