package publication

import (
	"context"

	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/sirupsen/logrus"
)

type AuthorResolver struct {
	Logger *logrus.Entry
}

func (r *AuthorResolver) Rank(ctx context.Context, obj *publication.Author) (*string, error) {
	s := ""
	return &s, nil
}
