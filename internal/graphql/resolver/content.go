package resolver

import (
	"context"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/content"
)

func (m *mutationResolver) CreateContent(ctx context.Context, input *models.CreateContentInput) (*content.Content, error) {
	panic("not implemented")
}
func (m *mutationResolver) UpdateContent(ctx context.Context, input *models.UpdateContentInput) (*content.Content, error) {
	panic("not implemented")
}
func (m *mutationResolver) DeleteContent(ctx context.Context, id string) (*models.DeleteContent, error) {
	panic("not implemented")
}

func (q *queryResolver) Content(ctx context.Context, id string) (*content.Content, error) {
	panic("not implemented")
}
func (q *queryResolver) ContentBySlug(ctx context.Context, slug string) (*content.Content, error) {
	panic("not implemented")
}
