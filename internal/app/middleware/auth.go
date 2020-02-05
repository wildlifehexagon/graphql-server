package middleware

import (
	"context"

	"net/http"
)

var AuthContextKey = "refreshToken"

type authResponseWriter struct {
	http.ResponseWriter
	cookieStr string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh-token",
		Value:    w.cookieStr,
		HttpOnly: true,
	})
	return w.ResponseWriter.Write(b)
}

func AuthMiddleWare(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		arw := authResponseWriter{w, ""}
		c, err := r.Cookie("refresh-token")
		// let unauthenticated users in
		if err != nil || c == nil {
			h.ServeHTTP(w, r)
		} else {
			arw.cookieStr = c.Value
			newCtx := context.WithValue(ctx, AuthContextKey, &w)
			h.ServeHTTP(w, r.WithContext(newCtx))
		}
	}
	return http.HandlerFunc(fn)
}
