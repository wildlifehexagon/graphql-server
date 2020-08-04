package resolver

import (
	"context"
	"fmt"
	"strings"
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

func (q *QueryResolver) Gene(ctx context.Context, gene string) (*models.Gene, error) {
	g := &models.Gene{}
	var id, name string
	cache := q.GetRedisRepository(key)
	exists, err := cache.HExists(geneHash, gene)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	if !exists {
		nferr := fmt.Errorf("gene %s does not exist", gene)
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "gene does not exist",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
		q.Logger.Error(nferr)
		return nil, nferr
	}
	val, err := cache.HGet(geneHash, gene)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	if strings.HasPrefix(gene, "DDB_G") {
		id = gene
		name = val
	} else {
		name = gene
		id = val
	}
	q.Logger.Debugf("retrieved %s for %s from cache", gene, val)
	g.ID = id
	g.Name = name
	return g, nil
}
