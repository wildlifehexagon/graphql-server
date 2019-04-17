package stock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/errorutils"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

type PlasmidResolver struct {
	Client     pb.StockServiceClient
	UserClient user.UserServiceClient
	Registry   registry.Registry
	Logger     *logrus.Entry
}

func (r *PlasmidResolver) ID(ctx context.Context, obj *models.Plasmid) (string, error) {
	return obj.Data.Id, nil
}
func (r *PlasmidResolver) CreatedAt(ctx context.Context, obj *models.Plasmid) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *PlasmidResolver) UpdatedAt(ctx context.Context, obj *models.Plasmid) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *PlasmidResolver) CreatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *PlasmidResolver) UpdatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		errorutils.AddGQLError(ctx, err)
		r.Logger.Error(err)
		return &user, err
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *PlasmidResolver) Summary(ctx context.Context, obj *models.Plasmid) (*string, error) {
	return &obj.Data.Attributes.Summary, nil
}
func (r *PlasmidResolver) EditableSummary(ctx context.Context, obj *models.Plasmid) (*string, error) {
	return &obj.Data.Attributes.EditableSummary, nil
}
func (r *PlasmidResolver) Depositor(ctx context.Context, obj *models.Plasmid) (string, error) {
	return obj.Data.Attributes.Depositor, nil
}
func (r *PlasmidResolver) Genes(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	g := obj.Data.Attributes.Genes
	pg := []*string{}
	for i := 0; i < len(g); i++ {
		pg = append(pg, &g[i])
	}
	return pg, nil
}
func (r *PlasmidResolver) Dbxrefs(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	d := obj.Data.Attributes.Dbxrefs
	pd := []*string{}
	for i := 0; i < len(d); i++ {
		pd = append(pd, &d[i])
	}
	return pd, nil
}
func (r *PlasmidResolver) Publications(ctx context.Context, obj *models.Plasmid) ([]*publication.Publication, error) {
	pubs := []*publication.Publication{}
	for _, id := range obj.Data.Attributes.Publications {
		endpoint := r.Registry.GetAPIEndpoint(registry.PUBLICATION)
		url := endpoint + "/" + id
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error in http get request %s", err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: "error fetching publication with this ID",
				Extensions: map[string]interface{}{
					"code":      "NotFound",
					"timestamp": time.Now(),
				},
			})
			r.Logger.Error(err)
			return nil, err
		}
		decoder := json.NewDecoder(res.Body)
		var pub PubJsonAPI
		err = decoder.Decode(&pub)
		if err != nil {
			return nil, fmt.Errorf("error decoding json %s", err)
		}
		attr := pub.Data.Attributes
		pd, err := time.Parse("2006-01-02", attr.PublishedDate)
		if err != nil {
			return nil, fmt.Errorf("could not parse published date %s", err)
		}
		p := &publication.Publication{
			Data: &publication.Publication_Data{
				Type: "publication",
				Id:   id,
				Attributes: &publication.PublicationAttributes{
					Doi:      attr.Doi,
					Title:    attr.Title,
					Abstract: attr.Abstract,
					Journal:  attr.Journal,
					PubDate:  aphgrpc.TimestampProto(pd),
					Pages:    attr.Page,
					Issn:     attr.Issn,
					PubType:  attr.PubType,
					Source:   attr.Source,
					Issue:    string(attr.Issue),
					Status:   attr.Status,
					Volume:   "", // field does not exist yet
				},
			},
		}
		var authors []*publication.Author
		for i, a := range attr.Authors {
			authors = append(authors, &publication.Author{
				FirstName: a.FirstName,
				LastName:  a.LastName,
				Rank:      int64(i),
				Initials:  a.Initials,
			})
		}
		p.Data.Attributes.Authors = authors
		r.Logger.Debugf("successfully found publication with ID %s", pub.Data.ID)
		pubs = append(pubs, p)
	}
	return pubs, nil
}
func (r *PlasmidResolver) ImageMap(ctx context.Context, obj *models.Plasmid) (*string, error) {
	return &obj.Data.Attributes.ImageMap, nil
}
func (r *PlasmidResolver) Sequence(ctx context.Context, obj *models.Plasmid) (*string, error) {
	return &obj.Data.Attributes.Sequence, nil
}
func (r *PlasmidResolver) Name(ctx context.Context, obj *models.Plasmid) (string, error) {
	return obj.Data.Attributes.Name, nil
}

/*
* Note: none of the below have been implemented yet.
 */
func (r *PlasmidResolver) InStock(ctx context.Context, obj *models.Plasmid) (bool, error) {
	return true, nil
}
func (r *PlasmidResolver) Keywords(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
func (r *PlasmidResolver) GenbankAccession(ctx context.Context, obj *models.Plasmid) (*string, error) {
	s := ""
	return &s, nil
}
