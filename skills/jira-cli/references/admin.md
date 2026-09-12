# Jira CLI - Admin Operations

**Note:** All admin operations require `admin_mode = true` in config.toml.

## Permissions

### List All Permissions
```bash
jiracli permission list [--config <path>]
```

**Output:**
```json
{
  "permissions": {
    "BROWSE_PROJECTS": {
      "id": 10,
      "key": "BROWSE_PROJECTS",
      "name": "Browse Projects",
      "type": "GLOBAL"
    }
  }
}
```

### Get My Permissions
```bash
jiracli permission my [--project-key <key>] [--project-id <id>] [--issue-key <key>] [--issue-id <id>] [--permissions <list>] [--config <path>]
```

**Flags:**
- `--project-key <key>`: Project key
- `--project-id <id>`: Project ID
- `--issue-key <key>`: Issue key
- `--issue-id <id>`: Issue ID
- `--permissions <list>`: Comma-separated list of permission keys

**Example:**
```bash
jiracli permission my
jiracli permission my --project-key PROJ
jiracli permission my --issue-key PROJ-123
jiracli permission my --permissions BROWSE_PROJECTS,EDIT_ISSUES
```

**Output:**
```json
{
  "permissions": {
    "BROWSE_PROJECTS": {
      "id": 10,
      "key": "BROWSE_PROJECTS",
      "name": "Browse Projects",
      "type": "PROJECT",
      "havePermission": true
    }
  }
}
```

### List Permission Schemes
```bash
jiracli permission scheme list [--start-at <num>] [--max-results <num>] [--expand <list>] [--config <path>]
```

**Flags:**
- `--expand <list>`: Properties to expand (permissions, user, group, projectRole, field, all)

**Example:**
```bash
jiracli permission scheme list
jiracli permission scheme list --expand permissions
```

**Output:**
```json
{
  "startAt": 0,
  "maxResults": 50,
  "total": 5,
  "permissionSchemes": [
    {
      "id": 10000,
      "name": "Default Permission Scheme",
      "description": "Default scheme"
    }
  ]
}
```

### Get Permission Scheme
```bash
jiracli permission scheme get <scheme-id> [--expand <list>] [--config <path>]
```

**Example:**
```bash
jiracli permission scheme get 10000
jiracli permission scheme get 10000 --expand permissions
```

### Create Permission Scheme (Admin)
```bash
jiracli permission scheme create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "Custom Scheme",
  "description": "Custom permission scheme",
  "permissions": [
    {
      "permission": "BROWSE_PROJECTS",
      "holder": {"type": "group", "parameter": "developers"}
    }
  ]
}
```

**Example:**
```bash
jiracli permission scheme create --json '{
  "name": "Custom Scheme",
  "description": "Custom permission scheme"
}'
```

### Delete Permission Scheme (Admin)
```bash
jiracli permission scheme delete <scheme-id> [--config <path>]
```

### Update Permission Scheme (Admin)
```bash
jiracli permission scheme update <scheme-id> --json <json-payload> [--config <path>]
```

### Add Permission Grant
```bash
jiracli permission grant add <scheme-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "permission": "EDIT_ISSUES",
  "holder": {"type": "group", "parameter": "developers"}
}
```

**Example:**
```bash
jiracli permission grant add 10000 --json '{
  "permission": "EDIT_ISSUES",
  "holder": {"type": "group", "parameter": "developers"}
}'
```

### Remove Permission Grant
```bash
jiracli permission grant remove <scheme-id> <permission-id> [--config <path>]
```

**Example:**
```bash
jiracli permission grant remove 10000 10001
```

## Reindex

### Trigger Reindex (Admin)
```bash
jiracli reindex trigger [--type <type>] [--config <path>]
```

**Flags:**
- `--type <type>`: Reindex type (BACKGROUND, FOREGROUND, PREFERRED)

**Example:**
```bash
jiracli reindex trigger
jiracli reindex trigger --type BACKGROUND
```

**Output:**
```json
{
  "reindexRequestId": 123,
  "progress": 0.0
}
```

### Get Reindex Status
```bash
jiracli reindex status --task-id <task-id> [--config <path>]
```

**Flags:**
- `--task-id <task-id>`: Task ID (required)

**Example:**
```bash
jiracli reindex status --task-id 123
```

**Output:**
```json
{
  "reindexRequestId": 123,
  "progress": 75.5
}
```

### Get Reindex Request Status
```bash
jiracli reindex request <request-id> [--config <path>]
```

**Example:**
```bash
jiracli reindex request 123
```

## Application Properties (Admin)

### List Properties
```bash
jiracli appprop list [--key <key>] [--permission-level <level>] [--config <path>]
```

**Flags:**
- `--key <key>`: Property key filter
- `--permission-level <level>`: Permission level (ADMIN, USERS)

**Example:**
```bash
jiracli appprop list
jiracli appprop list --key jira.title
```

**Output:**
```json
[
  {
    "id": "jira.title",
    "key": "jira.title",
    "value": "My Jira Instance",
    "name": "Jira Title",
    "type": "string"
  }
]
```

### Set Property (Admin)
```bash
jiracli appprop set <property-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"id": "jira.title", "value": "My Jira Instance"}
```

**Example:**
```bash
jiracli appprop set jira.title --json '{"id": "jira.title", "value": "My Jira"}'
```

## Configuration (Admin)

### Get Global Configuration
```bash
jiracli config get [--config <path>]
```

**Output:**
```json
{
  "self": "https://jira.example.com/rest/api/2/configuration",
  "votingEnabled": true,
  "watchingEnabled": true,
  "unassignedIssuesAllowed": true,
  "subTasksEnabled": true,
  "issueLinkingEnabled": true,
  "timeTrackingEnabled": true,
  "attachmentsEnabled": true
}
```

## Avatars

### Get System Avatars
```bash
jiracli avatar system <entity-type> [--config <path>]
```

**Arguments:**
- `<entity-type>`: Entity type (project, user, issuetype)

**Example:**
```bash
jiracli avatar system project
```

**Output:**
```json
{
  "system": [
    {
      "id": 10001,
      "owner": "system",
      "isSystemAvatar": true,
      "isSelected": false,
      "filename": "default.png"
    }
  ],
  "custom": []
}
```

### Get Custom Avatars
```bash
jiracli avatar custom <entity-type> <entity-id> [--config <path>]
```

**Example:**
```bash
jiracli avatar custom project 10000
```

### Delete Custom Avatar
```bash
jiracli avatar delete <entity-type> <avatar-id> [--config <path>]
```

**Example:**
```bash
jiracli avatar delete project 10001
```

### Update Custom Avatar
```bash
jiracli avatar update <entity-type> <avatar-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "id": 10001,
  "owner": "john.doe",
  "isSystemAvatar": false,
  "isSelected": true,
  "filename": "custom-avatar.png"
}
```

**Example:**
```bash
jiracli avatar update project 10001 --json '{
  "id": 10001,
  "isSelected": true
}'
```
