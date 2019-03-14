package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dictyBase/graphql-server/internal/registry"

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
	plasmid, err := q.GetStockClient(registry.STOCK).GetStock(ctx, &stock.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting plasmid with ID %d: %s", id, err)
	}
	q.Logger.Debugf("successfully found plasmid with ID %s", id)
	return plasmid, nil
}
func (q *QueryResolver) Strain(ctx context.Context, id string) (*stock.Stock, error) {
	strain, err := q.GetStockClient(registry.STOCK).GetStock(ctx, &stock.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting strain with ID %d: %s", id, err)
	}
	q.Logger.Debugf("successfully found strain with ID %s", id)
	return strain, nil
}
func (q *QueryResolver) ListStrains(ctx context.Context, cursor *string, limit *int, filter *string) (*models.StrainListWithCursor, error) {
	c, err := strconv.ParseInt(*cursor, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert string to int64 %s", err)
	}
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &stock.StockParameters{Cursor: c, Limit: int64(*limit), Filter: *filter})
	if err != nil {
		return nil, fmt.Errorf("error in getting list of strains %s", err)
	}
	strains := []stock.Stock{}
	for _, n := range list.Data {
		item := stock.Stock{
			Data: &stock.Stock_Data{
				Type: n.Type,
				Id:   n.Id,
				Attributes: &stock.StockAttributes{
					CreatedAt:       n.Attributes.CreatedAt,
					UpdatedAt:       n.Attributes.UpdatedAt,
					CreatedBy:       n.Attributes.CreatedBy,
					UpdatedBy:       n.Attributes.UpdatedBy,
					Summary:         n.Attributes.Summary,
					EditableSummary: n.Attributes.EditableSummary,
					Depositor:       n.Attributes.Depositor,
					Genes:           n.Attributes.Genes,
					Dbxrefs:         n.Attributes.Dbxrefs,
					Publications:    n.Attributes.Publications,
					StrainProperties: &stock.StrainProperties{
						SystematicName: n.Attributes.StrainProperties.SystematicName,
						Label:          n.Attributes.StrainProperties.Label,
						Species:        n.Attributes.StrainProperties.Species,
						Plasmid:        n.Attributes.StrainProperties.Plasmid,
						Parent:         n.Attributes.StrainProperties.Parent,
						Names:          n.Attributes.StrainProperties.Names,
					},
				},
			},
		}
		strains = append(strains, item)
	}
	return &models.StrainListWithCursor{
		Strains: strains,
		// NextCursor: "",
		// PreviousCursor: "",
		Limit:      limit,
		TotalCount: len(strains),
	}, nil
}
func (q *QueryResolver) ListPlasmids(ctx context.Context, cursor *string, limit *int, filter *string) (*models.PlasmidListWithCursor, error) {
	c, err := strconv.ParseInt(*cursor, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert string to int64 %s", err)
	}
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &stock.StockParameters{Cursor: c, Limit: int64(*limit), Filter: *filter})
	if err != nil {
		return nil, fmt.Errorf("error in getting list of strains %s", err)
	}
	plasmids := []stock.Stock{}
	for _, n := range list.Data {
		item := stock.Stock{
			Data: &stock.Stock_Data{
				Type: n.Type,
				Id:   n.Id,
				Attributes: &stock.StockAttributes{
					CreatedAt:       n.Attributes.CreatedAt,
					UpdatedAt:       n.Attributes.UpdatedAt,
					CreatedBy:       n.Attributes.CreatedBy,
					UpdatedBy:       n.Attributes.UpdatedBy,
					Summary:         n.Attributes.Summary,
					EditableSummary: n.Attributes.EditableSummary,
					Depositor:       n.Attributes.Depositor,
					Genes:           n.Attributes.Genes,
					Dbxrefs:         n.Attributes.Dbxrefs,
					Publications:    n.Attributes.Publications,
					PlasmidProperties: &stock.PlasmidProperties{
						ImageMap: n.Attributes.PlasmidProperties.ImageMap,
						Sequence: n.Attributes.PlasmidProperties.Sequence,
					},
				},
			},
		}
		plasmids = append(plasmids, item)
	}
	return &models.PlasmidListWithCursor{
		Plasmids: plasmids,
		// NextCursor: "",
		// PreviousCursor: "",
		Limit:      limit,
		TotalCount: len(plasmids),
	}, nil
}
