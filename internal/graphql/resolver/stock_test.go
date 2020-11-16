package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/stretchr/testify/assert"
)

func TestPlasmid(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	plasmidID := "DBP123456"
	p, err := q.Plasmid(context.Background(), plasmidID)
	assert.NoError(err, "expect no error from getting plasmid by ID")
	assert.Equal(p.ID, plasmidID, "should match plasmid ID")
	assert.Equal(p.CreatedBy, mocks.MockPlasmidAttributes.CreatedBy, "should match created_by")
	assert.Equal(p.UpdatedBy, mocks.MockPlasmidAttributes.UpdatedBy, "should match updated_by")
	assert.Equal(p.Summary, &mocks.MockPlasmidAttributes.Summary, "should match summary")
	assert.Equal(p.EditableSummary, &mocks.MockPlasmidAttributes.EditableSummary, "should match editable summary")
	assert.Equal(p.Depositor, &mocks.MockPlasmidAttributes.Depositor, "should match depositor (he's gold)")
	assert.ElementsMatch(p.Genes, sliceConverter(mocks.MockPlasmidAttributes.Genes), "should match genes list")
	assert.ElementsMatch(p.Dbxrefs, sliceConverter(mocks.MockPlasmidAttributes.Dbxrefs), "should match dbxrefs")
	assert.ElementsMatch(p.Publications, sliceConverter(mocks.MockPlasmidAttributes.Publications), "should match publications")
	assert.Equal(p.ImageMap, &mocks.MockPlasmidAttributes.ImageMap, "should match image map")
	assert.Equal(p.Sequence, &mocks.MockPlasmidAttributes.Sequence, "should match sequence")
	assert.Equal(p.Name, mocks.MockPlasmidAttributes.Name, "should match name")
}

func TestStrain(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	strainID := "DBS123456"
	p, err := q.Strain(context.Background(), strainID)
	assert.NoError(err, "expect no error from getting strain by ID")
	assert.Equal(p.ID, strainID, "should match strain ID")
	assert.Equal(p.CreatedBy, mocks.MockStrainAttributes.CreatedBy, "should match created_by")
	assert.Equal(p.UpdatedBy, mocks.MockStrainAttributes.UpdatedBy, "should match updated_by")
	assert.Equal(p.Summary, &mocks.MockStrainAttributes.Summary, "should match summary")
	assert.Equal(p.EditableSummary, &mocks.MockStrainAttributes.EditableSummary, "should match editable summary")
	assert.Equal(p.Depositor, &mocks.MockStrainAttributes.Depositor, "should match depositor (he's gold)")
	assert.ElementsMatch(p.Genes, sliceConverter(mocks.MockStrainAttributes.Genes), "should match genes list")
	assert.ElementsMatch(p.Dbxrefs, sliceConverter(mocks.MockStrainAttributes.Dbxrefs), "should match dbxrefs")
	assert.ElementsMatch(p.Publications, sliceConverter(mocks.MockStrainAttributes.Publications), "should match publications")
	assert.Equal(p.Label, mocks.MockStrainAttributes.Label, "should match label")
	assert.Equal(p.Species, mocks.MockStrainAttributes.Species, "should match species")
	assert.Equal(p.Plasmid, &mocks.MockStrainAttributes.Plasmid, "should match plasmid")
	assert.ElementsMatch(p.Names, sliceConverter(mocks.MockStrainAttributes.Names), "should match names")
}

func TestListPlasmids(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	filter := "type===plasmid"
	p, err := q.ListPlasmids(context.Background(), &cursor, &limit, &filter)
	assert.NoError(err, "expect no error from getting list of strains")
	assert.Equal(p.Limit, &limit, "should match limit")
	assert.Equal(p.PreviousCursor, cursor, "should match previous cursor")
	assert.Equal(p.NextCursor, 0, "should not have value for next cursor since less results than limit")
	assert.Equal(p.TotalCount, 3, "should match total count (length) of items")
	assert.Len(p.Plasmids, 3, "should have three plasmids")
}

func TestListPlasmidsWithAnnotation(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	s, err := q.ListPlasmidsWithAnnotation(context.Background(), &cursor, &limit, "plasmid_inventory", "plasmid inventory")
	assert.NoError(err, "expect no error from getting list of plasmids")
	assert.Equal(s.Limit, &limit, "should match limit")
	assert.Equal(s.PreviousCursor, cursor, "should match previous cursor")
	assert.Equal(s.NextCursor, 0, "should not have value for next cursor since less results than limit")
	assert.Equal(s.TotalCount, 4, "should match total count (length) of items")
	assert.Len(s.Plasmids, 4, "should have four plasmids")
}

func TestListStrains(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	filter := "type===strain"
	s, err := q.ListStrains(context.Background(), &cursor, &limit, &filter)
	assert.NoError(err, "expect no error from getting list of strains")
	assert.Equal(s.Limit, &limit, "should match limit")
	assert.Equal(s.PreviousCursor, cursor, "should match previous cursor")
	assert.Equal(s.NextCursor, 0, "should not have value for next cursor since less results than limit")
	assert.Equal(s.TotalCount, 3, "should match total count (length) of items")
	assert.Len(s.Strains, 3, "should have three strains")
}

func TestListStrainsWithAnnotation(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	s, err := q.ListStrainsWithAnnotation(context.Background(), &cursor, &limit, "phenotype", "delayed culmination")
	assert.NoError(err, "expect no error from getting list of strains")
	assert.Equal(s.Limit, &limit, "should match limit")
	assert.Equal(s.PreviousCursor, cursor, "should match previous cursor")
	assert.Equal(s.NextCursor, 0, "should not have value for next cursor since less results than limit")
	assert.Equal(s.TotalCount, 4, "should match total count (length) of items")
	assert.Len(s.Strains, 4, "should have four strains")
}

func TestCreatePlasmid(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	m := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	summary := "test summary"
	esummary := "editable test summary"
	depositor := "Kenny Bania"
	input := &models.CreatePlasmidInput{
		CreatedBy:       "art@vandelay.com",
		UpdatedBy:       "art@vandelay.com",
		Summary:         &summary,
		EditableSummary: &esummary,
		Depositor:       &depositor,
		InStock:         true,
	}
	p, err := m.CreatePlasmid(context.Background(), input)
	assert.NoError(err, "expect no error from creating new plasmid")
	assert.Equal(p.ID, "DBP123456", "should match plasmid ID")
	assert.Equal(p.CreatedBy, input.CreatedBy, "should match created_by")
	assert.Equal(p.UpdatedBy, input.UpdatedBy, "should match updated_by")
	assert.Equal(&p.Summary, &input.Summary, "should match summary")
	assert.Equal(&p.EditableSummary, &input.EditableSummary, "should match editable summary")
	assert.Equal(&p.Depositor, &input.Depositor, "should match depositor (he's gold)")
}

func TestCreateStrain(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	m := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	summary := "test summary"
	esummary := "editable test summary"
	depositor := "Kenny Bania"
	input := &models.CreateStrainInput{
		CreatedBy:       "art@vandelay.com",
		UpdatedBy:       "art@vandelay.com",
		Summary:         &summary,
		EditableSummary: &esummary,
		Depositor:       &depositor,
		SystematicName:  "test1",
		Label:           "test99",
		Species:         "human",
		InStock:         true,
	}
	p, err := m.CreateStrain(context.Background(), input)
	assert.NoError(err, "expect no error from creating new strain")
	assert.Equal(p.ID, "DBS123456", "should match strain ID")
	assert.Equal(p.CreatedBy, input.CreatedBy, "should match created_by")
	assert.Equal(p.UpdatedBy, input.UpdatedBy, "should match updated_by")
	assert.Equal(p.Summary, input.Summary, "should match summary")
	assert.Equal(p.EditableSummary, input.EditableSummary, "should match editable summary")
	assert.Equal(p.Depositor, input.Depositor, "should match depositor (he's gold)")
	assert.Equal(p.Label, input.Label, "should match label")
	assert.Equal(p.Species, input.Species, "should match species")
}

func TestUpdatePlasmid(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	m := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	summary := "updated summary"
	esummary := "editable updated summary"
	depositor := "Puddy"
	input := &models.UpdatePlasmidInput{
		UpdatedBy:       "h.e.@pennypacker.com",
		Summary:         &summary,
		EditableSummary: &esummary,
		Depositor:       &depositor,
	}
	p, err := m.UpdatePlasmid(context.Background(), "DBP123456", input)
	assert.NoError(err, "expect no error from creating new plasmid")
	assert.Equal(p.UpdatedBy, input.UpdatedBy, "should match updated updated_by")
	assert.Equal(p.Summary, input.Summary, "should match updated summary")
	assert.Equal(p.EditableSummary, input.EditableSummary, "should match updated editable summary")
	assert.Equal(p.Depositor, input.Depositor, "should match updated depositor (he's gold)")
	assert.ElementsMatch(p.Genes, sliceConverter(mocks.MockUpdatePlasmidAttributes.Genes), "should match existing genes list")
	assert.ElementsMatch(p.Dbxrefs, sliceConverter(mocks.MockUpdatePlasmidAttributes.Dbxrefs), "should match existing dbxrefs")
	assert.ElementsMatch(p.Publications, sliceConverter(mocks.MockUpdatePlasmidAttributes.Publications), "should match existing publications")
	assert.Equal(p.ImageMap, &mocks.MockUpdatePlasmidAttributes.ImageMap, "should match existing image map")
	assert.Equal(p.Sequence, &mocks.MockUpdatePlasmidAttributes.Sequence, "should match existing sequence")
	assert.Equal(p.Name, mocks.MockUpdatePlasmidAttributes.Name, "should match name")
}

func TestUpdateStrain(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	m := &MutationResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	summary := "updated summary"
	esummary := "editable updated summary"
	depositor := "Puddy"
	input := &models.UpdateStrainInput{
		UpdatedBy:       "h.e.@pennypacker.com",
		Summary:         &summary,
		EditableSummary: &esummary,
		Depositor:       &depositor,
	}
	p, err := m.UpdateStrain(context.Background(), "DBS123456", input)
	assert.NoError(err, "expect no error from creating new strain")
	assert.Equal(p.UpdatedBy, input.UpdatedBy, "should match updated updated_by")
	assert.Equal(p.Summary, input.Summary, "should match updated summary")
	assert.Equal(p.EditableSummary, input.EditableSummary, "should match updated editable summary")
	assert.Equal(p.Depositor, input.Depositor, "should match updated depositor (he's gold)")
	assert.ElementsMatch(p.Genes, sliceConverter(mocks.MockUpdateStrainAttributes.Genes), "should match existing genes list")
	assert.ElementsMatch(p.Dbxrefs, sliceConverter(mocks.MockUpdateStrainAttributes.Dbxrefs), "should match existing dbxrefs")
	assert.ElementsMatch(p.Publications, sliceConverter(mocks.MockUpdateStrainAttributes.Publications), "should match existing publications")
	assert.Equal(p.Label, mocks.MockUpdateStrainAttributes.Label, "should match existing label")
	assert.Equal(p.Species, mocks.MockUpdateStrainAttributes.Species, "should match existing species")
	assert.Equal(p.Plasmid, &mocks.MockUpdateStrainAttributes.Plasmid, "should match plasmid")
}

func TestGetOntology(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	assert.Equal(getOntology("phenotype"), registry.PhenoOntology, "should return phenotype ontology")
	assert.Equal(getOntology("characteristic"), registry.StrainCharOnto, "should return strain characteristics ontology")
	assert.Equal(getOntology("strain_inventory"), registry.StrainInvOnto, "should return strain inventory ontology")
	assert.Equal(getOntology("plasmid_inventory"), registry.PlasmidInvOnto, "should return plasmid inventory ontology")
	assert.Equal(getOntology("banana"), "invalid ontology", "should return invalid ontology")
}

func sliceConverter(s []string) []*string {
	c := []*string{}
	// need to use for loop here, not range
	// https://github.com/golang/go/issues/22791#issuecomment-345391395
	for i := 0; i < len(s); i++ {
		c = append(c, &s[i])
	}
	return c
}
