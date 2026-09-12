package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type ChangelogService struct {
	client *client.Client
}

func NewChangelogService(c *client.Client) *ChangelogService {
	return &ChangelogService{client: c}
}

type Changelog struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Values     []ChangelogEntry `json:"values"`
}

type ChangelogEntry struct {
	ID      string        `json:"id"`
	Author  *User         `json:"author"`
	Created string        `json:"created"`
	Items   []ChangelogItem `json:"items"`
}

type ChangelogItem struct {
	Field      string `json:"field"`
	FieldType  string `json:"fieldtype"`
	FieldID    string `json:"fieldId,omitempty"`
	From       string `json:"from,omitempty"`
	FromString string `json:"fromString,omitempty"`
	To         string `json:"to,omitempty"`
	ToString   string `json:"toString,omitempty"`
}

func (s *ChangelogService) Get(ctx context.Context, issueKeyOrID string, startAt, maxResults int) (*Changelog, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/changelog?startAt=%d&maxResults=%d",
		url.PathEscape(issueKeyOrID), startAt, maxResults)
	var changelog Changelog
	if err := s.client.Get(ctx, path, &changelog); err != nil {
		return nil, err
	}
	return &changelog, nil
}
