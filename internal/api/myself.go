package api

import (
	"context"

	"github.com/darkkoax/jiracli/internal/client"
)

type MyselfService struct {
	client *client.Client
}

func NewMyselfService(c *client.Client) *MyselfService {
	return &MyselfService{client: c}
}

type MyselfInfo struct {
	Self         string     `json:"self"`
	Key          string     `json:"key,omitempty"`
	Name         string     `json:"name,omitempty"`
	AccountID    string     `json:"accountId,omitempty"`
	EmailAddress string     `json:"emailAddress,omitempty"`
	DisplayName  string     `json:"displayName,omitempty"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone,omitempty"`
	Locale       string     `json:"locale,omitempty"`
	Groups       *UserGroups `json:"groups,omitempty"`
	ApplicationRoles *ApplicationRoles `json:"applicationRoles,omitempty"`
	Expand       string     `json:"expand,omitempty"`
	AvatarURLs   *AvatarURLs `json:"avatarUrls,omitempty"`
}

type LocaleInfo struct {
	Locale string `json:"locale"`
}

func (s *MyselfService) Get(ctx context.Context, expand []string) (*MyselfInfo, error) {
	if err := s.client.CheckEndpoint("myself"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/myself"
	if len(expand) > 0 {
		path += "?expand=" + joinStrings(expand)
	}
	var info MyselfInfo
	if err := s.client.Get(ctx, path, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *MyselfService) GetLocale(ctx context.Context) (*LocaleInfo, error) {
	if err := s.client.CheckEndpoint("myself"); err != nil {
		return nil, err
	}
	var info LocaleInfo
	if err := s.client.Get(ctx, "/rest/api/2/myself/locale", &info); err != nil {
		return nil, err
	}
	return &info, nil
}
