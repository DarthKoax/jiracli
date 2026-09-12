package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type ProjectService struct {
	client *client.Client
}

func NewProjectService(c *client.Client) *ProjectService {
	return &ProjectService{client: c}
}

type Project struct {
	ID          string            `json:"id,omitempty"`
	Key         string            `json:"key,omitempty"`
	Name        string            `json:"name,omitempty"`
	Self        string            `json:"self,omitempty"`
	Description string            `json:"description,omitempty"`
	Lead        *User             `json:"lead,omitempty"`
	ProjectType string            `json:"projectTypeKey,omitempty"`
	Roles       map[string]string `json:"roles,omitempty"`
	AvatarURLs  *AvatarURLs       `json:"avatarUrls,omitempty"`
}

type ProjectRole struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Self   string  `json:"self"`
	Actors []Actor `json:"actors,omitempty"`
}

type Actor struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	Type        string `json:"type"`
	Name        string `json:"name,omitempty"`
	AvatarURL   string `json:"actorUser,omitempty"`
}

type AvatarURLs struct {
	Four8x48  string `json:"48x48"`
	Two4x24   string `json:"24x24"`
	One6x16   string `json:"16x16"`
	Three2x32 string `json:"32x32"`
}

type ProjectComponent struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name"`
	Description         string `json:"description,omitempty"`
	Lead                *User  `json:"lead,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	RealAssignee        string `json:"realAssignee,omitempty"`
	Project             string `json:"project,omitempty"`
	ProjectID           int    `json:"projectId,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty"`
}

type ProjectVersion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Archived    bool   `json:"archived"`
	Released    bool   `json:"released"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Project     string `json:"project,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
	Self        string `json:"self,omitempty"`
}

func (s *ProjectService) GetAll(ctx context.Context, expand []string) ([]Project, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/project"
	if len(expand) > 0 {
		path += "?expand=" + joinStrings(expand)
	}
	var projects []Project
	if err := s.client.Get(ctx, path, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *ProjectService) Get(ctx context.Context, projectKeyOrID string, expand []string) (*Project, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s", url.PathEscape(projectKeyOrID))
	if len(expand) > 0 {
		path += "?expand=" + joinStrings(expand)
	}
	var project Project
	if err := s.client.Get(ctx, path, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *ProjectService) Create(ctx context.Context, project *Project) (*Project, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	var result Project
	if err := s.client.Post(ctx, "/rest/api/2/project", project, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ProjectService) Update(ctx context.Context, projectKeyOrID string, project *Project) error {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s", url.PathEscape(projectKeyOrID))
	return s.client.Put(ctx, path, project, nil)
}

func (s *ProjectService) Delete(ctx context.Context, projectKeyOrID string) error {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s", url.PathEscape(projectKeyOrID))
	return s.client.Delete(ctx, path)
}

func (s *ProjectService) GetRoles(ctx context.Context, projectKeyOrID string) ([]ProjectRole, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/role", url.PathEscape(projectKeyOrID))
	var roles []ProjectRole
	if err := s.client.Get(ctx, path, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *ProjectService) GetRole(ctx context.Context, projectKeyOrID string, roleID int) (*ProjectRole, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/role/%d", url.PathEscape(projectKeyOrID), roleID)
	var role ProjectRole
	if err := s.client.Get(ctx, path, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *ProjectService) UpdateRole(ctx context.Context, projectKeyOrID string, roleID int, role *ProjectRole) error {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/role/%d", url.PathEscape(projectKeyOrID), roleID)
	return s.client.Post(ctx, path, role, nil)
}

func (s *ProjectService) GetComponents(ctx context.Context, projectKeyOrID string) ([]ProjectComponent, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/components", url.PathEscape(projectKeyOrID))
	var components []ProjectComponent
	if err := s.client.Get(ctx, path, &components); err != nil {
		return nil, err
	}
	return components, nil
}

func (s *ProjectService) GetVersions(ctx context.Context, projectKeyOrID string) ([]ProjectVersion, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/versions", url.PathEscape(projectKeyOrID))
	var versions []ProjectVersion
	if err := s.client.Get(ctx, path, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *ProjectService) GetAvatars(ctx context.Context, projectKeyOrID string) (map[string]interface{}, error) {
	if err := s.client.CheckEndpoint("projects"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/project/%s/avatars", url.PathEscape(projectKeyOrID))
	var result map[string]interface{}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result, nil
}
