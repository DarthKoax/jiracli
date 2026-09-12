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
- `standardIssueTypes()`, `subTaskIssueTypes()`

## Advanced JQL Patterns

### Epic & Parent Queries

**Find issues belonging to an epic:**
```bash
# Single epic
jiracli search --jql '"Epic Link" = PROJ-100'

# Multiple epics
jiracli search --jql '"Epic Link" in (PROJ-100, PROJ-101)'

# All issues in any epic
jiracli search --jql '"Epic Link" is not EMPTY'

# Issues NOT in any epic
jiracli search --jql '"Epic Link" is EMPTY'
```

**IMPORTANT:** Use `"Epic Link"` (the display name with quotes), NOT the raw custom field ID (e.g., `customfield_10108`). JQL rejects raw custom field IDs.

**Find sub-tasks (children) of an issue:**
```bash
jiracli search --jql 'parent = PROJ-123'
```

**Find top-level issues (no parent):**
```bash
# parent is EMPTY is NOT supported. Use this instead:
jiracli search --jql 'issuetype != Sub-task'
```

**IMPORTANT:** `parent is EMPTY` returns an error. Use `issuetype != Sub-task` to find top-level issues.

**IMPORTANT:** Sub-tasks do NOT have an Epic Link. They inherit the epic through their parent. To get ALL issues in an epic including sub-tasks, use a two-step approach:
1. Find direct epic children: `"Epic Link" = PROJ-100`
2. For each child, find sub-tasks: `parent = PROJ-123`

### Status & Resolution Queries

```bash
# Single status (quotes needed for multi-word)
jiracli search --jql 'status = "To Do"'
jiracli search --jql 'status = "In Progress"'

# Multiple statuses
jiracli search --jql 'status in ("To Do", "In Progress")'

# Status categories
jiracli search --jql 'statusCategory = "To Do"'
jiracli search --jql 'statusCategory = Done'
jiracli search --jql 'statusCategory != Done'

# Unresolved issues
jiracli search --jql 'resolution = Unresolved'
```

### Issue Type Queries

```bash
# Standard issue types (Story, Bug, Task, etc.)
jiracli search --jql 'issuetype in standardIssueTypes()'

# Sub-tasks only
jiracli search --jql 'issuetype in subTaskIssueTypes()'

# Multiple types
jiracli search --jql 'issuetype in (Story, Bug)'

# Exclude types
jiracli search --jql 'issuetype not in (Sub-task, Epic)'
```

### Text Search

```bash
# Search summary only
jiracli search --jql 'summary ~ "login"'

# Search all text fields (summary, description, comments, etc.)
jiracli search --jql 'text ~ "authentication"'
```

### Date Queries

```bash
# Relative dates
jiracli search --jql 'created >= -7d'
jiracli search --jql 'updated >= -1d'

# Date functions
jiracli search --jql 'created >= startOfDay(-7d)'
```

### Sprint Queries

```bash
# Current sprint
jiracli search --jql 'sprint in openSprints()'

# Any sprint
jiracli search --jql 'sprint is not EMPTY'

# Specific sprint by name
jiracli search --jql 'sprint = "Sprint 1"'

# Specific sprint by ID
jiracli search --jql 'sprint = 1'
```

**IMPORTANT:** The `sprint` field can be queried but is NOT returned in search results even when requested via `--fields`. To see sprint details for issues, use `jiracli sprint issues <sprint-id>`.

### Assignee Queries

```bash
# Current user's issues
jiracli search --jql 'assignee = currentUser()'

# Specific user
jiracli search --jql 'assignee = john.doe'

# Unassigned issues
jiracli search --jql 'assignee is EMPTY'
```

### Ordering

```bash
# Single field
jiracli search --jql 'project = PROJ ORDER BY created DESC'

# Multiple fields
jiracli search --jql 'project = PROJ ORDER BY priority DESC, created ASC'

# By key
jiracli search --jql 'project = PROJ ORDER BY key DESC'
```

### Complex Combined Queries

```bash
# Stories in progress, ordered by priority
jiracli search --jql 'project = PROJ AND issuetype = Story AND status = "In Progress" ORDER BY priority DESC'

# Unresolved bugs assigned to me
jiracli search --jql 'project = PROJ AND issuetype = Bug AND resolution = Unresolved AND assignee = currentUser()'

# Issues in epic, not done, ordered by key
jiracli search --jql '"Epic Link" = PROJ-100 AND statusCategory != Done ORDER BY key DESC'
```

### JQL Gotchas

1. **Multi-word values need quotes:** `status = "In Progress"`, `statusCategory = "To Do"`
2. **`"Epic Link"` uses display name, not custom field ID** — `customfield_10108` will fail
3. **`parent is EMPTY` is not supported** — use `issuetype != Sub-task` instead
4. **Sub-tasks don't have Epic Link** — they're linked to epics only through their parent
5. **Sprint field can be queried but not displayed** in search results
6. **Component/label values must exist** — querying non-existent values returns errors, not empty results
7. **`text ~ "..."` searches all text fields**, while `summary ~ "..."` searches only summary
8. **`resolution = Unresolved`** is the standard way to find open/unresolved issues

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
