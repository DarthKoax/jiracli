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

**IMPORTANT:** The sprint must be in "active" state before completing. You cannot transition directly from "future" to "closed". First start the sprint:
```bash
# Start the sprint
jiracli sprint update 1 --json '{"state":"active","startDate":"2024-01-15T00:00:00.000Z","endDate":"2024-01-29T00:00:00.000Z"}'

# Then complete it
jiracli sprint complete 1 --json '{"completeDate": "2024-01-29T18:00:00.000Z"}'
```

## Common Patterns

### Find all active sprints across all boards
```bash
# List all boards first
jiracli board list

# Then check each board for active sprints
jiracli board sprints 1 --state active
jiracli board sprints 2 --state active
```

### Find issues in a specific sprint
```bash
# Using sprint issues command
jiracli sprint issues <sprint-id>

# Or using JQL search
jiracli search --jql "sprint = <sprint-id>"
```

### Find backlog issues for a board
```bash
jiracli board backlog <board-id>
```
Returns issues not assigned to any active/closed sprint.

### Filter board issues by JQL
```bash
jiracli board issues <board-id> --jql "status = 'To Do'"
jiracli board issues <board-id> --jql "assignee = currentUser()"
```
JQL is applied on top of the board's filter.

### Get sprint report data
No dedicated sprint report command. Approximate with:
```bash
# Get sprint metadata (dates, state, goal)
jiracli sprint get <sprint-id>

# Get all issues in sprint
jiracli sprint issues <sprint-id> --fields "summary,status,assignee,priority"
```

## Gotchas

| Issue | Detail |
|-------|--------|
| `board issues` / `sprint issues` / `board backlog` use `issues[]` key | Other list endpoints use `values[]` |
| `sprint complete` requires active state | Cannot transition from "future" to "closed" directly |
| `sprint update` requires all fields when changing state | Must include `startDate`, `endDate`, `name` |
| `sprint create` needs both `--board-id` and `originBoardId` in JSON | Redundant but both required |
| `board create` needs `filterId` | Not obvious from the help text |
| Closed sprints can't be re-opened | `sprint update` with `state: "active"` on a closed sprint returns 400 |
| Board issues JQL is additive | Applied on top of board's saved filter, not replacing it |
