package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/repository"
	"github.com/vektah/gqlparser/gqlerror"
)

const (
	RedisKey   = "redis"
	GenePrefix = "genelookup:"
	GeneHash   = "GENE2NAME/geneids"
)

func GetGeneFromCache(ctx context.Context, cache repository.Repository, gene string) (*models.Gene, error) {
	g := &models.Gene{}
	var id, name string
	exists, err := cache.HExists(GeneHash, gene)
	if err != nil {
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
		return nil, nferr
	}
	val, err := cache.HGet(GeneHash, gene)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(gene, "DDB_G") {
		id = gene
		name = val
	} else {
		name = gene
		id = val
	}
	g.ID = id
	g.Name = name
	return g, nil
}
