package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type DashboardService struct {
	client *client.Client
}

func NewDashboardService(c *client.Client) *DashboardService {
	return &DashboardService{client: c}
}

type Dashboard struct {
	ID               string            `json:"id,omitempty"`
	Self             string            `json:"self,omitempty"`
	Name             string            `json:"name"`
	Description      string            `json:"description,omitempty"`
	Owner            *User             `json:"owner,omitempty"`
	IsFavourite      bool              `json:"isFavourite"`
	FavouritedBy     []User            `json:"favouritedBy,omitempty"`
	SharePermissions []SharePermission `json:"sharePermissions,omitempty"`
	EditPermissions  []SharePermission `json:"editPermissions,omitempty"`
	View             string            `json:"view,omitempty"`
	IsWritable       bool              `json:"isWritable,omitempty"`
	SystemDashboard  bool              `json:"systemDashboard,omitempty"`
	Popularity       int               `json:"popularity,omitempty"`
	Rank             int               `json:"rank,omitempty"`
}

type DashboardSearchResult struct {
	StartAt    int         `json:"startAt"`
	MaxResults int         `json:"maxResults"`
	Total      int         `json:"total"`
	Dashboards []Dashboard `json:"dashboards"`
}

type DashboardGadget struct {
	ID     string `json:"id,omitempty"`
	Module string `json:"module,omitempty"`
	Color  string `json:"color,omitempty"`
	Position struct {
		Row int `json:"row"`
		Col int `json:"column"`
	} `json:"position"`
	Title      string                 `json:"title,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

func (s *DashboardService) GetAll(ctx context.Context, startAt, maxResults int, filter string) (*DashboardSearchResult, error) {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard?startAt=%d&maxResults=%d", startAt, maxResults)
	if filter != "" {
		path += "&filter=" + url.QueryEscape(filter)
	}
	var result DashboardSearchResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DashboardService) Get(ctx context.Context, dashboardID string) (*Dashboard, error) {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s", url.PathEscape(dashboardID))
	var dashboard Dashboard
	if err := s.client.Get(ctx, path, &dashboard); err != nil {
		return nil, err
	}
	return &dashboard, nil
}

func (s *DashboardService) Create(ctx context.Context, dashboard *Dashboard) (*Dashboard, error) {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return nil, err
	}
	var result Dashboard
	if err := s.client.Post(ctx, "/rest/api/2/dashboard", dashboard, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DashboardService) Update(ctx context.Context, dashboardID string, dashboard *Dashboard) (*Dashboard, error) {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s", url.PathEscape(dashboardID))
	var result Dashboard
	if err := s.client.Put(ctx, path, dashboard, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DashboardService) Delete(ctx context.Context, dashboardID string) error {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s", url.PathEscape(dashboardID))
	return s.client.Delete(ctx, path)
}

func (s *DashboardService) GetGadgets(ctx context.Context, dashboardID string) ([]DashboardGadget, error) {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s/gadget", url.PathEscape(dashboardID))
	var result struct {
		Gadgets []DashboardGadget `json:"gadgets"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Gadgets, nil
}

func (s *DashboardService) AddGadget(ctx context.Context, dashboardID string, gadget *DashboardGadget) error {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s/gadget", url.PathEscape(dashboardID))
	return s.client.Post(ctx, path, gadget, nil)
}

func (s *DashboardService) RemoveGadget(ctx context.Context, dashboardID, gadgetID string) error {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s/gadget/%s", url.PathEscape(dashboardID), url.PathEscape(gadgetID))
	return s.client.Delete(ctx, path)
}

func (s *DashboardService) UpdateGadget(ctx context.Context, dashboardID, gadgetID string, gadget *DashboardGadget) error {
	if err := s.client.CheckEndpoint("dashboards"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/dashboard/%s/gadget/%s", url.PathEscape(dashboardID), url.PathEscape(gadgetID))
	return s.client.Put(ctx, path, gadget, nil)
}
