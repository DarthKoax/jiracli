package api

import (
	"context"

	"github.com/darthkoax/jiracli/internal/client"
)

type HealthCheckService struct {
	client *client.Client
}

func NewHealthCheckService(c *client.Client) *HealthCheckService {
	return &HealthCheckService{client: c}
}

type HealthCheckResult struct {
	Name    string `json:"name"`
	Description string `json:"description"`
	Passed  bool   `json:"passed"`
	Failed  bool   `json:"failed,omitempty"`
	ErrorRate string `json:"errorRate,omitempty"`
	Severity string `json:"severity,omitempty"`
	Time    int64  `json:"time,omitempty"`
}

type HealthCheckResponse struct {
	State   string              `json:"state"`
	Checks  []HealthCheckResult `json:"checks"`
}

func (s *HealthCheckService) Get(ctx context.Context) (*HealthCheckResponse, error) {
	if err := s.client.CheckEndpoint("healthcheck"); err != nil {
		return nil, err
	}
	var result HealthCheckResponse
	if err := s.client.Get(ctx, "/rest/troubleshooting/1.0/check", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
