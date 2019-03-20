package models

import "github.com/dictyBase/go-genproto/dictybaseapis/stock"

type Strain struct {
	Data *stock.Strain_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}
type Plasmid struct {
	Data *stock.Plasmid_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (Strain) IsStock()  {}
func (Plasmid) IsStock() {}
