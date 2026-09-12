package api

import (
	"context"
	"fmt"

	"github.com/darkkoax/jiracli/internal/client"
)

type SearchService struct {
	client *client.Client
}

func NewSearchService(c *client.Client) *SearchService {
	return &SearchService{client: c}
}

type SearchRequest struct {
	JQL        string   `json:"jql"`
	StartAt    int      `json:"startAt,omitempty"`
	MaxResults int      `json:"maxResults,omitempty"`
	Fields     []string `json:"fields,omitempty"`
	Expand     []string `json:"expand,omitempty"`
}

type SearchResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

func (s *SearchService) Search(ctx context.Context, req *SearchRequest) (*SearchResult, error) {
	if err := s.client.CheckEndpoint("search"); err != nil {
		return nil, err
	}
	var result SearchResult
	if err := s.client.Post(ctx, "/rest/api/2/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SearchService) SearchGet(ctx context.Context, jql string, startAt, maxResults int, fields []string, expand []string) (*SearchResult, error) {
	if err := s.client.CheckEndpoint("search"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/search?jql=" + jql
	if startAt > 0 {
		path += "&startAt=" + itoa(startAt)
	}
	if maxResults > 0 {
		path += "&maxResults=" + itoa(maxResults)
	}
	if len(fields) > 0 {
		path += "&fields=" + joinStrings(fields)
	}
	if len(expand) > 0 {
		path += "&expand=" + joinStrings(expand)
	}

	var result SearchResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
