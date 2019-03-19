package resolver

import (
	"context"
	"fmt"

	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (m *MutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*pb.Stock, error) {
	attr := &pb.NewStockAttributes{}
	norm := normalizeCreateStrainAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.StrainProperties{}
	mapstructure.Decode(norm, prop)
	attr.StrainProperties = prop
	n, err := m.GetStockClient(registry.STOCK).CreateStock(ctx, &pb.NewStock{
		Data: &pb.NewStock_Data{
			Type:       "strain",
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating new strain %s", err)
	}
	// Note: InStock, Phenotypes, GeneticModification, MutagenesisMethod, Characteristics and Genotypes will need to be implemented later.
	m.Logger.Debugf("successfully created new strain with ID %s", n.Data.Id)
	return n, nil
}

func normalizeCreateStrainAttr(attr *models.CreateStrainInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			if k.Name() == "Descriptor" {
				newAttr["label"] = k.Value()
			} else {
				newAttr[k.Name()] = k.Value()
			}
		} else {
			switch k.Name() {
			case "Genes":
				newAttr[k.Name()] = nil
			case "Dbxrefs":
				newAttr[k.Name()] = nil
			case "Publications":
				newAttr[k.Name()] = nil
			case "Names":
				newAttr[k.Name()] = nil
			case "Phenotypes":
				newAttr[k.Name()] = nil
			case "Characteristics":
				newAttr[k.Name()] = nil
			case "Genotypes":
				newAttr[k.Name()] = nil
			default:
				newAttr[k.Name()] = ""
			}
		}
	}
	return newAttr
}

func (m *MutationResolver) CreatePlasmid(ctx context.Context, input *models.CreatePlasmidInput) (*pb.Stock, error) {
	attr := &pb.NewStockAttributes{}
	norm := normalizeCreatePlasmidAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.PlasmidProperties{}
	mapstructure.Decode(norm, prop)
	attr.PlasmidProperties = prop
	n, err := m.GetStockClient(registry.STOCK).CreateStock(ctx, &pb.NewStock{
		Data: &pb.NewStock_Data{
			Type:       "plasmid",
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating new plasmid %s", err)
	}
	// Note: InStock, Keywords and GenbankAccession will need to be implemented later.
	m.Logger.Debugf("successfully created new plasmid with ID %s", n.Data.Id)
	return n, nil
}

func normalizeCreatePlasmidAttr(attr *models.CreatePlasmidInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			newAttr[k.Name()] = k.Value()
		} else {
			switch k.Name() {
			case "Genes":
				newAttr[k.Name()] = nil
			case "Dbxrefs":
				newAttr[k.Name()] = nil
			case "Publications":
				newAttr[k.Name()] = nil
			case "Keywords":
				newAttr[k.Name()] = nil
			default:
				newAttr[k.Name()] = ""
			}
		}
	}
	return newAttr
}

func (m *MutationResolver) UpdateStrain(ctx context.Context, id string, input *models.UpdateStrainInput) (*pb.Stock, error) {
	_, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching strain with ID %s %s", id, err)
	}
	attr := &pb.StockUpdateAttributes{}
	norm := normalizeUpdateStrainAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.StrainUpdateProperties{}
	mapstructure.Decode(norm, prop)
	attr.StrainProperties = prop
	n, err := m.GetStockClient(registry.STOCK).UpdateStock(ctx, &pb.StockUpdate{
		Data: &pb.StockUpdate_Data{
			Type:       "strain",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error updating strain %s: %s", n.Data.Id, err)
	}
	u, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching strain with ID %s %s", id, err)
	}
	m.Logger.Debugf("successfully updated strain with ID %s", n.Data.Id)
	return u, nil
}

func normalizeUpdateStrainAttr(attr *models.UpdateStrainInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			if k.Name() == "Descriptor" {
				newAttr["label"] = k.Value()
			} else {
				newAttr[k.Name()] = k.Value()
			}
		}
	}
	return newAttr
}

func (m *MutationResolver) UpdatePlasmid(ctx context.Context, id string, input *models.UpdatePlasmidInput) (*pb.Stock, error) {
	_, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching plasmid with ID %s %s", id, err)
	}
	attr := &pb.StockUpdateAttributes{}
	norm := normalizeUpdatePlasmidAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.PlasmidProperties{}
	mapstructure.Decode(norm, prop)
	attr.PlasmidProperties = prop
	n, err := m.GetStockClient(registry.STOCK).UpdateStock(ctx, &pb.StockUpdate{
		Data: &pb.StockUpdate_Data{
			Type:       "plasmid",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error updating plasmid %s: %s", n.Data.Id, err)
	}
	u, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching plasmid with ID %s %s", id, err)
	}
	m.Logger.Debugf("successfully updated plasmid with ID %s", n.Data.Id)
	return u, nil
}

func normalizeUpdatePlasmidAttr(attr *models.UpdatePlasmidInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			newAttr[k.Name()] = k.Value()
		}
	}
	return newAttr
}

func (m *MutationResolver) DeleteStock(ctx context.Context, id string) (*models.DeleteStock, error) {
	if _, err := m.GetStockClient(registry.STOCK).RemoveStock(ctx, &pb.StockId{Id: id}); err != nil {
		return &models.DeleteStock{
			Success: false,
		}, fmt.Errorf("error deleting stock with ID %s: %s", id, err)
	}
	m.Logger.Debugf("successfully deleted stock with ID %s", id)
	return &models.DeleteStock{
		Success: true,
	}, nil
}

func (q *QueryResolver) Plasmid(ctx context.Context, id string) (*pb.Stock, error) {
	plasmid, err := q.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting plasmid with ID %s: %s", id, err)
	}
	q.Logger.Debugf("successfully found plasmid with ID %s", id)
	return plasmid, nil
}

func (q *QueryResolver) Strain(ctx context.Context, id string) (*pb.Stock, error) {
	strain, err := q.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error in getting strain with ID %s: %s", id, err)
	}
	q.Logger.Debugf("successfully found strain with ID %s", id)
	return strain, nil
}

func (q *QueryResolver) ListStrains(ctx context.Context, input *models.ListStockInput) (*models.StrainListWithCursor, error) {
	var cursor, limit int64
	var filter string
	if input.Cursor != nil {
		cursor = int64(*input.Cursor)
	} else {
		cursor = 0
	}
	if input.Limit != nil {
		limit = int64(*input.Limit)
	} else {
		limit = 10
	}
	if input.Filter != nil {
		filter = *input.Filter
	} else {
		filter = ""
	}
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &pb.StockParameters{Cursor: cursor, Limit: limit, Filter: filter})
	if err != nil {
		return nil, fmt.Errorf("error in getting list of strains %s", err)
	}
	strains := []pb.Stock{}
	for _, n := range list.Data {
		item := pb.Stock{
			Data: &pb.Stock_Data{
				Type: n.Type,
				Id:   n.Id,
				Attributes: &pb.StockAttributes{
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
					StrainProperties: &pb.StrainProperties{
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
	l := int(limit)
	return &models.StrainListWithCursor{
		Strains:        strains,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(cursor),
		Limit:          &l,
		TotalCount:     len(strains),
	}, nil
}

func (q *QueryResolver) ListPlasmids(ctx context.Context, input *models.ListStockInput) (*models.PlasmidListWithCursor, error) {
	var cursor, limit int64
	var filter string
	if input.Cursor != nil {
		cursor = int64(*input.Cursor)
	} else {
		cursor = 0
	}
	if input.Limit != nil {
		limit = int64(*input.Limit)
	} else {
		limit = 10
	}
	if input.Filter != nil {
		filter = *input.Filter
	} else {
		filter = ""
	}
	list, err := q.GetStockClient(registry.STOCK).ListPlasmids(ctx, &pb.StockParameters{Cursor: cursor, Limit: limit, Filter: filter})
	if err != nil {
		return nil, fmt.Errorf("error in getting list of strains %s", err)
	}
	plasmids := []pb.Stock{}
	for _, n := range list.Data {
		item := pb.Stock{
			Data: &pb.Stock_Data{
				Type: n.Type,
				Id:   n.Id,
				Attributes: &pb.StockAttributes{
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
					PlasmidProperties: &pb.PlasmidProperties{
						ImageMap: n.Attributes.PlasmidProperties.ImageMap,
						Sequence: n.Attributes.PlasmidProperties.Sequence,
					},
				},
			},
		}
		plasmids = append(plasmids, item)
	}
	l := int(limit)
	return &models.PlasmidListWithCursor{
		Plasmids:       plasmids,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(cursor),
		Limit:          &l,
		TotalCount:     len(plasmids),
	}, nil
}
