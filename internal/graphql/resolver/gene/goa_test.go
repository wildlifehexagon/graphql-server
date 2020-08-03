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

	"github.com/dictyBase/graphql-server/internal/repository/redis"
	"github.com/stretchr/testify/assert"
)

var redisAddr = fmt.Sprintf("%s:%s", os.Getenv("REDIS_MASTER_SERVICE_HOST"), os.Getenv("REDIS_MASTER_SERVICE_PORT"))

const (
	mockGeneID      = "DDB_G123456"
	mockGoID        = "GO:123456"
	mockUniprotID   = "U123456"
	mockGeneHash    = "GENE2NAME/mockids"
	mockGoHash      = "GO2NAME/mockids"
	mockUniprotHash = "UNIPROT2NAME/mock"
	mockValue       = "buzz"
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

func TestGetValFromHash(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := redis.NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.HSet(mockGeneHash, mockGeneID, mockValue)
	assert.NoError(err, "error in setting key")
	v := getValFromHash(mockGeneHash, mockGeneID, repo)
	assert.Equal(v, mockValue, "should match value from hash")
	nv := getValFromHash(mockGeneHash, "wrongID", repo)
	assert.Equal(nv, "", "should have empty string if value is missing")
}

func TestGetNameFromDB(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := redis.NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	// set up all of our hashes
	err = repo.HSet(mockGeneHash, mockGeneID, mockValue)
	assert.NoError(err, "error in setting key")
	err = repo.HSet(mockGoHash, mockGoID, mockValue)
	assert.NoError(err, "error in setting key")
	err = repo.HSet(mockUniprotHash, mockUniprotID, mockValue)
	assert.NoError(err, "error in setting key")
	// verify names returned
	gene := getNameFromDB("dictyBase", mockGeneID, repo)
	assert.Equal(gene, mockValue, "should match expected value")
	goa := getNameFromDB("GO", mockGoID, repo)
	assert.Equal(goa, mockValue, "should match expected value")
	uniprot := getNameFromDB("UniProtKB", mockUniprotID, repo)
	assert.Equal(uniprot, mockValue, "should match expected value")
	none := getNameFromDB("noDB", "misc", repo)
	assert.Equal(none, "", "should return empty string if wrong DB")
}
