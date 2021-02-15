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
	res, err := getDOIResp(ctx, doi)
	if err != nil {
		return pub, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	j, err := gabs.ParseJSONDecoder(decoder)
	if err != nil {
		return pub, err
	}
	authors := getDOIAuthors(j.Search("author").Children())
	d := j.Search("created", "date-time").Data().(string)
	pd, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return nil, fmt.Errorf("could not parse published date %s", err)
	}
	p := &pb.Publication{
		Data: &pb.Publication_Data{
			Id: "", // doi do not have pubmed IDs listed
			Attributes: &pb.PublicationAttributes{
				Doi:      doi,
				Title:    verifyStringProperty(j, "title"),
				Abstract: verifyStringProperty(j, "abstract"),
				Journal:  verifyStringProperty(j, "container-title-short"),
				PubDate:  aphgrpc.TimestampProto(pd),
				Volume:   verifyStringProperty(j, "volume"),
				Pages:    verifyStringProperty(j, "page"),
				Issn:     verifyArrayProperty(j, "ISSN"),
				PubType:  verifyStringProperty(j, "type"),
				Source:   verifyStringProperty(j, "source"),
				Issue:    verifyStringProperty(j, "issue"),
				Status:   verifyStringProperty(j, "subtype"),
				Authors:  authors,
			},
		},
	}
	return p, nil
}

// getDOIResp makes HTTP request with necessary
// headers for DOI and returns the response
func getDOIResp(ctx context.Context, doi string) (*http.Response, error) {
	r := &http.Response{}
	url, err := url.Parse(doi)
	if err != nil {
		return r, err
	}
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"Accept": {"application/vnd.citationstyles.csl+json"},
		},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

// getDOIAuthors converts DOI authors data into expected Author format
func getDOIAuthors(authors []*gabs.Container) []*pb.Author {
	a := []*pb.Author{}
	for _, v := range authors {
		n := &pb.Author{}
		for key, val := range v.ChildrenMap() {
			if key == "given" {
				n.FirstName = val.Data().(string)
			}
			if key == "family" {
				n.LastName = val.Data().(string)
			}
		}
		a = append(a, n)
	}
	return a
}

// verifyStringProperty checks if a property exists in the JSON and returns the
// value if true
func verifyStringProperty(j *gabs.Container, val string) string {
	if j.Exists(val) {
		return j.Search(val).Data().(string)
	}
	return ""
}

// verifyArrayProperty checks if a property exists in the JSON and then returns
// its first child as a string
func verifyArrayProperty(j *gabs.Container, val string) string {
	if j.Exists(val) {
		return j.Search(val).Children()[0].Data().(string)
	}
	return ""
}
