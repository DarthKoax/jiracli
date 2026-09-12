package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type GroupService struct {
	client *client.Client
}

func NewGroupService(c *client.Client) *GroupService {
	return &GroupService{client: c}
}

type Group struct {
	Name string `json:"name"`
	Self string `json:"self,omitempty"`
}

type GroupBasic struct {
	Name string `json:"name"`
	Self string `json:"self,omitempty"`
}

type GroupMembers struct {
	Size       int    `json:"size"`
	MaxResults int    `json:"max-results"`
	StartIndex int    `json:"start-index"`
	TotalSize  int    `json:"total-size"`
	Items      []User `json:"items"`
}

func (s *GroupService) Get(ctx context.Context, groupName string) (*Group, error) {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/group?groupname=%s", url.QueryEscape(groupName))
	var group Group
	if err := s.client.Get(ctx, path, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *GroupService) Create(ctx context.Context, groupName string) (*Group, error) {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return nil, err
	}
	if err := s.client.CheckAdmin("groups", "create"); err != nil {
		return nil, err
	}
	body := map[string]string{"name": groupName}
	var result Group
	if err := s.client.Post(ctx, "/rest/api/2/group", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GroupService) Delete(ctx context.Context, groupName string) error {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("groups", "delete"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/group?groupname=%s", url.QueryEscape(groupName))
	return s.client.Delete(ctx, path)
}

func (s *GroupService) GetMembers(ctx context.Context, groupName string, startAt, maxResults int) (*GroupMembers, error) {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/group/member?groupname=%s&startAt=%d&maxResults=%d",
		url.QueryEscape(groupName), startAt, maxResults)
	var members GroupMembers
	if err := s.client.Get(ctx, path, &members); err != nil {
		return nil, err
	}
	return &members, nil
}

func (s *GroupService) AddUser(ctx context.Context, groupName, username string) error {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/group/user?groupname=%s", url.QueryEscape(groupName))
	body := map[string]string{"name": username}
	return s.client.Post(ctx, path, body, nil)
}

func (s *GroupService) RemoveUser(ctx context.Context, groupName, username string) error {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/group/user?groupname=%s&username=%s",
		url.QueryEscape(groupName), url.QueryEscape(username))
	return s.client.Delete(ctx, path)
}

func (s *GroupService) Search(ctx context.Context, query string, startAt, maxResults int) ([]Group, error) {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/groups/picker?query=%s&maxResults=%d",
		url.QueryEscape(query), maxResults)
	var result struct {
		Groups []Group `json:"groups"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Groups, nil
}

func (s *GroupService) BulkGet(ctx context.Context, groupnames []string) ([]Group, error) {
	if err := s.client.CheckEndpoint("groups"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/groups?groupnames=" + url.QueryEscape(joinStrings(groupnames))
	var groups []Group
	if err := s.client.Get(ctx, path, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}
