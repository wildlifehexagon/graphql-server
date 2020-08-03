package gene

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

const (
	mockUniprotID = "A1XDC0"
)

func goaTestData() ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return []byte(""), fmt.Errorf("unable to get current dir %s", err)
	}
	path := filepath.Join(
		filepath.Dir(dir), "../../../testdata", "goas.json",
	)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, errors.New("unable to read test file")
	}
	return b, nil
}

func goasHandler(w http.ResponseWriter, r *http.Request) {
	b, err := goaTestData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func uniprotHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(mockUniprotID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestGetResp(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(goasHandler))
	defer ts.Close()
	assert := assert.New(t)
	g, err := getResp(context.Background(), ts.URL)
	assert.NoError(err, "should not have error when getting http response")
	assert.Equal(g.StatusCode, 200, "should have ok status code")
}

func TestFetchUniprotID(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(uniprotHandler))
	defer ts.Close()
	assert := assert.New(t)
	u, err := fetchUniprotID(context.Background(), ts.URL)
	assert.NoError(err, "should not have error when getting http response")
	assert.Equal(u, mockUniprotID, "should match uniprot ID")
}

func TestFetchGOAs(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(goasHandler))
	defer ts.Close()
	assert := assert.New(t)
	g, err := fetchGOAs(context.Background(), ts.URL)
	assert.NoError(err, "should not have error when getting http response")
	assert.Equal(g.NumberOfHits, 19, "should match number of hits")
	assert.Equal(len(g.Results), 19, "should match number of results in slice")
}
