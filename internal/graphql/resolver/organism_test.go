package resolver

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
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/stretchr/testify/assert"
)

func organismTestData() ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return []byte(""), fmt.Errorf("unable to get current dir %s", err)
	}
	path := filepath.Join(
		filepath.Dir(dir), "../../testdata", "organisms.json",
	)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, errors.New("unable to read test file")
	}
	return b, nil
}

func organismHandler(w http.ResponseWriter, r *http.Request) {
	b, err := organismTestData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestOrganism(t *testing.T) {
	t.Parallel()
	u := httptest.NewServer(http.HandlerFunc(organismHandler))
	defer u.Close()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{
			ConnMap: hashmap.New(),
		},
		Logger: mocks.TestLogger(),
	}
	id := "44689"
	q.AddAPIEndpoint("organism", u.URL)
	o, err := q.Organism(context.Background(), id)
	assert.NoError(err, "expect no error from getting organism information")
	assert.Equal(id, o.TaxonID, "should match taxon ID")
	assert.Equal("Dictyostelium discoideum", o.ScientificName, "should match scientific name")
	assert.Equal("Basu et. al, 2015", o.Citations[0].Authors, "should match authors")
	assert.Equal("'dictyBase 2015: Expanding data and annotations in a new software environment.'", o.Citations[0].Title, "should match title")
	assert.Equal("Genesis 53(8), 523–534.", o.Citations[0].Journal, "should match journal")
	assert.Equal("26088819", o.Citations[0].PubmedID, "should match pubmed ID")
}

func TestListOrganisms(t *testing.T) {
	t.Parallel()
	u := httptest.NewServer(http.HandlerFunc(organismHandler))
	defer u.Close()
	assert := assert.New(t)
	q := &QueryResolver{
		Registry: &mocks.MockRegistry{
			ConnMap: hashmap.New(),
		},
		Logger: mocks.TestLogger(),
	}
	q.AddAPIEndpoint("organism", u.URL)
	o, err := q.ListOrganisms(context.Background())
	assert.NoError(err, "expect no error from getting organism information")
	assert.Len(o, 4, "should match number of elements")
}
