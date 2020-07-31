package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
)

const (
	key        = "redis"
	genePrefix = "genelookup:"
	geneHash   = "GENE2NAME/geneids"
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
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "gene name does not exist",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
		q.Logger.Error(nferr)
		return nil, nferr
	}
	name, err = cache.HGet(geneHash, id)
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
		nferr := fmt.Errorf("gene name %s does not exist", name)
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "gene name does not exist",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
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
