package api

import (
	"context"

	"github.com/darkkoax/jiracli/internal/client"
)

type FieldService struct {
	client *client.Client
}

func NewFieldService(c *client.Client) *FieldService {
	return &FieldService{client: c}
}

type Field struct {
	ID          string      `json:"id"`
	Key         string      `json:"key,omitempty"`
	Name        string      `json:"name"`
	Custom      bool        `json:"custom"`
	Orderable   bool        `json:"orderable"`
	Navigable   bool        `json:"navigable"`
	Searchable  bool        `json:"searchable"`
	ClauseNames []string    `json:"clauseNames,omitempty"`
	Schema      *FieldSchema `json:"schema,omitempty"`
}

type FieldSchema struct {
	Type     string `json:"type"`
	Items    string `json:"items,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
	System   string `json:"system,omitempty"`
}

type CustomField struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	SearcherKey string `json:"searcherKey,omitempty"`
}

func (s *FieldService) GetAll(ctx context.Context) ([]Field, error) {
	if err := s.client.CheckEndpoint("fields"); err != nil {
		return nil, err
	}
	var fields []Field
	if err := s.client.Get(ctx, "/rest/api/2/field", &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func (s *FieldService) CreateCustom(ctx context.Context, field *CustomField) (*CustomField, error) {
	if err := s.client.CheckEndpoint("fields"); err != nil {
		return nil, err
	}
	var result CustomField
	if err := s.client.Post(ctx, "/rest/api/2/field", field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
