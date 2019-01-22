package resolver

import (
	"context"
	"strconv"

	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	UserClient user.UserServiceClient
	Logger     *logrus.Entry
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*models.User, error) {
	user := &models.User{
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
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*models.User, error) {
	// Steps:
	// 1) Get user by ID, check if it exists
	// 2) If it exists, then update with given input
	// 3) Save to database
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteItem, error) {
	// need to remove user by given id
	panic("not implemented")
	// return true
}
func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*models.Role, error) {
	role := &models.Role{
		Role:        input.Role,
		Description: input.Description,
	}
	// need to save to database
	return role, nil
}
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*models.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*models.Permission, error) {
	resource := input.Resource
	permission := &models.Permission{
		Permission:  input.Permission,
		Description: input.Description,
		Resource:    &resource,
	}
	// need to save to database
	return permission, nil
}
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*models.Permission, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	// Need to improve error handling
	// Also need to handle timestamps and roles
	// And verify that this indeed works
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	g, err := r.UserClient.GetUser(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("Error in getting user by ID: %s", err)
		return nil, err
	}
	attr := g.Data.Attributes
	return &models.User{
		ID:            strconv.Itoa(int(g.Data.Id)),
		FirstName:     attr.FirstName,
		LastName:      attr.LastName,
		Email:         attr.Email,
		Organization:  &attr.Organization,
		FirstAddress:  &attr.FirstAddress,
		SecondAddress: &attr.SecondAddress,
		City:          &attr.City,
		State:         &attr.State,
		Zipcode:       &attr.Zipcode,
		Country:       &attr.Country,
		Phone:         &attr.Phone,
		IsActive:      attr.IsActive,
		// CreatedAt:     attr.CreatedAt,
		// UpdatedAt:     attr.UpdatedAt,
		// Roles:         &attr.Roles,
	}, nil
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
