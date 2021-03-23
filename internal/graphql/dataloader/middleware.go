package dataloader

import (
	"context"
	"net/http"

	"github.com/dictyBase/graphql-server/internal/registry"
)

// DataloaderMiddleware stores Loaders as a request-scoped context value.
func DataloaderMiddleware(nr registry.Registry) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(ctx, nr)
			newCtx := context.WithValue(ctx, key, loaders)
			h.ServeHTTP(w, r.WithContext(newCtx))
		}
		return http.HandlerFunc(fn)
	}
}
