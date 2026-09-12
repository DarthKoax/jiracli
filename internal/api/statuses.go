package api

import (
	"context"

	"github.com/darthkoax/jiracli/internal/client"
)

type StatusService struct {
	client *client.Client
}

func NewStatusService(c *client.Client) *StatusService {
	return &StatusService{client: c}
}

type StatusDetail struct {
	Self        string       `json:"self"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	IconURL     string       `json:"iconUrl,omitempty"`
	StatusCategory *StatusCategory `json:"statusCategory,omitempty"`
}

type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

func (s *StatusService) GetAll(ctx context.Context) ([]StatusDetail, error) {
	if err := s.client.CheckEndpoint("statuses"); err != nil {
		return nil, err
	}
	var statuses []StatusDetail
	if err := s.client.Get(ctx, "/rest/api/2/status", &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}

func (s *StatusService) Get(ctx context.Context, statusID string) (*StatusDetail, error) {
	if err := s.client.CheckEndpoint("statuses"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/status/" + statusID
	var status StatusDetail
	if err := s.client.Get(ctx, path, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

func (s *StatusService) GetCategories(ctx context.Context) ([]StatusCategory, error) {
	if err := s.client.CheckEndpoint("statuses"); err != nil {
		return nil, err
	}
	var categories []StatusCategory
	if err := s.client.Get(ctx, "/rest/api/2/statuscategory", &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *StatusService) GetCategory(ctx context.Context, categoryID string) (*StatusCategory, error) {
	if err := s.client.CheckEndpoint("statuses"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/statuscategory/" + categoryID
	var category StatusCategory
	if err := s.client.Get(ctx, path, &category); err != nil {
		return nil, err
	}
	return &category, nil
}
