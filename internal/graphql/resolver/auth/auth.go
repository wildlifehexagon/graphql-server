package auth

import (
	"context"
	"strconv"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type AuthResolver struct {
	Client         auth.AuthServiceClient
	UserClient     user.UserServiceClient
	IdentityClient identity.IdentityServiceClient
	Logger         *logrus.Entry
}

func (r *AuthResolver) Identity(ctx context.Context, obj *auth.Auth) (*models.Identity, error) {
	attr := obj.Identity.Data.Attributes
	return &models.Identity{
		ID:         strconv.FormatInt(obj.Identity.Data.Id, 10),
		Identifier: attr.Identifier,
		Provider:   attr.Provider,
		UserID:     strconv.FormatInt(attr.UserId, 10),
		CreatedAt:  aphgrpc.ProtoTimeStamp(attr.CreatedAt),
		UpdatedAt:  aphgrpc.ProtoTimeStamp(attr.UpdatedAt),
	}, nil
}
