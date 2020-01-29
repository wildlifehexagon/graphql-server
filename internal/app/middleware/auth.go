package auth

import (
	"context"

	"net/http"
)

var contextKey = "refreshToken"

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

func AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arw := authResponseWriter{w, ""}
		c, err := r.Cookie("refresh-token")
		// let unauthenticated users in
		if err != nil || c == nil {
			next.ServeHTTP(w, r)
			return
		}
		arw.cookieStr = c.Value
		// store cookie in context
		ctx := context.WithValue(r.Context(), contextKey, &w)
		// call with new context
		r = r.WithContext(ctx)
		// execute next
		next(&arw, r)
	}
}
