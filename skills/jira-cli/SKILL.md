---
name: jira-cli
description: Use when working with Jira Data Center/Server via the jiracli CLI - getting, creating, updating, deleting, transitioning, or assigning issues, adding comments/worklogs/watchers/votes, JQL search, saved filters, dashboards, projects, components, versions, fields, issue types, screens, boards, sprints, users, groups, permissions, reindex, or avatars.
---

# Jira CLI

This skill teaches you how to use the `jiracli` binary to interact with Jira Data Center / Server.

## Critical Rules

**NEVER read, modify, or reference the `config.toml` file.**
**NEVER explain how the CLI is configured internally.**

You only know HOW to use the tool, not HOW it works or HOW to configure it.

## Command Structure

```bash
jiracli <resource> <action> [flags] [json-payload]
```

The CLI accepts JSON payloads via stdin or command-line arguments.

### Global Flags

- `--config <path>`: Path to config file (uses default if not specified)
- `--help`: Show help information

## Help System

Every command and subcommand has comprehensive help available:

```bash
jiracli help                        # Root help
jiracli help <command>              # Command help
jiracli <command> help <action>     # Action help
```

Help includes:
- Command description and usage
- All available flags with descriptions
- JSON payload structure (for POST/PUT operations)
- Examples
- Expected output format
- Common errors

## Error Handling

When operations are blocked, the CLI returns descriptive errors:
- Endpoint disabled: "X endpoint is disabled in configuration..."
- Method disabled: "X method is disabled in configuration..."
- Admin required: "X requires admin mode to be enabled..."

**When blocked by method gating:**

If you encounter a "method is disabled" error, you can override the restriction by prefixing the command with the appropriate environment variable:

```bash
JIRA_ALLOW_GET=true jiracli <command>
JIRA_ALLOW_POST=true jiracli <command>
JIRA_ALLOW_PUT=true jiracli <command>
JIRA_ALLOW_DELETE=true jiracli <command>
```

Example:
```bash
# If POST is disabled and you need to create an issue:
JIRA_ALLOW_POST=true jiracli issue create --json '{"fields":{...}}'

# If DELETE is disabled and you need to delete an issue:
JIRA_ALLOW_DELETE=true jiracli issue delete PROJ-123
```

**When blocked by admin mode:**

If you encounter a "requires admin mode" error, you can override it with:
```bash
JIRA_ADMIN_MODE=true jiracli <command>
```

**When blocked:**
1. Try using the appropriate environment variable override
2. If the override doesn't work or you're unsure, inform the user the operation is blocked
3. Provide the full command with JSON payload so they can execute it manually if they have appropriate permissions
4. Ask the user for guidance on how to proceed

## Output Format

All commands return JSON output. Use standard JSON parsing tools (jq, etc.) to extract specific fields.

## Reference Files

Detailed command reference is split into topic files under `references/`.
**Load ONLY the reference file(s) needed for the current task - never all of them.**

| Task involves | Read |
|---------------|------|
| Setup, connection test, current user, version | `references/setup.md` |
| Issues: get/create/update/delete, transitions, comments, worklogs, watchers, votes, assign | `references/issues.md` |
| JQL search, saved filters, share permissions | `references/search.md` |
| Projects, components, versions, fields, issue types, priorities, statuses, resolutions, screens, dashboards | `references/projects.md` |
| Agile boards and sprints | `references/agile.md` |
| Users and groups | `references/users-groups.md` |
| Admin: permission schemes, reindex, application properties, global configuration, avatars | `references/admin.md` |
| Jira terminology or concept definitions | `references/glossary.md` |

## Best Practices

1. **Use help commands** - Check `jiracli <command> help <action>` for detailed usage
2. **Validate inputs** - Ensure issue keys, project keys, and IDs are correct before operations
3. **Handle pagination** - Use startAt and maxResults for large result sets
4. **Check transitions** - Always get available transitions before attempting to transition an issue
5. **Use JQL** - Leverage Jira Query Language for complex searches
6. **Epic queries** - Use `"Epic Link" = KEY` (with quotes), NOT custom field IDs
7. **Parent/child queries** - Use `parent = KEY` for sub-tasks; use `issuetype != Sub-task` for top-level issues (NOT `parent is EMPTY`)
8. **Sub-tasks and epics** - Sub-tasks don't have Epic Link; they inherit epic through their parent
9. **Status categories** - Use `statusCategory = Done` for broad status grouping
10. **Unresolved issues** - Use `resolution = Unresolved` to find open issues
