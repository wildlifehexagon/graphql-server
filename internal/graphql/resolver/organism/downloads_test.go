package organism

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

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/stretchr/testify/assert"
)

var mockOrganismModel = &models.Organism{
	TaxonID:        "44689",
	ScientificName: "Dictyostelium discoideum",
}

func downloadTestData() ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return []byte(""), fmt.Errorf("unable to get current dir %s", err)
	}
	path := filepath.Join(
		filepath.Dir(dir), "../../../testdata", "44689.json",
	)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, errors.New("unable to read test file")
	}
	return b, nil
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	b, err := downloadTestData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestDownloads(t *testing.T) {
	t.Parallel()
	u := httptest.NewServer(http.HandlerFunc(downloadHandler))
	defer u.Close()
	assert := assert.New(t)
	r := &OrganismResolver{
		Logger:       mocks.TestLogger(),
		DownloadsURL: u.URL,
	}
	o, err := r.Downloads(context.Background(), mockOrganismModel)
	assert.NoError(err, "expect no error from getting download information")
	assert.Len(o, 8, "should have expected number of elements")
	assert.Len(o[0].Items, 5, "should match number of items in first element")
	assert.Equal(o[0].Title, "Gene Information", "should match title")
}
