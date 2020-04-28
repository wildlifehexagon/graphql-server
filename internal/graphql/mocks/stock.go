package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

var MockPlasmidAttributes = &stock.PlasmidAttributes{
	CreatedAt:       ptypes.TimestampNow(),
	UpdatedAt:       ptypes.TimestampNow(),
	CreatedBy:       "art@vandelay.com",
	UpdatedBy:       "art@vandelay.com",
	Summary:         "test summary",
	EditableSummary: "editable test summary",
	Depositor:       "Kenny Bania",
	Genes:           []string{"sadA"},
	Dbxrefs:         []string{"test1"},
	Publications:    []string{"99999"},
	ImageMap:        "https://eric.dictybase.dev/test.jpg",
	Sequence:        "ABCDEF",
	Name:            "pTest",
}

var MockStrainAttributes = &stock.StrainAttributes{
	CreatedAt:       ptypes.TimestampNow(),
	UpdatedAt:       ptypes.TimestampNow(),
	CreatedBy:       "art@vandelay.com",
	UpdatedBy:       "art@vandelay.com",
	Summary:         "test summary",
	EditableSummary: "editable test summary",
	Depositor:       "Kenny Bania",
	Genes:           []string{"sadA"},
	Dbxrefs:         []string{"test1"},
	Publications:    []string{"99999"},
	Label:           "test99",
	Species:         "human",
	Plasmid:         "pTest",
	Names:           []string{"fusilli"},
}

var mockStrainList = &stock.StrainCollection_Data{
	Type:       "strain",
	Id:         "DBS123456",
	Attributes: MockStrainAttributes,
}

var mockPlasmidList = &stock.PlasmidCollection_Data{
	Type:       "plasmid",
	Id:         "DBP123456",
	Attributes: MockPlasmidAttributes,
}

func mockStrainCollection() *stock.StrainCollection {
	var strains []*stock.StrainCollection_Data
	strains = append(strains, mockStrainList, mockStrainList, mockStrainList)
	return &stock.StrainCollection{
		Data: strains,
		Meta: &stock.Meta{
			NextCursor: 10000,
			Limit:      10,
			Total:      int64(len(strains)),
		},
	}
}

func mockPlasmidCollection() *stock.PlasmidCollection {
	var plasmids []*stock.PlasmidCollection_Data
	plasmids = append(plasmids, mockPlasmidList, mockPlasmidList, mockPlasmidList)
	return &stock.PlasmidCollection{
		Data: plasmids,
		Meta: &stock.Meta{
			NextCursor: 10000,
			Limit:      10,
			Total:      int64(len(plasmids)),
		},
	}
}

func mockPlasmid() *stock.Plasmid {
	return &stock.Plasmid{
		Data: &stock.Plasmid_Data{
			Type:       "plasmid",
			Id:         "DBP123456",
			Attributes: MockPlasmidAttributes,
		},
	}
}

func mockStrain() *stock.Strain {
	return &stock.Strain{
		Data: &stock.Strain_Data{
			Type:       "strain",
			Id:         "DBS123456",
			Attributes: MockStrainAttributes,
		},
	}
}

func mockedStockClient() *clients.StockServiceClient {
	mockedStockClient := new(clients.StockServiceClient)
	mockedStockClient.On(
		"GetPlasmid",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.StockId"),
	).Return(mockPlasmid(), nil).On(
		"GetStrain",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.StockId"),
	).Return(mockStrain(), nil).On(
		"ListStrains",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.StockParameters"),
	).Return(mockStrainCollection(), nil).On(
		"ListPlasmids",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.StockParameters"),
	).Return(mockPlasmidCollection(), nil).On(
		"CreateStrain",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.NewStrain"),
	).Return(mockStrain(), nil).On(
		"CreatePlasmid",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.NewPlasmid"),
	).Return(mockPlasmid(), nil)
	return mockedStockClient
}
