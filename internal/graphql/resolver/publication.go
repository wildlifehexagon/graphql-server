package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dictyBase/graphql-server/internal/registry"

	"github.com/dictyBase/graphql-server/internal/graphql/models"
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
	Issue         int64     `json:"issue,omitempty"`
	PublishedDate string    `json:"publication_date"`
	Authors       []*Author `json:"authors"`
}

type Author struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Initials  string `json:"initials"`
}

// Publication is the resolver for getting an individual publication by ID.
func (q *QueryResolver) Publication(ctx context.Context, id string) (*models.Publication, error) {
	endpoint := q.GetAPIEndpoint(registry.PUBLICATION)
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
	// convert pub_date to expected time.Time format
	pd, err := time.Parse("2006-01-02", attr.PublishedDate)
	if err != nil {
		return nil, fmt.Errorf("could not parse published date %s", err)
	}
	// convert issue to expected string format
	issue := strconv.Itoa(int(attr.Issue))
	p := &models.Publication{
		ID:       pub.Data.ID,
		Doi:      &attr.Doi,
		Title:    &attr.Title,
		Abstract: &attr.Abstract,
		Journal:  &attr.Journal,
		PubDate:  &pd,
		Pages:    &attr.Page,
		Issn:     &attr.Issn,
		PubType:  &attr.PubType,
		Source:   &attr.Source,
		Issue:    &issue,
		Status:   &attr.Status,
		// Volume: nil,
	}
	var authors []*models.Author
	for _, a := range attr.Authors {
		authors = append(authors, &models.Author{
			FirstName: &a.FirstName,
			LastName:  &a.LastName,
			// Rank:      nil,
			Initials: &a.Initials,
		})
	}
	p.Authors = authors
	q.Logger.Debugf("successfully found publication with ID %s", pub.Data.ID)
	return p, nil
}
