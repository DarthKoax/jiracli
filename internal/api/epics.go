package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type EpicService struct {
	client *client.Client
}

func NewEpicService(c *client.Client) *EpicService {
	return &EpicService{client: c}
}

type EpicIssuesResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	IsLast     bool    `json:"isLast"`
	Values     []Issue `json:"values"`
}

func (s *EpicService) GetIssues(ctx context.Context, epicKeyOrID string, startAt, maxResults int, jql string, fields []string) (*EpicIssuesResult, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/epic/%s/issue?startAt=%d&maxResults=%d",
		url.PathEscape(epicKeyOrID), startAt, maxResults)
	if jql != "" {
		path += "&jql=" + url.QueryEscape(jql)
	}
	if len(fields) > 0 {
		path += "&fields=" + joinStrings(fields)
	}
	var result EpicIssuesResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *EpicService) RemoveIssues(ctx context.Context, epicKeyOrID string, issueKeys []string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/agile/1.0/epic/%s/remove", url.PathEscape(epicKeyOrID))
	payload := map[string]interface{}{"issues": issueKeys}
	return s.client.Post(ctx, path, payload, nil)
}
