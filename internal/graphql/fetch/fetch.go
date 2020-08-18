package fetch

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetResp(ctx context.Context, url string) (*http.Response, error) {
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
