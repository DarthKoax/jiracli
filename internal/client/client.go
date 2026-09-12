package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/darthkoax/jiracli/internal/config"
)

var adminEndpoints = map[string]bool{
	"reindex":               true,
	"applicationproperties": true,
	"configuration":         true,
}

var adminOperations = map[string]bool{
	"users.create":   true,
	"users.delete":   true,
	"groups.create":  true,
	"groups.delete":  true,
	"screens.create": true,
	"screens.delete": true,
	"screens.update": true,
	"permissions.create_scheme": true,
	"permissions.delete_scheme": true,
	"permissions.update_scheme": true,
}

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiToken   string
	methods    config.MethodsConfig
	adminMode  bool
	endpoints  config.EndpointsConfig
}

func New(cfg *config.Config) (*Client, error) {
	tlsConfig := &tls.Config{}

	if cfg.Jira.CustomCACert != "" {
		caCert, err := os.ReadFile(cfg.Jira.CustomCACert)
		if err != nil {
			return nil, fmt.Errorf("reading custom CA cert: %w", err)
		}
		pool, err := x509.SystemCertPool()
		if err != nil {
			pool = x509.NewCertPool()
		}
		if !pool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to append custom CA cert")
		}
		tlsConfig.RootCAs = pool
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(cfg.Jira.Timeout) * time.Second,
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    strings.TrimRight(cfg.Jira.BaseURL, "/"),
		apiToken:   cfg.Jira.APIToken,
		methods:    cfg.Methods,
		adminMode:  cfg.AdminMode,
		endpoints:  cfg.Endpoints,
	}, nil
}

func (c *Client) CheckEndpoint(name string) error {
	switch name {
	case "issues":
		if !c.endpoints.Issues {
			return fmt.Errorf("issues endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the issues endpoint in config.toml or contact an administrator")
		}
	case "projects":
		if !c.endpoints.Projects {
			return fmt.Errorf("projects endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the projects endpoint in config.toml or contact an administrator")
		}
	case "users":
		if !c.endpoints.Users {
			return fmt.Errorf("users endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the users endpoint in config.toml or contact an administrator")
		}
	case "groups":
		if !c.endpoints.Groups {
			return fmt.Errorf("groups endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the groups endpoint in config.toml or contact an administrator")
		}
	case "search":
		if !c.endpoints.Search {
			return fmt.Errorf("search endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the search endpoint in config.toml or contact an administrator")
		}
	case "filters":
		if !c.endpoints.Filters {
			return fmt.Errorf("filters endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the filters endpoint in config.toml or contact an administrator")
		}
	case "dashboards":
		if !c.endpoints.Dashboards {
			return fmt.Errorf("dashboards endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the dashboards endpoint in config.toml or contact an administrator")
		}
	case "fields":
		if !c.endpoints.Fields {
			return fmt.Errorf("fields endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the fields endpoint in config.toml or contact an administrator")
		}
	case "components":
		if !c.endpoints.Components {
			return fmt.Errorf("components endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the components endpoint in config.toml or contact an administrator")
		}
	case "versions":
		if !c.endpoints.Versions {
			return fmt.Errorf("versions endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the versions endpoint in config.toml or contact an administrator")
		}
	case "priorities":
		if !c.endpoints.Priorities {
			return fmt.Errorf("priorities endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the priorities endpoint in config.toml or contact an administrator")
		}
	case "statuses":
		if !c.endpoints.Statuses {
			return fmt.Errorf("statuses endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the statuses endpoint in config.toml or contact an administrator")
		}
	case "resolutions":
		if !c.endpoints.Resolutions {
			return fmt.Errorf("resolutions endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the resolutions endpoint in config.toml or contact an administrator")
		}
	case "issuetypes":
		if !c.endpoints.IssueTypes {
			return fmt.Errorf("issuetypes endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the issuetypes endpoint in config.toml or contact an administrator")
		}
	case "permissions":
		if !c.endpoints.Permissions {
			return fmt.Errorf("permissions endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the permissions endpoint in config.toml or contact an administrator")
		}
	case "myself":
		if !c.endpoints.Myself {
			return fmt.Errorf("myself endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the myself endpoint in config.toml or contact an administrator")
		}
	case "screens":
		if !c.endpoints.Screens {
			return fmt.Errorf("screens endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the screens endpoint in config.toml or contact an administrator")
		}
	case "reindex":
		if !c.endpoints.Reindex {
			return fmt.Errorf("reindex endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the reindex endpoint in config.toml or contact an administrator")
		}
	case "applicationproperties":
		if !c.endpoints.ApplicationProperties {
			return fmt.Errorf("applicationproperties endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the applicationproperties endpoint in config.toml or contact an administrator")
		}
	case "configuration":
		if !c.endpoints.Configuration {
			return fmt.Errorf("configuration endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the configuration endpoint in config.toml or contact an administrator")
		}
	case "avatars":
		if !c.endpoints.Avatars {
			return fmt.Errorf("avatars endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the avatars endpoint in config.toml or contact an administrator")
		}
	case "boards":
		if !c.endpoints.Boards {
			return fmt.Errorf("boards endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the boards endpoint in config.toml or contact an administrator")
		}
	case "sprints":
		if !c.endpoints.Sprints {
			return fmt.Errorf("sprints endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the sprints endpoint in config.toml or contact an administrator")
		}
	case "attachments":
		if !c.endpoints.Attachments {
			return fmt.Errorf("attachments endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the attachments endpoint in config.toml or contact an administrator")
		}
	case "comments":
		if !c.endpoints.Comments {
			return fmt.Errorf("comments endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the comments endpoint in config.toml or contact an administrator")
		}
	case "issuelinks":
		if !c.endpoints.IssueLinks {
			return fmt.Errorf("issuelinks endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the issuelinks endpoint in config.toml or contact an administrator")
		}
	case "healthcheck":
		if !c.endpoints.HealthCheck {
			return fmt.Errorf("healthcheck endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the healthcheck endpoint in config.toml or contact an administrator")
		}
	case "webhooks":
		if !c.endpoints.Webhooks {
			return fmt.Errorf("webhooks endpoint is disabled in configuration. This operation cannot be performed. Please ask the user to enable the webhooks endpoint in config.toml or contact an administrator")
		}
	default:
		return fmt.Errorf("unknown endpoint: %s", name)
	}
	return nil
}

func (c *Client) CheckAdmin(endpoint, operation string) error {
	key := endpoint + "." + operation
	if adminEndpoints[endpoint] || adminOperations[key] {
		if !c.adminMode {
			return fmt.Errorf("%s.%s requires admin mode to be enabled. This operation cannot be performed. You should not attempt to bypass this restriction. Please ask the user to enable admin_mode in config.toml or contact an administrator", endpoint, operation)
		}
	}
	return nil
}

func (c *Client) Do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		if !c.methods.AllowGet {
			return fmt.Errorf("GET method is disabled in configuration. This operation cannot be performed. You should not attempt to bypass this restriction. Please ask the user to enable GET requests in config.toml or contact an administrator")
		}
	case http.MethodPost:
		if !c.methods.AllowPost {
			return fmt.Errorf("POST method is disabled in configuration. This operation cannot be performed. You should not attempt to bypass this restriction. Please ask the user to enable POST requests in config.toml or contact an administrator")
		}
	case http.MethodPut:
		if !c.methods.AllowPut {
			return fmt.Errorf("PUT method is disabled in configuration. This operation cannot be performed. You should not attempt to bypass this restriction. Please ask the user to enable PUT requests in config.toml or contact an administrator")
		}
	case http.MethodDelete:
		if !c.methods.AllowDelete {
			return fmt.Errorf("DELETE method is disabled in configuration. This operation cannot be performed. You should not attempt to bypass this restriction. Please ask the user to enable DELETE requests in config.toml or contact an administrator")
		}
	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("unmarshaling response: %w", err)
		}
	}

	return nil
}

func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	return c.Do(ctx, http.MethodGet, path, nil, result)
}

func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodPost, path, body, result)
}

func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodPut, path, body, result)
}

func (c *Client) Delete(ctx context.Context, path string) error {
	return c.Do(ctx, http.MethodDelete, path, nil, nil)
}
