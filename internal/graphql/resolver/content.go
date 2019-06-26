package resolver

import (
	"context"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

func (m *MutationResolver) CreateContent(ctx context.Context, input *models.CreateContentInput) (*pb.Content, error) {
	panic("not implemented")
}
func (m *MutationResolver) UpdateContent(ctx context.Context, input *models.UpdateContentInput) (*pb.Content, error) {
	panic("not implemented")
}
func (m *MutationResolver) DeleteContent(ctx context.Context, id string) (*models.DeleteContent, error) {
	panic("not implemented")
}

func (q *QueryResolver) Content(ctx context.Context, id string) (*pb.Content, error) {
	panic("not implemented")
}
func (q *QueryResolver) ContentBySlug(ctx context.Context, slug string) (*pb.Content, error) {
	panic("not implemented")
}
