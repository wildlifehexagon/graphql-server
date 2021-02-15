package fetch

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func doiTestData() ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return []byte(""), fmt.Errorf("unable to get current dir %s", err)
	}
	path := filepath.Join(
		filepath.Dir(dir), "../../testdata", "gwdi_doi.json",
	)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, errors.New("unable to read test file")
	}
	return b, nil
}

func doiHandler(w http.ResponseWriter, r *http.Request) {
	b, err := doiTestData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestFetchDOI(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(doiHandler))
	defer ts.Close()
	assert := assert.New(t)
	g, err := FetchDOI(context.Background(), ts.URL)
	assert.NoError(err, "should not have error when getting http response")
	// expected empty values
	assert.Empty(g.Data.Id, "should have empty string for ID")
	assert.Empty(g.Data.Attributes.Issn, "should have empty string for ISSN")
	assert.Empty(g.Data.Attributes.Journal, "should have empty string for journal")
	assert.Empty(g.Data.Attributes.Volume, "should have empty string for volume")
	assert.Empty(g.Data.Attributes.Pages, "should have empty string for pages")
	assert.Empty(g.Data.Attributes.Issue, "should have empty string for issue")
	// verify authors length
	assert.Len(g.Data.Attributes.Authors, 9, "should match number of authors")
	// expected string values
	assert.Equal(g.Data.Attributes.Status, "preprint", "should match status")
	assert.Equal(g.Data.Attributes.Source, "Crossref", "should match source")
	assert.Equal(g.Data.Attributes.Doi, ts.URL, "should match doi")
	assert.Equal(g.Data.Attributes.PubType, "posted-content", "should match pub type")
	assert.Equal(g.Data.Attributes.Abstract, "<jats:title>Abstract</jats:title><jats:p>Genomes can be sequenced with relative ease, but ascribing gene function remains a major challenge. Genetically tractable model systems are crucial to meet this challenge. One powerful model is the social amoeba<jats:italic>Dictyostelium discoideum</jats:italic>, a eukaryotic microbe widely used to study diverse questions in cell, developmental and evolutionary biology. However, its utility is hampered by the inefficiency with which sequence, transcriptome or proteome variation can be linked to phenotype. To address this, we have developed methods (REMI-seq) to (1) generate a near genome-wide resource of individual mutants (2) allow large-scale parallel phenotyping. We demonstrate that integrating these resources allows novel regulators of cell migration, phagocytosis and macropinocytosis to be rapidly identified. Therefore, these methods and resources provide a step change for high throughput gene discovery in a key model system, and the study of genes affecting traits associated with higher eukaryotes.</jats:p>", "should match doi")
}
