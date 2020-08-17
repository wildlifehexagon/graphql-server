package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (q *QueryResolver) Organism(ctx context.Context, taxonID string) (*models.Organism, error) {
	panic("not implemented")
}
