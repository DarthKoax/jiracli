package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type SprintService struct {
	client *client.Client
}

func NewSprintService(c *client.Client) *SprintService {
	return &SprintService{client: c}
}

type Sprint struct {
	ID            int    `json:"id,omitempty"`
	Self          string `json:"self,omitempty"`
	Name          string `json:"name"`
	State         string `json:"state,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	CompleteDate  string `json:"completeDate,omitempty"`
	OriginBoardID int    `json:"originBoardId,omitempty"`
	Goal          string `json:"goal,omitempty"`
}

type PaginatedSprints struct {
	MaxResults int      `json:"maxResults"`
	StartAt    int      `json:"startAt"`
	Total      int      `json:"total"`
	IsLast     bool     `json:"isLast"`
	Values     []Sprint `json:"values"`
}

type PaginatedEpics struct {
	MaxResults int    `json:"maxResults"`
	StartAt    int    `json:"startAt"`
	Total      int    `json:"total"`
	IsLast     bool   `json:"isLast"`
	Values     []Epic `json:"values"`
}

type Epic struct {
	ID      int    `json:"id,omitempty"`
	Key     string `json:"key,omitempty"`
	Self    string `json:"self,omitempty"`
	Name    string `json:"name"`
	Summary string `json:"summary,omitempty"`
	Done    bool   `json:"done"`
}

func (s *SprintService) Get(ctx context.Context, sprintID int) (*Sprint, error) {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d", sprintID)
	var sprint Sprint
	if err := s.client.Get(ctx, path, &sprint); err != nil {
		return nil, err
	}
	return &sprint, nil
}

func (s *SprintService) Create(ctx context.Context, boardID int, sprint *Sprint) (*Sprint, error) {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return nil, err
	}
	_ = boardID
	var result Sprint
	if err := s.client.Post(ctx, "/rest/agile/1.0/sprint", sprint, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SprintService) Update(ctx context.Context, sprintID int, sprint *Sprint) (*Sprint, error) {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d", sprintID)
	var result Sprint
	if err := s.client.Put(ctx, path, sprint, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SprintService) Delete(ctx context.Context, sprintID int) error {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d", sprintID)
	return s.client.Delete(ctx, path)
}

func (s *SprintService) GetIssues(ctx context.Context, sprintID int, startAt, maxResults int, jql string, fields []string) (*PaginatedIssues, error) {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d/issue?startAt=%d&maxResults=%d", sprintID, startAt, maxResults)
	if jql != "" {
		path += "&jql=" + url.QueryEscape(jql)
	}
	if len(fields) > 0 {
		path += "&fields=" + url.QueryEscape(joinStrings(fields))
	}
	var result PaginatedIssues
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SprintService) MoveIssues(ctx context.Context, sprintID int, issueKeys []string) error {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d/issue", sprintID)
	body := map[string]interface{}{"issues": issueKeys}
	return s.client.Post(ctx, path, body, nil)
}

func (s *SprintService) SwapIssues(ctx context.Context, sprintID int, issueKeys []string) error {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d/issueswap", sprintID)
	body := map[string]interface{}{"issues": issueKeys}
	return s.client.Post(ctx, path, body, nil)
}

func (s *SprintService) Complete(ctx context.Context, sprintID int, completeDate string) (*Sprint, error) {
	if err := s.client.CheckEndpoint("sprints"); err != nil {
		return nil, err
	}
	
	// First get the current sprint details
	sprint, err := s.Get(ctx, sprintID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sprint details: %w", err)
	}
	
	// Update with closed state and complete date
	path := fmt.Sprintf("/rest/agile/1.0/sprint/%d", sprintID)
	body := map[string]interface{}{
		"name":          sprint.Name,
		"state":         "closed",
		"completeDate":  completeDate,
		"originBoardId": sprint.OriginBoardID,
	}
	if sprint.StartDate != "" {
		body["startDate"] = sprint.StartDate
	}
	if sprint.EndDate != "" {
		body["endDate"] = sprint.EndDate
	}
	
	var result Sprint
	if err := s.client.Put(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
