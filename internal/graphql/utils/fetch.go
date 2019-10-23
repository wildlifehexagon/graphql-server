package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/apihelpers/aphgrpc"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

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

func FetchPublication(ctx context.Context, reg registry.Registry, id string) (*pb.Publication, error) {
	logger := *logrus.New()
	endpoint := reg.GetAPIEndpoint(registry.PUBLICATION)
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
		logger.Error("error fetching publication with this ID")
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	var pub PubJsonAPI
	err = decoder.Decode(&pub)
	if err != nil {
		return nil, fmt.Errorf("error decoding json %s", err)
	}
	attr := pub.Data.Attributes
	// convert pub_date to expected time.Time format
	pd, err := time.Parse("2006-01-02", attr.PublishedDate)
	if err != nil {
		return nil, fmt.Errorf("could not parse published date %s", err)
	}
	p := &pb.Publication{
		Data: &pb.Publication_Data{
			Type: "publication",
			Id:   id,
			Attributes: &pb.PublicationAttributes{
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
	var authors []*pb.Author
	for i, a := range attr.Authors {
		authors = append(authors, &pb.Author{
			FirstName: a.FirstName,
			LastName:  a.LastName,
			Rank:      int64(i),
			Initials:  a.Initials,
		})
	}
	p.Data.Attributes.Authors = authors
	logger.Debugf("successfully found publication with ID %s", pub.Data.ID)
	return p, nil
}
