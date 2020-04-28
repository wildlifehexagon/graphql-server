package errorutils

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

// AddGQLError adds a custom error to the GraphQL response.
func AddGQLError(ctx context.Context, err error) {
	errStatus, _ := status.FromError(err)
	code := fmt.Sprint(errStatus.Code())
	graphql.AddError(ctx, &gqlerror.Error{
		Message: errStatus.Message(),
		Extensions: map[string]interface{}{
			"code":      code,
			"timestamp": time.Now(),
		},
	})
}
