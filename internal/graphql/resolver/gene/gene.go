package gene

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type GeneResolver struct {
	Logger *logrus.Entry
}

func (g *GeneResolver) Goas(ctx context.Context, obj *models.Gene) ([]*models.GOAnnotation, error) {
	goas := []*models.GOAnnotation{}
	return goas, nil
}
