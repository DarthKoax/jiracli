package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/darthkoax/jiracli/internal/config"
)

func testConfig(baseURL string) *config.Config {
	return &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  baseURL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods: config.MethodsConfig{
			AllowGet:    true,
			AllowPost:   true,
			AllowPut:    true,
			AllowDelete: true,
		},
		AdminMode: false,
		Endpoints: config.DefaultEndpoints(),
	}
}

func TestNew(t *testing.T) {
	cfg := testConfig("https://jira.example.com")
	c, err := New(cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if c == nil {
		t.Fatal("New() returned nil client")
	}
}

func TestNew_CustomCACert(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "ca*.crt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	cfg := testConfig("https://jira.example.com")
	cfg.Jira.CustomCACert = tmpfile.Name()

	_, err = New(cfg)
	if err == nil {
		t.Error("New() expected error for empty CA cert file, got nil")
	}
}

func TestNew_InvalidCACert(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "ca*.crt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte("invalid cert")); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	cfg := testConfig("https://jira.example.com")
	cfg.Jira.CustomCACert = tmpfile.Name()

	_, err = New(cfg)
	if err == nil {
		t.Error("New() expected error for invalid CA cert, got nil")
	}
}

func TestClient_Do_MethodGating(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	tests := []struct {
		name    string
		method  string
		allowed bool
	}{
		{"GET allowed", http.MethodGet, true},
		{"POST allowed", http.MethodPost, true},
		{"PUT allowed", http.MethodPut, true},
		{"DELETE allowed", http.MethodDelete, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ReadOnlyMode && tt.method != http.MethodGet {
				t.Skip("skipping write method test in readonly mode")
			}
			cfg := testConfig(server.URL)
			switch tt.method {
			case http.MethodGet:
				cfg.Methods.AllowGet = tt.allowed
			case http.MethodPost:
				cfg.Methods.AllowPost = tt.allowed
			case http.MethodPut:
				cfg.Methods.AllowPut = tt.allowed
			case http.MethodDelete:
				cfg.Methods.AllowDelete = tt.allowed
			}

			c, err := New(cfg)
			if err != nil {
				t.Fatal(err)
			}

			err = c.Do(context.Background(), tt.method, "/test", nil, nil)
			if tt.allowed && err != nil {
				t.Errorf("Do() error = %v, want nil", err)
			}
			if !tt.allowed && err == nil {
				t.Error("Do() expected error for disabled method, got nil")
			}
		})
	}
}

func TestClient_Do_MethodDisabled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	cfg.Methods.AllowGet = false

	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Error("Get() expected error for disabled GET, got nil")
	}
}

func TestClient_Do_AuthHeader(t *testing.T) {
	var receivedAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	cfg.Jira.APIToken = "my-secret-token"

	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Get(context.Background(), "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	if receivedAuth != "Bearer my-secret-token" {
		t.Errorf("Authorization = %v, want %v", receivedAuth, "Bearer my-secret-token")
	}
}

func TestClient_Do_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"errorMessages":["Issue not found"]}`))
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Error("Get() expected error for 404, got nil")
	}
}

func TestClient_Do_JSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key": "value"})
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]string
	err = c.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Fatal(err)
	}

	if result["key"] != "value" {
		t.Errorf("result[key] = %v, want %v", result["key"], "value")
	}
}

func TestClient_CheckEndpoint(t *testing.T) {
	cfg := testConfig("https://jira.example.com")
	cfg.Endpoints.Issues = true
	cfg.Endpoints.Boards = false

	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.CheckEndpoint("issues")
	if err != nil {
		t.Errorf("CheckEndpoint(issues) error = %v, want nil", err)
	}

	err = c.CheckEndpoint("boards")
	if err == nil {
		t.Error("CheckEndpoint(boards) expected error for disabled endpoint, got nil")
	}
}

func TestClient_CheckEndpoint_AllEndpoints(t *testing.T) {
	cfg := testConfig("https://jira.example.com")
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	endpoints := []string{
		"issues", "projects", "users", "groups", "search", "filters",
		"dashboards", "fields", "components", "versions", "priorities",
		"statuses", "resolutions", "issuetypes", "permissions", "myself",
		"screens", "reindex", "applicationproperties", "configuration",
		"avatars", "boards", "sprints",
	}

	for _, ep := range endpoints {
		err = c.CheckEndpoint(ep)
		if err != nil {
			t.Errorf("CheckEndpoint(%s) error = %v, want nil", ep, err)
		}
	}

	err = c.CheckEndpoint("unknown")
	if err == nil {
		t.Error("CheckEndpoint(unknown) expected error, got nil")
	}
}

func TestClient_CheckAdmin(t *testing.T) {
	cfg := testConfig("https://jira.example.com")
	cfg.AdminMode = false

	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.CheckAdmin("reindex", "trigger")
	if err == nil {
		t.Error("CheckAdmin(reindex, trigger) expected error when admin mode disabled, got nil")
	}

	err = c.CheckAdmin("issues", "create")
	if err != nil {
		t.Errorf("CheckAdmin(issues, create) error = %v, want nil", err)
	}

	cfg.AdminMode = true
	c, err = New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.CheckAdmin("reindex", "trigger")
	if err != nil {
		t.Errorf("CheckAdmin(reindex, trigger) error = %v, want nil when admin mode enabled", err)
	}
}

func TestClient_CheckAdmin_AdminOperations(t *testing.T) {
	cfg := testConfig("https://jira.example.com")
	cfg.AdminMode = false

	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	adminOps := []struct {
		endpoint  string
		operation string
	}{
		{"users", "create"},
		{"users", "delete"},
		{"groups", "create"},
		{"groups", "delete"},
		{"screens", "create"},
		{"screens", "delete"},
		{"screens", "update"},
		{"permissions", "create_scheme"},
		{"permissions", "delete_scheme"},
		{"permissions", "update_scheme"},
		{"reindex", "trigger"},
		{"applicationproperties", "set"},
		{"configuration", "get"},
	}

	for _, op := range adminOps {
		err = c.CheckAdmin(op.endpoint, op.operation)
		if err == nil {
			t.Errorf("CheckAdmin(%s, %s) expected error when admin mode disabled, got nil", op.endpoint, op.operation)
		}
	}
}

func TestClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method = %v, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": "ok"})
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]string
	err = c.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Fatal(err)
	}
	if result["result"] != "ok" {
		t.Errorf("result = %v, want ok", result)
	}
}

func TestClient_Post(t *testing.T) {
	if ReadOnlyMode {
		t.Skip("skipping POST test in readonly mode")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %v, want POST", r.Method)
		}
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["key"] != "value" {
			t.Errorf("body[key] = %v, want value", body["key"])
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Post(context.Background(), "/test", map[string]string{"key": "value"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Put(t *testing.T) {
	if ReadOnlyMode {
		t.Skip("skipping PUT test in readonly mode")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method = %v, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Put(context.Background(), "/test", map[string]string{"key": "value"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Delete(t *testing.T) {
	if ReadOnlyMode {
		t.Skip("skipping DELETE test in readonly mode")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method = %v, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Delete(context.Background(), "/test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew_NoCustomCA(t *testing.T) {
	cfg := testConfig("https://jira.example.com")

	c, err := New(cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if c == nil {
		t.Fatal("New() returned nil")
	}
}

func TestReadOnlyMode_Default(t *testing.T) {
	if ReadOnlyMode {
		t.Skip("skipping default-build test: binary compiled with -tags readonly")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := testConfig(server.URL)
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if err := c.Post(context.Background(), "/test", nil, nil); err != nil {
		t.Errorf("Post() should succeed in default build, got: %v", err)
	}
	if err := c.Put(context.Background(), "/test", nil, nil); err != nil {
		t.Errorf("Put() should succeed in default build, got: %v", err)
	}
	if err := c.Delete(context.Background(), "/test"); err != nil {
		t.Errorf("Delete() should succeed in default build, got: %v", err)
	}
}
