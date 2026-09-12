package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type BoardService struct {
	client *client.Client
}

func NewBoardService(c *client.Client) *BoardService {
	return &BoardService{client: c}
}

type Board struct {
	ID       int       `json:"id,omitempty"`
	Self     string    `json:"self,omitempty"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	FilterID int       `json:"filterId,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type Location struct {
	Type        string `json:"type,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	AvatarURI   string `json:"avatarURI,omitempty"`
	Name        string `json:"name,omitempty"`
}

type BoardSearchResult struct {
	MaxResults int     `json:"maxResults"`
	StartAt    int     `json:"startAt"`
	Total      int     `json:"total"`
	IsLast     bool    `json:"isLast"`
	Values     []Board `json:"values"`
}

type BoardConfig struct {
	ID                int                `json:"id,omitempty"`
	Self              string             `json:"self,omitempty"`
	Name              string             `json:"name,omitempty"`
	Filter            *FilterRef         `json:"filter,omitempty"`
	SubQuery          *SubQuery          `json:"subQuery,omitempty"`
	ColumnConfig      *ColumnConfig      `json:"columnConfig,omitempty"`
	Estimation        *Estimation        `json:"estimation,omitempty"`
	RapidViewId       int                `json:"rapidViewId,omitempty"`
	WorkingDaysConfig *WorkingDaysConfig `json:"workingDaysConfig,omitempty"`
}

type FilterRef struct {
	ID   string `json:"id"`
	Self string `json:"self"`
}

type SubQuery struct {
	Query string `json:"query"`
}

type ColumnConfig struct {
	Columns        []Column `json:"columns"`
	ConstraintType string   `json:"constraintType,omitempty"`
}

type Column struct {
	Name     string   `json:"name"`
	Statuses []Status `json:"statuses,omitempty"`
	Min      int      `json:"min,omitempty"`
	Max      int      `json:"max,omitempty"`
}

type Estimation struct {
	Type  string    `json:"type"`
	Field *FieldRef `json:"field,omitempty"`
}

type FieldRef struct {
	FieldID string `json:"fieldId"`
}

type WorkingDaysConfig struct {
	NonWorkingDays []string `json:"nonWorkingDays,omitempty"`
}

type PaginatedIssues struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	IsLast     bool    `json:"isLast"`
	Values     []Issue `json:"values"`
}

func (s *BoardService) GetAll(ctx context.Context, startAt, maxResults int, boardType, name, projectKeyOrID string) (*BoardSearchResult, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board?startAt=%d&maxResults=%d", startAt, maxResults)
	if boardType != "" {
		path += "&type=" + url.QueryEscape(boardType)
	}
	if name != "" {
		path += "&name=" + url.QueryEscape(name)
	}
	if projectKeyOrID != "" {
		path += "&projectKeyOrId=" + url.QueryEscape(projectKeyOrID)
	}
	var result BoardSearchResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BoardService) Get(ctx context.Context, boardID int) (*Board, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d", boardID)
	var board Board
	if err := s.client.Get(ctx, path, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *BoardService) Create(ctx context.Context, board *Board) (*Board, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	var result Board
	if err := s.client.Post(ctx, "/rest/agile/1.0/board", board, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BoardService) Delete(ctx context.Context, boardID int) error {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d", boardID)
	return s.client.Delete(ctx, path)
}

func (s *BoardService) GetConfig(ctx context.Context, boardID int) (*BoardConfig, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d/configuration", boardID)
	var config BoardConfig
	if err := s.client.Get(ctx, path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *BoardService) GetIssues(ctx context.Context, boardID int, startAt, maxResults int, jql string, fields []string) (*PaginatedIssues, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d/issue?startAt=%d&maxResults=%d", boardID, startAt, maxResults)
	if jql != "" {
		path += "&jql=" + url.QueryEscape(jql)
	}
	if len(fields) > 0 {
		path += "&fields=" + joinStrings(fields)
	}
	var result PaginatedIssues
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BoardService) GetEpics(ctx context.Context, boardID int, startAt, maxResults int, done bool) (*PaginatedEpics, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d/epic?startAt=%d&maxResults=%d&done=%t", boardID, startAt, maxResults, done)
	var result PaginatedEpics
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BoardService) GetSprints(ctx context.Context, boardID int, startAt, maxResults int, state string) (*PaginatedSprints, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d/sprint?startAt=%d&maxResults=%d", boardID, startAt, maxResults)
	if state != "" {
		path += "&state=" + url.QueryEscape(state)
	}
	var result PaginatedSprints
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BoardService) GetBacklog(ctx context.Context, boardID int, startAt, maxResults int, jql string, fields []string) (*PaginatedIssues, error) {
	if err := s.client.CheckEndpoint("boards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/agile/1.0/board/%d/backlog?startAt=%d&maxResults=%d", boardID, startAt, maxResults)
	if jql != "" {
		path += "&jql=" + url.QueryEscape(jql)
	}
	if len(fields) > 0 {
		path += "&fields=" + joinStrings(fields)
	}
	var result PaginatedIssues
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
