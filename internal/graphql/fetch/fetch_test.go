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

func TestGetResp(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(goasHandler))
	defer ts.Close()
	assert := assert.New(t)
	g, err := GetResp(context.Background(), ts.URL)
	assert.NoError(err, "should not have error when getting http response")
	assert.Equal(g.StatusCode, 200, "should have ok status code")
}
