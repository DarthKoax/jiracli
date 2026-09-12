# jiracli

A Go wrapper library and CLI for the **Jira Data Center / Server REST API**. Provides typed, configurable access to all major Jira DC endpoints with Bearer token authentication and custom CA certificate support.

## Features

- **Full API Coverage**: 23 service wrappers covering Issues, Projects, Users, Groups, Search, Filters, Dashboards, Boards, Sprints, and more
- **Configurable Access Control**: Method gating (GET/POST/PUT/DELETE) and per-endpoint toggles
- **Admin Mode**: Separate admin-only endpoint protection for sensitive operations
- **Custom TLS**: Support for custom CA certificates for self-signed or internal CAs
- **Bearer Token Auth**: Simple API token authentication
- **TOML Configuration**: Human-readable config file format
- **Environment Overrides**: All settings overridable via environment variables

## Installation

### From Source

```bash
git clone https://github.com/darkkoax/jiracli.git
cd jiracli
go build -o jiracli ./cmd/jiracli
```

### Install to GOPATH

```bash
go install github.com/darkkoax/jiracli/cmd/jiracli@latest
```

## Quick Start

### 1. Initialize Configuration

```bash
jiracli init
```

This creates a default config file at `~/.config/jiracli/config.toml`.

### 2. Edit Configuration

Open the config file and set your Jira instance details:

```toml
[jira]
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
# ... enable/disable endpoints as needed
```

### 3. Test Connection

```bash
jiracli connect
```

Expected output:
```
Connected to Jira DC as: John Doe (john.doe@example.com)
```

## Configuration

### Config File Location

Default: `~/.config/jiracli/config.toml`

Override with `JIRA_CONFIG_DIR` environment variable:
```bash
export JIRA_CONFIG_DIR="/custom/path"
jiracli init
```

### Configuration Options

| Key | Description | Default |
|-----|-------------|---------|
| `jira.base_url` | Jira DC instance URL (required) | - |
| `jira.api_token` | Bearer API token (required) | - |
| `jira.custom_ca_cert` | Path to custom CA certificate bundle | `""` |
| `jira.timeout` | HTTP timeout in seconds | `30` |
| `admin_mode` | Enable admin-only endpoints | `false` |
| `methods.allow_get` | Enable GET requests | `true` |
| `methods.allow_post` | Enable POST requests | `true` |
| `methods.allow_put` | Enable PUT requests | `true` |
| `methods.allow_delete` | Enable DELETE requests | `true` |
| `endpoints.*` | Enable/disable individual endpoints | `true` |

### Environment Variable Overrides

All configuration options can be overridden via environment variables:

```bash
JIRA_BASE_URL="https://jira.example.com"
JIRA_API_TOKEN="your-token"
JIRA_CUSTOM_CA_CERT="/path/to/ca.crt"
JIRA_TIMEOUT="60"
JIRA_ADMIN_MODE="true"
JIRA_ALLOW_GET="true"
JIRA_ALLOW_POST="true"
JIRA_ALLOW_PUT="true"
JIRA_ALLOW_DELETE="true"
JIRA_ENDPOINT_ISSUES="true"
JIRA_ENDPOINT_BOARDS="false"
# ... etc
```

## Agent setup.

### Opencode
### Skill config
Copy the included skill to your harness skills directory.
```
cp -rp skills/jira-cli ~/.config/opencode/skills
```
#### Agent Permission
Apply the following to your opencode.json permissions block if you want to 
manual review write operations of your agent.
```json
 "permission": {
 14     "bash": {
 17       "jiracli *": "allow",
 18       "JIRA_ALLOW_*=true jiracli *": "ask"
 19     }
 20   },
```

## CLI Commands

### `jiracli init`

Create a default configuration file.

```bash
jiracli init [--dir <path>]
```

| Flag | Description |
|------|-------------|
| `--dir <path>` | Directory to create config in (default: `~/.config/jiracli`) |

### `jiracli connect`

Connect to Jira and verify authentication.

```bash
jiracli connect [--config <path>]
```

| Flag | Description |
|------|-------------|
| `--config <path>` | Path to config file (default: `~/.config/jiracli/config.toml`) |

### `jiracli version`

Print version information.

```bash
jiracli version
```

### `jiracli help`

Show help message.

```bash
jiracli help
```

## API Coverage

### Core Jira REST API (`/rest/api/2/`)

| Service | Endpoints |
|---------|-----------|
| **Issues** | GET/POST/PUT/DELETE issue, transitions, comments, worklogs, watchers, votes, assign |
| **Projects** | GET/POST/PUT/DELETE project, roles, components, versions, avatars |
| **Users** | GET/POST/DELETE user, search, picker, columns |
| **Groups** | GET/POST/DELETE group, members, search, bulk |
| **Search** | POST JQL search, GET search |
| **Filters** | GET/POST/PUT/DELETE filter, favourites, my filters, search, share permissions |
| **Dashboards** | GET/POST/PUT/DELETE dashboard, gadgets |
| **Fields** | GET all fields, POST custom field |
| **Components** | GET/POST/PUT/DELETE component, related issue counts |
| **Versions** | GET/POST/PUT/DELETE version, related counts, merge, move unresolved |
| **Priorities** | GET all, GET by ID |
| **Statuses** | GET all, GET by ID, GET categories |
| **Resolutions** | GET all, GET by ID |
| **Issue Types** | GET/POST/PUT/DELETE issue type, alternatives |
| **Permissions** | GET all permissions, my permissions, permission schemes CRUD, grants |
| **Myself** | GET current user, locale |
| **Screens** | GET/POST/PUT/DELETE screen, tabs, tab fields, move, available fields |
| **Reindex** | POST trigger, GET status (admin) |
| **Application Properties** | GET/PUT properties (admin) |
| **Configuration** | GET global config (admin) |
| **Avatars** | GET system/custom avatars, DELETE/PUT custom |

### Jira Software REST API (`/rest/agile/1.0/`)

| Service | Endpoints |
|---------|-----------|
| **Boards** | GET/POST/DELETE board, config, issues, epics, sprints, backlog |
| **Sprints** | GET/POST/PUT/DELETE sprint, issues, move, swap, complete |

## Admin Mode

Some endpoints require admin privileges. Enable `admin_mode = true` in config to access:

- **Reindex**: Trigger and monitor reindex operations
- **Application Properties**: Get/set application properties
- **Configuration**: Get global Jira configuration
- **User Management**: Create/delete users
- **Group Management**: Create/delete groups
- **Screen Management**: Create/update/delete screens
- **Permission Schemes**: Create/update/delete permission schemes

## Access Control

### Method Gating

Disable HTTP methods to create read-only or append-only modes:

```toml
[methods]
allow_get = true
allow_post = false    # Read-only mode
allow_put = false
allow_delete = false
```

### Endpoint Gating

Disable specific endpoints to restrict functionality:

```toml
[endpoints]
issues = true
boards = false        # Disable board operations
sprints = false       # Disable sprint operations
```

## Development

### Build

```bash
go build ./...
```

### Test

```bash
go test ./... -v
```

### Test with Coverage

```bash
go test ./... -cover
```

### Lint

```bash
go vet ./...
```

## Project Structure

```
.
├── cmd/jiracli/          # CLI entry point
├── internal/
│   ├── api/              # API service wrappers (23 services)
│   ├── client/           # HTTP client with TLS, auth, method gating
│   └── config/           # TOML config loading
├── agents/               # Agent skill documentation
├── config.toml           # Example configuration
├── AGENTS.md             # Developer documentation
└── README.md             # This file
```

## Architecture

### Client

The HTTP client (`internal/client/`) handles:
- Bearer token authentication
- Custom CA certificate loading
- Method gating (GET/POST/PUT/DELETE)
- Endpoint gating
- Admin mode checks
- TLS configuration

### Services

Each API resource has a dedicated service (`internal/api/`):
- Type-safe request/response structures
- Endpoint and admin checks before each operation
- Consistent error handling

### Config

Configuration (`internal/config/`) supports:
- TOML file parsing
- Environment variable overrides
- Default values
- Validation

## Error Handling

- **Config errors**: Returned at load time with descriptive messages
- **TLS/CA errors**: Returned at client creation time
- **API errors**: Return HTTP status code and response body
- **Method/endpoint disabled**: Return error before HTTP request

## Limitations

- Multipart file uploads (attachments) not yet implemented
- OAuth authentication not supported (Bearer token only)
- Some advanced endpoints excluded (see AGENTS.md for full list)

## License

See LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass: `go test ./... -v`
5. Ensure code passes lint: `go vet ./...`
6. Submit a pull request

See AGENTS.md for detailed development guidelines and test coverage requirements.
