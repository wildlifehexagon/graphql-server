package gene

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/dictyBase/graphql-server/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	geneHash    = "GENE2NAME/geneids"
	goHash      = "GO2NAME/goids"
	uniprotHash = "UNIPROT2NAME/uniprot"
)

type GeneResolver struct {
	Registry registry.Registry
	Logger   *logrus.Entry
	Redis    repository.Repository
}

type quickGo struct {
	NumberOfHits int      `json:"numberOfHits"`
	Results      []result `json:"results"`
	PageInfo     pageInfo `json:"pageInfo"`
}

type result struct {
	ID            string      `json:"id"`
	GeneProductID string      `json:"geneProductId"`
	Qualifier     string      `json:"qualifier"`
	GoID          string      `json:"goId"`
	GoName        string      `json:"goName"`
	GoEvidence    string      `json:"goEvidence"`
	GoAspect      string      `json:"goAspect"`
	EvidenceCode  string      `json:"evidenceCode"`
	Reference     string      `json:"reference"`
	WithFrom      []with      `json:"withFrom"`
	TaxonID       int         `json:"taxonId"`
	TaxonName     string      `json:"taxonName"`
	AssignedBy    string      `json:"assignedBy"`
	Extensions    []extension `json:"extensions"`
	Symbol        string      `json:"symbol"`
	Date          string      `json:"date"`
	// TargetSets    []string    `json:"targetSets"`
	// Synonyms      []string    `json:"synonyms"`
	// Name          string      `json:"name"`
}

type with struct {
	ConnectedXRefs []withXRef `json:"connectedXrefs"`
}

type extension struct {
	ConnectedXRefs []extensionXRef `json:"connectedXrefs"`
}

type withXRef struct {
	DB string `json:"db"`
	ID string `json:"id"`
}

type extensionXRef struct {
	DB       string `json:"db"`
	ID       string `json:"id"`
	Relation string `json:"relation"`
}

type pageInfo struct {
	ResultsPerPage int `json:"resultsPerPage"`
	Current        int `json:"current"`
	Total          int `json:"total"`
}

func getResp(ctx context.Context, url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		return res, fmt.Errorf("error in http get request with %s", err)
	}
	if res.StatusCode == 404 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "404 error fetching data",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
		return res, fmt.Errorf("404 error fetching data %s", err)
	}
	if res.StatusCode != 200 {
		return res, fmt.Errorf("error fetching data with status code %d", res.StatusCode)
	}
	return res, nil
}

func fetchUniprotID(ctx context.Context, id string) (string, error) {
	url := fmt.Sprintf("https://www.uniprot.org/uniprot?query=%s&columns=id&format=list", id)
	res, err := getResp(ctx, url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body %s", err)
	}
	return strings.TrimSpace(string(r)), nil
}

func fetchGOAs(ctx context.Context, id string) (*quickGo, error) {
	goa := new(quickGo)
	url := fmt.Sprintf("https://www.ebi.ac.uk/QuickGO/services/annotation/search?includeFields=goName&limit=100&geneProductId=%s", id)
	res, err := getResp(ctx, url)
	if err != nil {
		return goa, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(goa); err != nil {
		return nil, fmt.Errorf("error in decoding json %s", err)
	}
	return goa, nil
}

func getValFromHash(hash, key string, cache repository.Repository) string {
	exists, _ := cache.HExists(hash, key)
	if exists {
		name, _ := cache.HGet(hash, key)
		return name
	}
	return ""
}

func getNameFromDB(db, id string, cache repository.Repository) string {
	switch db {
	case "dictyBase":
		return getValFromHash(geneHash, id, cache)
	case "GO":
		key := fmt.Sprintf("%s:%s", db, id)
		return getValFromHash(goHash, key, cache)
	case "UniProtKB":
		return getValFromHash(uniprotHash, id, cache)
	}
	return ""
}

func getWith(with []with, repo repository.Repository) []*models.With {
	wm := []*models.With{}
	for _, v := range with {
		for _, w := range v.ConnectedXRefs {
			wm = append(wm, &models.With{
				ID:   w.ID,
				Db:   w.DB,
				Name: getNameFromDB(w.DB, w.ID, repo),
			})
		}
	}
	return wm
}

func getExtensions(extensions []extension, repo repository.Repository) []*models.Extension {
	ext := []*models.Extension{}
	for _, v := range extensions {
		for _, e := range v.ConnectedXRefs {
			ext = append(ext, &models.Extension{
				ID:       e.ID,
				Db:       e.DB,
				Relation: e.Relation,
				Name:     getNameFromDB(e.DB, e.ID, repo),
			})
		}
	}
	return ext
}

func (g *GeneResolver) Goas(ctx context.Context, obj *models.Gene) ([]*models.GOAnnotation, error) {
	goas := []*models.GOAnnotation{}
	id, err := fetchUniprotID(ctx, obj.ID)
	if err != nil {
		return goas, err
	}
	gn, err := fetchGOAs(ctx, id)
	if err != nil {
		return goas, err
	}
	for _, val := range gn.Results {
		with := []*models.With{}
		ext := []*models.Extension{}
		if val.WithFrom != nil {
			with = getWith(val.WithFrom, g.Redis)
		}
		if val.Extensions != nil {
			ext = getExtensions(val.Extensions, g.Redis)
		}
		goas = append(goas, &models.GOAnnotation{
			ID:           val.ID,
			Type:         val.GoAspect,
			Date:         val.Date,
			EvidenceCode: val.GoEvidence,
			GoTerm:       val.GoName,
			Qualifier:    val.Qualifier,
			Publication:  val.Reference,
			With:         with,
			Extensions:   ext,
			AssignedBy:   val.AssignedBy,
		})
	}
	return goas, nil
}
