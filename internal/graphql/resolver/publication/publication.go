package publication

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/sirupsen/logrus"
)

type PublicationResolver struct {
	Logger *logrus.Entry
}

func (r *PublicationResolver) ID(ctx context.Context, obj *publication.Publication) (string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Doi(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Title(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Abstract(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Journal(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) PubDate(ctx context.Context, obj *publication.Publication) (*time.Time, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Volume(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Pages(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Issn(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) PubType(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Source(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Issue(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Status(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}
func (r *PublicationResolver) Authors(ctx context.Context, obj *publication.Publication) ([]*publication.Author, error) {
	panic("not implemented")
}
