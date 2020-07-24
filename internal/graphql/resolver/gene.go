package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (q *QueryResolver) GeneByID(ctx context.Context, id string) (*models.Gene, error) {
	g := &models.Gene{}
	return g, nil
}

func (q *QueryResolver) GeneByName(ctx context.Context, name string) (*models.Gene, error) {
	g := &models.Gene{}
	return g, nil
}
