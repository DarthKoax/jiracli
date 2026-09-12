package api

import (
	"context"
	"fmt"

	"github.com/darkkoax/jiracli/internal/client"
)

type AvatarService struct {
	client *client.Client
}

func NewAvatarService(c *client.Client) *AvatarService {
	return &AvatarService{client: c}
}

type Avatar struct {
	ID          interface{}          `json:"id"`
	Owner       string               `json:"owner,omitempty"`
	IsSystem    bool                 `json:"isSystemAvatar"`
	IsSelected  bool                 `json:"isSelected"`
	Filename    string               `json:"filename,omitempty"`
	ContentType string               `json:"contentType,omitempty"`
	Portable    bool                 `json:"isDeletable,omitempty"`
	Selectable  bool                 `json:"isSelectable,omitempty"`
	URLs        map[string]interface{} `json:"urls,omitempty"`
}

type AvatarList struct {
	System []Avatar `json:"system"`
	Custom []Avatar `json:"custom"`
}

func (s *AvatarService) GetSystemAvatars(ctx context.Context, entityType string) (*AvatarList, error) {
	if err := s.client.CheckEndpoint("avatars"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/avatar/%s/system", entityType)
	var result AvatarList
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AvatarService) GetCustomAvatars(ctx context.Context, entityType, entityID string) (*AvatarList, error) {
	if err := s.client.CheckEndpoint("avatars"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/avatar/%s/%s", entityType, entityID)
	var result AvatarList
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AvatarService) DeleteCustomAvatar(ctx context.Context, entityType, avatarID string) error {
	if err := s.client.CheckEndpoint("avatars"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/avatar/%s/%s", entityType, avatarID)
	return s.client.Delete(ctx, path)
}

func (s *AvatarService) UpdateCustomAvatar(ctx context.Context, entityType, avatarID string, avatar *Avatar) error {
	if err := s.client.CheckEndpoint("avatars"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/avatar/%s/%s", entityType, avatarID)
	return s.client.Put(ctx, path, avatar, nil)
}
