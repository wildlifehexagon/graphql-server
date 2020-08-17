package models

type Organism struct {
	TaxonID        string `json:"taxon_id"`
	ScientificName string `json:"scientific_name"`
}
