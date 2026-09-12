package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type PermissionService struct {
	client *client.Client
}

func NewPermissionService(c *client.Client) *PermissionService {
	return &PermissionService{client: c}
}

type Permissions struct {
	Permissions map[string]PermissionScheme `json:"permissions"`
}

type PermissionScheme struct {
	ID       int    `json:"id"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Description string `json:"description,omitempty"`
}

type MyPermissions struct {
	Permissions map[string]MyPermissionDetail `json:"permissions"`
}

type MyPermissionDetail struct {
	ID       int    `json:"id"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	HavePermission bool `json:"havePermission"`
}

type PermissionSchemeFull struct {
	ID          int                `json:"id,omitempty"`
	Self        string             `json:"self,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Permissions []PermissionGrant  `json:"permissions,omitempty"`
}

type PermissionGrant struct {
	ID         int            `json:"id,omitempty"`
	Self       string         `json:"self,omitempty"`
	Holder     *PermissionHolder `json:"holder,omitempty"`
	Permission string         `json:"permission"`
}

type PermissionHolder struct {
	Type      string `json:"type"`
	Parameter string `json:"parameter,omitempty"`
	Expand    string `json:"expand,omitempty"`
}

func (s *PermissionService) GetAll(ctx context.Context) (*Permissions, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	var perms Permissions
	if err := s.client.Get(ctx, "/rest/api/2/permissions", &perms); err != nil {
		return nil, err
	}
	return &perms, nil
}

func (s *PermissionService) GetMyPermissions(ctx context.Context, projectKey, projectID, issueKey, issueID, permID string) (*MyPermissions, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/mypermissions?"
	params := url.Values{}
	if projectKey != "" {
		params.Set("projectKey", projectKey)
	}
	if projectID != "" {
		params.Set("projectId", projectID)
	}
	if issueKey != "" {
		params.Set("issueKey", issueKey)
	}
	if issueID != "" {
		params.Set("issueId", issueID)
	}
	if permID != "" {
		params.Set("permissions", permID)
	}
	path += params.Encode()

	var perms MyPermissions
	if err := s.client.Get(ctx, path, &perms); err != nil {
		return nil, err
	}
	return &perms, nil
}

func (s *PermissionService) GetAllSchemes(ctx context.Context, startAt, maxResults int, expand string) (*PermissionSchemeSearchResult, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme?startAt=%d&maxResults=%d", startAt, maxResults)
	if expand != "" {
		path += "&expand=" + url.QueryEscape(expand)
	}
	var result PermissionSchemeSearchResult
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PermissionSchemeSearchResult struct {
	StartAt         int                  `json:"startAt"`
	MaxResults      int                  `json:"maxResults"`
	Total           int                  `json:"total"`
	PermissionSchemes []PermissionSchemeFull `json:"permissionSchemes"`
}

func (s *PermissionService) GetScheme(ctx context.Context, schemeID string, expand string) (*PermissionSchemeFull, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme/%s", url.PathEscape(schemeID))
	if expand != "" {
		path += "?expand=" + url.QueryEscape(expand)
	}
	var scheme PermissionSchemeFull
	if err := s.client.Get(ctx, path, &scheme); err != nil {
		return nil, err
	}
	return &scheme, nil
}

func (s *PermissionService) CreateScheme(ctx context.Context, scheme *PermissionSchemeFull) (*PermissionSchemeFull, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("permissions", "create_scheme"); err != nil {
		return nil, err
	}
	var result PermissionSchemeFull
	if err := s.client.Post(ctx, "/rest/api/2/permissionscheme", scheme, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PermissionService) DeleteScheme(ctx context.Context, schemeID string) error {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("permissions", "delete_scheme"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme/%s", url.PathEscape(schemeID))
	return s.client.Delete(ctx, path)
}

func (s *PermissionService) UpdateScheme(ctx context.Context, schemeID string, scheme *PermissionSchemeFull) (*PermissionSchemeFull, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("permissions", "update_scheme"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme/%s", url.PathEscape(schemeID))
	var result PermissionSchemeFull
	if err := s.client.Put(ctx, path, scheme, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PermissionService) AddPermissionGrant(ctx context.Context, schemeID string, grant *PermissionGrant) (*PermissionSchemeFull, error) {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme/%s/permission", url.PathEscape(schemeID))
	var result PermissionSchemeFull
	if err := s.client.Post(ctx, path, grant, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PermissionService) RemovePermissionGrant(ctx context.Context, schemeID string, permissionID string) error {
	if err := s.client.CheckEndpoint("permissions"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/permissionscheme/%s/permission/%s",
		url.PathEscape(schemeID), url.PathEscape(permissionID))
	return s.client.Delete(ctx, path)
}
