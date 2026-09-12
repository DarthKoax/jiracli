package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type ScreenService struct {
	client *client.Client
}

func NewScreenService(c *client.Client) *ScreenService {
	return &ScreenService{client: c}
}

type Screen struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Self        string `json:"self,omitempty"`
}

type ScreenTab struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Self string `json:"self,omitempty"`
}

type ScreenTabField struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Self string `json:"self,omitempty"`
}

type ScreenScheme struct {
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Screens     map[string]interface{} `json:"screens,omitempty"`
}

func (s *ScreenService) GetAll(ctx context.Context, startAt, maxResults int, queryString string) (*ScreenSearchResult, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens?startAt=%d&maxResults=%d", startAt, maxResults)
	if queryString != "" {
		path += "&queryString=" + url.QueryEscape(queryString)
	}
	var result ScreenSearchResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ScreenSearchResult struct {
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Total      int      `json:"total"`
	IsLast     bool     `json:"isLast"`
	Values     []Screen `json:"values"`
}

func (s *ScreenService) Get(ctx context.Context, screenID int) (*Screen, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d", screenID)
	var screen Screen
	if err := s.client.Get(ctx, path, &screen); err != nil {
		return nil, err
	}
	return &screen, nil
}

func (s *ScreenService) Create(ctx context.Context, screen *Screen) (*Screen, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("screens", "create"); err != nil {
		return nil, err
	}
	var result Screen
	if err := s.client.Post(ctx, "/rest/api/2/screens", screen, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ScreenService) Update(ctx context.Context, screenID int, screen *Screen) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("screens", "update"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d", screenID)
	return s.client.Put(ctx, path, screen, nil)
}

func (s *ScreenService) Delete(ctx context.Context, screenID int) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("screens", "delete"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d", screenID)
	return s.client.Delete(ctx, path)
}

func (s *ScreenService) GetTabs(ctx context.Context, screenID int) ([]ScreenTab, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs", screenID)
	var tabs []ScreenTab
	if err := s.client.Get(ctx, path, &tabs); err != nil {
		return nil, err
	}
	return tabs, nil
}

func (s *ScreenService) CreateTab(ctx context.Context, screenID int, tab *ScreenTab) (*ScreenTab, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs", screenID)
	var result ScreenTab
	if err := s.client.Post(ctx, path, tab, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ScreenService) UpdateTab(ctx context.Context, screenID, tabID int, tab *ScreenTab) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d", screenID, tabID)
	return s.client.Put(ctx, path, tab, nil)
}

func (s *ScreenService) DeleteTab(ctx context.Context, screenID, tabID int) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d", screenID, tabID)
	return s.client.Delete(ctx, path)
}

func (s *ScreenService) MoveTab(ctx context.Context, screenID, tabID, position int) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/move/%d", screenID, tabID, position)
	return s.client.Post(ctx, path, nil, nil)
}

func (s *ScreenService) GetTabFields(ctx context.Context, screenID, tabID int) ([]ScreenTabField, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/fields", screenID, tabID)
	var fields []ScreenTabField
	if err := s.client.Get(ctx, path, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func (s *ScreenService) AddTabField(ctx context.Context, screenID, tabID int, fieldID string) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/fields", screenID, tabID)
	body := map[string]string{"fieldId": fieldID}
	return s.client.Post(ctx, path, body, nil)
}

func (s *ScreenService) RemoveTabField(ctx context.Context, screenID, tabID int, fieldID string) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/fields/%s", screenID, tabID, url.PathEscape(fieldID))
	return s.client.Delete(ctx, path)
}

func (s *ScreenService) MoveTabField(ctx context.Context, screenID, tabID int, fieldID string, position int) error {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/fields/%s/move/%d", screenID, tabID, url.PathEscape(fieldID), position)
	return s.client.Post(ctx, path, nil, nil)
}

func (s *ScreenService) GetAvailableFields(ctx context.Context, screenID, tabID int) ([]ScreenTabField, error) {
	if err := s.client.CheckEndpoint("screens"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/screens/%d/tabs/%d/fields/available", screenID, tabID)
	var fields []ScreenTabField
	if err := s.client.Get(ctx, path, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}
