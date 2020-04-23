package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

func mockPlasmid() *stock.Plasmid {
	return &stock.Plasmid{}
}

func mockStrain() *stock.Strain {
	return &stock.Strain{}
}

func mockedStockClient() *clients.StockServiceClient {
	mockedStockClient := new(clients.StockServiceClient)
	return mockedStockClient
}
