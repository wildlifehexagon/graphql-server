package stock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dictyBase/apihelpers/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type StrainResolver struct {
	Client     pb.StockServiceClient
	UserClient user.UserServiceClient
	Registry   registry.Registry
	Logger     *logrus.Entry
}

func (r *StrainResolver) ID(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Id, nil
}
func (r *StrainResolver) CreatedAt(ctx context.Context, obj *models.Strain) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.CreatedAt)
	return &time, nil
}
func (r *StrainResolver) UpdatedAt(ctx context.Context, obj *models.Strain) (*time.Time, error) {
	time := aphgrpc.ProtoTimeStamp(obj.Data.Attributes.UpdatedAt)
	return &time, nil
}
func (r *StrainResolver) CreatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.CreatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return &user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *StrainResolver) UpdatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	user := user.User{}
	email := obj.Data.Attributes.UpdatedBy
	g, err := r.UserClient.GetUserByEmail(ctx, &jsonapi.GetEmailRequest{Email: email})
	if err != nil {
		return &user, fmt.Errorf("error in getting user by email %s: %s", email, err)
	}
	r.Logger.Debugf("successfully found user with email %s", email)
	return g, nil
}
func (r *StrainResolver) Summary(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.Summary, nil
}
func (r *StrainResolver) EditableSummary(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.EditableSummary, nil
}
func (r *StrainResolver) Depositor(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.Depositor, nil
}
func (r *StrainResolver) Genes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	genes := []*string{}
	for _, n := range obj.Data.Attributes.Genes {
		genes = append(genes, &n)
	}
	return genes, nil
}
func (r *StrainResolver) Dbxrefs(ctx context.Context, obj *models.Strain) ([]*string, error) {
	dbxrefs := []*string{}
	for _, n := range obj.Data.Attributes.Dbxrefs {
		dbxrefs = append(dbxrefs, &n)
	}
	return dbxrefs, nil
}
func (r *StrainResolver) Publications(ctx context.Context, obj *models.Strain) ([]*publication.Publication, error) {
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
			return nil, fmt.Errorf("error fetching publication with ID %s", id)
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
func (r *StrainResolver) SystematicName(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.StrainProperties.SystematicName, nil
}
func (r *StrainResolver) Descriptor(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.StrainProperties.Label, nil
}
func (r *StrainResolver) Species(ctx context.Context, obj *models.Strain) (string, error) {
	return obj.Data.Attributes.StrainProperties.Species, nil
}
func (r *StrainResolver) Plasmid(ctx context.Context, obj *models.Strain) (*string, error) {
	return &obj.Data.Attributes.StrainProperties.Plasmid, nil
}
func (r *StrainResolver) Parent(ctx context.Context, obj *models.Strain) (*models.Strain, error) {
	parent := obj.Data.Attributes.StrainProperties.Parent
	strain, err := r.Client.GetStrain(ctx, &pb.StockId{Id: parent})
	if err != nil {
		return nil, fmt.Errorf("error in getting parent strain with ID %s: %s", parent, err)
	}
	r.Logger.Debugf("successfully found parent strain with ID %s", parent)
	return &models.Strain{
		Data: strain.Data,
	}, nil
}
func (r *StrainResolver) Names(ctx context.Context, obj *models.Strain) ([]*string, error) {
	names := []*string{}
	for _, n := range obj.Data.Attributes.StrainProperties.Names {
		names = append(names, &n)
	}
	return names, nil
}

/*
* Note: none of the below have been implemented yet.
 */
func (r *StrainResolver) InStock(ctx context.Context, obj *models.Strain) (bool, error) {
	panic("not implemented")
}
func (r *StrainResolver) Phenotypes(ctx context.Context, obj *models.Strain) ([]*models.Phenotype, error) {
	return []*models.Phenotype{}, nil
}
func (r *StrainResolver) GeneticModification(ctx context.Context, obj *models.Strain) (*string, error) {
	s := ""
	return &s, nil
}
func (r *StrainResolver) MutagenesisMethod(ctx context.Context, obj *models.Strain) (*string, error) {
	s := ""
	return &s, nil
}
func (r *StrainResolver) Characteristics(ctx context.Context, obj *models.Strain) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}
func (r *StrainResolver) Genotypes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	s := ""
	return []*string{&s}, nil
}

type PubJsonAPI struct {
	Data  *PubData `json:"data"`
	Links *Links   `json:"links"`
}

type Links struct {
	Self string `json:"self"`
}

type PubData struct {
	Type       string       `json:"type"`
	ID         string       `json:"id"`
	Attributes *Publication `json:"attributes"`
}

type Publication struct {
	Abstract      string      `json:"abstract"`
	Doi           string      `json:"doi,omitempty"`
	FullTextURL   string      `json:"full_text_url,omitempty"`
	PubmedURL     string      `json:"pubmed_url"`
	Journal       string      `json:"journal"`
	Issn          string      `json:"issn,omitempty"`
	Page          string      `json:"page,omitempty"`
	Pubmed        string      `json:"pubmed"`
	Title         string      `json:"title"`
	Source        string      `json:"source"`
	Status        string      `json:"status"`
	PubType       string      `json:"pub_type"`
	Issue         json.Number `json:"issue,omitempty"`
	PublishedDate string      `json:"publication_date"`
	Authors       []*Author   `json:"authors"`
}

type Author struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Initials  string `json:"initials"`
}
