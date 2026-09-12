package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type NotifyService struct {
	client *client.Client
}

func NewNotifyService(c *client.Client) *NotifyService {
	return &NotifyService{client: c}
}

type NotifyPayload struct {
	TextBody        string        `json:"textBody"`
	HTMLBody        string        `json:"htmlBody,omitempty"`
	Subject         string        `json:"subject,omitempty"`
	To              *NotifyTo     `json:"to,omitempty"`
	Restrict        *NotifyRestrict `json:"restrict,omitempty"`
}

type NotifyTo struct {
	Reporter       bool     `json:"reporter,omitempty"`
	Assignee       bool     `json:"assignee,omitempty"`
	Watchers       bool     `json:"watchers,omitempty"`
	Voters         bool     `json:"voters,omitempty"`
	Users          []*User  `json:"users,omitempty"`
	Groups         []*Group `json:"groups,omitempty"`
}

type NotifyRestrict struct {
	Groups         []*Group `json:"groups,omitempty"`
	Permissions    []string `json:"permissions,omitempty"`
}

func (s *NotifyService) Send(ctx context.Context, issueKeyOrID string, payload *NotifyPayload) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/notify", url.PathEscape(issueKeyOrID))
	return s.client.Post(ctx, path, payload, nil)
}
