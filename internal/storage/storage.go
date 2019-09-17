package storage

import (
	"context"
)

type Storage interface {
	Get(context.Context, string) (string, bool)
	Add(context.Context, string, string)
}
