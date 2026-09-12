package api

import (
	"context"
	"fmt"

	"github.com/darthkoax/jiracli/internal/client"
)

type ReindexService struct {
	client *client.Client
}

func NewReindexService(c *client.Client) *ReindexService {
	return &ReindexService{client: c}
}

type ReindexResult struct {
	ReindexRequestID int    `json:"reindexRequestId,omitempty"`
	Progress         float64 `json:"progress,omitempty"`
	Self             string `json:"self,omitempty"`
}

type ReindexRequest struct {
	Type string `json:"type,omitempty"`
}

func (s *ReindexService) Trigger(ctx context.Context, reindexType string) (*ReindexResult, error) {
	if err := s.client.CheckEndpoint("reindex"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("reindex", "trigger"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/reindex"
	if reindexType != "" {
		path += "?type=" + reindexType
	}
	var result ReindexResult
	if err := s.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ReindexService) GetStatus(ctx context.Context, taskID string) (*ReindexResult, error) {
	if err := s.client.CheckEndpoint("reindex"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/reindex?taskId=%s", taskID)
	var result ReindexResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ReindexService) GetRequestStatus(ctx context.Context, requestID int) (*ReindexResult, error) {
	if err := s.client.CheckEndpoint("reindex"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/reindex/%d", requestID)
	var result ReindexResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
