package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

func mockPlasmid() *stock.Plasmid {
	return &stock.Plasmid{
		Data: &stock.Plasmid_Data{
			Type: "plasmid",
			Id:   "DBP123456",
			Attributes: &stock.PlasmidAttributes{
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
			},
		},
	}
}

// func mockStrain() *stock.Strain {
// 	return &stock.Strain{}
// }

func mockedStockClient() *clients.StockServiceClient {
	mockedStockClient := new(clients.StockServiceClient)
	mockedStockClient.On(
		"GetPlasmid",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("*stock.StockId"),
	).Return(mockPlasmid(), nil)
	return mockedStockClient
}
