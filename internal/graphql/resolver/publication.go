package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

// Publication is the resolver for getting an individual publication by ID.
func (q *QueryResolver) Publication(ctx context.Context, id string) (*models.Publication, error) {
	panic("not implemented")
}
