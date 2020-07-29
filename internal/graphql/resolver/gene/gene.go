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
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type GeneResolver struct {
	Logger *logrus.Entry
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

func fetchUniprotIDs(ctx context.Context, id string) (string, error) {
	url := fmt.Sprintf("https://www.uniprot.org/uniprot?query=%s&columns=id&format=list", id)
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error in uniprot http get request %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "error fetching uniprot with this ID",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
		return "", fmt.Errorf("error fetching uniprot with this id %s", err)
	}

	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body %s", err)
	}
	return strings.TrimSpace(string(r)), nil
}

func fetchGOAs(ctx context.Context, id string) (*quickGo, error) {
	url := fmt.Sprintf("https://www.ebi.ac.uk/QuickGO/services/annotation/search?includeFields=goName&limit=100&geneProductId=%s", id)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error in http get request %s", err)
	}
	if res.StatusCode != 200 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "error fetching go annotations",
			Extensions: map[string]interface{}{
				"code":      "NotFound",
				"timestamp": time.Now(),
			},
		})
		return nil, err
	}
	defer res.Body.Close()
	goa := new(quickGo)
	if err := json.NewDecoder(res.Body).Decode(goa); err != nil {
		return nil, fmt.Errorf("error in decoding json %s", err)
	}
	return goa, nil
}

func (g *GeneResolver) Goas(ctx context.Context, obj *models.Gene) ([]*models.GOAnnotation, error) {
	goas := []*models.GOAnnotation{}
	id, err := fetchUniprotIDs(ctx, obj.ID)
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
			for _, v := range val.WithFrom[0].ConnectedXRefs {
				with = append(with, &models.With{
					ID: v.ID,
					Db: v.DB,
					// Name?
				})
			}
		}
		if val.Extensions != nil {
			for _, v := range val.Extensions[0].ConnectedXRefs {
				ext = append(ext, &models.Extension{
					ID:       v.ID,
					Db:       v.DB,
					Relation: v.Relation,
					// Name?
				})
			}
		}
		goas = append(goas, &models.GOAnnotation{
			ID:           val.ID,
			Type:         val.GoAspect,
			Date:         val.Date,
			EvidenceCode: val.EvidenceCode,
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
