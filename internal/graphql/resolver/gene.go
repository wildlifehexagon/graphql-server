package resolver

import (
	"context"
	"fmt"

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
	var name string
	cache := q.GetRedisRepository(key)
	exists, err := cache.HExists(geneHash, id)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	if !exists {
		nferr := fmt.Errorf("gene id %s does not exist", id)
		errorutils.AddGQLError(ctx, nferr)
		q.Logger.Error(nferr)
		return nil, nferr
	}
	id, err = cache.HGet(geneHash, id)
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
	var id string
	cache := q.GetRedisRepository(key)
	exists, err := cache.HExists(geneHash, name)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	if !exists {
		nferr := fmt.Errorf("gene name %s does not exist", id)
		errorutils.AddGQLError(ctx, nferr)
		q.Logger.Error(nferr)
		return nil, nferr
	}
	id, err = cache.HGet(geneHash, name)
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
