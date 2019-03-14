package publication

import (
	"context"
	"strconv"

	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/sirupsen/logrus"
)

type AuthorResolver struct {
	Logger *logrus.Entry
}

func (r *AuthorResolver) Rank(ctx context.Context, obj *publication.Author) (*string, error) {
	rank := strconv.Itoa(int(obj.Rank))
	return &rank, nil
}
