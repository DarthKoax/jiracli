package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type UserService struct {
	client *client.Client
}

func NewUserService(c *client.Client) *UserService {
	return &UserService{client: c}
}

type User struct {
	Self             string            `json:"self,omitempty"`
	Key              string            `json:"key,omitempty"`
	Name             string            `json:"name,omitempty"`
	AccountID        string            `json:"accountId,omitempty"`
	EmailAddress     string            `json:"emailAddress,omitempty"`
	DisplayName      string            `json:"displayName,omitempty"`
	Active           bool              `json:"active"`
	TimeZone         string            `json:"timeZone,omitempty"`
	Locale           string            `json:"locale,omitempty"`
	Groups           *UserGroups       `json:"groups,omitempty"`
	ApplicationRoles *ApplicationRoles `json:"applicationRoles,omitempty"`
	AvatarURLs       *AvatarURLs       `json:"avatarUrls,omitempty"`
}

type UserGroups struct {
	Size  int          `json:"size"`
	Items []GroupBasic `json:"items"`
}

type ApplicationRoles struct {
	Size  int               `json:"size"`
	Items []ApplicationRole `json:"items"`
}

type ApplicationRole struct {
	Key               string   `json:"key"`
	Name              string   `json:"name"`
	Groups            []string `json:"groups"`
	SelectedByDefault bool     `json:"selectedByDefault"`
	Defined           bool     `json:"defined"`
}

func (s *UserService) Get(ctx context.Context, username string, expand []string) (*User, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/user?username=%s", url.QueryEscape(username))
	if len(expand) > 0 {
		path += "&expand=" + joinStrings(expand)
	}
	var user User
	if err := s.client.Get(ctx, path, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetByAccountID(ctx context.Context, accountID string) (*User, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/user?accountId=%s", url.QueryEscape(accountID))
	var user User
	if err := s.client.Get(ctx, path, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Search(ctx context.Context, query string, startAt, maxResults int) ([]User, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/user/search?username=%s&startAt=%d&maxResults=%d",
		url.QueryEscape(query), startAt, maxResults)
	var users []User
	if err := s.client.Get(ctx, path, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) FindUsersForPicker(ctx context.Context, query string, maxResults int) ([]User, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/user/picker?query=%s&maxResults=%d",
		url.QueryEscape(query), maxResults)
	var result struct {
		Users []struct {
			User User `json:"user"`
		} `json:"users"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	users := make([]User, len(result.Users))
	for i, u := range result.Users {
		users[i] = u.User
	}
	return users, nil
}

func (s *UserService) Create(ctx context.Context, user *User) (*User, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("users", "create"); err != nil {
		return nil, err
	}
	var result User
	if err := s.client.Post(ctx, "/rest/api/2/user", user, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserService) Delete(ctx context.Context, username string) error {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("users", "delete"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/user?username=%s", url.QueryEscape(username))
	return s.client.Delete(ctx, path)
}

func (s *UserService) AssignToProject(ctx context.Context, projectKey, username string) error {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/user/assignable/search?project=%s&username=%s",
		url.QueryEscape(projectKey), url.QueryEscape(username))
	return s.client.Get(ctx, path, nil)
}

func (s *UserService) GetColumns(ctx context.Context) ([]map[string]interface{}, error) {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return nil, err
	}
	var columns []map[string]interface{}
	if err := s.client.Get(ctx, "/rest/api/2/user/columns", &columns); err != nil {
		return nil, err
	}
	return columns, nil
}

func (s *UserService) SetColumns(ctx context.Context, columns []map[string]interface{}) error {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return err
	}
	return s.client.Put(ctx, "/rest/api/2/user/columns", columns, nil)
}

func (s *UserService) ResetColumns(ctx context.Context) error {
	if err := s.client.CheckEndpoint("users"); err != nil {
		return err
	}
	return s.client.Delete(ctx, "/rest/api/2/user/columns")
}
