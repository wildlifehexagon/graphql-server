package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/stretchr/testify/assert"
)

func TestPlasmid(t *testing.T) {
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	plasmidID := "DBP123456"
	p, err := q.Plasmid(context.Background(), plasmidID)
	assert.NoError(err, "expect no error from getting plasmid by ID")
	assert.Exactly(p.Data.Id, plasmidID, "should match plasmid ID")
	assert.Exactly(p.Data.Attributes.CreatedBy, mocks.MockPlasmidAttributes.CreatedBy, "should match created_by")
	assert.Exactly(p.Data.Attributes.UpdatedBy, mocks.MockPlasmidAttributes.UpdatedBy, "should match updated_by")
	assert.Exactly(p.Data.Attributes.Summary, mocks.MockPlasmidAttributes.Summary, "should match summary")
	assert.Exactly(p.Data.Attributes.EditableSummary, mocks.MockPlasmidAttributes.EditableSummary, "should match editable summary")
	assert.Exactly(p.Data.Attributes.Depositor, mocks.MockPlasmidAttributes.Depositor, "should match depositor (he's gold)")
	assert.ElementsMatch(p.Data.Attributes.Genes, mocks.MockPlasmidAttributes.Genes, "should match genes list")
	assert.ElementsMatch(p.Data.Attributes.Dbxrefs, mocks.MockPlasmidAttributes.Dbxrefs, "should match dbxrefs")
	assert.ElementsMatch(p.Data.Attributes.Publications, mocks.MockPlasmidAttributes.Publications, "should match publications")
	assert.Exactly(p.Data.Attributes.ImageMap, mocks.MockPlasmidAttributes.ImageMap, "should match image map")
	assert.Exactly(p.Data.Attributes.Sequence, mocks.MockPlasmidAttributes.Sequence, "should match sequence")
	assert.Exactly(p.Data.Attributes.Name, mocks.MockPlasmidAttributes.Name, "should match name")
}

func TestStrain(t *testing.T) {
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	strainID := "DBS123456"
	p, err := q.Strain(context.Background(), strainID)
	assert.NoError(err, "expect no error from getting strain by ID")
	assert.Exactly(p.Data.Id, strainID, "should match strain ID")
	assert.Exactly(p.Data.Attributes.CreatedBy, mocks.MockStrainAttributes.CreatedBy, "should match created_by")
	assert.Exactly(p.Data.Attributes.UpdatedBy, mocks.MockStrainAttributes.UpdatedBy, "should match updated_by")
	assert.Exactly(p.Data.Attributes.Summary, mocks.MockStrainAttributes.Summary, "should match summary")
	assert.Exactly(p.Data.Attributes.EditableSummary, mocks.MockStrainAttributes.EditableSummary, "should match editable summary")
	assert.Exactly(p.Data.Attributes.Depositor, mocks.MockStrainAttributes.Depositor, "should match depositor (he's gold)")
	assert.ElementsMatch(p.Data.Attributes.Genes, mocks.MockStrainAttributes.Genes, "should match genes list")
	assert.ElementsMatch(p.Data.Attributes.Dbxrefs, mocks.MockStrainAttributes.Dbxrefs, "should match dbxrefs")
	assert.ElementsMatch(p.Data.Attributes.Publications, mocks.MockStrainAttributes.Publications, "should match publications")
	assert.Exactly(p.Data.Attributes.Label, mocks.MockStrainAttributes.Label, "should match label")
	assert.Exactly(p.Data.Attributes.Species, mocks.MockStrainAttributes.Species, "should match species")
	assert.Exactly(p.Data.Attributes.Plasmid, mocks.MockStrainAttributes.Plasmid, "should match plasmid")
	assert.ElementsMatch(p.Data.Attributes.Names, mocks.MockStrainAttributes.Names, "should match names")
}

func TestListPlasmids(t *testing.T) {
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	filter := "type===plasmid"
	p, err := q.ListPlasmids(context.Background(), &models.ListStockInput{
		Cursor: &cursor,
		Limit:  &limit,
		Filter: &filter,
	})
	assert.NoError(err, "expect no error from getting list of strains")
	assert.Exactly(p.Limit, &limit, "should match limit")
	assert.Exactly(p.PreviousCursor, 0, "should match previous cursor")
	assert.Exactly(p.NextCursor, 10000, "should match next cursor")
	assert.Exactly(p.TotalCount, 3, "should match total count (length) of items")
	assert.Len(p.Plasmids, 3, "should have three plasmids")
}

func TestListStrains(t *testing.T) {
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	cursor := 0
	limit := 10
	filter := "type===strain"
	s, err := q.ListStrains(context.Background(), &models.ListStockInput{
		Cursor: &cursor,
		Limit:  &limit,
		Filter: &filter,
	})
	assert.NoError(err, "expect no error from getting list of strains")
	assert.Exactly(s.Limit, &limit, "should match limit")
	assert.Exactly(s.PreviousCursor, 0, "should match previous cursor")
	assert.Exactly(s.NextCursor, 10000, "should match next cursor")
	assert.Exactly(s.TotalCount, 3, "should match total count (length) of items")
	assert.Len(s.Strains, 3, "should have three strains")
}

func CreatePlasmid(t *testing.T) {
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
	assert.Exactly(p.Data.Id, "DBP123456", "should match plasmid ID")
	assert.Exactly(p.Data.Attributes.CreatedBy, input.CreatedBy, "should match created_by")
	assert.Exactly(p.Data.Attributes.UpdatedBy, input.UpdatedBy, "should match updated_by")
	assert.Exactly(p.Data.Attributes.Summary, input.Summary, "should match summary")
	assert.Exactly(p.Data.Attributes.EditableSummary, input.EditableSummary, "should match editable summary")
	assert.Exactly(p.Data.Attributes.Depositor, input.Depositor, "should match depositor (he's gold)")
}

func CreateStrain(t *testing.T) {
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
	assert.Exactly(p.Data.Id, "DBS123456", "should match strain ID")
	assert.Exactly(p.Data.Attributes.CreatedBy, input.CreatedBy, "should match created_by")
	assert.Exactly(p.Data.Attributes.UpdatedBy, input.UpdatedBy, "should match updated_by")
	assert.Exactly(p.Data.Attributes.Summary, input.Summary, "should match summary")
	assert.Exactly(p.Data.Attributes.EditableSummary, input.EditableSummary, "should match editable summary")
	assert.Exactly(p.Data.Attributes.Depositor, input.Depositor, "should match depositor (he's gold)")
	assert.Exactly(p.Data.Attributes.Label, input.Label, "should match label")
	assert.Exactly(p.Data.Attributes.Species, input.Species, "should match species")
}
