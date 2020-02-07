package middleware

import (
	"context"
	"time"

	"net/http"
)

type contextKey string

// String output the details of context key
func (c contextKey) String() string {
	return "context key " + string(c)
}

var (
	AuthContextKey = contextKey("refreshToken")
	CookieStr      = "refresh-token"
)

type authResponseWriter struct {
	http.ResponseWriter
	RefreshToken string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	if w.RefreshToken == "logout" {
		http.SetCookie(w, &http.Cookie{
			Name:     CookieStr,
			Value:    w.RefreshToken,
			HttpOnly: true,
			Expires:  time.Unix(0, 0), // expired
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     CookieStr,
			Value:    w.RefreshToken,
			HttpOnly: true,
			Expires:  time.Now().AddDate(0, 1, 0), // one month
		})
	}
	return w.ResponseWriter.Write(b)
}

func AuthMiddleWare(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		arw := authResponseWriter{w, ""}
		w = &arw
		// get refresh token from cookie
		c, err := r.Cookie(CookieStr)
		if err != nil || c == nil {
			newCtx := context.WithValue(ctx, AuthContextKey, w)
			h.ServeHTTP(w, r.WithContext(newCtx))
			return
		}
		arw.RefreshToken = c.Value
		newCtx := context.WithValue(ctx, AuthContextKey, w)
		h.ServeHTTP(w, r.WithContext(newCtx))
	}
	return http.HandlerFunc(fn)
}

// WriterFromContext finds the HTTP response writer from the context.
func WriterFromContext(ctx context.Context) *authResponseWriter {
	return ctx.Value(AuthContextKey).(*authResponseWriter)
}
