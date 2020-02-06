package middleware

import (
	"context"

	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
type contextKey struct {
	name string
}

var (
	AuthContextKey = &contextKey{"refreshToken"}
	WriterKey      = &contextKey{"authResponseWriter"}
	CookieStr      = "refresh-token"
)

type authResponseWriter struct {
	http.ResponseWriter
	refreshTokenFromCookie string
	refreshTokenToResolver string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     CookieStr,
	// 	Value:    string(b), // this is refresh token from auth service
	// 	HttpOnly: true,
	// })
	return w.ResponseWriter.Write(b)
}

func AuthMiddleWare(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		arw := authResponseWriter{w, "", ""}
		// get refresh token from cookie
		c, err := r.Cookie(CookieStr)
		if err != nil || c == nil {
			newCtx := context.WithValue(ctx, AuthContextKey, &arw.refreshTokenToResolver)
			newCtxWithWriter := context.WithValue(newCtx, WriterKey, &arw)
			h.ServeHTTP(w, r.WithContext(newCtxWithWriter))
			return
		}
		arw.refreshTokenFromCookie = c.Value
		arw.refreshTokenToResolver = c.Value
		newCtx := context.WithValue(ctx, AuthContextKey, &arw.refreshTokenToResolver)
		newCtxWithWriter := context.WithValue(newCtx, WriterKey, &arw)
		h.ServeHTTP(&arw, r.WithContext(newCtxWithWriter))
	}
	return http.HandlerFunc(fn)
}

// TokenFromContext finds the refresh token from the context.
func TokenFromContext(ctx context.Context) *string {
	return ctx.Value(AuthContextKey).(*string)
}

// WriterFromContext finds the HTTP response writer from the context.
func WriterFromContext(ctx context.Context) *authResponseWriter {
	return ctx.Value(WriterKey).(*authResponseWriter)
}
