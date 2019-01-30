package models

import (
	"errors"
	"io"
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/vektah/gqlgen/graphql"
)

// need to verify - what format for time output?
func MarshalTimestamp(t timestamp.Timestamp) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Seconds, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (*timestamp.Timestamp, error) {
	if tmpStr, ok := v.(int); ok {
		return &timestamp.Timestamp{
			Seconds: int64(tmpStr),
			Nanos:   0,
		}, nil
	}
	return &timestamp.Timestamp{}, errors.New("time should be a protobuf timestamp")
}
