package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/graphql/utils"
)

type organisms struct {
	Data []organismData `json:"data"`
}

type organismData struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes organismAttributes `json:"attributes"`
}

type organismAttributes struct {
	TaxonID        string     `json:"taxon_id"`
	ScientificName string     `json:"scientific_name"`
	Citations      []citation `json:"citations"`
}

type citation struct {
	Authors string `json:"authors"`
	Title   string `json:"title"`
	Journal string `json:"journal"`
	Link    string `json:"link"`
}

type downloads struct {
	Data []downloadsData `json:"data"`
}

type downloadsData struct {
	Type       string              `json:"type"`
	ID         string              `json:"id"`
	Attributes downloadsAttributes `json:"attributes"`
}

type downloadsAttributes struct {
	Title string          `json:"title"`
	Items []downloadsItem `json:"items"`
}

type downloadsItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func fetchOrganisms(ctx context.Context, url string) (*organisms, error) {
	o := new(organisms)
	res, err := utils.GetResp(ctx, url)
	if err != nil {
		return o, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(o); err != nil {
		return o, fmt.Errorf("error in decoding json %s", err)
	}
	return o, nil
}

func (q *QueryResolver) Organism(ctx context.Context, taxonID string) (*models.Organism, error) {
	o := &models.Organism{}
	url := "https://github.com/dictyBase/migration-data/blob/master/downloads/organisms-with-citations.staging.json"
	d, err := fetchOrganisms(ctx, url)
	if err != nil {
		return o, err
	}
	for _, val := range d.Data {
		if val.ID == taxonID {
			o.TaxonID = val.Attributes.TaxonID
			o.ScientificName = val.Attributes.ScientificName
		}
	}
	return o, nil
}
