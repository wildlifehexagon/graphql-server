package models

import (
	"io"
	"strconv"
	"time"

	"github.com/vektah/gqlgen/graphql"
)

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.UnixNano(), 10))
	})
}

// Note: UnmarshalTimestamp is only required if the scalar appears as an input.
// That is not the case here.
