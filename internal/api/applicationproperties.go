package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type ApplicationPropertiesService struct {
	client *client.Client
}

func NewApplicationPropertiesService(c *client.Client) *ApplicationPropertiesService {
	return &ApplicationPropertiesService{client: c}
}

type ApplicationProperty struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Value string `json:"value"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

func (s *ApplicationPropertiesService) GetAll(ctx context.Context, key, permissionLevel string) ([]ApplicationProperty, error) {
	if err := s.client.CheckEndpoint("applicationproperties"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/application-properties?"
	params := url.Values{}
	if key != "" {
		params.Set("key", key)
	}
	if permissionLevel != "" {
		params.Set("permissionLevel", permissionLevel)
	}
	path += params.Encode()

	var props []ApplicationProperty
	if err := s.client.Get(ctx, path, &props); err != nil {
		return nil, err
	}
	return props, nil
}

func (s *ApplicationPropertiesService) Set(ctx context.Context, propertyID string, value string) error {
	if err := s.client.CheckEndpoint("applicationproperties"); err != nil {
		return err
	}
	if err := s.client.CheckAdmin("applicationproperties", "set"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/application-properties/%s", url.PathEscape(propertyID))
	body := map[string]string{"id": propertyID, "value": value}
	return s.client.Put(ctx, path, body, nil)
}
