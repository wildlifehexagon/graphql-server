package dataloader

//go:generate go run github.com/vektah/dataloaden StrainLoader string *github.com/dictyBase/graphql-server/internal/graphql/models.Strain

import (
	"context"
	"time"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
)

type contextKey string

const key = contextKey("dataloaders")

type Loaders struct {
	StrainById *StrainLoader
}

func newLoaders(ctx context.Context, nr registry.Registry) *Loaders {
	return &Loaders{
		StrainById: newStrainById(ctx, nr),
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}

func newStrainById(ctx context.Context, nr registry.Registry) *StrainLoader {
	return NewStrainLoader(StrainLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(ids []string) ([]*models.Strain, []error) {
			strains := []*models.Strain{}
			// add logic here...
			return strains, nil
		},
	})
}
