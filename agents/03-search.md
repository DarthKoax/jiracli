# Jira CLI - Search & Filters

## Search Issues

```bash
jiracli search --jql <jql-query> [--start-at <num>] [--max-results <num>] [--fields <list>] [--expand <list>] [--config <path>]
```

**Flags:**
- `--jql <query>`: JQL query string (required)
- `--start-at <num>`: Index of first result (default: 0)
- `--max-results <num>`: Maximum number of results (default: 50)
- `--fields <list>`: Comma-separated list of fields to return
- `--expand <list>`: Comma-separated list of properties to expand (schema, names, operations, editmeta, changelog)
- `--config <path>`: Path to config file

**Description:**
Searches for issues using Jira Query Language (JQL). Maps to POST /rest/api/2/search.

**Examples:**
```bash
# Simple search
jiracli search --jql "project = PROJ"

# Search with status filter
jiracli search --jql "project = PROJ AND status = Open"

# Search with pagination
jiracli search --jql "project = PROJ" --start-at 0 --max-results 20

# Search with specific fields
jiracli search --jql "assignee = john.doe" --fields "summary,status,priority"

# Complex JQL query
jiracli search --jql "project = PROJ AND status != Closed AND priority in (Highest, High) ORDER BY created DESC"

# Search with date range
jiracli search --jql "project = PROJ AND created >= -7d"

# Search by component
jiracli search --jql "project = PROJ AND component = Backend"

# Search by label
jiracli search --jql "project = PROJ AND labels = urgent"
```

**Output:**
```json
{
  "startAt": 0,
  "maxResults": 50,
  "total": 123,
  "issues": [
    {
      "id": "10001",
      "key": "PROJ-123",
      "self": "https://jira.example.com/rest/api/2/issue/10001",
      "fields": {
        "summary": "Issue summary",
        "status": {"name": "Open"},
        "assignee": {"name": "john.doe"},
        "priority": {"name": "High"}
      }
    }
  ]
}
```

**JQL Reference:**

Common operators:
- `=` (equals)
- `!=` (not equals)
- `>`, `<`, `>=`, `<=` (comparison)
- `IN` (in list)
- `NOT IN` (not in list)
- `~` (contains)
- `IS` / `IS NOT` (null checks)

Common fields:
- `project`, `status`, `assignee`, `reporter`, `priority`, `issuetype`
- `created`, `updated`, `due`, `resolution`
- `labels`, `components`, `versions`, `fixVersions`
- `summary`, `description`, `comment`

Common functions:
- `now()`, `startOfDay()`, `endOfDay()`, `startOfWeek()`, `endOfWeek()`
- `currentUser()`, `membersOf("group")`
- `openSprints()`, `closedSprints()`, `linkedIssues(KEY)`

## Filters

### Get Filter
```bash
jiracli filter get <filter-id> [--expand <list>] [--config <path>]
```

**Flags:**
- `--expand <list>`: Comma-separated list of properties to expand (sharePermissions, editPermissions, subscriptions)

**Example:**
```bash
jiracli filter get 10000
jiracli filter get 10000 --expand sharePermissions
```

**Output:**
```json
{
  "id": "10000",
  "self": "https://jira.example.com/rest/api/2/filter/10000",
  "name": "My Open Issues",
  "description": "All open issues in PROJ",
  "owner": {"name": "john.doe", "displayName": "John Doe"},
  "jql": "project = PROJ AND status != Closed",
  "viewUrl": "https://jira.example.com/issues/?filter=10000",
  "searchUrl": "https://jira.example.com/rest/api/2/search?jql=project+%3D+PROJ+AND+status+%21%3D+Closed",
  "favourite": true,
  "sharePermissions": [
    {"type": "project", "project": {"id": "10000", "key": "PROJ"}}
  ]
}
```

### Create Filter
```bash
jiracli filter create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "My Open Issues",
  "description": "All open issues in PROJ",
  "jql": "project = PROJ AND status != Closed",
  "favourite": true,
  "sharePermissions": [
    {"type": "project", "project": {"id": "10000"}},
    {"type": "group", "group": {"name": "developers"}}
  ]
}
```

**Example:**
```bash
jiracli filter create --json '{
  "name": "My Filter",
  "jql": "project = PROJ AND status = Open",
  "favourite": true
}'
```

### Update Filter
```bash
jiracli filter update <filter-id> --json <json-payload> [--config <path>]
```

**Example:**
```bash
jiracli filter update 10000 --json '{
  "name": "Updated Filter Name",
  "jql": "project = PROJ AND status = Open",
  "favourite": true
}'
```

### Delete Filter
```bash
jiracli filter delete <filter-id> [--config <path>]
```

### List Favourite Filters
```bash
jiracli filter list [--config <path>]
```

**Example:**
```bash
jiracli filter list
```

**Output:**
```json
[
  {
    "id": "10000",
    "name": "My Open Issues",
    "jql": "project = PROJ AND status != Closed",
    "favourite": true
  }
]
```

### List My Filters
```bash
jiracli filter my [--start-at <num>] [--max-results <num>] [--config <path>]
```

**Flags:**
- `--start-at <num>`: Index of first result (default: 0)
- `--max-results <num>`: Maximum number of results (default: 50)

### Search Filters
```bash
jiracli filter search [--filter-name <name>] [--owner <username>] [--groupname <name>] [--start-at <num>] [--max-results <num>] [--config <path>]
```

**Flags:**
- `--filter-name <name>`: Filter name to search for
- `--owner <username>`: Filter owner username
- `--groupname <name>`: Group name
- `--start-at <num>`: Index of first result
- `--max-results <num>`: Maximum number of results

**Example:**
```bash
jiracli filter search --filter-name "My Filter"
jiracli filter search --owner john.doe
```

### Add Share Permission
```bash
jiracli filter permission add <filter-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"type": "group", "group": {"name": "testers"}}
```

**Example:**
```bash
jiracli filter permission add 10000 --json '{"type": "group", "group": {"name": "testers"}}'
```

### Set Favourite
```bash
jiracli filter favourite <filter-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"favourite": true}
```

**Example:**
```bash
jiracli filter favourite 10000 --json '{"favourite": true}'
```
