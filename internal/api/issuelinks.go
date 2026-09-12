package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type IssueLinkService struct {
	client *client.Client
}

func NewIssueLinkService(c *client.Client) *IssueLinkService {
	return &IssueLinkService{client: c}
}

type IssueLink struct {
	ID           string `json:"id"`
	Self         string `json:"self"`
	Type         *IssueLinkType `json:"type"`
	InwardIssue  *Issue `json:"inwardIssue"`
	OutwardIssue *Issue `json:"outwardIssue"`
	Comment      *Comment `json:"comment,omitempty"`
}

type IssueLinkType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
	Self    string `json:"self"`
}

type IssueLinkCreatePayload struct {
	Type         *IssueLinkTypeRef `json:"type"`
	InwardIssue  *IssueRef         `json:"inwardIssue"`
	OutwardIssue *IssueRef         `json:"outwardIssue"`
	Comment      *Comment          `json:"comment,omitempty"`
}

type IssueLinkTypeRef struct {
	Name string `json:"name"`
}

type IssueRef struct {
	Key string `json:"key"`
}

func (s *IssueLinkService) Get(ctx context.Context, linkID string) (*IssueLink, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issueLink/%s", url.PathEscape(linkID))
	var link IssueLink
	if err := s.client.Get(ctx, path, &link); err != nil {
		return nil, err
	}
	return &link, nil
}

func (s *IssueLinkService) Create(ctx context.Context, payload *IssueLinkCreatePayload) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	return s.client.Post(ctx, "/rest/api/2/issueLink", payload, nil)
}

func (s *IssueLinkService) Delete(ctx context.Context, linkID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issueLink/%s", url.PathEscape(linkID))
	return s.client.Delete(ctx, path)
}

func (s *IssueLinkService) AddComment(ctx context.Context, linkID string, comment *Comment) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issueLink/%s/comment", url.PathEscape(linkID))
	return s.client.Post(ctx, path, comment, nil)
}
