package fetch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/dictyBase/aphgrpc"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/sirupsen/logrus"
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
	Abstract      string    `json:"abstract"`
	Doi           string    `json:"doi,omitempty"`
	FullTextURL   string    `json:"full_text_url,omitempty"`
	PubmedURL     string    `json:"pubmed_url"`
	Journal       string    `json:"journal"`
	Issn          string    `json:"issn,omitempty"`
	Page          string    `json:"page,omitempty"`
	Pubmed        string    `json:"pubmed"`
	Title         string    `json:"title"`
	Source        string    `json:"source"`
	Status        string    `json:"status"`
	PubType       string    `json:"pub_type"`
	Issue         string    `json:"issue"`
	Volume        string    `json:"volume"`
	PublishedDate string    `json:"publication_date"`
	Authors       []*Author `json:"authors"`
}

type Author struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Initials  string `json:"initials"`
}

func FetchPublication(ctx context.Context, endpoint, id string) (*pb.Publication, error) {
	logger := *logrus.New()
	url := endpoint + "/" + id
	res, err := GetResp(ctx, url)
	if err != nil {
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
				Issue:    attr.Issue,
				Volume:   attr.Volume,
				Status:   attr.Status,
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

func FetchDOI(ctx context.Context, doi string) (*pb.Publication, error) {
	pub := &pb.Publication{}
	reqURL, _ := url.Parse(fmt.Sprintf("https://doi.org/%s", doi))
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Accept": {"application/vnd.citationstyles.csl+json"},
		},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return pub, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	j, err := gabs.ParseJSONDecoder(decoder)
	if err != nil {
		return pub, err
	}
	authors := []*publication.Author{}
	a := j.Search("author")
	for _, v := range a.Children() {
		n := &publication.Author{}
		for key, val := range v.ChildrenMap() {
			if key == "given" {
				n.FirstName = val.Data().(string)
			}
			if key == "family" {
				n.LastName = val.Data().(string)
			}
		}
		authors = append(authors, n)
	}
	d := j.Search("created", "date-time").Data().(string)
	pd, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return nil, fmt.Errorf("could not parse published date %s", err)
	}
	p := &publication.Publication{
		Data: &publication.Publication_Data{
			Id: "",
			Attributes: &publication.PublicationAttributes{
				Doi:      doi,
				Title:    j.Search("title").Data().(string),
				Abstract: j.Search("abstract").Data().(string),
				Journal:  "",
				PubDate:  aphgrpc.TimestampProto(pd),
				Volume:   "",
				Pages:    "",
				Issn:     "",
				PubType:  j.Search("type").Data().(string),
				Source:   j.Search("source").Data().(string),
				Issue:    "",
				Status:   "",
				Authors:  authors,
			},
		},
	}
	return p, nil
}
