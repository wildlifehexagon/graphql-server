package resolver

import (
	"context"

	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (r *MutationResolver) Login(ctx context.Context, input *models.LoginInput) (*auth.Auth, error) {
	panic("not implemented")
}
func (r *MutationResolver) Logout(ctx context.Context) (*models.Logout, error) {
	panic("not implemented")
}
func (r *QueryResolver) GetRefreshToken(ctx context.Context, token string) (*models.Token, error) {
	panic("not implemented")
}
