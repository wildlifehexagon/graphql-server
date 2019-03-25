package models

import "github.com/dictyBase/go-genproto/dictybaseapis/stock"

type Strain struct {
	Data *stock.Strain_Data
}
type Plasmid struct {
	Data *stock.Plasmid_Data
}

func (Strain) IsStock()  {}
func (Plasmid) IsStock() {}
