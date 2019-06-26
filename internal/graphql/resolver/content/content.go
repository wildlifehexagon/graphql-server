package content

import (
	"context"
	"time"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/sirupsen/logrus"
)

type ContentResolver struct {
	Client pb.ContentServiceClient
	Logger *logrus.Entry
}

func (r *ContentResolver) ID(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) Name(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) Slug(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) CreatedBy(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) UpdatedBy(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) CreatedAt(ctx context.Context, obj *content.Content) (*time.Time, error) {
	panic("not implemented")
}
func (r *ContentResolver) UpdatedAt(ctx context.Context, obj *content.Content) (*time.Time, error) {
	panic("not implemented")
}
func (r *ContentResolver) Content(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
func (r *ContentResolver) Namespace(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}
