package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/darkkoax/jiracli/internal/client"
	"github.com/darkkoax/jiracli/internal/config"
)

func testClient(t *testing.T, handler http.Handler) (*client.Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods: config.MethodsConfig{
			AllowGet:    true,
			AllowPost:   true,
			AllowPut:    true,
			AllowDelete: true,
		},
		AdminMode: true,
		Endpoints: config.DefaultEndpoints(),
	}
	c, err := client.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return c, server
}

func TestIssueService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/2/issue/PROJ-1" {
			t.Errorf("Path = %v, want /rest/api/2/issue/PROJ-1", r.URL.Path)
		}
		json.NewEncoder(w).Encode(Issue{Key: "PROJ-1", ID: "10001"})
	}))
	defer server.Close()

	svc := NewIssueService(c)
	issue, err := svc.Get(context.Background(), "PROJ-1", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if issue.Key != "PROJ-1" {
		t.Errorf("Key = %v, want PROJ-1", issue.Key)
	}
}

func TestIssueService_Create(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %v, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Issue{Key: "PROJ-2", ID: "10002"})
	}))
	defer server.Close()

	svc := NewIssueService(c)
	issue, err := svc.Create(context.Background(), &Issue{
		Fields: map[string]interface{}{"summary": "Test"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if issue.Key != "PROJ-2" {
		t.Errorf("Key = %v, want PROJ-2", issue.Key)
	}
}

func TestIssueService_Update(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method = %v, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	svc := NewIssueService(c)
	err := svc.Update(context.Background(), "PROJ-1", &Issue{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueService_Delete(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method = %v, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	svc := NewIssueService(c)
	err := svc.Delete(context.Background(), "PROJ-1", false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueService_GetTransitions(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"transitions": []Transition{
				{ID: "1", Name: "Done"},
			},
		})
	}))
	defer server.Close()

	svc := NewIssueService(c)
	transitions, err := svc.GetTransitions(context.Background(), "PROJ-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(transitions) != 1 {
		t.Errorf("len(transitions) = %v, want 1", len(transitions))
	}
}

func TestIssueService_DoTransition(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %v, want POST", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	svc := NewIssueService(c)
	err := svc.DoTransition(context.Background(), "PROJ-1", "1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueService_AddComment(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Comment{ID: "100", Body: "test comment"})
	}))
	defer server.Close()

	svc := NewIssueService(c)
	comment, err := svc.AddComment(context.Background(), "PROJ-1", &Comment{Body: "test comment"})
	if err != nil {
		t.Fatal(err)
	}
	if comment.Body != "test comment" {
		t.Errorf("Body = %v, want test comment", comment.Body)
	}
}

func TestIssueService_EndpointDisabled(t *testing.T) {
	_, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not reach server when endpoint is disabled")
	}))
	defer server.Close()

	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods: config.MethodsConfig{
			AllowGet:  true,
			AllowPost: true,
		},
		Endpoints: config.DefaultEndpoints(),
	}
	cfg.Endpoints.Issues = false

	c2, err := client.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	svc := NewIssueService(c2)
	_, err = svc.Get(context.Background(), "PROJ-1", nil, nil)
	if err == nil {
		t.Error("Expected error for disabled issues endpoint, got nil")
	}
}

func TestProjectService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]Project{
			{Key: "PROJ1", Name: "Project 1"},
			{Key: "PROJ2", Name: "Project 2"},
		})
	}))
	defer server.Close()

	svc := NewProjectService(c)
	projects, err := svc.GetAll(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 2 {
		t.Errorf("len(projects) = %v, want 2", len(projects))
	}
}

func TestProjectService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Project{Key: "PROJ", Name: "Test Project"})
	}))
	defer server.Close()

	svc := NewProjectService(c)
	project, err := svc.Get(context.Background(), "PROJ", nil)
	if err != nil {
		t.Fatal(err)
	}
	if project.Key != "PROJ" {
		t.Errorf("Key = %v, want PROJ", project.Key)
	}
}

func TestUserService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(User{Name: "testuser", DisplayName: "Test User"})
	}))
	defer server.Close()

	svc := NewUserService(c)
	user, err := svc.Get(context.Background(), "testuser", nil)
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "testuser" {
		t.Errorf("Name = %v, want testuser", user.Name)
	}
}

func TestUserService_Create_AdminRequired(t *testing.T) {
	_, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(User{Name: "newuser"})
	}))
	defer server.Close()

	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods: config.MethodsConfig{
			AllowGet:  true,
			AllowPost: true,
		},
		AdminMode: false,
		Endpoints: config.DefaultEndpoints(),
	}
	c2, _ := client.New(cfg)

	svc := NewUserService(c2)
	_, err := svc.Create(context.Background(), &User{Name: "newuser"})
	if err == nil {
		t.Error("Expected error for admin operation without admin mode, got nil")
	}
}

func TestGroupService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Group{Name: "testgroup"})
	}))
	defer server.Close()

	svc := NewGroupService(c)
	group, err := svc.Get(context.Background(), "testgroup")
	if err != nil {
		t.Fatal(err)
	}
	if group.Name != "testgroup" {
		t.Errorf("Name = %v, want testgroup", group.Name)
	}
}

func TestSearchService_Search(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(SearchResult{
			Total: 1,
			Issues: []Issue{
				{Key: "PROJ-1"},
			},
		})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	result, err := svc.Search(context.Background(), &SearchRequest{JQL: "project = PROJ"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("Total = %v, want 1", result.Total)
	}
}

func TestFilterService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Filter{ID: "10000", Name: "My Filter"})
	}))
	defer server.Close()

	svc := NewFilterService(c)
	filter, err := svc.Get(context.Background(), "10000", nil)
	if err != nil {
		t.Fatal(err)
	}
	if filter.Name != "My Filter" {
		t.Errorf("Name = %v, want My Filter", filter.Name)
	}
}

func TestDashboardService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Dashboard{ID: "10000", Name: "My Dashboard"})
	}))
	defer server.Close()

	svc := NewDashboardService(c)
	dashboard, err := svc.Get(context.Background(), "10000")
	if err != nil {
		t.Fatal(err)
	}
	if dashboard.Name != "My Dashboard" {
		t.Errorf("Name = %v, want My Dashboard", dashboard.Name)
	}
}

func TestBoardService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Board{ID: 1, Name: "My Board"})
	}))
	defer server.Close()

	svc := NewBoardService(c)
	board, err := svc.Get(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if board.Name != "My Board" {
		t.Errorf("Name = %v, want My Board", board.Name)
	}
}

func TestBoardService_GetIssues(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/agile/1.0/board/1/issue" {
			t.Errorf("Path = %v, want /rest/agile/1.0/board/1/issue", r.URL.Path)
		}
		json.NewEncoder(w).Encode(PaginatedIssues{
			Total: 2,
			Issues: []Issue{
				{Key: "PROJ-1"},
				{Key: "PROJ-2"},
			},
		})
	}))
	defer server.Close()

	svc := NewBoardService(c)
	result, err := svc.GetIssues(context.Background(), 1, 0, 50, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 2 {
		t.Errorf("Total = %v, want 2", result.Total)
	}
	if len(result.Issues) != 2 {
		t.Errorf("len(Issues) = %v, want 2", len(result.Issues))
	}
}

func TestBoardService_GetBacklog(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/agile/1.0/board/1/backlog" {
			t.Errorf("Path = %v, want /rest/agile/1.0/board/1/backlog", r.URL.Path)
		}
		json.NewEncoder(w).Encode(PaginatedIssues{
			Total: 3,
			Issues: []Issue{
				{Key: "PROJ-1"},
				{Key: "PROJ-2"},
				{Key: "PROJ-3"},
			},
		})
	}))
	defer server.Close()

	svc := NewBoardService(c)
	result, err := svc.GetBacklog(context.Background(), 1, 0, 50, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 3 {
		t.Errorf("Total = %v, want 3", result.Total)
	}
}

func TestBoardService_GetSprints(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/agile/1.0/board/1/sprint" {
			t.Errorf("Path = %v, want /rest/agile/1.0/board/1/sprint", r.URL.Path)
		}
		state := r.URL.Query().Get("state")
		if state != "active" {
			t.Errorf("state = %v, want active", state)
		}
		json.NewEncoder(w).Encode(PaginatedSprints{
			Total: 1,
			Values: []Sprint{
				{ID: 1, Name: "Sprint 1", State: "active"},
			},
		})
	}))
	defer server.Close()

	svc := NewBoardService(c)
	result, err := svc.GetSprints(context.Background(), 1, 0, 50, "active")
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("Total = %v, want 1", result.Total)
	}
}

func TestBoardService_EndpointDisabled(t *testing.T) {
	_, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not reach server when endpoint is disabled")
	}))
	defer server.Close()

	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods:   config.MethodsConfig{AllowGet: true},
		Endpoints: config.DefaultEndpoints(),
	}
	cfg.Endpoints.Boards = false

	c2, _ := client.New(cfg)
	svc := NewBoardService(c2)
	_, err := svc.Get(context.Background(), 1)
	if err == nil {
		t.Error("Expected error for disabled boards endpoint, got nil")
	}
}

func TestSprintService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Sprint{ID: 1, Name: "Sprint 1"})
	}))
	defer server.Close()

	svc := NewSprintService(c)
	sprint, err := svc.Get(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if sprint.Name != "Sprint 1" {
		t.Errorf("Name = %v, want Sprint 1", sprint.Name)
	}
}

func TestSprintService_GetIssues(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/agile/1.0/sprint/1/issue" {
			t.Errorf("Path = %v, want /rest/agile/1.0/sprint/1/issue", r.URL.Path)
		}
		json.NewEncoder(w).Encode(PaginatedIssues{
			Total: 2,
			Issues: []Issue{
				{Key: "PROJ-1"},
				{Key: "PROJ-2"},
			},
		})
	}))
	defer server.Close()

	svc := NewSprintService(c)
	result, err := svc.GetIssues(context.Background(), 1, 0, 50, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 2 {
		t.Errorf("Total = %v, want 2", result.Total)
	}
	if len(result.Issues) != 2 {
		t.Errorf("len(Issues) = %v, want 2", len(result.Issues))
	}
}

func TestSprintService_MoveIssues(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %v, want POST", r.Method)
		}
		if r.URL.Path != "/rest/agile/1.0/sprint/1/issue" {
			t.Errorf("Path = %v, want /rest/agile/1.0/sprint/1/issue", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	svc := NewSprintService(c)
	err := svc.MoveIssues(context.Background(), 1, []string{"PROJ-1", "PROJ-2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprintService_Create(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %v, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Sprint{ID: 2, Name: "Sprint 2", State: "future"})
	}))
	defer server.Close()

	svc := NewSprintService(c)
	sprint, err := svc.Create(context.Background(), 1, &Sprint{Name: "Sprint 2"})
	if err != nil {
		t.Fatal(err)
	}
	if sprint.Name != "Sprint 2" {
		t.Errorf("Name = %v, want Sprint 2", sprint.Name)
	}
}

func TestSprintService_Update(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method = %v, want PUT", r.Method)
		}
		json.NewEncoder(w).Encode(Sprint{ID: 1, Name: "Sprint 1 Updated"})
	}))
	defer server.Close()

	svc := NewSprintService(c)
	sprint, err := svc.Update(context.Background(), 1, &Sprint{Name: "Sprint 1 Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if sprint.Name != "Sprint 1 Updated" {
		t.Errorf("Name = %v, want Sprint 1 Updated", sprint.Name)
	}
}

func TestSprintService_Delete(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method = %v, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	svc := NewSprintService(c)
	err := svc.Delete(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprintService_Complete(t *testing.T) {
	requestCount := 0
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount == 1 {
			// First request: GET sprint details
			if r.Method != http.MethodGet {
				t.Errorf("First request method = %v, want GET", r.Method)
			}
			json.NewEncoder(w).Encode(Sprint{ID: 1, Name: "Sprint 1", State: "active", OriginBoardID: 1})
		} else {
			// Second request: PUT to complete
			if r.Method != http.MethodPut {
				t.Errorf("Second request method = %v, want PUT", r.Method)
			}
			if r.URL.Path != "/rest/agile/1.0/sprint/1" {
				t.Errorf("Path = %v, want /rest/agile/1.0/sprint/1", r.URL.Path)
			}
			json.NewEncoder(w).Encode(Sprint{ID: 1, Name: "Sprint 1", State: "closed"})
		}
	}))
	defer server.Close()

	svc := NewSprintService(c)
	sprint, err := svc.Complete(context.Background(), 1, "2024-01-29T18:00:00.000Z")
	if err != nil {
		t.Fatal(err)
	}
	if sprint.State != "closed" {
		t.Errorf("State = %v, want closed", sprint.State)
	}
}

func TestFieldService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]Field{
			{ID: "summary", Name: "Summary"},
			{ID: "description", Name: "Description"},
		})
	}))
	defer server.Close()

	svc := NewFieldService(c)
	fields, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Errorf("len(fields) = %v, want 2", len(fields))
	}
}

func TestComponentService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Component{ID: "10000", Name: "Backend"})
	}))
	defer server.Close()

	svc := NewComponentService(c)
	comp, err := svc.Get(context.Background(), "10000")
	if err != nil {
		t.Fatal(err)
	}
	if comp.Name != "Backend" {
		t.Errorf("Name = %v, want Backend", comp.Name)
	}
}

func TestVersionService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Version{ID: "10000", Name: "1.0"})
	}))
	defer server.Close()

	svc := NewVersionService(c)
	version, err := svc.Get(context.Background(), "10000", nil)
	if err != nil {
		t.Fatal(err)
	}
	if version.Name != "1.0" {
		t.Errorf("Name = %v, want 1.0", version.Name)
	}
}

func TestPriorityService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]Priority{
			{ID: "1", Name: "Highest"},
			{ID: "2", Name: "High"},
		})
	}))
	defer server.Close()

	svc := NewPriorityService(c)
	priorities, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(priorities) != 2 {
		t.Errorf("len(priorities) = %v, want 2", len(priorities))
	}
}

func TestStatusService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]StatusDetail{
			{ID: "1", Name: "Open"},
			{ID: "2", Name: "Closed"},
		})
	}))
	defer server.Close()

	svc := NewStatusService(c)
	statuses, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(statuses) != 2 {
		t.Errorf("len(statuses) = %v, want 2", len(statuses))
	}
}

func TestResolutionService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]Resolution{
			{ID: "1", Name: "Fixed"},
			{ID: "2", Name: "Won't Fix"},
		})
	}))
	defer server.Close()

	svc := NewResolutionService(c)
	resolutions, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(resolutions) != 2 {
		t.Errorf("len(resolutions) = %v, want 2", len(resolutions))
	}
}

func TestIssueTypeService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]IssueType{
			{ID: "1", Name: "Bug"},
			{ID: "2", Name: "Story"},
		})
	}))
	defer server.Close()

	svc := NewIssueTypeService(c)
	issueTypes, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(issueTypes) != 2 {
		t.Errorf("len(issueTypes) = %v, want 2", len(issueTypes))
	}
}

func TestPermissionService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Permissions{
			Permissions: map[string]PermissionScheme{
				"BROWSE": {ID: 1, Key: "BROWSE", Name: "Browse"},
			},
		})
	}))
	defer server.Close()

	svc := NewPermissionService(c)
	perms, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(perms.Permissions) != 1 {
		t.Errorf("len(perms) = %v, want 1", len(perms.Permissions))
	}
}

func TestMyselfService_Get(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(MyselfInfo{
			DisplayName:  "Test User",
			EmailAddress: "test@example.com",
		})
	}))
	defer server.Close()

	svc := NewMyselfService(c)
	info, err := svc.Get(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if info.DisplayName != "Test User" {
		t.Errorf("DisplayName = %v, want Test User", info.DisplayName)
	}
}

func TestScreenService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ScreenSearchResult{
			Total: 1,
			Values: []Screen{
				{ID: 1, Name: "Default Screen"},
			},
		})
	}))
	defer server.Close()

	svc := NewScreenService(c)
	result, err := svc.GetAll(context.Background(), 0, 50, "")
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("Total = %v, want 1", result.Total)
	}
}

func TestReindexService_GetStatus(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ReindexResult{Progress: 100.0})
	}))
	defer server.Close()

	svc := NewReindexService(c)
	result, err := svc.GetStatus(context.Background(), "123")
	if err != nil {
		t.Fatal(err)
	}
	if result.Progress != 100.0 {
		t.Errorf("Progress = %v, want 100.0", result.Progress)
	}
}

func TestReindexService_Trigger_AdminRequired(t *testing.T) {
	_, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not reach server when admin mode disabled")
	}))
	defer server.Close()

	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods:   config.MethodsConfig{AllowGet: true, AllowPost: true},
		AdminMode: false,
		Endpoints: config.DefaultEndpoints(),
	}
	c2, _ := client.New(cfg)

	svc := NewReindexService(c2)
	_, err := svc.Trigger(context.Background(), "")
	if err == nil {
		t.Error("Expected error for admin operation without admin mode, got nil")
	}
}

func TestApplicationPropertiesService_GetAll(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]ApplicationProperty{
			{ID: "jira.title", Key: "jira.title", Value: "Jira"},
		})
	}))
	defer server.Close()

	svc := NewApplicationPropertiesService(c)
	props, err := svc.GetAll(context.Background(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(props) != 1 {
		t.Errorf("len(props) = %v, want 1", len(props))
	}
}

func TestConfigurationService_Get_AdminRequired(t *testing.T) {
	_, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not reach server when admin mode disabled")
	}))
	defer server.Close()

	cfg := &config.Config{
		Jira: config.JiraConfig{
			BaseURL:  server.URL,
			APIToken: "test-token",
			Timeout:  30,
		},
		Methods:   config.MethodsConfig{AllowGet: true},
		AdminMode: false,
		Endpoints: config.DefaultEndpoints(),
	}
	c2, _ := client.New(cfg)

	svc := NewConfigurationService(c2)
	_, err := svc.Get(context.Background())
	if err == nil {
		t.Error("Expected error for admin operation without admin mode, got nil")
	}
}

func TestAvatarService_GetSystemAvatars(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(AvatarList{
			System: []Avatar{{ID: 1, IsSystem: true}},
		})
	}))
	defer server.Close()

	svc := NewAvatarService(c)
	avatars, err := svc.GetSystemAvatars(context.Background(), "project")
	if err != nil {
		t.Fatal(err)
	}
	if len(avatars.System) != 1 {
		t.Errorf("len(System) = %v, want 1", len(avatars.System))
	}
}

func TestSearchService_SearchGet(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method = %v, want GET", r.Method)
		}
		jql := r.URL.Query().Get("jql")
		if jql != "project = PROJ AND status = Open" {
			t.Errorf("jql = %v, want 'project = PROJ AND status = Open'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{
			Total: 1,
			Issues: []Issue{
				{Key: "PROJ-1"},
			},
		})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	result, err := svc.SearchGet(context.Background(), "project = PROJ AND status = Open", 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("Total = %v, want 1", result.Total)
	}
}

func TestSearchService_SearchGet_URLEncoding(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != "project = TEST AND summary ~ \"special chars & symbols\"" {
			t.Errorf("jql = %v, want 'project = TEST AND summary ~ \"special chars & symbols\"'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 0})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), "project = TEST AND summary ~ \"special chars & symbols\"", 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_SearchGet_WithPagination(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startAt := r.URL.Query().Get("startAt")
		maxResults := r.URL.Query().Get("maxResults")
		if startAt != "10" {
			t.Errorf("startAt = %v, want 10", startAt)
		}
		if maxResults != "5" {
			t.Errorf("maxResults = %v, want 5", maxResults)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 0})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), "project = PROJ", 10, 5, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_SearchGet_WithFieldsAndExpand(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fields := r.URL.Query().Get("fields")
		expand := r.URL.Query().Get("expand")
		if fields != "summary,status" {
			t.Errorf("fields = %v, want summary,status", fields)
		}
		if expand != "changelog" {
			t.Errorf("expand = %v, want changelog", expand)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 0})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), "project = PROJ", 0, 50, []string{"summary", "status"}, []string{"changelog"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMyselfService_GetWithExpand(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expand := r.URL.Query().Get("expand")
		if expand != "groups,applicationRoles" {
			t.Errorf("expand = %v, want groups,applicationRoles", expand)
		}
		json.NewEncoder(w).Encode(MyselfInfo{
			DisplayName:  "Test User",
			EmailAddress: "test@example.com",
		})
	}))
	defer server.Close()

	svc := NewMyselfService(c)
	info, err := svc.Get(context.Background(), []string{"groups", "applicationRoles"})
	if err != nil {
		t.Fatal(err)
	}
	if info.DisplayName != "Test User" {
		t.Errorf("DisplayName = %v, want Test User", info.DisplayName)
	}
}

func TestSearchService_EpicLinkQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `"Epic Link" = PROJ-100` {
			t.Errorf("jql = %v, want '\"Epic Link\" = PROJ-100'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{
			Total: 3,
			Issues: []Issue{
				{Key: "PROJ-101"},
				{Key: "PROJ-102"},
				{Key: "PROJ-103"},
			},
		})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	result, err := svc.SearchGet(context.Background(), `"Epic Link" = PROJ-100`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 3 {
		t.Errorf("Total = %v, want 3", result.Total)
	}
}

func TestSearchService_EpicLinkInQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `"Epic Link" in (PROJ-100, PROJ-101)` {
			t.Errorf("jql = %v, want '\"Epic Link\" in (PROJ-100, PROJ-101)'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 5})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `"Epic Link" in (PROJ-100, PROJ-101)`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_ParentQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `parent = PROJ-123` {
			t.Errorf("jql = %v, want 'parent = PROJ-123'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{
			Total: 2,
			Issues: []Issue{
				{Key: "PROJ-124"},
				{Key: "PROJ-125"},
			},
		})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	result, err := svc.SearchGet(context.Background(), `parent = PROJ-123`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 2 {
		t.Errorf("Total = %v, want 2", result.Total)
	}
}

func TestSearchService_StatusCategoryQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `statusCategory = Done` {
			t.Errorf("jql = %v, want 'statusCategory = Done'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 10})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `statusCategory = Done`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_ResolutionQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `resolution = Unresolved` {
			t.Errorf("jql = %v, want 'resolution = Unresolved'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 25})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `resolution = Unresolved`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_SprintQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `sprint in openSprints()` {
			t.Errorf("jql = %v, want 'sprint in openSprints()'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 8})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `sprint in openSprints()`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_SubTaskIssueTypesQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `issuetype in subTaskIssueTypes()` {
			t.Errorf("jql = %v, want 'issuetype in subTaskIssueTypes()'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 5})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `issuetype in subTaskIssueTypes()`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_ComplexCombinedQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		expected := `project = PROJ AND issuetype = Story AND status = "In Progress" ORDER BY priority DESC`
		if jql != expected {
			t.Errorf("jql = %v, want '%s'", jql, expected)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 7})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `project = PROJ AND issuetype = Story AND status = "In Progress" ORDER BY priority DESC`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_TextSearchQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `text ~ "authentication"` {
			t.Errorf("jql = %v, want 'text ~ \"authentication\"'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 3})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `text ~ "authentication"`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_DateQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `created >= -7d` {
			t.Errorf("jql = %v, want 'created >= -7d'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 15})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `created >= -7d`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchService_AssigneeQuery(t *testing.T) {
	c, server := testClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jql := r.URL.Query().Get("jql")
		if jql != `assignee = currentUser()` {
			t.Errorf("jql = %v, want 'assignee = currentUser()'", jql)
		}
		json.NewEncoder(w).Encode(SearchResult{Total: 12})
	}))
	defer server.Close()

	svc := NewSearchService(c)
	_, err := svc.SearchGet(context.Background(), `assignee = currentUser()`, 0, 50, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}
