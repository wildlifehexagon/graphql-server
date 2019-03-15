package resolver

import (
	"context"
	"fmt"
	"strconv"

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
	n, err := m.GetStockClient(registry.STOCK).CreateStock(ctx, &pb.NewStock{
		Data: &pb.NewStock_Data{
			Type: "strain",
			Attributes: &pb.NewStockAttributes{
				CreatedBy:       attr.CreatedBy,
				UpdatedBy:       attr.UpdatedBy,
				Summary:         attr.Summary,
				EditableSummary: attr.EditableSummary,
				Depositor:       attr.Depositor,
				Genes:           attr.Genes,
				Dbxrefs:         attr.Dbxrefs,
				Publications:    attr.Publications,
				StrainProperties: &pb.StrainProperties{
					SystematicName: prop.SystematicName,
					Label:          prop.Label,
					Species:        prop.Species,
					Plasmid:        prop.Plasmid,
					Parent:         prop.Parent,
					Names:          prop.Names,
				},
			},
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
			if k.Name() == "Genes" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Dbxrefs" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Publications" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Names" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Phenotypes" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Characteristics" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Genotypes" {
				newAttr[k.Name()] = nil
			} else {
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
	n, err := m.GetStockClient(registry.STOCK).CreateStock(ctx, &pb.NewStock{
		Data: &pb.NewStock_Data{
			Type: "plasmid",
			Attributes: &pb.NewStockAttributes{
				CreatedBy:       attr.CreatedBy,
				UpdatedBy:       attr.UpdatedBy,
				Summary:         attr.Summary,
				EditableSummary: attr.EditableSummary,
				Depositor:       attr.Depositor,
				Genes:           attr.Genes,
				Dbxrefs:         attr.Dbxrefs,
				Publications:    attr.Publications,
				PlasmidProperties: &pb.PlasmidProperties{
					ImageMap: prop.ImageMap,
					Sequence: prop.Sequence,
				},
			},
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
			if k.Name() == "Genes" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Dbxrefs" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Publications" {
				newAttr[k.Name()] = nil
			} else if k.Name() == "Keywords" {
				newAttr[k.Name()] = nil
			} else {
				newAttr[k.Name()] = ""
			}
		}
	}
	return newAttr
}

func (m *MutationResolver) UpdateStrain(ctx context.Context, id string, input *models.UpdateStrainInput) (*pb.Stock, error) {
	g, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching strain with ID %s %s", id, err)
	}
	attr := getUpdateStrainAttributes(input, g)
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

func getUpdateStrainAttributes(input *models.UpdateStrainInput, f *pb.Stock) *pb.StockUpdateAttributes {
	// Need to use this to check if input field exists
	// If it does, overwrite the old field content
	// If not, make sure old field remains
	attr := &pb.StockUpdateAttributes{}
	// if input.ID != nil {
	// 	attr.ID = input.ID
	// } else {
	// 	attr.ID = f.Data.Id
	// }
	return attr
}

func (m *MutationResolver) UpdatePlasmid(ctx context.Context, id string, input *models.UpdatePlasmidInput) (*pb.Stock, error) {
	g, err := m.GetStockClient(registry.STOCK).GetStock(ctx, &pb.StockId{Id: id})
	if err != nil {
		return nil, fmt.Errorf("error fetching plasmid with ID %s %s", id, err)
	}
	attr := getUpdatePlasmidAttributes(input, g)
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

func getUpdatePlasmidAttributes(input *models.UpdatePlasmidInput, f *pb.Stock) *pb.StockUpdateAttributes {
	// Need to use this to check if input field exists
	// If it does, overwrite the old field content
	// If not, make sure old field remains
	attr := &pb.StockUpdateAttributes{}
	// if input.ID != nil {
	// 	attr.ID = input.ID
	// } else {
	// 	attr.ID = f.Data.Id
	// }
	return attr
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

func (q *QueryResolver) ListStrains(ctx context.Context, cursor *string, limit *int, filter *string) (*models.StrainListWithCursor, error) {
	var c int64
	if len(*cursor) > 0 {
		c, _ = strconv.ParseInt(*cursor, 10, 64)
		fmt.Println(c)
	}
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &pb.StockParameters{Cursor: c, Limit: int64(*limit), Filter: *filter})
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
	return &models.StrainListWithCursor{
		Strains: strains,
		// NextCursor: "",
		// PreviousCursor: "",
		Limit:      limit,
		TotalCount: len(strains),
	}, nil
}

func (q *QueryResolver) ListPlasmids(ctx context.Context, cursor *string, limit *int, filter *string) (*models.PlasmidListWithCursor, error) {
	var c int64
	if len(*cursor) > 0 {
		c, _ = strconv.ParseInt(*cursor, 10, 64)
		fmt.Println(c)
	}
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &pb.StockParameters{Cursor: c, Limit: int64(*limit), Filter: *filter})
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
	return &models.PlasmidListWithCursor{
		Plasmids: plasmids,
		// NextCursor: "",
		// PreviousCursor: "",
		Limit:      limit,
		TotalCount: len(plasmids),
	}, nil
}
