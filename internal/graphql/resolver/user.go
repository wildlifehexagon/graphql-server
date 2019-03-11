package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*pb.User, error) {
	attr := &pb.UserAttributes{}
	a := normalizeCreateUserAttr(input)
	mapstructure.Decode(a, attr)
	n, err := r.UserClient.CreateUser(ctx, &pb.CreateUserRequest{
		Data: &pb.CreateUserRequest_Data{
			Type: "user",
			Attributes: &pb.UserAttributes{
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

func (r *mutationResolver) CreateUserRoleRelationship(ctx context.Context, userId string, roleId string) (*pb.User, error) {
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", userId, err)
		return nil, err
	}
	rid, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", roleId, err)
		return nil, err
	}
	rr, err := r.UserClient.CreateRoleRelationship(ctx, &jsonapi.DataCollection{
		Id: uid,
		Data: []*jsonapi.Data{
			{
				Type: "role",
				Id:   rid,
			},
		}})
	if err != nil {
		r.Logger.Errorf("error in creating role relationship with user %s", err)
		return nil, err
	}
	r.Logger.Infof("successfully created user ID %d relationship role with ID %d %s", uid, rid, rr)
	g, err := r.UserClient.GetUser(ctx, &jsonapi.GetRequest{Id: uid})
	if err != nil {
		r.Logger.Errorf("error in getting user by ID %d: %s", uid, err)
		return nil, err
	}
	return g, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*pb.User, error) {
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
	n, err := r.UserClient.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		Id: i,
		Data: &pb.UpdateUserRequest_Data{
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

func getUpdateUserAttributes(input *models.UpdateUserInput, f *pb.User) *pb.UserAttributes {
	attr := &pb.UserAttributes{}
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

func (r *queryResolver) User(ctx context.Context, id string) (*pb.User, error) {
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

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*pb.User, error) {
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		r.Logger.Errorf("error in getting user by email %s: %s", email, err)
		return nil, err
	}
	r.Logger.Infof("successfully found user with email %s", email)
	return g, nil
}

func (r *queryResolver) ListUsers(ctx context.Context, pagenum string, pagesize string, filter string) (*models.UserList, error) {
	users := []pb.User{}
	pn, err := strconv.ParseInt(pagenum, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", pagenum, err)
		return nil, err
	}
	ps, err := strconv.ParseInt(pagesize, 10, 64)
	if err != nil {
		r.Logger.Errorf("error in parsing string %s to int %s", pagesize, err)
		return nil, err
	}
	g, err := r.UserClient.ListUsers(ctx, &jsonapi.ListRequest{
		Pagenum:  pn,
		Pagesize: ps,
		Filter:   filter,
	})
	if err != nil {
		r.Logger.Errorf("error in getting list of users %s", err)
		return nil, err
	}
	for _, n := range g.Data {
		item := pb.User{
			Data: &pb.UserData{
				Type: "user",
				Id:   n.Id,
				Attributes: &pb.UserAttributes{
					FirstName:     n.Attributes.FirstName,
					LastName:      n.Attributes.LastName,
					Email:         n.Attributes.LastName,
					Organization:  n.Attributes.Organization,
					GroupName:     n.Attributes.GroupName,
					FirstAddress:  n.Attributes.FirstAddress,
					SecondAddress: n.Attributes.SecondAddress,
					City:          n.Attributes.City,
					State:         n.Attributes.State,
					Zipcode:       n.Attributes.Zipcode,
					Country:       n.Attributes.Country,
					Phone:         n.Attributes.Phone,
					IsActive:      n.Attributes.IsActive,
					CreatedAt:     n.Attributes.CreatedAt,
					UpdatedAt:     n.Attributes.UpdatedAt,
				},
			},
		}
		users = append(users, item)
	}
	r.Logger.Infof("successfully retrieved list of %d users", len(users))
	return &models.UserList{
		TotalCount: len(users),
		Users:      users,
	}, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *pb.User) (string, error) {
	return strconv.FormatInt(obj.Data.Id, 10), nil
}
func (r *userResolver) FirstName(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.FirstName, nil
}
func (r *userResolver) LastName(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.LastName, nil
}
func (r *userResolver) Email(ctx context.Context, obj *pb.User) (string, error) {
	return obj.Data.Attributes.Email, nil
}
func (r *userResolver) Organization(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Organization, nil
}
func (r *userResolver) GroupName(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.GroupName, nil
}
func (r *userResolver) FirstAddress(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.FirstAddress, nil
}
func (r *userResolver) SecondAddress(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.SecondAddress, nil
}
func (r *userResolver) City(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.City, nil
}
func (r *userResolver) State(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.State, nil
}
func (r *userResolver) Zipcode(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Zipcode, nil
}
func (r *userResolver) Country(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Country, nil
}
func (r *userResolver) Phone(ctx context.Context, obj *pb.User) (*string, error) {
	return &obj.Data.Attributes.Phone, nil
}
func (r *userResolver) IsActive(ctx context.Context, obj *pb.User) (bool, error) {
	return obj.Data.Attributes.IsActive, nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *pb.User) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt), nil
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *pb.User) (time.Time, error) {
	return aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt), nil
}
func (r *userResolver) Roles(ctx context.Context, obj *pb.User) ([]pb.Role, error) {
	roles := []pb.Role{}
	rr, err := r.UserClient.GetRelatedRoles(ctx, &jsonapi.RelationshipRequest{Id: obj.Data.Id})
	if err != nil {
		r.Logger.Errorf("error getting list of related roles for user ID %d: %s", obj.Data.Id, err)
		return roles, err
	}
	for _, n := range rr.Data {
		item := pb.Role{
			Data: &pb.RoleData{
				Type: "role",
				Id:   n.Id,
				Attributes: &pb.RoleAttributes{
					Role:        n.Attributes.Role,
					Description: n.Attributes.Description,
					CreatedAt:   n.Attributes.CreatedAt,
					UpdatedAt:   n.Attributes.UpdatedAt,
				},
			},
		}
		roles = append(roles, item)
	}
	r.Logger.Infof("successfully retrieved list of %d roles for user ID %d", len(roles), obj.Data.Id)
	return roles, nil
}
