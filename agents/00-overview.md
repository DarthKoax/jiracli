# Jira CLI Agent Skill - Overview

## Critical Rules

**NEVER read, modify, or reference the `config.toml` file.**
**NEVER discuss environment variables or configuration settings.**
**NEVER explain how the CLI is configured internally.**

You only know HOW to use the tool, not HOW it works or HOW to configure it.

## Available Commands

The CLI accepts JSON payloads via stdin or command-line arguments. All output is JSON.

### Global Flags

- `--config <path>`: Path to config file (uses default if not specified)
- `--help`: Show help information

### Command Structure

```bash
jiracli <resource> <action> [flags] [json-payload]
```

## Error Handling

When operations are blocked due to configuration restrictions, the CLI will return descriptive error messages. These messages indicate:
- Which endpoint or method is disabled
- That the operation cannot be performed
- That you should not attempt to bypass the restriction
- That you should ask the user for advice if required

**When you receive a blocked error:**
1. Do NOT attempt to retry or work around the restriction
2. Inform the user that the operation is blocked
3. Provide the full command with JSON payload so they can execute it manually if they have appropriate permissions
4. Ask the user for guidance on how to proceed

## Help System

Every command and subcommand has comprehensive help available:

```bash
jiracli help                                    # Root help
jiracli help <command>                          # Command help
jiracli <command> help <action>                 # Action help
```

Help includes:
- Command description and usage
- All available flags with descriptions
- JSON payload structure (for POST/PUT operations)
- Examples
- Expected output format
- Common errors

## Topic Files

This skill is split into multiple topic files for better context management:

- `01-setup.md` - Setup and connection verification
- `02-issues.md` - Issue management (get, create, update, delete, comments, worklogs, etc.)
- `03-search.md` - Search and filters
- `04-projects.md` - Project management
- `05-agile.md` - Boards and sprints
- `06-users-groups.md` - User and group management
- `07-admin.md` - Admin operations (reindex, permissions, etc.)

Only load the topic files you need for your current task.

## Output Format

All commands return JSON output. Use standard JSON parsing tools (jq, etc.) to extract specific fields.

## Best Practices

1. **Use help commands** - Always check `jiracli <command> help <action>` for detailed usage
2. **Validate inputs** - Ensure issue keys, project keys, and IDs are correct before operations
3. **Handle pagination** - Use startAt and maxResults for large result sets
4. **Check transitions** - Always get available transitions before attempting to transition an issue
5. **Use JQL for complex searches** - Leverage the full power of Jira Query Language
