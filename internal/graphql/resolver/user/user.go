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
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	UserClient       user.UserServiceClient
	RoleClient       user.RoleServiceClient
	PermissionClient user.PermissionServiceClient
	Logger           *logrus.Entry
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*user.User, error) {
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
		r.Logger.Errorf("error creating new user %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully created new user with ID %d", n.Data.Id)
	return n, nil
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

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*user.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	f, err := r.UserClient.GetUser(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("error fetching user with ID %s %s", id, err)
		return nil, err
	}
	attr := getUpdateUserAttributes(input, f)
	attr.Email = f.Data.Attributes.Email
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
		r.Logger.Errorf("error updating user %d: %s", n.Data.Id, err)
		return nil, err
	}
	o, err := r.UserClient.GetUser(context.Background(), &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("error fetching recently updated user: %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully updated user with ID %d", n.Data.Id)
	return o, nil
}

func getUpdateUserAttributes(input *models.UpdateUserInput, f *user.User) *user.UserAttributes {
	attr := &user.UserAttributes{}
	if input.FirstName != nil {
		attr.FirstName = *input.FirstName
	} else {
		attr.FirstName = f.Data.Attributes.FirstName
	}
	if input.LastName != nil {
		attr.LastName = *input.LastName
	} else {
		attr.LastName = f.Data.Attributes.LastName
	}
	if input.Organization != nil {
		attr.Organization = *input.Organization
	} else {
		attr.Organization = f.Data.Attributes.Organization
	}
	if input.GroupName != nil {
		attr.GroupName = *input.GroupName
	} else {
		attr.GroupName = f.Data.Attributes.GroupName
	}
	if input.FirstAddress != nil {
		attr.FirstAddress = *input.FirstAddress
	} else {
		attr.FirstAddress = f.Data.Attributes.FirstAddress
	}
	if input.SecondAddress != nil {
		attr.SecondAddress = *input.SecondAddress
	} else {
		attr.SecondAddress = f.Data.Attributes.SecondAddress
	}
	if input.City != nil {
		attr.City = *input.City
	} else {
		attr.City = f.Data.Attributes.City
	}
	if input.State != nil {
		attr.State = *input.State
	} else {
		attr.State = f.Data.Attributes.State
	}
	if input.Zipcode != nil {
		attr.Zipcode = *input.Zipcode
	} else {
		attr.Zipcode = f.Data.Attributes.Zipcode
	}
	if input.Country != nil {
		attr.Country = *input.Country
	} else {
		attr.Country = f.Data.Attributes.Country
	}
	if input.Phone != nil {
		attr.Phone = *input.Phone
	} else {
		attr.Phone = f.Data.Attributes.Phone
	}
	if input.IsActive != nil {
		attr.IsActive = *input.IsActive
	} else {
		attr.IsActive = f.Data.Attributes.IsActive
	}
	return attr
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteItem, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	if _, err := r.UserClient.DeleteUser(context.Background(), &jsonapi.DeleteRequest{Id: i}); err != nil {
		r.Logger.Errorf("error deleting user with ID %s: %s", id, err)
		return &models.DeleteItem{
			Success: false,
		}, err
	}
	r.Logger.Infof("successfully deleted user with ID %s", id)
	return &models.DeleteItem{
		Success: true,
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", id, err)
		return nil, err
	}
	g, err := r.UserClient.GetUser(ctx, &jsonapi.GetRequest{Id: i})
	if err != nil {
		r.Logger.Errorf("error in getting user by ID %d: %s", i, err)
		return nil, err
	}
	r.Logger.Infof("successfully found user with ID %s", id)
	return g, nil
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*user.User, error) {
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		r.Logger.Errorf("error in getting user by email %s: %s", email, err)
		return nil, err
	}
	r.Logger.Infof("successfully found user with email %s", email)
	return g, nil
}

func (r *queryResolver) ListUsers(ctx context.Context, cursor *string, limit *int, filter *string) (*models.UserListWithCursor, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *userResolver) FirstName(ctx context.Context, obj *user.User) (string, error) {
	return obj.Data.Attributes.FirstName, nil
}
func (r *userResolver) LastName(ctx context.Context, obj *user.User) (string, error) {
	return obj.Data.Attributes.LastName, nil
}
func (r *userResolver) Email(ctx context.Context, obj *user.User) (string, error) {
	return obj.Data.Attributes.Email, nil
}
func (r *userResolver) Organization(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.Organization, nil
}
func (r *userResolver) GroupName(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.GroupName, nil
}
func (r *userResolver) FirstAddress(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.FirstAddress, nil
}
func (r *userResolver) SecondAddress(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.SecondAddress, nil
}
func (r *userResolver) City(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.City, nil
}
func (r *userResolver) State(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.State, nil
}
func (r *userResolver) Zipcode(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.Zipcode, nil
}
func (r *userResolver) Country(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.Country, nil
}
func (r *userResolver) Phone(ctx context.Context, obj *user.User) (*string, error) {
	return &obj.Data.Attributes.Phone, nil
}
func (r *userResolver) IsActive(ctx context.Context, obj *user.User) (bool, error) {
	return obj.Data.Attributes.IsActive, nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *user.User) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *user.User) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *userResolver) Roles(ctx context.Context, obj *user.User) ([]user.Role, error) {
	panic("not implemented")
}
