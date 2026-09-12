# Jira CLI - Setup & Connection

## Initialize Configuration

```bash
jiracli init [--dir <path>]
```

**Flags:**
- `--dir <path>`: Directory to create config in (default: ~/.config/darthkoax/jiracli)

**Description:**
Creates a default configuration file at the specified location. The config file contains placeholders for your Jira instance URL and API token.

**Example:**
```bash
jiracli init
jiracli init --dir /custom/path
```

**Output:**
```
Configuration file created at: ~/.config/darthkoax/jiracli/config.toml
Edit the file to set your Jira base_url and api_token.
```

## Test Connection

```bash
jiracli connect [--config <path>]
```

**Flags:**
- `--config <path>`: Path to config file (default: ~/.config/darthkoax/jiracli/config.toml)

**Description:**
Connects to the Jira Data Center instance using the configuration file and verifies authentication by fetching the current user's information.

**Example:**
```bash
jiracli connect
jiracli connect --config /path/to/config.toml
```

**Output:**
```
Connected to Jira DC as: John Doe (john.doe@example.com)
Jira instance: https://jira.example.com
```

**Errors:**
- Config file not found: Run `jiracli init` first
- Invalid credentials: Check API token
- Network error: Verify Jira instance is accessible

## Get Current User

```bash
jiracli myself [--expand <list>] [--config <path>]
```

**Flags:**
- `--expand <list>`: Comma-separated list of properties to expand (groups, applicationRoles)
- `--config <path>`: Path to config file

**Description:**
Retrieves information about the currently authenticated user. Maps to GET /rest/api/2/myself.

**Example:**
```bash
jiracli myself
jiracli myself --expand groups,applicationRoles
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
  "timeZone": "America/New_York",
  "groups": {
    "size": 3,
    "items": [
      {"name": "developers"},
      {"name": "users"}
    ]
  }
}
```

## Version Information

```bash
jiracli version
```

**Description:**
Prints the CLI version and GitHub repository information.

**Output:**
```
jiracli version 0.1.0
GitHub: https://github.com/darkkoax/jiracli
```
