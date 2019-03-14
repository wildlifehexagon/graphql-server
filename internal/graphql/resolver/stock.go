package resolver

import (
	"context"

	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (m *MutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*stock.Stock, error) {
	panic("not implemented")
}
func (m *MutationResolver) CreatePlasmid(ctx context.Context, input *models.CreatePlasmidInput) (*stock.Stock, error) {
	panic("not implemented")
}
func (m *MutationResolver) UpdateStrain(ctx context.Context, id string, input *models.UpdateStrainInput) (*stock.Stock, error) {
	panic("not implemented")
}
func (m *MutationResolver) UpdatePlasmid(ctx context.Context, id string, input *models.UpdatePlasmidInput) (*stock.Stock, error) {
	panic("not implemented")
}
func (m *MutationResolver) DeleteStock(ctx context.Context, id string) (*models.DeleteStock, error) {
	panic("not implemented")
}

func (q *QueryResolver) Plasmid(ctx context.Context, id string) (*stock.Stock, error) {
	panic("not implemented")
}
func (q *QueryResolver) Strain(ctx context.Context, id string) (*stock.Stock, error) {
	panic("not implemented")
}
func (q *QueryResolver) ListStrains(ctx context.Context, cursor *string, limit *int, filter *string) (*models.StrainListWithCursor, error) {
	panic("not implemented")
}
func (q *QueryResolver) ListPlasmids(ctx context.Context, cursor *string, limit *int, filter *string) (*models.PlasmidListWithCursor, error) {
	panic("not implemented")
}
