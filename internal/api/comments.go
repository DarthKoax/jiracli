package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type StandaloneCommentService struct {
	client *client.Client
}

func NewStandaloneCommentService(c *client.Client) *StandaloneCommentService {
	return &StandaloneCommentService{client: c}
}

type StandaloneComment struct {
	ID           string `json:"id"`
	Self         string `json:"self"`
	Body         string `json:"body"`
	Author       *User  `json:"author"`
	UpdateAuthor *User  `json:"updateAuthor"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
	Visibility   *struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility,omitempty"`
	JSDAuthor *struct {
		Name string `json:"name,omitempty"`
		Key  string `json:"key,omitempty"`
	} `json:"jsdAuthor,omitempty"`
}

type CommentListResult struct {
	Comments   []StandaloneComment `json:"comments"`
	MaxResults int                 `json:"maxResults"`
	StartAt    int                 `json:"startAt"`
	Total      int                 `json:"total"`
}

func (s *StandaloneCommentService) Get(ctx context.Context, commentID string) (*StandaloneComment, error) {
	if err := s.client.CheckEndpoint("comments"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/comment/%s", url.PathEscape(commentID))
	var comment StandaloneComment
	if err := s.client.Get(ctx, path, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *StandaloneCommentService) Update(ctx context.Context, commentID string, body string) (*StandaloneComment, error) {
	if err := s.client.CheckEndpoint("comments"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/comment/%s", url.PathEscape(commentID))
	payload := map[string]string{"body": body}
	var comment StandaloneComment
	if err := s.client.Put(ctx, path, payload, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *StandaloneCommentService) Delete(ctx context.Context, commentID string) error {
	if err := s.client.CheckEndpoint("comments"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/comment/%s", url.PathEscape(commentID))
	return s.client.Delete(ctx, path)
}

func (s *StandaloneCommentService) List(ctx context.Context, commentIDs []string) (*CommentListResult, error) {
	if err := s.client.CheckEndpoint("comments"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/comment/list"
	payload := map[string]interface{}{"ids": commentIDs}
	var result CommentListResult
	if err := s.client.Post(ctx, path, payload, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
