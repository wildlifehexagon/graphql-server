package publication

import (
	"context"
	"time"

	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/sirupsen/logrus"
)

type PublicationResolver struct {
	Logger *logrus.Entry
}

func (r *PublicationResolver) ID(ctx context.Context, obj *publication.Publication) (string, error) {
	return obj.Data.Id, nil
}
func (r *PublicationResolver) Doi(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Doi, nil
}
func (r *PublicationResolver) Title(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Title, nil
}
func (r *PublicationResolver) Abstract(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Abstract, nil
}
func (r *PublicationResolver) Journal(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Journal, nil
}
func (r *PublicationResolver) PubDate(ctx context.Context, obj *publication.Publication) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.PubDate)
	return &time, nil
}
func (r *PublicationResolver) Volume(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Volume, nil
}
func (r *PublicationResolver) Pages(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Pages, nil
}
func (r *PublicationResolver) Issn(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Issn, nil
}
func (r *PublicationResolver) PubType(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.PubType, nil
}
func (r *PublicationResolver) Source(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Source, nil
}
func (r *PublicationResolver) Issue(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Issue, nil
}
func (r *PublicationResolver) Status(ctx context.Context, obj *publication.Publication) (*string, error) {
	return &obj.Data.Attributes.Status, nil
}
func (r *PublicationResolver) Authors(ctx context.Context, obj *publication.Publication) ([]*publication.Author, error) {
	return obj.Data.Attributes.Authors, nil
}
