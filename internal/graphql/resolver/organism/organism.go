package organism

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/graphql/utils"
	"github.com/sirupsen/logrus"
)

type OrganismResolver struct {
	Logger *logrus.Entry
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

func fetchDownloads(ctx context.Context, url string) (*downloads, error) {
	d := new(downloads)
	res, err := utils.GetResp(ctx, url)
	if err != nil {
		return d, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(d); err != nil {
		return d, fmt.Errorf("error in decoding json %s", err)
	}
	return d, nil
}

func (r *OrganismResolver) Downloads(ctx context.Context, obj *models.Organism) ([]*models.Download, error) {
	ds := []*models.Download{}
	items := []*models.DownloadItem{}
	res, err := fetchDownloads(ctx, fmt.Sprintf("https://raw.githubusercontent.com/dictyBase/migration-data/master/downloads/%s.staging.json", obj.TaxonID))
	if err != nil {
		return ds, err
	}
	for _, val := range res.Data {
		for _, item := range val.Attributes.Items {
			di := &models.DownloadItem{
				Title: item.Title,
				URL:   item.URL,
			}
			items = append(items, di)
		}
		d := &models.Download{
			Title: val.Attributes.Title,
			Items: items,
		}
		ds = append(ds, d)
	}
	return ds, nil
}
