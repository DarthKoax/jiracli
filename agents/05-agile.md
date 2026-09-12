# Jira CLI - Agile (Boards & Sprints)

## Boards

### List Boards
```bash
jiracli board list [--start-at <num>] [--max-results <num>] [--type <type>] [--name <name>] [--project-key <key>] [--config <path>]
```

**Flags:**
- `--type <type>`: Board type (scrum, kanban)
- `--name <name>`: Board name filter
- `--project-key <key>`: Project key filter

**Example:**
```bash
jiracli board list
jiracli board list --type scrum --project-key PROJ
```

**Output:**
```json
{
  "maxResults": 50,
  "startAt": 0,
  "total": 5,
  "isLast": true,
  "values": [
    {
      "id": 1,
      "self": "https://jira.example.com/rest/agile/1.0/board/1",
      "name": "My Scrum Board",
      "type": "scrum"
    }
  ]
}
```

### Get Board
```bash
jiracli board get <board-id> [--config <path>]
```

### Create Board
```bash
jiracli board create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "My Scrum Board",
  "type": "scrum",
  "filterId": 10000,
  "location": {
    "type": "project",
    "projectKey": "PROJ"
  }
}
```

### Delete Board
```bash
jiracli board delete <board-id> [--config <path>]
```

### Get Board Configuration
```bash
jiracli board config <board-id> [--config <path>]
```

**Output:**
```json
{
  "id": 1,
  "name": "My Scrum Board",
  "filter": {"id": "10000"},
  "columnConfig": {
    "columns": [
      {"name": "To Do", "statuses": [{"id": "10000"}]},
      {"name": "In Progress", "statuses": [{"id": "10001"}]},
      {"name": "Done", "statuses": [{"id": "10002"}]}
    ]
  },
  "estimation": {"type": "field", "field": {"fieldId": "storyPoints"}}
}
```

### Get Board Issues
```bash
jiracli board issues <board-id> [--start-at <num>] [--max-results <num>] [--jql <jql>] [--fields <list>] [--config <path>]
```

**Example:**
```bash
jiracli board issues 1
jiracli board issues 1 --jql "sprint in openSprints()"
jiracli board issues 1 --jql "assignee = john.doe" --fields "summary,status"
```

### Get Board Epics
```bash
jiracli board epics <board-id> [--start-at <num>] [--max-results <num>] [--done <bool>] [--config <path>]
```

**Example:**
```bash
jiracli board epics 1
jiracli board epics 1 --done false
```

### Get Board Sprints
```bash
jiracli board sprints <board-id> [--start-at <num>] [--max-results <num>] [--state <state>] [--config <path>]
```

**Flags:**
- `--state <state>`: Sprint state (active, closed, future)

**Example:**
```bash
jiracli board sprints 1
jiracli board sprints 1 --state active
```

### Get Board Backlog
```bash
jiracli board backlog <board-id> [--start-at <num>] [--max-results <num>] [--jql <jql>] [--fields <list>] [--config <path>]
```

**Example:**
```bash
jiracli board backlog 1
jiracli board backlog 1 --jql "type = Story"
```

## Sprints

### Get Sprint
```bash
jiracli sprint get <sprint-id> [--config <path>]
```

**Output:**
```json
{
  "id": 1,
  "self": "https://jira.example.com/rest/agile/1.0/sprint/1",
  "name": "Sprint 1",
  "state": "active",
  "startDate": "2024-01-15T00:00:00.000Z",
  "endDate": "2024-01-29T00:00:00.000Z",
  "originBoardId": 1,
  "goal": "Complete login feature"
}
```

### Create Sprint
```bash
jiracli sprint create --board-id <board-id> --json <json-payload> [--config <path>]
```

**Flags:**
- `--board-id <board-id>`: Board ID (required)

**JSON Payload:**
```json
{
  "name": "Sprint 1",
  "startDate": "2024-01-15T00:00:00.000Z",
  "endDate": "2024-01-29T00:00:00.000Z",
  "goal": "Complete login feature",
  "originBoardId": 1
}
```

**Example:**
```bash
jiracli sprint create --board-id 1 --json '{
  "name": "Sprint 1",
  "startDate": "2024-01-15T00:00:00.000Z",
  "endDate": "2024-01-29T00:00:00.000Z"
}'
```

### Update Sprint
```bash
jiracli sprint update <sprint-id> --json <json-payload> [--config <path>]
```

**Example:**
```bash
jiracli sprint update 1 --json '{
  "name": "Sprint 1 - Extended",
  "endDate": "2024-02-05T00:00:00.000Z",
  "goal": "Updated sprint goal"
}'
```

### Delete Sprint
```bash
jiracli sprint delete <sprint-id> [--config <path>]
```

### Get Sprint Issues
```bash
jiracli sprint issues <sprint-id> [--start-at <num>] [--max-results <num>] [--jql <jql>] [--fields <list>] [--config <path>]
```

**Example:**
```bash
jiracli sprint issues 1
jiracli sprint issues 1 --jql "type = Bug"
```

### Move Issues to Sprint
```bash
jiracli sprint move <sprint-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"issues": ["PROJ-1", "PROJ-2", "PROJ-3"]}
```

**Example:**
```bash
jiracli sprint move 1 --json '{"issues": ["PROJ-1", "PROJ-2"]}'
```

### Swap Issues Between Sprints
```bash
jiracli sprint swap <sprint-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"issues": ["PROJ-1", "PROJ-2"]}
```

### Complete Sprint
```bash
jiracli sprint complete <sprint-id> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"completeDate": "2024-01-29T18:00:00.000Z"}
```

**Example:**
```bash
jiracli sprint complete 1 --json '{"completeDate": "2024-01-29T18:00:00.000Z"}'
```
