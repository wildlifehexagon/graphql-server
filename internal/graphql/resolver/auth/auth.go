package auth

import (
	"context"

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
	panic("not implemented")
}
