package auth

import (
	"strconv"
	"context"

	"github.com/dictyBase/apihelpers/aphgrpc"
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
	return &models.Identity{
		ID:         strconv.FormatInt(obj.Identity.Data.Id, 10),
		Identifier: obj.Identity.Data.Attributes.Identifier,
		Provider:   obj.Identity.Data.Attributes.Provider,
		UserID:     strconv.FormatInt(obj.Identity.Data.Attributes.UserId, 10),
		CreatedAt:  aphgrpc.ProtoTimeStamp(obj.Identity.Data.Attributes.CreatedAt),
		UpdatedAt:  aphgrpc.ProtoTimeStamp(obj.Identity.Data.Attributes.UpdatedAt),
	}, nil
}
