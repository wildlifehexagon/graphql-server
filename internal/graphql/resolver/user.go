package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

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
	attr := &user.UserAttributes{}
	a := normalizeCreateUserAttr(input)
	mapstructure.Decode(a, attr)
	n, err := r.UserClient.CreateUser(context.Background(), &user.CreateUserRequest{
		Data: &user.CreateUserRequest_Data{
			Type: "user",
			Attributes: &user.UserAttributes{
				FirstName:     attr.FirstName,
				LastName:      attr.LastName,
				Email:         attr.Email,
				Organization:  attr.Organization,
				GroupName:     attr.GroupName,
				FirstAddress:  attr.FirstAddress,
				SecondAddress: attr.SecondAddress,
				City:          attr.City,
				State:         attr.State,
				Zipcode:       attr.Zipcode,
				Country:       attr.Country,
				Phone:         attr.Phone,
				IsActive:      attr.IsActive,
			},
		},
	})
	if err != nil {
		r.Logger.Errorf("Error creating new user from mutation resolver: %s", err)
		return nil, err
	}
	id := strconv.FormatInt(n.Data.Id, 10)
	user := &models.User{
		ID:            id,
		FirstName:     n.Data.Attributes.FirstName,
		LastName:      n.Data.Attributes.LastName,
		Email:         n.Data.Attributes.Email,
		Organization:  &n.Data.Attributes.Organization,
		GroupName:     &n.Data.Attributes.GroupName,
		FirstAddress:  &n.Data.Attributes.FirstAddress,
		SecondAddress: &n.Data.Attributes.SecondAddress,
		City:          &n.Data.Attributes.City,
		State:         &n.Data.Attributes.State,
		Zipcode:       &n.Data.Attributes.Zipcode,
		Country:       &n.Data.Attributes.Country,
		Phone:         &n.Data.Attributes.Phone,
		IsActive:      n.Data.Attributes.IsActive,
		CreatedAt:     aphgrpc.ProtoTimeStamp(n.Data.Attributes.CreatedAt),
		UpdatedAt:     aphgrpc.ProtoTimeStamp(n.Data.Attributes.UpdatedAt),
		// Roles:     &n.Data.Attributes.Roles,
	}
	return user, nil
}

func normalizeCreateUserAttr(attr *models.CreateUserInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			newAttr[k.Name()] = k.Value()
		} else {
			newAttr[k.Name()] = ""
		}
	}
	return newAttr
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*models.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	attr := &user.UserAttributes{}
	if input.FirstName != nil {
		attr.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		attr.LastName = *input.LastName
	}
	if input.Organization != nil {
		attr.Organization = *input.Organization
	}
	if input.GroupName != nil {
		attr.GroupName = *input.GroupName
	}
	if input.FirstAddress != nil {
		attr.FirstAddress = *input.FirstAddress
	}
	if input.SecondAddress != nil {
		attr.SecondAddress = *input.SecondAddress
	}
	if input.City != nil {
		attr.City = *input.City
	}
	if input.State != nil {
		attr.State = *input.State
	}
	if input.Zipcode != nil {
		attr.Zipcode = *input.Zipcode
	}
	if input.Country != nil {
		attr.Country = *input.Country
	}
	if input.Phone != nil {
		attr.Phone = *input.Phone
	}
	if input.IsActive != nil {
		attr.IsActive = *input.IsActive
	}
	attr.UpdatedAt = aphgrpc.TimestampProto(time.Now())
	n, err := r.UserClient.UpdateUser(context.Background(), &user.UpdateUserRequest{
		Id: i,
		Data: &user.UpdateUserRequest_Data{
			Id:         i,
			Type:       "user",
			Attributes: attr,
		},
	})
	if err != nil {
		r.Logger.Errorf("Error updating user %d from mutation resolver: %s", n.Data.Id, err)
		return nil, err
	}
	o, err := r.UserClient.GetUser(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("Error fetching recently updated user from mutation resolver: %s", err)
		return nil, err
	}
	uid := strconv.FormatInt(o.Data.Id, 10)
	user := &models.User{
		ID:            uid,
		FirstName:     o.Data.Attributes.FirstName,
		LastName:      o.Data.Attributes.LastName,
		Email:         o.Data.Attributes.Email,
		Organization:  &o.Data.Attributes.Organization,
		GroupName:     &o.Data.Attributes.GroupName,
		FirstAddress:  &o.Data.Attributes.FirstAddress,
		SecondAddress: &o.Data.Attributes.SecondAddress,
		City:          &o.Data.Attributes.City,
		State:         &o.Data.Attributes.State,
		Zipcode:       &o.Data.Attributes.Zipcode,
		Country:       &o.Data.Attributes.Country,
		Phone:         &o.Data.Attributes.Phone,
		IsActive:      o.Data.Attributes.IsActive,
		CreatedAt:     aphgrpc.ProtoTimeStamp(o.Data.Attributes.CreatedAt),
		UpdatedAt:     aphgrpc.ProtoTimeStamp(o.Data.Attributes.UpdatedAt),
		// Roles:         &o.Data.Attributes.Roles,
	}
	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteItem, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	if _, err := r.UserClient.DeleteUser(context.Background(), &jsonapi.DeleteRequest{Id: i}); err != nil {
		r.Logger.Errorf("Error deleting user from mutation resolver: %s", err)
		return nil, err
	}

	return &models.DeleteItem{
		Success: true,
	}, nil
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

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	g, err := r.UserClient.GetUser(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("Error in getting user by ID %d: %s", i, err)
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
		CreatedAt:     aphgrpc.ProtoTimeStamp(attr.CreatedAt),
		UpdatedAt:     aphgrpc.ProtoTimeStamp(attr.UpdatedAt),
		// Roles:         &attr.Roles,
	}, nil
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	g, err := r.UserClient.GetUserByEmail(context.Background(), &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		r.Logger.Errorf("Error in getting user by email: %s", err)
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
		CreatedAt:     aphgrpc.ProtoTimeStamp(attr.CreatedAt),
		UpdatedAt:     aphgrpc.ProtoTimeStamp(attr.UpdatedAt),
		// Roles:         &attr.Roles,
	}, nil
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
