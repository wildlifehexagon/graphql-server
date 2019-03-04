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

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*user.User, error) {
	panic("not implemented")
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
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteItem, error) {
	panic("not implemented")
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
		r.Logger.Errorf("error in getting user by email %d: %s", email, err)
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
func (r *userResolver) Roles(ctx context.Context, obj *user.User) ([]models.Role, error) {
	panic("not implemented")
}
