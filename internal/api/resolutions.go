package api

import (
	"context"

	"github.com/darthkoax/jiracli/internal/client"
)

type ResolutionService struct {
	client *client.Client
}

func NewResolutionService(c *client.Client) *ResolutionService {
	return &ResolutionService{client: c}
}

type Resolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (s *ResolutionService) GetAll(ctx context.Context) ([]Resolution, error) {
	if err := s.client.CheckEndpoint("resolutions"); err != nil {
		return nil, err
	}
	var resolutions []Resolution
	if err := s.client.Get(ctx, "/rest/api/2/resolution", &resolutions); err != nil {
		return nil, err
	}
	return resolutions, nil
}

func (s *ResolutionService) Get(ctx context.Context, resolutionID string) (*Resolution, error) {
	if err := s.client.CheckEndpoint("resolutions"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/resolution/" + resolutionID
	var resolution Resolution
	if err := s.client.Get(ctx, path, &resolution); err != nil {
		return nil, err
	}
	return &resolution, nil
}
