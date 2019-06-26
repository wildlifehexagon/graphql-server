package resolver

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
)

func (m *MutationResolver) CreateContent(ctx context.Context, input *models.CreateContentInput) (*pb.Content, error) {
	cid, err := strconv.ParseInt(input.CreatedBy, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", input.CreatedBy, err)
	}
	n, err := m.GetContentClient(registry.CONTENT).StoreContent(ctx, &pb.StoreContentRequest{
		Data: &pb.StoreContentRequest_Data{
			Type: "contents",
			Attributes: &pb.NewContentAttributes{
				Name:      input.Name,
				CreatedBy: cid,
				Content:   input.Content,
				Namespace: input.Namespace,
			},
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully created new content with ID %d", n.Data.Id)
	return n, nil
}
func (m *MutationResolver) UpdateContent(ctx context.Context, input *models.UpdateContentInput) (*pb.Content, error) {
	cid, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", input.ID, err)
	}
	uid, err := strconv.ParseInt(input.UpdatedBy, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", input.UpdatedBy, err)
	}
	n, err := m.GetContentClient(registry.CONTENT).UpdateContent(ctx, &pb.UpdateContentRequest{
		Id: cid,
		Data: &pb.UpdateContentRequest_Data{
			Type: "contents",
			Id:   cid,
			Attributes: &pb.ExistingContentAttributes{
				UpdatedBy: uid,
				Content:   input.Content,
			},
		},
	})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	o, err := m.GetContentClient(registry.CONTENT).GetContent(ctx, &pb.ContentIdRequest{Id: n.Data.Id})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		m.Logger.Error(err)
		return nil, err
	}
	m.Logger.Debugf("successfully updated content with ID %d", o.Data.Id)
	return o, nil
}
func (m *MutationResolver) DeleteContent(ctx context.Context, id string) (*models.DeleteContent, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	if _, err := m.GetContentClient(registry.CONTENT).DeleteContent(ctx, &pb.ContentIdRequest{Id: cid}); err != nil {
		return &models.DeleteContent{
			Success: false,
		}, err
	}
	m.Logger.Debugf("successfully deleted content with ID %s", id)
	return &models.DeleteContent{
		Success: true,
	}, nil
}

func (q *QueryResolver) Content(ctx context.Context, id string) (*pb.Content, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error in parsing string %s to int %s", id, err)
	}
	content, err := q.GetContentClient(registry.CONTENT).GetContent(ctx, &pb.ContentIdRequest{Id: cid})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found content with ID %s", id)
	return content, nil
}
func (q *QueryResolver) ContentBySlug(ctx context.Context, slug string) (*pb.Content, error) {
	content, err := q.GetContentClient(registry.CONTENT).GetContentBySlug(ctx, &pb.ContentRequest{Slug: slug})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		q.Logger.Error(err)
		return nil, err
	}
	q.Logger.Debugf("successfully found content with slug %s", slug)
	return content, nil
}
