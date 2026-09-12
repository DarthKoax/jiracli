# Jira CLI - Issue Management

## Get Issue

```bash
jiracli issue get <issue-key> [--fields <list>] [--expand <list>] [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID

**Flags:**
- `--fields <list>`: Comma-separated list of fields to return (default: all fields)
- `--expand <list>`: Comma-separated list of properties to expand (renderedFields, names, schema, transitions, operations, editmeta, changelog)
- `--config <path>`: Path to config file

**Description:**
Retrieves detailed information about a Jira issue. Maps to GET /rest/api/2/issue/{issueIdOrKey}.

**Examples:**
```bash
jiracli issue get PROJ-123
jiracli issue get PROJ-123 --fields "summary,status,assignee,priority"
jiracli issue get PROJ-123 --expand "renderedFields,changelog"
```

**Output:**
```json
{
  "id": "10001",
  "key": "PROJ-123",
  "self": "https://jira.example.com/rest/api/2/issue/10001",
  "fields": {
    "summary": "Issue summary",
    "status": {"name": "Open"},
    "assignee": {"name": "john.doe", "displayName": "John Doe"},
    "priority": {"name": "High"},
    "description": "Issue description",
    "created": "2024-01-15T10:00:00.000+0000",
    "updated": "2024-01-16T15:30:00.000+0000"
  }
}
```

## Create Issue

```bash
jiracli issue create --json <json-payload> [--config <path>]
```

**Flags:**
- `--json <payload>`: JSON object with issue fields (required)
- `--config <path>`: Path to config file

**Description:**
Creates a new issue in Jira. Maps to POST /rest/api/2/issue.

**JSON Payload Structure:**
```json
{
  "fields": {
    "project": {"key": "PROJ"},
    "summary": "Issue summary",
    "description": "Issue description",
    "issuetype": {"name": "Bug"},
    "priority": {"name": "High"},
    "labels": ["label1", "label2"],
    "components": [{"name": "Backend"}],
    "versions": [{"name": "1.0"}],
    "fixVersions": [{"name": "1.1"}],
    "assignee": {"name": "john.doe"},
    "reporter": {"name": "jane.smith"},
    "customfield_10001": "custom value"
  }
}
```

**Examples:**
```bash
# Simple bug
jiracli issue create --json '{
  "fields": {
    "project": {"key": "PROJ"},
    "summary": "Fix login error",
    "issuetype": {"name": "Bug"},
    "priority": {"name": "High"}
  }
}'

# Story with all fields
jiracli issue create --json '{
  "fields": {
    "project": {"key": "PROJ"},
    "summary": "Implement new feature",
    "description": "Detailed description here",
    "issuetype": {"name": "Story"},
    "priority": {"name": "Medium"},
    "labels": ["feature", "backend"],
    "components": [{"name": "API"}],
    "assignee": {"name": "john.doe"},
    "reporter": {"name": "jane.smith"}
  }
}'
```

**Output:**
```json
{
  "id": "10002",
  "key": "PROJ-124",
  "self": "https://jira.example.com/rest/api/2/issue/10002"
}
```

## Update Issue

```bash
jiracli issue update <issue-key> --json <json-payload> [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID

**Flags:**
- `--json <payload>`: JSON object with fields to update (required)
- `--config <path>`: Path to config file

**Description:**
Updates an existing issue. Only fields included in the JSON payload will be updated. Maps to PUT /rest/api/2/issue/{issueIdOrKey}.

**Examples:**
```bash
# Update summary
jiracli issue update PROJ-123 --json '{
  "fields": {
    "summary": "Updated summary"
  }
}'

# Update multiple fields
jiracli issue update PROJ-123 --json '{
  "fields": {
    "summary": "New summary",
    "priority": {"name": "Highest"},
    "labels": ["urgent", "critical"],
    "assignee": {"name": "john.doe"}
  }
}'

# Clear a field
jiracli issue update PROJ-123 --json '{
  "fields": {
    "assignee": null
  }
}'
```

**Output:**
```
Issue PROJ-123 updated successfully
```

## Delete Issue

```bash
jiracli issue delete <issue-key> [--delete-subtasks] [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID

**Flags:**
- `--delete-subtasks`: Also delete subtasks of this issue
- `--config <path>`: Path to config file

**Description:**
Deletes an issue from Jira. This action cannot be undone. Maps to DELETE /rest/api/2/issue/{issueIdOrKey}.

**Examples:**
```bash
jiracli issue delete PROJ-123
jiracli issue delete PROJ-123 --delete-subtasks
```

**Output:**
```
Issue PROJ-123 deleted successfully
```

## Get Transitions

```bash
jiracli issue transitions <issue-key> [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID

**Flags:**
- `--config <path>`: Path to config file

**Description:**
Retrieves the list of available transitions for an issue. Maps to GET /rest/api/2/issue/{issueIdOrKey}/transitions.

**Example:**
```bash
jiracli issue transitions PROJ-123
```

**Output:**
```json
[
  {
    "id": "11",
    "name": "Start Progress",
    "to": {
      "id": "3",
      "name": "In Progress",
      "description": "Work has started"
    }
  },
  {
    "id": "21",
    "name": "Done",
    "to": {
      "id": "5",
      "name": "Done",
      "description": "Work is complete"
    }
  }
]
```

## Do Transition

```bash
jiracli issue transition <issue-key> <transition-id> [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID
- `<transition-id>`: Transition ID (from 'jiracli issue transitions')

**Flags:**
- `--config <path>`: Path to config file

**Description:**
Performs a workflow transition on an issue. Maps to POST /rest/api/2/issue/{issueIdOrKey}/transitions.

**Examples:**
```bash
jiracli issue transition PROJ-123 11
jiracli issue transition PROJ-123 21
```

**Output:**
```
Issue PROJ-123 transitioned successfully
```

## Comments

### List Comments
```bash
jiracli issue comment list <issue-key> [--config <path>]
```

### Add Comment
```bash
jiracli issue comment add <issue-key> --body <text> [--config <path>]
```

**Flags:**
- `--body <text>`: Comment text (required)

**Example:**
```bash
jiracli issue comment add PROJ-123 --body "Working on this issue"
```

**Output:**
```json
{
  "id": "10001",
  "body": "Working on this issue",
  "author": {"name": "john.doe", "displayName": "John Doe"},
  "created": "2024-01-15T10:00:00.000+0000"
}
```

### Update Comment
```bash
jiracli issue comment update <issue-key> <comment-id> --body <text> [--config <path>]
```

### Delete Comment
```bash
jiracli issue comment delete <issue-key> <comment-id> [--config <path>]
```

## Worklogs

### List Worklogs
```bash
jiracli issue worklog list <issue-key> [--config <path>]
```

### Add Worklog
```bash
jiracli issue worklog add <issue-key> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "timeSpent": "2h 30m",
  "comment": "Implemented the fix and wrote tests",
  "started": "2024-01-15T10:00:00.000+0000"
}
```

**Example:**
```bash
jiracli issue worklog add PROJ-123 --json '{
  "timeSpent": "2h",
  "comment": "Implementation work"
}'
```

### Update Worklog
```bash
jiracli issue worklog update <issue-key> <worklog-id> --json <json-payload> [--config <path>]
```

### Delete Worklog
```bash
jiracli issue worklog delete <issue-key> <worklog-id> [--config <path>]
```

## Watchers

### List Watchers
```bash
jiracli issue watcher list <issue-key> [--config <path>]
```

### Add Watcher
```bash
jiracli issue watcher add <issue-key> <username> [--config <path>]
```

**Example:**
```bash
jiracli issue watcher add PROJ-123 john.doe
```

### Remove Watcher
```bash
jiracli issue watcher remove <issue-key> <username> [--config <path>]
```

## Votes

### Get Votes
```bash
jiracli issue vote get <issue-key> [--config <path>]
```

**Output:**
```json
{
  "self": "https://jira.example.com/rest/api/2/issue/PROJ-123/votes",
  "votes": 5,
  "hasVoted": true
}
```

### Add Vote
```bash
jiracli issue vote add <issue-key> [--config <path>]
```

### Remove Vote
```bash
jiracli issue vote remove <issue-key> [--config <path>]
```

## Assign Issue

```bash
jiracli issue assign <issue-key> <username> [--config <path>]
```

**Arguments:**
- `<issue-key>`: Issue key (e.g., PROJ-123) or issue ID
- `<username>`: Username to assign the issue to

**Example:**
```bash
jiracli issue assign PROJ-123 john.doe
```

**Output:**
```
Issue PROJ-123 assigned to john.doe
```

**Note:** Use empty string to unassign: `jiracli issue assign PROJ-123 ""`
