# Jira CLI Agent Skill

## Overview

This skill teaches you how to use the `jiracli` binary to interact with Jira Data Center / Server.

## Critical Rules

**NEVER read, modify, or reference the `config.toml` file.**
**NEVER discuss environment variables or configuration settings.**
**NEVER explain how the CLI is configured internally.**

You only know HOW to use the tool, not HOW it works or HOW to configure it.

## Command Structure

```bash
jiracli <command> <action> [flags]
```

All commands support hierarchical help:
- `jiracli help` - Root help
- `jiracli <command> help` - Command help
- `jiracli <command> <action> help` - Action help

## Error Handling

When operations are blocked, the CLI returns descriptive errors:
- Endpoint disabled: "X endpoint is disabled in configuration..."
- Method disabled: "X method is disabled in configuration..."
- Admin required: "X requires admin mode to be enabled..."

**When blocked:**
1. Do NOT retry or work around
2. Inform the user the operation is blocked
3. Provide the full command with JSON payload for manual execution
4. Ask for guidance

## Topic Files

Detailed command reference is split into topic files:
- `01-setup.md` - Setup and connection
- `02-issues.md` - Issue management
- `03-search.md` - Search and filters
- `04-projects.md` - Project management
- `05-agile.md` - Boards and sprints
- `06-users-groups.md` - User and group management
- `07-admin.md` - Admin operations
- `glossary.md` - Jira terminology

Load only the topic files you need for your current task.

## Output Format

All commands return JSON output. Use standard JSON parsing tools (jq, etc.) to extract specific fields.

## Best Practices

1. **Use help commands** - Check `jiracli <command> help <action>` for detailed usage
2. **Validate inputs** - Ensure issue keys, project keys, and IDs are correct
3. **Handle pagination** - Use startAt and maxResults for large result sets
4. **Check transitions** - Get available transitions before transitioning an issue
5. **Use JQL** - Leverage Jira Query Language for complex searches
