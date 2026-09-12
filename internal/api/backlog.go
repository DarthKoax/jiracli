package api

import (
	"context"

	"github.com/darkkoax/jiracli/internal/client"
)

type BacklogService struct {
	client *client.Client
}

func NewBacklogService(c *client.Client) *BacklogService {
	return &BacklogService{client: c}
}

func (s *BacklogService) MoveToBacklog(ctx context.Context, issueKeys []string) error {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return err
	}
	payload := map[string]interface{}{"issues": issueKeys}
	return s.client.Post(ctx, "/rest/agile/1.0/backlog/issue", payload, nil)
}

func (s *BacklogService) MoveIssueToBacklog(ctx context.Context, issueKeys []string) error {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return err
	}
	payload := map[string]interface{}{"issues": issueKeys}
	return s.client.Post(ctx, "/rest/agile/1.0/issue/moveToBacklog", payload, nil)
}
