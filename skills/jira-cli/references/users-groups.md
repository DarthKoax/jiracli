# Jira CLI - Users & Groups

## Users

### Get User
```bash
jiracli user get --username <username> [--expand <list>] [--config <path>]
```

**Flags:**
- `--username <username>`: Username (required)
- `--expand <list>`: Comma-separated list of properties to expand (groups, applicationRoles)

**Example:**
```bash
jiracli user get --username john.doe
jiracli user get --username john.doe --expand groups,applicationRoles
```

**Output:**
```json
{
  "self": "https://jira.example.com/rest/api/2/user?username=john.doe",
  "key": "john.doe",
  "name": "john.doe",
  "emailAddress": "john.doe@example.com",
  "displayName": "John Doe",
  "active": true,
  "timeZone": "America/New_York"
}
```

### Get User by Account ID
```bash
jiracli user get --account-id <account-id> [--config <path>]
```

**Example:**
```bash
jiracli user get --account-id 5b10ac8d82e05b22cc7d4ef5
```

### Search Users
```bash
jiracli user search --username <query> [--start-at <num>] [--max-results <num>] [--config <path>]
```

**Flags:**
- `--username <query>`: Search query (required)
- `--start-at <num>`: Index of first result (default: 0)
- `--max-results <num>`: Maximum number of results (default: 50)

**Example:**
```bash
jiracli user search --username john
jiracli user search --username john --max-results 10
```

**Output:**
```json
[
  {
    "name": "john.doe",
    "displayName": "John Doe",
    "emailAddress": "john.doe@example.com",
    "active": true
  }
]
```

### Find Users for Picker
```bash
jiracli user picker --query <query> [--max-results <num>] [--config <path>]
```

**Flags:**
- `--query <query>`: Search query (required)
- `--max-results <num>`: Maximum number of results (default: 50)

**Example:**
```bash
jiracli user picker --query john --max-results 10
```

### Create User (Admin)
```bash
jiracli user create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{
  "name": "new.user",
  "password": "SecurePass123!",
  "emailAddress": "new.user@example.com",
  "displayName": "New User",
  "notification": "true"
}
```

**Example:**
```bash
jiracli user create --json '{
  "name": "new.user",
  "password": "SecurePass123!",
  "emailAddress": "new.user@example.com",
  "displayName": "New User"
}'
```

**Note:** Requires admin_mode = true in config.

### Delete User (Admin)
```bash
jiracli user delete --username <username> [--config <path>]
```

**Example:**
```bash
jiracli user delete --username old.user
```

**Note:** Requires admin_mode = true in config.

### Get User Columns
```bash
jiracli user columns [--config <path>]
```

**Output:**
```json
[
  {"id": "issuekey"},
  {"id": "summary"},
  {"id": "status"},
  {"id": "assignee"}
]
```

### Set User Columns
```bash
jiracli user columns set --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
[
  {"id": "issuekey"},
  {"id": "summary"},
  {"id": "status"},
  {"id": "assignee"}
]
```

**Example:**
```bash
jiracli user columns set --json '[
  {"id": "issuekey"},
  {"id": "summary"},
  {"id": "status"}
]'
```

### Reset User Columns
```bash
jiracli user columns reset [--config <path>]
```

## Groups

### Get Group
```bash
jiracli group get --groupname <groupname> [--config <path>]
```

**Flags:**
- `--groupname <groupname>`: Group name (required)

**Example:**
```bash
jiracli group get --groupname developers
```

**Output:**
```json
{
  "name": "developers",
  "self": "https://jira.example.com/rest/api/2/group?groupname=developers"
}
```

### Create Group (Admin)
```bash
jiracli group create --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"name": "new-group"}
```

**Example:**
```bash
jiracli group create --json '{"name": "new-group"}'
```

**Note:** Requires admin_mode = true in config.

### Delete Group (Admin)
```bash
jiracli group delete --groupname <groupname> [--config <path>]
```

**Example:**
```bash
jiracli group delete --groupname old-group
```

**Note:** Requires admin_mode = true in config.

### Get Group Members
```bash
jiracli group members --groupname <groupname> [--start-at <num>] [--max-results <num>] [--config <path>]
```

**Flags:**
- `--groupname <groupname>`: Group name (required)
- `--start-at <num>`: Index of first result (default: 0)
- `--max-results <num>`: Maximum number of results (default: 50)

**Example:**
```bash
jiracli group members --groupname developers
jiracli group members --groupname developers --max-results 100
```

**Output:**
```json
{
  "size": 3,
  "max-results": 50,
  "start-index": 0,
  "total-size": 3,
  "items": [
    {
      "name": "john.doe",
      "displayName": "John Doe",
      "emailAddress": "john.doe@example.com",
      "active": true
    }
  ]
}
```

### Add User to Group
```bash
jiracli group user add --groupname <groupname> --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
{"name": "john.doe"}
```

**Example:**
```bash
jiracli group user add --groupname developers --json '{"name": "john.doe"}'
```

### Remove User from Group
```bash
jiracli group user remove --groupname <groupname> --username <username> [--config <path>]
```

**Example:**
```bash
jiracli group user remove --groupname developers --username john.doe
```

### Search Groups
```bash
jiracli group search --query <query> [--max-results <num>] [--config <path>]
```

**Flags:**
- `--query <query>`: Search query (required)
- `--max-results <num>`: Maximum number of results (default: 50)

**Example:**
```bash
jiracli group search --query dev
jiracli group search --query dev --max-results 10
```

**Output:**
```json
[
  {
    "name": "developers",
    "self": "https://jira.example.com/rest/api/2/group?groupname=developers"
  }
]
```

### Bulk Get Groups
```bash
jiracli group bulk --json <json-payload> [--config <path>]
```

**JSON Payload:**
```json
["developers", "testers", "admins"]
```

**Example:**
```bash
jiracli group bulk --json '["developers", "testers"]'
```

**Output:**
```json
[
  {"name": "developers"},
  {"name": "testers"}
]
```
