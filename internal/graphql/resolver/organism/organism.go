package organism

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
)

type OrganismResolver struct {
	Logger *logrus.Entry
}

func (r *OrganismResolver) Citations(ctx context.Context, obj *models.Organism) ([]*models.Citation, error) {
	panic("not implemented")
}

func (r *OrganismResolver) Downloads(ctx context.Context, obj *models.Organism) ([]*models.Download, error) {
	panic("not implemented")
}
