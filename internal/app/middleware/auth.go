package middleware

import (
	"context"

	"net/http"
)

const (
	AuthContextKey = "refreshToken"
	CookieStr      = "refresh-token"
)

type authResponseWriter struct {
	http.ResponseWriter
	refreshTokenFromCookie string
	refreshTokenToResolver string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	if w.refreshTokenToResolver != w.refreshTokenFromCookie {
		http.SetCookie(w, &http.Cookie{
			Name:     CookieStr,
			Value:    w.refreshTokenFromCookie,
			HttpOnly: true,
		})
	}
	return w.ResponseWriter.Write(b)
}

func AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		arw := authResponseWriter{w, "", ""}
		// get refresh token from cookie
		c, err := r.Cookie(CookieStr)
		if err != nil || c == nil {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		arw.refreshTokenFromCookie = c.Value
		arw.refreshTokenToResolver = c.Value
		newCtx := context.WithValue(ctx, AuthContextKey, &arw.refreshTokenToResolver)
		// executing next
		next(&arw, r.WithContext(newCtx))
	}
}
