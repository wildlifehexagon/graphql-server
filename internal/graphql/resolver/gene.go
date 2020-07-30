package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

const (
	key         = "redis"
	genePrefix  = "genelookup:"
	geneHash    = "GENE2NAME/geneids"
	uniprotHash = "UNIPROT2NAME/uniprot"
)

func (q *QueryResolver) GeneByID(ctx context.Context, id string) (*models.Gene, error) {
	g := &models.Gene{}
	cache := q.GetRedisRepository(key)
	name, err := cache.HGet(geneHash, id)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("retrieved %s for %s from cache", name, id)
	g.ID = id
	g.Name = name
	return g, nil
}

func (q *QueryResolver) GeneByName(ctx context.Context, name string) (*models.Gene, error) {
	g := &models.Gene{}
	cache := q.GetRedisRepository(key)
	id, err := cache.HGet(geneHash, name)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("retrieved %s for %s from cache", id, name)
	g.ID = id
	g.Name = name
	return g, nil
}
