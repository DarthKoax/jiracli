# AGENTS.md — jiracli

## Project Overview

Go wrapper library for the **Jira Data Center / Server REST API**. Provides typed, configurable access to all major Jira DC endpoints with Bearer token authentication and custom CA certificate support.

## Build & Test

```bash
go mod tidy
go build ./...
go vet ./...
go test ./... -v
```

## Test Coverage Requirements

**All code changes must maintain comprehensive test coverage.**

### Testing Standards

1. **Unit Tests Required**: Every public function must have corresponding unit tests
2. **Test Coverage Target**: Maintain minimum 80% code coverage across all packages
3. **Test All Paths**: Cover both success and error paths in all functions
4. **Mock External Dependencies**: Use httptest for HTTP client testing
5. **Test Data Validation**: Verify config validation, endpoint gating, and admin checks

### Required Test Categories

- **Config Tests**: TOML parsing, defaults, environment variable overrides, validation
- **Client Tests**: HTTP method gating, admin mode checks, endpoint gating, authentication, TLS/CA handling
- **Service Tests**: All API service methods (23 services), endpoint disabled scenarios, admin operation blocking
- **Integration Tests**: End-to-end workflows with mock servers

### Test File Organization

- Test files must be co-located with source files (e.g., `config_test.go` alongside `config.go`)
- Use table-driven tests where appropriate for multiple scenarios
- Use `httptest.NewServer` for HTTP client testing
- Clean up test resources (temp files, servers) with `defer`

### Running Tests

```bash
# Run all tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run specific package tests
go test ./internal/config -v
go test ./internal/client -v
go test ./internal/api -v

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Before Committing

Always verify:
- [ ] `go build ./...` succeeds
- [ ] `go vet ./...` passes
- [ ] `go test ./...` passes with no failures
- [ ] New functionality has corresponding tests
- [ ] Error paths are tested
- [ ] Admin/endpoint gating is tested

## Configuration

All configuration is in `config.toml`:

| Key | Description |
|-----|-------------|
| `jira.base_url` | Jira DC instance URL (required) |
| `jira.api_token` | Bearer API token (required) |
| `jira.custom_ca_cert` | Path to custom CA certificate bundle |
| `jira.timeout` | HTTP timeout in seconds (default: 30) |
| `admin_mode` | Enable admin-only endpoints (default: false) |
| `methods.allow_get` | Enable GET requests (default: true) |
| `methods.allow_post` | Enable POST requests (default: true) |
| `methods.allow_put` | Enable PUT requests (default: true) |
| `methods.allow_delete` | Enable DELETE requests (default: true) |
| `endpoints.*` | Enable/disable individual API endpoints (default: all true) |

### Environment Variable Overrides

All config values can be overridden via environment variables:

| Env Var | Config Key |
|---------|------------|
| `JIRA_BASE_URL` | `jira.base_url` |
| `JIRA_API_TOKEN` | `jira.api_token` |
| `JIRA_CUSTOM_CA_CERT` | `jira.custom_ca_cert` |
| `JIRA_TIMEOUT` | `jira.timeout` |
| `JIRA_ADMIN_MODE` | `admin_mode` |
| `JIRA_ALLOW_GET` | `methods.allow_get` |
| `JIRA_ALLOW_POST` | `methods.allow_post` |
| `JIRA_ALLOW_PUT` | `methods.allow_put` |
| `JIRA_ALLOW_DELETE` | `methods.allow_delete` |
| `JIRA_ENDPOINT_<NAME>` | `endpoints.<name>` (e.g., `JIRA_ENDPOINT_BOARDS`) |
| `JIRA_CONFIG_DIR` | Override default config directory |

## Architecture

```
internal/
  config/   — TOML config loading, defaults, env overrides, init
  client/   — HTTP client (TLS, auth, method gating, endpoint gating, admin checks)
  api/      — Typed service wrappers per Jira resource (23 services)
cmd/jiracli/ — CLI entry point (init, connect, version, help subcommands)
agents/     — Agent skill documentation (jira-cli-skill.md)
```

### CLI Subcommands

| Command | Description |
|---------|-------------|
| `init` | Create default config at `~/.config/darthkoax/jiracli/config.toml` |
| `connect` | Connect to Jira and verify authentication |
| `version` | Print version information |
| `help` | Show usage information |

### Default Config Path

- Default: `~/.config/darthkoax/jiracli/config.toml`
- Override directory: `JIRA_CONFIG_DIR` environment variable
- Override file: `--config <path>` flag on `connect` command

## Included APIs (v1)

### Core Jira REST API (`/rest/api/2/`)

| Service | File | Endpoints |
|---------|------|-----------|
| **Issues** | `api/issues.go` | GET/POST/PUT/DELETE issue, transitions, comments, worklogs, watchers, votes, assign |
| **Projects** | `api/projects.go` | GET/POST/PUT/DELETE project, roles, components, versions, avatars |
| **Users** | `api/users.go` | GET/POST/DELETE user, search, picker, columns |
| **Groups** | `api/groups.go` | GET/POST/DELETE group, members, search, bulk |
| **Search** | `api/search.go` | POST JQL search, GET search |
| **Filters** | `api/filters.go` | GET/POST/PUT/DELETE filter, favourites, my filters, search, share permissions |
| **Dashboards** | `api/dashboards.go` | GET/POST/PUT/DELETE dashboard, gadgets |
| **Fields** | `api/fields.go` | GET all fields, POST custom field |
| **Components** | `api/components.go` | GET/POST/PUT/DELETE component, related issue counts |
| **Versions** | `api/versions.go` | GET/POST/PUT/DELETE version, related counts, merge, move unresolved |
| **Priorities** | `api/priorities.go` | GET all, GET by ID |
| **Statuses** | `api/statuses.go` | GET all, GET by ID, GET categories |
| **Resolutions** | `api/resolutions.go` | GET all, GET by ID |
| **Issue Types** | `api/issuetypes.go` | GET/POST/PUT/DELETE issue type, alternatives |
| **Permissions** | `api/permissions.go` | GET all permissions, my permissions, permission schemes CRUD, grants |
| **Myself** | `api/myself.go` | GET current user, locale |
| **Screens** | `api/screens.go` | GET/POST/PUT/DELETE screen, tabs, tab fields, move, available fields |
| **Reindex** | `api/reindex.go` | POST trigger, GET status |
| **Application Properties** | `api/applicationproperties.go` | GET/PUT properties |
| **Configuration** | `api/configuration.go` | GET global config |
| **Avatars** | `api/avatars.go` | GET system/custom avatars, DELETE/PUT custom |

### Jira Software REST API (`/rest/agile/1.0/`)

| Service | File | Endpoints |
|---------|------|-----------|
| **Boards** | `api/boards.go` | GET/POST/DELETE board, config, issues, epics, sprints, backlog |
| **Sprints** | `api/sprints.go` | GET/POST/PUT/DELETE sprint, issues, move, swap, complete |

## Excluded APIs — Not in v1

The following Jira DC REST API endpoints are **not included** in this release and must be documented for future work:

### Core Jira REST API (`/rest/api/2/`)

| Endpoint | Notes |
|----------|-------|
| `/rest/api/2/issue/{key}/attachments` | Multipart file upload requires `multipart/form-data` — stub exists but not functional |
| `/rest/api/2/issue/{key}/remotelink` | Remote issue links (link to external issues) |
| `/rest/api/2/issue/{key}/properties` | Issue entity properties (arbitrary JSON) |
| `/rest/api/2/issue/createmeta` | Create metadata (field configs for issue creation) |
| `/rest/api/2/issue/{key}/editmeta` | Edit metadata (field configs for issue editing) |
| `/rest/api/2/securitylevel` | Issue security levels |
| `/rest/api/2/password` | Password validation/change |
| `/rest/api/2/project/{key}/avatar` | POST custom project avatar (multipart upload) |
| `/rest/api/2/project/{key}/role/{id}` | Full CRUD for project roles (only GET implemented) |
| `/rest/api/2/projectValidate` | Project key validation |
| `/rest/api/2/project/{key}/statuses` | Valid statuses for project |
| `/rest/api/2/project/{key}/type` | Project type update |
| `/rest/api/2/user/assignable/multiProjectSearch` | Multi-project assignable user search |
| `/rest/api/2/user/bulk` | Bulk user retrieval |
| `/rest/api/2/user/picker` | User picker (partially implemented) |
| `/rest/api/2/group/userpicker` | Group user picker |
| `/rest/api/2/group/bulk` | Bulk group retrieval |
| `/rest/api/2/workflow` | Workflow CRUD |
| `/rest/api/2/workflowscheme` | Workflow scheme CRUD |
| `/rest/api/2/workflow/transitions/{id}/properties` | Workflow transition properties |
| `/rest/api/2/issuelinks` | Issue links (link/unlink issues) |
| `/rest/api/2/field/{fieldKey}/option` | Custom field option management |
| `/rest/api/2/field/{fieldKey}/option/{optionId}` | Custom field option CRUD |
| `/rest/api/2/field/{fieldKey}/option/suggestions/edit` | Editable option suggestions |
| `/rest/api/2/field/{fieldKey}/option/suggestions/view` | Viewable option suggestions |
| `/rest/api/2/jql/autocompletedata` | JQL autocomplete data |
| `/rest/api/2/jql/autocompletedata/suggestions` | JQL autocomplete suggestions |
| `/rest/api/2/auditRecord` | Audit records |
| `/rest/api/2/cluster/nodes` | Cluster node management |
| `/rest/api/2/tasks/{taskId}` | Background task status |
| `/rest/api/2/settings` | Global settings (base URL, etc.) |
| `/rest/api/2/locale` | System locale |
| `/rest/api/2/favouriteFilters` | Deprecated — use `/filter/favourite` |
| `/rest/api/2/dashboard/{id}/items` | Dashboard items (non-gadget) |
| `/rest/api/2/attachment/{id}` | Individual attachment operations |
| `/rest/api/2/attachment/meta` | Attachment limits/configuration |
| `/rest/api/2/comment/{commentId}` | Standalone comment operations |
| `/rest/api/2/comment/list` | Bulk comment retrieval |

### Jira Software REST API (`/rest/agile/1.0/`)

| Endpoint | Notes |
|----------|-------|
| `/rest/agile/1.0/board/{boardId}/version` | Board versions |
| `/rest/agile/1.0/board/{boardId}/features` | Board features |
| `/rest/agile/1.0/board/{boardId}/properties` | Board entity properties |
| `/rest/agile/1.0/epic/{epicId}` | Full epic CRUD (only list/get via board implemented) |
| `/rest/agile/1.0/epic/rank` | Rank epics |
| `/rest/agile/1.0/sprint/{sprintId}/issueswap` | Sprint issue swap (stub exists) |
| `/rest/agile/1.0/issue/rank` | Rank issues |
| `/rest/agile/1.0/board/{boardId}/reports/burndownchart` | Burndown chart data |
| `/rest/agile/1.0/board/{boardId}/reports/velocitychart` | Velocity chart data |
| `/rest/agile/1.0/board/{boardId}/reports/controlchart` | Control chart data |
| `/rest/agile/1.0/board/{boardId}/reports/cumulativeflowdiagram` | CFD data |

### Jira Service Desk REST API (`/rest/servicedeskapi/`)

The entire Jira Service Desk API is **not included** in v1:

| Endpoint | Notes |
|----------|-------|
| `/rest/servicedeskapi/servicedesk` | Service desk CRUD |
| `/rest/servicedeskapi/customer` | Customer management |
| `/rest/servicedeskapi/request` | Service desk requests |
| `/rest/servicedeskapi/request/{id}/comment` | Request comments |
| `/rest/servicedeskapi/request/{id}/attachment` | Request attachments |
| `/rest/servicedeskapi/request/{id}/status` | Request status transitions |
| `/rest/servicedeskapi/queue` | Queue management |
| `/rest/servicedeskapi/queue/{id}/issue` | Issues in queues |
| `/rest/servicedeskapi/organization` | Organization management |
| `/rest/servicedeskapi/servicedesk/{id}/organization` | SD-organization associations |
| `/rest/servicedeskapi/approval` | Approval management |
| `/rest/servicedeskapi/servicedesk/{id}/requesttype` | Request type management |
| `/rest/servicedeskapi/servicedesk/{id}/requesttype/{id}/field` | Request type fields |

### Other Excluded APIs

| API | Notes |
|-----|-------|
| Jira Portfolio (`/rest/portfolio/2.0/`) | Advanced roadmaps / portfolio planning |
| Jira Ops (`/rest/jira-ops/1.0/`) | Operations management |
| Plugin/Connect API (`/rest/atlassian-connect/1/`) | Connect app management |
| OAuth API | OAuth token management |
| Admin REST API | Jira DC admin-only endpoints |
| Issue Navigator / CSV export | Export formats beyond JSON |

## Method Gating

All HTTP methods are gated by the `[methods]` section in `config.toml`. When a method is disabled, the client returns an error before making any HTTP request. This allows read-only or append-only operational modes.

## Error Handling

- API errors return `fmt.Errorf` with the HTTP status code and response body.
- Config validation errors are returned at load time.
- TLS/CA errors are returned at client creation time.
