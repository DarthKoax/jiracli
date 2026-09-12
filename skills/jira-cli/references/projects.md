# Jira CLI - Project Management

## Projects

### List Projects
```bash
jiracli project list [--expand <list>] [--config <path>]
```

**Flags:**
- `--expand <list>`: Comma-separated list of properties to expand (lead, description, projectCategory)

**Example:**
```bash
jiracli project list
jiracli project list --expand lead,description
```

### Get Project
```bash
jiracli project get <project-key> [--expand <list>] [--config <path>]
```

**Example:**
```bash
jiracli project get PROJ
```

**Output:**
```json
{
  "id": "10000",
  "key": "PROJ",
  "name": "My Project",
  "projectTypeKey": "software",
  "lead": {"name": "john.doe", "displayName": "John Doe"},
  "description": "Project description"
}
```

### Create Project
```bash
jiracli project create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "key": "PROJ",
  "name": "My Project",
  "projectTypeKey": "software",
  "lead": "john.doe",
  "description": "Project description"
}
```

### Update Project
```bash
jiracli project update <project-key> --json <json-payload> [--config <path>]
```

### Delete Project
```bash
jiracli project delete <project-key> [--config <path>]
```

### Get Project Roles
```bash
jiracli project roles <project-key> [--config <path>]
```

### Get Project Components
```bash
jiracli project components <project-key> [--config <path>]
```

### Get Project Versions
```bash
jiracli project versions <project-key> [--config <path>]
```

## Components

### Get Component
```bash
jiracli component get <component-id> [--config <path>]
```

### Create Component
```bash
jiracli component create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "Backend",
  "description": "Backend services",
  "leadUserName": "john.doe",
  "assigneeType": "COMPONENT_LEAD",
  "project": "PROJ"
}
```

### Update Component
```bash
jiracli component update <component-id> --json <json-payload> [--config <path>]
```

### Delete Component
```bash
jiracli component delete <component-id> [--replace-with <component-id>] [--config <path>]
```

### Get Related Issue Count
```bash
jiracli component count <component-id> [--config <path>]
```

## Versions

### Get Version
```bash
jiracli version-resource get <version-id> [--expand <list>] [--config <path>]
```

### Create Version
```bash
jiracli version-resource create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "1.0.0",
  "description": "First release",
  "project": "PROJ",
  "startDate": "2024-01-01",
  "releaseDate": "2024-03-01",
  "released": false,
  "archived": false
}
```

### Update Version
```bash
jiracli version-resource update <version-id> --json <json-payload> [--config <path>]
```

### Delete Version
```bash
jiracli version-resource delete <version-id> [--move-fix-issues-to <version-id>] [--move-affected-issues-to <version-id>] [--config <path>]
```

### Get Related Issue Counts
```bash
jiracli version-resource count <version-id> [--config <path>]
```

### Move Unresolved Issues
```bash
jiracli version-resource move <version-id> --json '{"moveTo": "<version-id>"}' [--config <path>]
```

### Merge Versions
```bash
jiracli version-resource merge <version-id> <move-to-version-id> [--config <path>]
```

## Fields

### List Fields
```bash
jiracli field list [--config <path>]
```

**Output:**
```json
[
  {
    "id": "summary",
    "name": "Summary",
    "custom": false,
    "orderable": true,
    "navigable": true,
    "searchable": true,
    "clauseNames": ["summary"]
  }
]
```

### Create Custom Field
```bash
jiracli field create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "Customer ID",
  "description": "Customer identifier",
  "type": "com.atlassian.jira.plugin.system.customfieldtypes:textfield",
  "searcherKey": "com.atlassian.jira.plugin.system.customfieldtypes:textsearcher"
}
```

## Issue Types

### List Issue Types
```bash
jiracli issuetype list [--config <path>]
```

### Get Issue Type
```bash
jiracli issuetype get <issue-type-id> [--config <path>]
```

### Create Issue Type
```bash
jiracli issuetype create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "Epic",
  "description": "Large user story",
  "type": "standard"
}
```

### Update Issue Type
```bash
jiracli issuetype update <issue-type-id> --json <json-payload> [--config <path>]
```

### Delete Issue Type
```bash
jiracli issuetype delete <issue-type-id> [--alternative-issue-type-id <id>] [--config <path>]
```

### Get Alternative Issue Types
```bash
jiracli issuetype alternatives <issue-type-id> [--config <path>]
```

## Priorities, Statuses, Resolutions

### List Priorities
```bash
jiracli priority list [--config <path>]
```

### Get Priority
```bash
jiracli priority get <priority-id> [--config <path>]
```

### List Statuses
```bash
jiracli status list [--config <path>]
```

### Get Status
```bash
jiracli status get <status-id> [--config <path>]
```

### List Status Categories
```bash
jiracli status categories [--config <path>]
```

### List Resolutions
```bash
jiracli resolution list [--config <path>]
```

### Get Resolution
```bash
jiracli resolution get <resolution-id> [--config <path>]
```

## Screens

### List Screens
```bash
jiracli screen list [--start-at <num>] [--max-results <num>] [--query-string <query>] [--config <path>]
```

### Get Screen
```bash
jiracli screen get <screen-id> [--config <path>]
```

### Create Screen (Admin)
```bash
jiracli screen create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "Custom Screen",
  "description": "Screen for custom fields"
}
```

### Update Screen (Admin)
```bash
jiracli screen update <screen-id> --json <json-payload> [--config <path>]
```

### Delete Screen (Admin)
```bash
jiracli screen delete <screen-id> [--config <path>]
```

### Get Screen Tabs
```bash
jiracli screen tabs <screen-id> [--config <path>]
```

### Create Tab
```bash
jiracli screen tab create <screen-id> --json '{"name": "Custom Tab"}' [--config <path>]
```

### Update Tab
```bash
jiracli screen tab update <screen-id> <tab-id> --json '{"name": "Updated Tab"}' [--config <path>]
```

### Delete Tab
```bash
jiracli screen tab delete <screen-id> <tab-id> [--config <path>]
```

### Move Tab
```bash
jiracli screen tab move <screen-id> <tab-id> <position> [--config <path>]
```

### Get Tab Fields
```bash
jiracli screen tab fields <screen-id> <tab-id> [--config <path>]
```

### Add Field to Tab
```bash
jiracli screen tab field add <screen-id> <tab-id> --json '{"fieldId": "customfield_10001"}' [--config <path>]
```

### Remove Field from Tab
```bash
jiracli screen tab field remove <screen-id> <tab-id> <field-id> [--config <path>]
```

### Move Tab Field
```bash
jiracli screen tab field move <screen-id> <tab-id> <field-id> <position> [--config <path>]
```

### Get Available Fields
```bash
jiracli screen tab fields available <screen-id> <tab-id> [--config <path>]
```

## Dashboards

### List Dashboards
```bash
jiracli dashboard list [--start-at <num>] [--max-results <num>] [--filter <filter>] [--config <path>]
```

**Flags:**
- `--filter <filter>`: Filter type (favourite, my, shared)

### Get Dashboard
```bash
jiracli dashboard get <dashboard-id> [--config <path>]
```

### Create Dashboard
```bash
jiracli dashboard create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "My Dashboard",
  "description": "Dashboard for tracking progress",
  "sharePermissions": [{"type": "global"}]
}
```

### Update Dashboard
```bash
jiracli dashboard update <dashboard-id> --json <json-payload> [--config <path>]
```

### Delete Dashboard
```bash
jiracli dashboard delete <dashboard-id> [--config <path>]
```

### Get Dashboard Gadgets
```bash
jiracli dashboard gadgets <dashboard-id> [--config <path>]
```

### Add Gadget
```bash
jiracli dashboard gadget add <dashboard-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "module": "com.atlassian.jira.gadgets:assignedtome",
  "color": "blue",
  "position": {"row": 0, "column": 0},
  "title": "My Issues",
  "properties": {"numIssues": 10, "filterId": "10000"}
}
```

### Remove Gadget
```bash
jiracli dashboard gadget remove <dashboard-id> <gadget-id> [--config <path>]
```

### Update Gadget
```bash
jiracli dashboard gadget update <dashboard-id> <gadget-id> --json <json-payload> [--config <path>]
```
