package api

import (
	"context"
	"fmt"

	"github.com/darkkoax/jiracli/internal/client"
)

type WebhookService struct {
	client *client.Client
}

func NewWebhookService(c *client.Client) *WebhookService {
	return &WebhookService{client: c}
}

type Webhook struct {
	ID                 int      `json:"id"`
	Self               string   `json:"self"`
	Name               string   `json:"name"`
	URL                string   `json:"url"`
	Enabled            bool     `json:"enabled"`
	Events             []string `json:"events"`
	Filter             string   `json:"filter,omitempty"`
	ExcludeBody        bool     `json:"excludeBody,omitempty"`
	Secret             string   `json:"secret,omitempty"`
	LastUpdatedUser    *User    `json:"lastUpdatedUser,omitempty"`
	LastUpdatedDisplayName string `json:"lastUpdatedDisplayName,omitempty"`
	Created            string   `json:"created,omitempty"`
	Updated            string   `json:"updated,omitempty"`
}

type WebhookCreatePayload struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Events  []string `json:"events"`
	Filter  string   `json:"filter,omitempty"`
	Secret  string   `json:"secret,omitempty"`
}

type WebhookListResult struct {
	Webhooks []Webhook `json:"webhooks"`
}

func (s *WebhookService) GetAll(ctx context.Context) (*WebhookListResult, error) {
	if err := s.client.CheckEndpoint("webhooks"); err != nil {
		return nil, err
	}
	var result WebhookListResult
	if err := s.client.Get(ctx, "/rest/webhooks/1.0/webhook", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WebhookService) Get(ctx context.Context, webhookID int) (*Webhook, error) {
	if err := s.client.CheckEndpoint("webhooks"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/webhooks/1.0/webhook/%d", webhookID)
	var webhook Webhook
	if err := s.client.Get(ctx, path, &webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) Create(ctx context.Context, payload *WebhookCreatePayload) (*Webhook, error) {
	if err := s.client.CheckEndpoint("webhooks"); err != nil {
		return nil, err
	}
	var webhook Webhook
	if err := s.client.Post(ctx, "/rest/webhooks/1.0/webhook", payload, &webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) Update(ctx context.Context, webhookID int, payload *WebhookCreatePayload) (*Webhook, error) {
	if err := s.client.CheckEndpoint("webhooks"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/webhooks/1.0/webhook/%d", webhookID)
	var webhook Webhook
	if err := s.client.Put(ctx, path, payload, &webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) Delete(ctx context.Context, webhookID int) error {
	if err := s.client.CheckEndpoint("webhooks"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/webhooks/1.0/webhook/%d", webhookID)
	return s.client.Delete(ctx, path)
}
