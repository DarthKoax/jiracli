package api

import (
	"context"

	"github.com/darthkoax/jiracli/internal/client"
)

type PriorityService struct {
	client *client.Client
}

func NewPriorityService(c *client.Client) *PriorityService {
	return &PriorityService{client: c}
}

type Priority struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	StatusColor string `json:"statusColor,omitempty"`
}

func (s *PriorityService) GetAll(ctx context.Context) ([]Priority, error) {
	if err := s.client.CheckEndpoint("priorities"); err != nil {
		return nil, err
	}
	var priorities []Priority
	if err := s.client.Get(ctx, "/rest/api/2/priority", &priorities); err != nil {
		return nil, err
	}
	return priorities, nil
}

func (s *PriorityService) Get(ctx context.Context, priorityID string) (*Priority, error) {
	if err := s.client.CheckEndpoint("priorities"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/priority/" + priorityID
	var priority Priority
	if err := s.client.Get(ctx, path, &priority); err != nil {
		return nil, err
	}
	return &priority, nil
}
