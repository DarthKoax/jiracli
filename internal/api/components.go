package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type ComponentService struct {
	client *client.Client
}

func NewComponentService(c *client.Client) *ComponentService {
	return &ComponentService{client: c}
}

type Component struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name"`
	Description         string `json:"description,omitempty"`
	Lead                *User  `json:"lead,omitempty"`
	LeadUserName        string `json:"leadUserName,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	Assignee            *User  `json:"assignee,omitempty"`
	RealAssignee        *User  `json:"realAssignee,omitempty"`
	Project             string `json:"project,omitempty"`
	ProjectID           int    `json:"projectId,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty"`
	Self                string `json:"self,omitempty"`
}

type ComponentIssueCount struct {
	Self            string `json:"self"`
	IssueCount      int    `json:"issueCount"`
	Lead            *User  `json:"lead,omitempty"`
	Description     string `json:"description,omitempty"`
}

func (s *ComponentService) Get(ctx context.Context, componentID string) (*Component, error) {
	if err := s.client.CheckEndpoint("components"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/component/%s", url.PathEscape(componentID))
	var component Component
	if err := s.client.Get(ctx, path, &component); err != nil {
		return nil, err
	}
	return &component, nil
}

func (s *ComponentService) Create(ctx context.Context, component *Component) (*Component, error) {
	if err := s.client.CheckEndpoint("components"); err != nil {
		return nil, err
	}
	var result Component
	if err := s.client.Post(ctx, "/rest/api/2/component", component, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ComponentService) Update(ctx context.Context, componentID string, component *Component) error {
	if err := s.client.CheckEndpoint("components"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/component/%s", url.PathEscape(componentID))
	return s.client.Put(ctx, path, component, nil)
}

func (s *ComponentService) Delete(ctx context.Context, componentID string, replaceWith string) error {
	if err := s.client.CheckEndpoint("components"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/component/%s", url.PathEscape(componentID))
	if replaceWith != "" {
		path += "?replaceWith=" + url.QueryEscape(replaceWith)
	}
	return s.client.Delete(ctx, path)
}

func (s *ComponentService) GetRelatedIssueCount(ctx context.Context, componentID string) (*ComponentIssueCount, error) {
	if err := s.client.CheckEndpoint("components"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/component/%s/relatedIssueCounts", url.PathEscape(componentID))
	var result ComponentIssueCount
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
