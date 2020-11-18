package resolver

import (
	"context"
	"fmt"

	anno "github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/stock"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func (m *MutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*models.Strain, error) {
	attr := &pb.NewStrainAttributes{}
	norm := normalizeCreateStrainAttr(input)
	err := mapstructure.Decode(norm, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
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
	strainID := n.Data.Id
	// Note: InStock, Phenotypes, GeneticModification, MutagenesisMethod, Characteristics, SystematicName and Genotypes will need to be implemented later.
	m.Logger.Debugf("successfully created new strain with ID %s", strainID)
	return stock.ConvertToStrainModel(strainID, n.Data.Attributes), nil
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
	err := mapstructure.Decode(norm, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
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
	plasmidID := n.Data.Id
	// Note: InStock, Keywords and GenbankAccession will need to be implemented later.
	m.Logger.Debugf("successfully created new plasmid with ID %s", plasmidID)
	return stock.ConvertToPlasmidModel(plasmidID, n.Data.Attributes), nil
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
	err = mapstructure.Decode(norm, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
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
	strainID := n.Data.Id
	m.Logger.Debugf("successfully updated strain with ID %s", strainID)
	return stock.ConvertToStrainModel(strainID, n.Data.Attributes), nil
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
	err = mapstructure.Decode(norm, attr)
	if err != nil {
		m.Logger.Error(err)
		return nil, err
	}
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
	plasmidID := n.Data.Id
	m.Logger.Debugf("successfully updated plasmid with ID %s", plasmidID)
	return stock.ConvertToPlasmidModel(plasmidID, n.Data.Attributes), nil
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
	n, err := q.GetStockClient(registry.STOCK).GetPlasmid(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	plasmidID := n.Data.Id
	q.Logger.Debugf("successfully found plasmid with ID %s", plasmidID)
	return stock.ConvertToPlasmidModel(plasmidID, n.Data.Attributes), nil
}

func (q *QueryResolver) Strain(ctx context.Context, id string) (*models.Strain, error) {
	n, err := q.GetStockClient(registry.STOCK).GetStrain(ctx, &pb.StockId{Id: id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	strainID := n.Data.Id
	q.Logger.Debugf("successfully found strain with ID %s", strainID)
	return stock.ConvertToStrainModel(strainID, n.Data.Attributes), nil
}

func (q *QueryResolver) ListStrains(ctx context.Context, cursor *int, limit *int, filter *string) (*models.StrainListWithCursor, error) {
	c := getCursor(cursor)
	l := getLimit(limit)
	f := getFilter(filter)
	list, err := q.GetStockClient(registry.STOCK).ListStrains(ctx, &pb.StockParameters{Cursor: c, Limit: l, Filter: f})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	strains := []*models.Strain{}
	for _, n := range list.Data {
		attr := n.Attributes
		item := stock.ConvertToStrainModel(n.Id, attr)
		strains = append(strains, item)
	}
	lm := int(list.Meta.Limit)
	q.Logger.Debugf("successfully retrieved list of %v strains", list.Meta.Total)
	return &models.StrainListWithCursor{
		Strains:        strains,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(c),
		Limit:          &lm,
		TotalCount:     int(list.Meta.Total),
	}, nil
}

func (q *QueryResolver) ListPlasmids(ctx context.Context, cursor *int, limit *int, filter *string) (*models.PlasmidListWithCursor, error) {
	c := getCursor(cursor)
	l := getLimit(limit)
	f := getFilter(filter)
	list, err := q.GetStockClient(registry.STOCK).ListPlasmids(ctx, &pb.StockParameters{Cursor: c, Limit: l, Filter: f})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	plasmids := []*models.Plasmid{}
	for _, n := range list.Data {
		attr := n.Attributes
		item := stock.ConvertToPlasmidModel(n.Id, attr)
		plasmids = append(plasmids, item)
	}
	lm := int(list.Meta.Limit)
	q.Logger.Debugf("successfully retrieved list of %v plasmids", list.Meta.Total)
	return &models.PlasmidListWithCursor{
		Plasmids:       plasmids,
		NextCursor:     int(list.Meta.NextCursor),
		PreviousCursor: int(c),
		Limit:          &lm,
		TotalCount:     int(list.Meta.Total),
	}, nil
}

func (q *QueryResolver) ListStrainsWithAnnotation(ctx context.Context, cursor *int, limit *int, typeArg string, annotation string) (*models.StrainListWithCursor, error) {
	strains := []*models.Strain{}
	c := getCursor(cursor)
	l := getLimit(limit)
	o := getOntology(typeArg)
	a, err := q.GetAnnotationClient(registry.ANNOTATION).ListAnnotations(ctx, &anno.ListParameters{
		Cursor: c,
		Limit:  l,
		Filter: fmt.Sprintf("ontology==%s;tag==%s", o, annotation),
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	for _, v := range a.Data {
		strain, err := q.Strain(ctx, v.Attributes.EntryId)
		if err != nil {
			// errorutils.AddGQLError(ctx, err)
			q.Logger.Error(err)
		}
		strains = append(strains, strain)
	}
	/**
	  Some phenotypes list the same strain ID more than once. Consider a new approach
	  to de-duping this list while also keeping the Meta data from the annotations list.
	*/
	lm := int(a.Meta.Limit)
	return &models.StrainListWithCursor{
		Strains:        strains,
		NextCursor:     int(a.Meta.NextCursor),
		PreviousCursor: int(c),
		Limit:          &lm,
		TotalCount:     len(a.Data),
	}, nil
}

func (q *QueryResolver) ListPlasmidsWithAnnotation(ctx context.Context, cursor *int, limit *int, typeArg string, annotation string) (*models.PlasmidListWithCursor, error) {
	plasmids := []*models.Plasmid{}
	c := getCursor(cursor)
	l := getLimit(limit)
	o := getOntology(typeArg)
	a, err := q.GetAnnotationClient(registry.ANNOTATION).ListAnnotations(ctx, &anno.ListParameters{
		Cursor: c,
		Limit:  l,
		Filter: fmt.Sprintf("ontology==%s;tag==%s", o, annotation),
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	for _, v := range a.Data {
		plasmid, err := q.Plasmid(ctx, v.Attributes.EntryId)
		if err != nil {
			// errorutils.AddGQLError(ctx, err)
			q.Logger.Error(err)
		}
		plasmids = append(plasmids, plasmid)
	}
	lm := int(a.Meta.Limit)
	return &models.PlasmidListWithCursor{
		Plasmids:       plasmids,
		NextCursor:     int(a.Meta.NextCursor),
		PreviousCursor: int(c),
		Limit:          &lm,
		TotalCount:     len(a.Data),
	}, nil
}

func getCursor(c *int) int64 {
	var cursor int64
	if c != nil {
		cursor = int64(*c)
	} else {
		cursor = 0
	}
	return cursor
}

func getLimit(l *int) int64 {
	var limit int64
	if l != nil {
		limit = int64(*l)
	} else {
		limit = 10
	}
	return limit
}

func getFilter(f *string) string {
	var filter string
	if f != nil {
		filter = *f
	} else {
		filter = ""
	}
	return filter
}

func getOntology(onto string) string {
	var o string
	switch onto {
	case "phenotype":
		o = registry.PhenoOntology
	case "characteristic":
		o = registry.StrainCharOnto
	case "strain_inventory":
		o = registry.StrainInvOnto
	case "plasmid_inventory":
		o = registry.PlasmidInvOnto
	default:
		o = "invalid ontology"
	}
	return o
}
