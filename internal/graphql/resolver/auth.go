package resolver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/graphql-server/internal/app/middleware"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
)

func (m *MutationResolver) Login(ctx context.Context, input *models.LoginInput) (*pb.Auth, error) {
	a := &pb.Auth{}
	// 1. Call service login method
	l, err := m.GetAuthClient(registry.AUTH).Login(ctx, &pb.NewLogin{
		ClientId:    input.ClientID,
		Scopes:      input.Scopes,
		State:       input.State,
		RedirectUrl: input.RedirectURL,
		Code:        input.Code,
		Provider:    input.Provider,
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return a, err
	}
	// 2, Set refresh token cookie with response
	arw := middleware.WriterFromContext(ctx)
	cookie := &http.Cookie{
		Name:     middleware.CookieStr,
		Value:    l.RefreshToken,
		HttpOnly: true,
		Expires: time.Now().AddDate(0, 1, 0), // one month
	}
	http.SetCookie(arw, cookie)
	// 3. Convert rest of response to Auth model
	a = &pb.Auth{
		Token:        l.Token,
		RefreshToken: cookie.Value,
		User:         l.User,
		Identity:     l.Identity,
	}
	return a, nil
}
func (m *MutationResolver) Logout(ctx context.Context) (*models.Logout, error) {
	// 1. Check for refresh token
	if rt := middleware.TokenFromContext(ctx); *rt == "" {
		err := fmt.Errorf("refresh token does not exist")
		errorutils.AddGQLError(ctx, err)
		return nil, err
	}
	// 2. Create expired cookie
	arw := middleware.WriterFromContext(ctx)
	cookie := http.Cookie{
		Name:     middleware.CookieStr,
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(arw, &cookie)
	// 3. Call Logout service method
	_, err := m.GetAuthClient(registry.AUTH).Logout(ctx, &pb.NewRefreshToken{
		RefreshToken: cookie.Value,
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	return &models.Logout{
		Success: true,
	}, nil
}
func (q *QueryResolver) GetRefreshToken(ctx context.Context, token string) (*models.Token, error) {
	tkn := &models.Token{}
	// 1. Get the refresh token from the cookie
	cookie := middleware.TokenFromContext(ctx)
	// 2. If it doesn't exist, send back empty token
	if *cookie == "" {
		return tkn, nil
	}
	// 3. Pass refresh token and JWT into GetRefreshToken method
	t, err := q.GetAuthClient(registry.AUTH).GetRefreshToken(ctx, &pb.NewToken{
		RefreshToken: *cookie,
		Token:        token,
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	// 4. Set new refresh token cookie from response
	arw := middleware.WriterFromContext(ctx)
	nc := &http.Cookie{
		Name:     middleware.CookieStr,
		Value:    t.RefreshToken,
		HttpOnly: true,
		Expires: time.Now().AddDate(0, 1, 0), // one month
	}
	http.SetCookie(arw, nc)
	// 5. Return JWT
	return &models.Token{
		Token: t.Token,
	}, nil
}
