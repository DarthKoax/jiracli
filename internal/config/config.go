package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Jira      JiraConfig      `toml:"jira"`
	Methods   MethodsConfig   `toml:"methods"`
	AdminMode bool            `toml:"admin_mode"`
	Endpoints EndpointsConfig `toml:"endpoints"`
}

type JiraConfig struct {
	BaseURL      string `toml:"base_url"`
	APIToken     string `toml:"api_token"`
	CustomCACert string `toml:"custom_ca_cert"`
	Timeout      int    `toml:"timeout"`
}

type MethodsConfig struct {
	AllowGet    bool `toml:"allow_get"`
	AllowPost   bool `toml:"allow_post"`
	AllowPut    bool `toml:"allow_put"`
	AllowDelete bool `toml:"allow_delete"`
}

type EndpointsConfig struct {
	Issues                bool `toml:"issues"`
	Projects              bool `toml:"projects"`
	Users                 bool `toml:"users"`
	Groups                bool `toml:"groups"`
	Search                bool `toml:"search"`
	Filters               bool `toml:"filters"`
	Dashboards            bool `toml:"dashboards"`
	Fields                bool `toml:"fields"`
	Components            bool `toml:"components"`
	Versions              bool `toml:"versions"`
	Priorities            bool `toml:"priorities"`
	Statuses              bool `toml:"statuses"`
	Resolutions           bool `toml:"resolutions"`
	IssueTypes            bool `toml:"issuetypes"`
	Permissions           bool `toml:"permissions"`
	Myself                bool `toml:"myself"`
	Screens               bool `toml:"screens"`
	Reindex               bool `toml:"reindex"`
	ApplicationProperties bool `toml:"applicationproperties"`
	Configuration         bool `toml:"configuration"`
	Avatars               bool `toml:"avatars"`
	Boards                bool `toml:"boards"`
	Sprints               bool `toml:"sprints"`
	Attachments           bool `toml:"attachments"`
	Comments              bool `toml:"comments"`
	IssueLinks            bool `toml:"issuelinks"`
	HealthCheck           bool `toml:"healthcheck"`
	Webhooks              bool `toml:"webhooks"`
}

var allEndpointKeys = []string{
	"issues", "projects", "users", "groups", "search", "filters",
	"dashboards", "fields", "components", "versions", "priorities",
	"statuses", "resolutions", "issuetypes", "permissions", "myself",
	"screens", "reindex", "applicationproperties", "configuration",
	"avatars", "boards", "sprints", "attachments", "comments",
	"issuelinks", "healthcheck", "webhooks",
}

func DefaultEndpoints() EndpointsConfig {
	return EndpointsConfig{
		Issues:                true,
		Projects:              true,
		Users:                 true,
		Groups:                true,
		Search:                true,
		Filters:               true,
		Dashboards:            true,
		Fields:                true,
		Components:            true,
		Versions:              true,
		Priorities:            true,
		Statuses:              true,
		Resolutions:           true,
		IssueTypes:            true,
		Permissions:           true,
		Myself:                true,
		Screens:               true,
		Reindex:               true,
		ApplicationProperties: true,
		Configuration:         true,
		Avatars:               true,
		Boards:                true,
		Sprints:               true,
		Attachments:           true,
		Comments:              true,
		IssueLinks:            true,
		HealthCheck:           true,
		Webhooks:              true,
	}
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var raw map[string]interface{}
	if _, err := toml.Decode(string(data), &raw); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	if cfg.Jira.BaseURL == "" {
		return nil, fmt.Errorf("jira.base_url is required")
	}
	if cfg.Jira.APIToken == "" {
		return nil, fmt.Errorf("jira.api_token is required")
	}
	if cfg.Jira.Timeout == 0 {
		cfg.Jira.Timeout = 30
	}

	endpointsSection, hasEndpoints := raw["endpoints"].(map[string]interface{})
	if !hasEndpoints {
		cfg.Endpoints = DefaultEndpoints()
	} else {
		cfg.Endpoints = DefaultEndpoints()
		for _, key := range allEndpointKeys {
			if val, ok := endpointsSection[key]; ok {
				if b, ok := val.(bool); ok {
					setEndpoint(&cfg.Endpoints, key, b)
				}
			}
		}
	}

	applyEnvOverrides(&cfg)

	return &cfg, nil
}

func LoadFromEnv() (*Config, error) {
	baseURL := os.Getenv("JIRA_BASE_URL")
	apiToken := os.Getenv("JIRA_API_TOKEN")

	if baseURL == "" {
		return nil, fmt.Errorf("JIRA_BASE_URL environment variable is required")
	}
	if apiToken == "" {
		return nil, fmt.Errorf("JIRA_API_TOKEN environment variable is required")
	}

	cfg := &Config{
		Jira: JiraConfig{
			BaseURL:      baseURL,
			APIToken:     apiToken,
			CustomCACert: os.Getenv("JIRA_CUSTOM_CA_CERT"),
			Timeout:      30,
		},
		Methods: MethodsConfig{
			AllowGet:    true,
			AllowPost:   true,
			AllowPut:    true,
			AllowDelete: true,
		},
		AdminMode: false,
		Endpoints: DefaultEndpoints(),
	}

	if timeout := os.Getenv("JIRA_TIMEOUT"); timeout != "" {
		if t, err := strconv.Atoi(timeout); err == nil {
			cfg.Jira.Timeout = t
		}
	}

	applyEnvOverrides(cfg)

	return cfg, nil
}

func setEndpoint(e *EndpointsConfig, key string, val bool) {
	switch key {
	case "issues":
		e.Issues = val
	case "projects":
		e.Projects = val
	case "users":
		e.Users = val
	case "groups":
		e.Groups = val
	case "search":
		e.Search = val
	case "filters":
		e.Filters = val
	case "dashboards":
		e.Dashboards = val
	case "fields":
		e.Fields = val
	case "components":
		e.Components = val
	case "versions":
		e.Versions = val
	case "priorities":
		e.Priorities = val
	case "statuses":
		e.Statuses = val
	case "resolutions":
		e.Resolutions = val
	case "issuetypes":
		e.IssueTypes = val
	case "permissions":
		e.Permissions = val
	case "myself":
		e.Myself = val
	case "screens":
		e.Screens = val
	case "reindex":
		e.Reindex = val
	case "applicationproperties":
		e.ApplicationProperties = val
	case "configuration":
		e.Configuration = val
	case "avatars":
		e.Avatars = val
	case "boards":
		e.Boards = val
	case "sprints":
		e.Sprints = val
	case "attachments":
		e.Attachments = val
	case "comments":
		e.Comments = val
	case "issuelinks":
		e.IssueLinks = val
	case "healthcheck":
		e.HealthCheck = val
	case "webhooks":
		e.Webhooks = val
	}
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("JIRA_BASE_URL"); v != "" {
		cfg.Jira.BaseURL = v
	}
	if v := os.Getenv("JIRA_API_TOKEN"); v != "" {
		cfg.Jira.APIToken = v
	}
	if v := os.Getenv("JIRA_CUSTOM_CA_CERT"); v != "" {
		cfg.Jira.CustomCACert = v
	}
	if v := os.Getenv("JIRA_TIMEOUT"); v != "" {
		if t, err := strconv.Atoi(v); err == nil {
			cfg.Jira.Timeout = t
		}
	}
	if v := os.Getenv("JIRA_ADMIN_MODE"); v != "" {
		cfg.AdminMode = v == "true" || v == "1"
	}
	if v := os.Getenv("JIRA_ALLOW_GET"); v != "" {
		cfg.Methods.AllowGet = v == "true" || v == "1"
	}
	if v := os.Getenv("JIRA_ALLOW_POST"); v != "" {
		cfg.Methods.AllowPost = v == "true" || v == "1"
	}
	if v := os.Getenv("JIRA_ALLOW_PUT"); v != "" {
		cfg.Methods.AllowPut = v == "true" || v == "1"
	}
	if v := os.Getenv("JIRA_ALLOW_DELETE"); v != "" {
		cfg.Methods.AllowDelete = v == "true" || v == "1"
	}

	for _, key := range allEndpointKeys {
		envKey := "JIRA_ENDPOINT_" + strings.ToUpper(key)
		if v := os.Getenv(envKey); v != "" {
			setEndpoint(&cfg.Endpoints, key, v == "true" || v == "1")
		}
	}
}

func DefaultConfigDir() string {
	if v := os.Getenv("JIRA_CONFIG_DIR"); v != "" {
		return v
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ".config/jiracli"
	}
	return filepath.Join(home, ".config", "jiracli")
}

func DefaultConfigPath() string {
	return filepath.Join(DefaultConfigDir(), "config.toml")
}

func DefaultConfigContent() string {
	return `[jira]
base_url = "https://jira.example.com"
api_token = "YOUR_API_TOKEN_HERE"
custom_ca_cert = ""
timeout = 30

admin_mode = false

[methods]
allow_get = true
allow_post = true
allow_put = true
allow_delete = true

[endpoints]
issues = true
projects = true
users = true
groups = true
search = true
filters = true
dashboards = true
fields = true
components = true
versions = true
priorities = true
statuses = true
resolutions = true
issuetypes = true
permissions = true
myself = true
screens = true
reindex = true
applicationproperties = true
configuration = true
avatars = true
boards = true
sprints = true
`
}

func Init(dir string) (string, error) {
	if dir == "" {
		dir = DefaultConfigDir()
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("creating config directory: %w", err)
	}

	configPath := filepath.Join(dir, "config.toml")

	if _, err := os.Stat(configPath); err == nil {
		return configPath, fmt.Errorf("config file already exists at %s", configPath)
	}

	if err := os.WriteFile(configPath, []byte(DefaultConfigContent()), 0600); err != nil {
		return "", fmt.Errorf("writing config file: %w", err)
	}

	return configPath, nil
}
