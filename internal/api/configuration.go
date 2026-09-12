package api

import (
	"context"

	"github.com/darthkoax/jiracli/internal/client"
)

type ConfigurationService struct {
	client *client.Client
}

func NewConfigurationService(c *client.Client) *ConfigurationService {
	return &ConfigurationService{client: c}
}

type Configuration struct {
	Self              string `json:"self"`
	VotingEnabled     bool   `json:"votingEnabled"`
	WatchingEnabled   bool   `json:"watchingEnabled"`
	UnassignedIssuesAllowed bool `json:"unassignedIssuesAllowed"`
	SubTasksEnabled   bool   `json:"subTasksEnabled"`
	IssueLinkingEnabled bool `json:"issueLinkingEnabled"`
	TimeTrackingEnabled bool `json:"timeTrackingEnabled"`
	AttachmentsEnabled bool  `json:"attachmentsEnabled"`
}

func (s *ConfigurationService) Get(ctx context.Context) (*Configuration, error) {
	if err := s.client.CheckEndpoint("configuration"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("configuration", "get"); err != nil {
		return nil, err
	}
	var config Configuration
	if err := s.client.Get(ctx, "/rest/api/2/configuration", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
