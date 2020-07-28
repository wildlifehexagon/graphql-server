package storage

import (
	"context"
)

type Storage interface {
	Get(context.Context, string) (string, bool)
	Set(context.Context, string, string)
}
