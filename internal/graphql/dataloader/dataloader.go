package dataloader

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
