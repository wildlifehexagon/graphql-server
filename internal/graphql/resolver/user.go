package resolver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
)

func (m *MutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*pb.User, error) {
	attr := &pb.UserAttributes{}
	a := normalizeCreateUserAttr(input)
	err := mapstructure.Decode(a, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
	n, err := m.GetUserClient(registry.USER).CreateUser(ctx, &pb.CreateUserRequest{
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
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully created new user with ID %d", n.Data.Id)
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

func (m *MutationResolver) CreateUserRoleRelationship(ctx context.Context, userId string, roleId string) (*pb.User, error) {
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", userId, err)
	}
	rid, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", roleId, err)
	}
	rr, err := m.GetUserClient(registry.USER).CreateRoleRelationship(ctx, &jsonapi.DataCollection{
		Id: uid,
		Data: []*jsonapi.Data{
			{
				Type: "role",
				Id:   rid,
			},
		}})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully created user ID %d relationship role with ID %d %s", uid, rid, rr)
	g, err := m.GetUserClient(registry.USER).GetUser(ctx, &jsonapi.GetRequest{Id: uid})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	return g, nil
}

func (m *MutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*pb.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	f, err := m.GetUserClient(registry.USER).GetUser(ctx, &jsonapi.GetRequest{Id: i})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	attr := getUpdateUserAttributes(input, f)
	attr.Email = f.Data.Attributes.Email
	attr.UpdatedAt = aphgrpc.TimestampProto(time.Now())
	n, err := m.GetUserClient(registry.USER).UpdateUser(ctx, &pb.UpdateUserRequest{
		Id: i,
		Data: &pb.UpdateUserRequest_Data{
			Id:         i,
			Type:       "user",
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	o, err := m.GetUserClient(registry.USER).GetUser(ctx, &jsonapi.GetRequest{Id: n.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully updated user with ID %d", n.Data.Id)
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

func (m *MutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteUser, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	if _, err := m.GetUserClient(registry.USER).DeleteUser(ctx, &jsonapi.DeleteRequest{Id: i}); err != nil {
		m.Logger.Error(err)
		return &models.DeleteUser{
			Success: false,
		}, err
	}
	m.Logger.Debugf("successfully deleted user with ID %s", id)
	return &models.DeleteUser{
		Success: true,
	}, nil
}

func (q *QueryResolver) User(ctx context.Context, id string) (*pb.User, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	g, err := q.GetUserClient(registry.USER).GetUser(ctx, &jsonapi.GetRequest{Id: i})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found user with ID %s", id)
	return g, nil
}

func (q *QueryResolver) UserByEmail(ctx context.Context, email string) (*pb.User, error) {
	g, err := q.GetUserClient(registry.USER).GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}

func (q *QueryResolver) ListUsers(ctx context.Context, pagenum string, pagesize string, filter string) (*models.UserList, error) {
	users := []*pb.User{}
	pn, err := strconv.ParseInt(pagenum, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", pagenum, err)
	}
	ps, err := strconv.ParseInt(pagesize, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", pagesize, err)
	}
	g, err := q.GetUserClient(registry.USER).ListUsers(ctx, &jsonapi.ListRequest{
		Pagenum:  pn,
		Pagesize: ps,
		Filter:   filter,
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	for _, n := range g.Data {
		item := &pb.User{
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
	q.Logger.Debugf("successfully retrieved list of %d users", len(users))
	return &models.UserList{
		TotalCount: len(users),
		Users:      users,
	}, nil
}
