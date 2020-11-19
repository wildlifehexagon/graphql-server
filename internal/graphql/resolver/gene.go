package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/cache"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (q *QueryResolver) Gene(ctx context.Context, gene string) (*models.Gene, error) {
	g := &models.Gene{}
	redis := q.GetRedisRepository(cache.RedisKey)
	gn, err := cache.GetGeneFromCache(ctx, redis, gene)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return g, err
	}
	return gn, nil
}
