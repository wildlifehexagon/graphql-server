package resolver

import (
	"context"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func (m *MutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*models.Strain, error) {
	attr := &pb.NewStrainAttributes{}
	norm := normalizeCreateStrainAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.StrainProperties{}
	mapstructure.Decode(norm, prop)
	attr.StrainProperties = prop
	n, err := m.GetStockClient(registry.STOCK).CreateStrain(ctx, &pb.NewStrain{
		Data: &pb.NewStrain_Data{
			Type:       "strain",
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	// Note: InStock, Phenotypes, GeneticModification, MutagenesisMethod, Characteristics and Genotypes will need to be implemented later.
	m.Logger.Debugf("successfully created new strain with ID %s", n.Data.Id)
	return &models.Strain{
		Data: n.Data,
	}, nil
}

func normalizeCreateStrainAttr(attr *models.CreateStrainInput) map[string]interface{} {
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

func (m *MutationResolver) CreatePlasmid(ctx context.Context, input *models.CreatePlasmidInput) (*models.Plasmid, error) {
	attr := &pb.NewPlasmidAttributes{}
	norm := normalizeCreatePlasmidAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.PlasmidProperties{}
	mapstructure.Decode(norm, prop)
	attr.PlasmidProperties = prop
	n, err := m.GetStockClient(registry.STOCK).CreatePlasmid(ctx, &pb.NewPlasmid{
		Data: &pb.NewPlasmid_Data{
			Type:       "plasmid",
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	// Note: InStock, Keywords and GenbankAccession will need to be implemented later.
	m.Logger.Debugf("successfully created new plasmid with ID %s", n.Data.Id)
	return &models.Plasmid{
		Data: n.Data,
	}, nil
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

func (m *MutationResolver) UpdateStrain(ctx context.Context, id string, input *models.UpdateStrainInput) (*models.Strain, error) {
	_, err := m.GetStockClient(registry.STOCK).GetStrain(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	attr := &pb.StrainUpdateAttributes{}
	norm := normalizeUpdateStrainAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.StrainUpdateProperties{}
	mapstructure.Decode(norm, prop)
	attr.StrainProperties = prop
	n, err := m.GetStockClient(registry.STOCK).UpdateStrain(ctx, &pb.StrainUpdate{
		Data: &pb.StrainUpdate_Data{
			Type:       "strain",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	u, err := m.GetStockClient(registry.STOCK).GetStrain(ctx, &pb.StockId{Id: n.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully updated strain with ID %s", u.Data.Id)
	return &models.Strain{
		Data: u.Data,
	}, nil
}

func normalizeUpdateStrainAttr(attr *models.UpdateStrainInput) map[string]interface{} {
	fields := structs.Fields(attr)
	newAttr := make(map[string]interface{})
	for _, k := range fields {
		if !k.IsZero() {
			newAttr[k.Name()] = k.Value()
		}
	}
	return newAttr
}

func (m *MutationResolver) UpdatePlasmid(ctx context.Context, id string, input *models.UpdatePlasmidInput) (*models.Plasmid, error) {
	_, err := m.GetStockClient(registry.STOCK).GetPlasmid(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	attr := &pb.PlasmidUpdateAttributes{}
	norm := normalizeUpdatePlasmidAttr(input)
	mapstructure.Decode(norm, attr)
	prop := &pb.PlasmidProperties{}
	mapstructure.Decode(norm, prop)
	attr.PlasmidProperties = prop
	n, err := m.GetStockClient(registry.STOCK).UpdatePlasmid(ctx, &pb.PlasmidUpdate{
		Data: &pb.PlasmidUpdate_Data{
			Type:       "plasmid",
			Id:         id,
			Attributes: attr,
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	u, err := m.GetStockClient(registry.STOCK).GetPlasmid(ctx, &pb.StockId{Id: n.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully updated plasmid with ID %s", u.Data.Id)
	return &models.Plasmid{
		Data: u.Data,
	}, nil
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
		}, err
	}
	m.Logger.Debugf("successfully deleted stock with ID %s", id)
	return &models.DeleteStock{
		Success: true,
	}, nil
}

func (q *QueryResolver) Plasmid(ctx context.Context, id string) (*models.Plasmid, error) {
	plasmid, err := q.GetStockClient(registry.STOCK).GetPlasmid(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found plasmid with ID %s", id)
	return &models.Plasmid{
		Data: plasmid.Data,
	}, nil
}

func (q *QueryResolver) Strain(ctx context.Context, id string) (*models.Strain, error) {
	strain, err := q.GetStockClient(registry.STOCK).GetStrain(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found strain with ID %s", id)
	return &models.Strain{
		Data: strain.Data,
	}, nil
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
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	strains := []models.Strain{}

	for _, n := range list.Data {
		attr := n.Attributes
		attr.StrainProperties = n.Attributes.StrainProperties

		item := models.Strain{
			Data: &pb.Strain_Data{
				Type:       n.Type,
				Id:         n.Id,
				Attributes: attr,
			},
		}
		strains = append(strains, item)
	}
	l := int(limit)
	q.Logger.Debugf("successfully retrieved list of %v strains", l)
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
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	plasmids := []models.Plasmid{}
	for _, n := range list.Data {
		item := models.Plasmid{
			Data: &pb.Plasmid_Data{
				Type: n.Type,
				Id:   n.Id,
				Attributes: &pb.PlasmidAttributes{
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
	q.Logger.Debugf("successfully retrieved list of %v plasmids", l)
	return &models.PlasmidListWithCursor{
		Plasmids:       plasmids,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(cursor),
		Limit:          &l,
		TotalCount:     len(plasmids),
	}, nil
}
