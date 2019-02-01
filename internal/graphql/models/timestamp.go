package models

import (
	"io"
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/vektah/gqlgen/graphql"
)

func MarshalTimestamp(t timestamp.Timestamp) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Seconds, 10))
	})
}

// Note: UnmarshalTimestamp is only required if the scalar appears as an input.
// That is not the case here.
