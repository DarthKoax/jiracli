package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type FilterService struct {
	client *client.Client
}

func NewFilterService(c *client.Client) *FilterService {
	return &FilterService{client: c}
}

type Filter struct {
	ID               string               `json:"id,omitempty"`
	Self             string               `json:"self,omitempty"`
	Name             string               `json:"name"`
	Description      string               `json:"description,omitempty"`
	Owner            *User                `json:"owner,omitempty"`
	JQL              string               `json:"jql"`
	ViewURL          string               `json:"viewUrl,omitempty"`
	SearchURL        string               `json:"searchUrl,omitempty"`
	Favourite        bool                 `json:"favourite"`
	SharePermissions []SharePermission    `json:"sharePermissions,omitempty"`
	EditPermissions  []SharePermission    `json:"editPermissions,omitempty"`
	Subscriptions    interface{}          `json:"subscriptions,omitempty"`
}

type SharePermission struct {
	ID      int          `json:"id,omitempty"`
	Type    string       `json:"type"`
	Project *Project     `json:"project,omitempty"`
	Role    *ProjectRole `json:"role,omitempty"`
	Group   *Group       `json:"group,omitempty"`
}

type FilterSubscription struct {
	ID   int   `json:"id,omitempty"`
	User *User `json:"user,omitempty"`
}

func (s *FilterService) Get(ctx context.Context, filterID string, expand []string) (*Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s", url.PathEscape(filterID))
	if len(expand) > 0 {
		path += "?expand=" + joinStrings(expand)
	}
	var filter Filter
	if err := s.client.Get(ctx, path, &filter); err != nil {
		return nil, err
	}
	return &filter, nil
}

func (s *FilterService) Create(ctx context.Context, filter *Filter) (*Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	var result Filter
	if err := s.client.Post(ctx, "/rest/api/2/filter", filter, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *FilterService) Update(ctx context.Context, filterID string, filter *Filter) (*Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s", url.PathEscape(filterID))
	var result Filter
	if err := s.client.Put(ctx, path, filter, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *FilterService) Delete(ctx context.Context, filterID string) error {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s", url.PathEscape(filterID))
	return s.client.Delete(ctx, path)
}

func (s *FilterService) GetFavourite(ctx context.Context) ([]Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	var filters []Filter
	if err := s.client.Get(ctx, "/rest/api/2/filter/favourite", &filters); err != nil {
		return nil, err
	}
	return filters, nil
}

func (s *FilterService) GetMy(ctx context.Context, startAt, maxResults int, expand []string) ([]Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/filter/my?startAt=%d&maxResults=%d", startAt, maxResults)
	if len(expand) > 0 {
		path += "&expand=" + joinStrings(expand)
	}
	var filters []Filter
	if err := s.client.Get(ctx, path, &filters); err != nil {
		return nil, err
	}
	return filters, nil
}

func (s *FilterService) Search(ctx context.Context, filterName, ownerName, groupName string, startAt, maxResults int) ([]Filter, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/filter/search?"
	params := url.Values{}
	if filterName != "" {
		params.Set("filterName", filterName)
	}
	if ownerName != "" {
		params.Set("owner", ownerName)
	}
	if groupName != "" {
		params.Set("groupname", groupName)
	}
	if startAt > 0 {
		params.Set("startAt", itoa(startAt))
	}
	if maxResults > 0 {
		params.Set("maxResults", itoa(maxResults))
	}
	path += params.Encode()

	var result struct {
		Filters []Filter `json:"filters"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Filters, nil
}

func (s *FilterService) AddSharePermission(ctx context.Context, filterID string, perm *SharePermission) ([]SharePermission, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s/permission", url.PathEscape(filterID))
	var result []SharePermission
	if err := s.client.Post(ctx, path, perm, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *FilterService) GetSharePermissions(ctx context.Context, filterID string) ([]SharePermission, error) {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s/permission", url.PathEscape(filterID))
	var perms []SharePermission
	if err := s.client.Get(ctx, path, &perms); err != nil {
		return nil, err
	}
	return perms, nil
}

func (s *FilterService) DeleteSharePermission(ctx context.Context, filterID string, permissionID int) error {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s/permission/%d", url.PathEscape(filterID), permissionID)
	return s.client.Delete(ctx, path)
}

func (s *FilterService) SetFavourite(ctx context.Context, filterID string, favourite bool) error {
	if err := s.client.CheckEndpoint("filters"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/filter/%s/favourite", url.PathEscape(filterID))
	body := map[string]bool{"favourite": favourite}
	return s.client.Put(ctx, path, body, nil)
}
