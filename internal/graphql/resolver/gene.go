package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

const (
	genePrefix      = "genelookup:"
	geneCacheKey    = "GENE2NAME/geneids"
	uniprotCacheKey = "UNIPROT2NAME/uniprot"
)

func (q *QueryResolver) GeneByID(ctx context.Context, id string) (*models.Gene, error) {
	g := &models.Gene{
		ID: id,
		// Name: "test1",
	}
	return g, nil
}

func (q *QueryResolver) GeneByName(ctx context.Context, name string) (*models.Gene, error) {
	goas := []*models.GOAnnotation{}
	goas = append(goas, &models.GOAnnotation{
		ID:           "GO:987654",
		Type:         "cellular_component",
		Date:         "20181129",
		EvidenceCode: "IDA",
		GoTerm:       "cell cortex",
		Qualifier:    "colocalizes_with",
		Publication:  "PMID:12499361",
		AssignedBy:   "dictyBase",
	})
	goas = append(goas, &models.GOAnnotation{
		ID:           "GO:123456",
		Type:         "molecular_function",
		Date:         "20200718",
		EvidenceCode: "IEA",
		GoTerm:       "guanyl-nucleotide exchange factor activity",
		Qualifier:    "enables",
		Publication:  "GO_REF:0000043",
		AssignedBy:   "UniProt",
	})
	g := &models.Gene{
		ID:   "DDB_G123456",
		Name: "test1",
		Goas: goas,
	}

	// 1. check if in cache first
	// 2. return cached value OR issue fetch request
	return g, nil
}
