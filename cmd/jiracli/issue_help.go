package main

import "fmt"

func printIssueHelp() {
	fmt.Print(`jiracli issue - Manage issues

USAGE:
  jiracli issue <action> [flags]
  jiracli issue help <action>

ACTIONS:
  get           Get issue details
  create        Create a new issue
  update        Update an existing issue
  delete        Delete an issue
  transitions   Get available transitions for an issue
  transition    Perform a transition on an issue
  comment       Manage comments (list, add, update, delete)
  worklog       Manage worklogs (list, add, update, delete)
  watcher       Manage watchers (list, add, remove)
  vote          Manage votes (get, add, remove)
  assign        Assign an issue to a user

EXAMPLES:
  jiracli issue get PROJ-123
  jiracli issue create --json '{"fields":{...}}'
  jiracli issue update PROJ-123 --json '{"fields":{...}}'
  jiracli issue delete PROJ-123
  jiracli issue transitions PROJ-123
  jiracli issue transition PROJ-123 21
  jiracli issue comment list PROJ-123
  jiracli issue comment add PROJ-123 --body "Working on this"
  jiracli issue worklog add PROJ-123 --json '{"timeSpent":"2h"}'
  jiracli issue watcher add PROJ-123 john.doe
  jiracli issue vote add PROJ-123
  jiracli issue assign PROJ-123 john.doe

For more information about an action, run:
  jiracli issue help <action>
`)
}

func printIssueGetHelp() {
	fmt.Print(`jiracli issue get - Get issue details

USAGE:
  jiracli issue get <issue-key> [flags]

DESCRIPTION:
  Retrieves detailed information about a Jira issue including all fields,
  comments, worklogs, and other metadata.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getIssue

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --fields <list>      Comma-separated list of fields to return
                       Default: all fields (*all)
                       Example: --fields "summary,status,assignee"

  --expand <list>      Comma-separated list of properties to expand
                       Options: renderedFields, names, schema, transitions,
                               operations, editmeta, changelog
                       Example: --expand "renderedFields,transitions"

  --config <path>      Path to config file
                       Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue get PROJ-123
  jiracli issue get PROJ-123 --fields "summary,status,assignee,priority"
  jiracli issue get PROJ-123 --expand "renderedFields,changelog"
  jiracli issue get PROJ-123 --fields "summary" --expand "transitions"

OUTPUT:
  Returns JSON object with issue details:
  {
    "id": "10001",
    "key": "PROJ-123",
    "self": "https://jira.example.com/rest/api/2/issue/10001",
    "fields": {
      "summary": "Issue summary",
      "status": {"name": "Open"},
      "assignee": {"name": "john.doe", "displayName": "John Doe"},
      ...
    }
  }

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueCreateHelp() {
	fmt.Print(`jiracli issue create - Create a new issue

USAGE:
  jiracli issue create --json <json-payload> [flags]

DESCRIPTION:
  Creates a new issue in Jira with the specified fields. The JSON payload
  must include at minimum: project, summary, and issuetype.

  Maps to: POST /rest/api/2/issue
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-createIssue

FLAGS:
  --json <payload>    JSON object with issue fields (required)
                      Must include: fields.project, fields.summary, fields.issuetype
                      See JSON structure below

  --config <path>     Path to config file
                      Default: ~/.config/jiracli/config.toml

JSON PAYLOAD STRUCTURE:
  {
    "fields": {
      "project": {"key": "PROJ"},              // Required: project key
      "summary": "Issue summary",              // Required: issue summary
      "description": "Issue description",      // Optional: description
      "issuetype": {"name": "Bug"},            // Required: issue type name or ID
      "priority": {"name": "High"},            // Optional: priority name or ID
      "labels": ["label1", "label2"],          // Optional: labels
      "components": [{"name": "Backend"}],     // Optional: components
      "versions": [{"name": "1.0"}],           // Optional: affected versions
      "fixVersions": [{"name": "1.1"}],        // Optional: fix versions
      "assignee": {"name": "john.doe"},        // Optional: assignee username
      "reporter": {"name": "jane.smith"},      // Optional: reporter username
      "customfield_10001": "custom value"      // Optional: custom fields
    }
  }

EXAMPLES:
  # Create a simple bug
  jiracli issue create --json '{
    "fields": {
      "project": {"key": "PROJ"},
      "summary": "Fix login error",
      "issuetype": {"name": "Bug"},
      "priority": {"name": "High"}
    }
  }'

  # Create a story with all fields
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

OUTPUT:
  Returns JSON object with created issue:
  {
    "id": "10002",
    "key": "PROJ-124",
    "self": "https://jira.example.com/rest/api/2/issue/10002"
  }

ERRORS:
  - Missing required fields: "API error (status 400)"
  - Invalid project: "API error (status 400)"
  - Invalid issue type: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

NOTES:
  - Custom fields use the format: customfield_XXXXX
  - Use 'jiracli field list' to see available fields and their IDs
  - Use 'jiracli issuetype list' to see available issue types
  - Use 'jiracli priority list' to see available priorities
`)
}

func printIssueUpdateHelp() {
	fmt.Print(`jiracli issue update - Update an existing issue

USAGE:
  jiracli issue update <issue-key> --json <json-payload> [flags]

DESCRIPTION:
  Updates an existing issue with the specified fields. Only the fields
  included in the JSON payload will be updated.

  Maps to: PUT /rest/api/2/issue/{issueIdOrKey}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-editIssue

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --json <payload>    JSON object with fields to update (required)
                      Only include fields you want to change

  --config <path>     Path to config file
                      Default: ~/.config/jiracli/config.toml

JSON PAYLOAD STRUCTURE:
  {
    "fields": {
      "summary": "Updated summary",
      "description": "Updated description",
      "priority": {"name": "Critical"},
      "labels": ["updated", "urgent"],
      "assignee": {"name": "new.user"}
    }
  }

EXAMPLES:
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

  # Clear a field (set to null)
  jiracli issue update PROJ-123 --json '{
    "fields": {
      "assignee": null
    }
  }'

OUTPUT:
  On success: "Issue PROJ-123 updated successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Invalid field: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - PUT method disabled: "PUT method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

NOTES:
  - Some fields cannot be updated (e.g., created, creator)
  - Use 'jiracli field list' to see which fields are editable
  - To transition an issue, use 'jiracli issue transition' instead
`)
}

func printIssueDeleteHelp() {
	fmt.Print(`jiracli issue delete - Delete an issue

USAGE:
  jiracli issue delete <issue-key> [flags]

DESCRIPTION:
  Deletes an issue from Jira. This action cannot be undone.

  Maps to: DELETE /rest/api/2/issue/{issueIdOrKey}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-deleteIssue

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --delete-subtasks    Also delete subtasks of this issue
                       Default: false (subtasks are not deleted)

  --config <path>      Path to config file
                       Default: ~/.config/jiracli/config.toml

EXAMPLES:
  # Delete a simple issue
  jiracli issue delete PROJ-123

  # Delete issue and all subtasks
  jiracli issue delete PROJ-123 --delete-subtasks

OUTPUT:
  On success: "Issue PROJ-123 deleted successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Issue has subtasks: "API error (status 400)" (use --delete-subtasks)
  - Permission denied: "API error (status 403)"
  - DELETE method disabled: "DELETE method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

WARNING:
  This action cannot be undone. Use with caution.
`)
}

func printIssueTransitionsHelp() {
	fmt.Print(`jiracli issue transitions - Get available transitions

USAGE:
  jiracli issue transitions <issue-key> [flags]

DESCRIPTION:
  Retrieves the list of available transitions for an issue. Each transition
  has an ID that can be used with 'jiracli issue transition'.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}/transitions
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getTransitions

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue transitions PROJ-123

OUTPUT:
  Returns JSON array of transitions:
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

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

NOTES:
  - Use the transition ID with 'jiracli issue transition'
  - Available transitions depend on current status and permissions
  - Some transitions may require additional fields (e.g., resolution)
`)
}

func printIssueTransitionHelp() {
	fmt.Print(`jiracli issue transition - Perform a transition

USAGE:
  jiracli issue transition <issue-key> <transition-id> [flags]

DESCRIPTION:
  Performs a workflow transition on an issue, changing its status.

  Maps to: POST /rest/api/2/issue/{issueIdOrKey}/transitions
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-doTransition

ARGUMENTS:
  <issue-key>       The issue key (e.g., PROJ-123) or issue ID
  <transition-id>   The transition ID (from 'jiracli issue transitions')

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/jiracli/config.toml

EXAMPLES:
  # Transition to "In Progress"
  jiracli issue transition PROJ-123 11

  # Transition to "Done"
  jiracli issue transition PROJ-123 21

OUTPUT:
  On success: "Issue PROJ-123 transitioned successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Invalid transition: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - Missing required fields: "API error (status 400)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

NOTES:
  - Use 'jiracli issue transitions' to get available transition IDs
  - Some transitions require additional fields (e.g., resolution)
  - The transition must be valid for the current status
`)
}

func printIssueCommentHelp() {
	fmt.Print(`jiracli issue comment - Manage comments

USAGE:
  jiracli issue comment <action> [flags]

ACTIONS:
  list      List all comments on an issue
  add       Add a comment to an issue
  update    Update an existing comment
  delete    Delete a comment

EXAMPLES:
  jiracli issue comment list PROJ-123
  jiracli issue comment add PROJ-123 --body "Working on this"
  jiracli issue comment update PROJ-123 10001 --body "Updated comment"
  jiracli issue comment delete PROJ-123 10001

For more information about an action, run:
  jiracli issue comment help <action>
`)
}

func printIssueWorklogHelp() {
	fmt.Print(`jiracli issue worklog - Manage worklogs

USAGE:
  jiracli issue worklog <action> [flags]

ACTIONS:
  list      List all worklogs on an issue
  add       Add a worklog entry to an issue
  update    Update an existing worklog entry
  delete    Delete a worklog entry

EXAMPLES:
  jiracli issue worklog list PROJ-123
  jiracli issue worklog add PROJ-123 --json '{"timeSpent":"2h","comment":"Implementation"}'
  jiracli issue worklog update PROJ-123 10001 --json '{"timeSpent":"3h"}'
  jiracli issue worklog delete PROJ-123 10001

For more information about an action, run:
  jiracli issue worklog help <action>
`)
}

func printIssueWatcherHelp() {
	fmt.Print(`jiracli issue watcher - Manage watchers

USAGE:
  jiracli issue watcher <action> [flags]

ACTIONS:
  list      List all watchers of an issue
  add       Add a watcher to an issue
  remove    Remove a watcher from an issue

EXAMPLES:
  jiracli issue watcher list PROJ-123
  jiracli issue watcher add PROJ-123 john.doe
  jiracli issue watcher remove PROJ-123 john.doe

For more information about an action, run:
  jiracli issue watcher help <action>
`)
}

func printIssueVoteHelp() {
	fmt.Print(`jiracli issue vote - Manage votes

USAGE:
  jiracli issue vote <action> [flags]

ACTIONS:
  get       Get vote information for an issue
  add       Vote for an issue
  remove    Remove vote from an issue

EXAMPLES:
  jiracli issue vote get PROJ-123
  jiracli issue vote add PROJ-123
  jiracli issue vote remove PROJ-123

For more information about an action, run:
  jiracli issue vote help <action>
`)
}

func printIssueAssignHelp() {
	fmt.Print(`jiracli issue assign - Assign an issue

USAGE:
  jiracli issue assign <issue-key> <username> [flags]

DESCRIPTION:
  Assigns an issue to a user.

  Maps to: PUT /rest/api/2/issue/{issueIdOrKey}/assignee
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-assign

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <username>     The username to assign the issue to

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue assign PROJ-123 john.doe

OUTPUT:
  On success: "Issue PROJ-123 assigned to john.doe"

ERRORS:
  - Issue not found: "API error (status 404)"
  - User not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - PUT method disabled: "PUT method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

NOTES:
  - Use 'jiracli user search' to find valid usernames
  - Set to empty string to unassign: jiracli issue assign PROJ-123 ""
`)
}

// Comment sub-action help
func printIssueCommentListHelp() {
	fmt.Print(`jiracli issue comment list - List all comments on an issue

USAGE:
  jiracli issue comment list <issue-key> [flags]

DESCRIPTION:
  Retrieves all comments for a specific issue.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}/comment
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getComments

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue comment list PROJ-123
  jiracli issue comment list 10001

OUTPUT:
  Returns JSON array of comments:
  [
    {
      "id": "10001",
      "body": "This is a comment",
      "author": {"name": "john.doe", "displayName": "John Doe"},
      "created": "2024-01-15T10:00:00.000+0000",
      "updated": "2024-01-15T10:00:00.000+0000"
    }
  ]

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueCommentAddHelp() {
	fmt.Print(`jiracli issue comment add - Add a comment to an issue

USAGE:
  jiracli issue comment add <issue-key> --body <text> [flags]

DESCRIPTION:
  Adds a new comment to an issue.

  Maps to: POST /rest/api/2/issue/{issueIdOrKey}/comment
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-addComment

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --body <text>    Comment text (required)
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue comment add PROJ-123 --body "This is a comment"
  jiracli issue comment add 10001 --body "Working on this issue"

OUTPUT:
  Returns JSON object with created comment:
  {
    "id": "10001",
    "body": "This is a comment",
    "author": {"name": "john.doe", "displayName": "John Doe"},
    "created": "2024-01-15T10:00:00.000+0000"
  }

ERRORS:
  - Issue not found: "API error (status 404)"
  - Missing body: "Error: --body flag is required"
  - Permission denied: "API error (status 403)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueCommentUpdateHelp() {
	fmt.Print(`jiracli issue comment update - Update an existing comment

USAGE:
  jiracli issue comment update <issue-key> <comment-id> --body <text> [flags]

DESCRIPTION:
  Updates an existing comment on an issue.

  Maps to: PUT /rest/api/2/issue/{issueIdOrKey}/comment/{id}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-updateComment

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <comment-id>   The comment ID to update

FLAGS:
  --body <text>    New comment text (required)
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue comment update PROJ-123 10001 --body "Updated comment text"

OUTPUT:
  On success: "Comment 10001 updated successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Comment not found: "API error (status 404)"
  - Missing body: "Error: --body flag is required"
  - Permission denied: "API error (status 403)"
  - PUT method disabled: "PUT method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueCommentDeleteHelp() {
	fmt.Print(`jiracli issue comment delete - Delete a comment

USAGE:
  jiracli issue comment delete <issue-key> <comment-id> [flags]

DESCRIPTION:
  Deletes a comment from an issue.

  Maps to: DELETE /rest/api/2/issue/{issueIdOrKey}/comment/{id}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-deleteComment

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <comment-id>   The comment ID to delete

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue comment delete PROJ-123 10001

OUTPUT:
  On success: "Comment 10001 deleted successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Comment not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - DELETE method disabled: "DELETE method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

WARNING:
  This action cannot be undone.
`)
}

// Worklog sub-action help
func printIssueWorklogListHelp() {
	fmt.Print(`jiracli issue worklog list - List all worklogs on an issue

USAGE:
  jiracli issue worklog list <issue-key> [flags]

DESCRIPTION:
  Retrieves all worklogs for a specific issue.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}/worklog
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getWorklogs

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue worklog list PROJ-123

OUTPUT:
  Returns JSON array of worklogs:
  [
    {
      "id": "10001",
      "timeSpent": "2h",
      "timeSpentSeconds": 7200,
      "comment": "Implementation work",
      "started": "2024-01-15T10:00:00.000+0000",
      "author": {"name": "john.doe", "displayName": "John Doe"}
    }
  ]

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueWorklogAddHelp() {
	fmt.Print(`jiracli issue worklog add - Add a worklog entry to an issue

USAGE:
  jiracli issue worklog add <issue-key> --json <json-payload> [flags]

DESCRIPTION:
  Adds a new worklog entry to an issue.

  Maps to: POST /rest/api/2/issue/{issueIdOrKey}/worklog
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-addWorklog

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --json <payload>  JSON object with worklog data (required)
  --config <path>   Path to config file
                    Default: ~/.config/jiracli/config.toml

JSON PAYLOAD:
  {
    "timeSpent": "2h 30m",
    "comment": "Implementation work",
    "started": "2024-01-15T10:00:00.000+0000"
  }

EXAMPLES:
  jiracli issue worklog add PROJ-123 --json '{
    "timeSpent": "2h",
    "comment": "Implementation work"
  }'

OUTPUT:
  Returns JSON object with created worklog:
  {
    "id": "10001",
    "timeSpent": "2h",
    "timeSpentSeconds": 7200,
    "comment": "Implementation work"
  }

ERRORS:
  - Issue not found: "API error (status 404)"
  - Invalid JSON: "Error parsing JSON: ..."
  - Missing JSON: "Error: --json flag is required"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueWorklogUpdateHelp() {
	fmt.Print(`jiracli issue worklog update - Update an existing worklog entry

USAGE:
  jiracli issue worklog update <issue-key> <worklog-id> --json <json-payload> [flags]

DESCRIPTION:
  Updates an existing worklog entry on an issue.

  Maps to: PUT /rest/api/2/issue/{issueIdOrKey}/worklog/{id}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-updateWorklog

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <worklog-id>   The worklog ID to update

FLAGS:
  --json <payload>  JSON object with updated worklog data (required)
  --config <path>   Path to config file
                    Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue worklog update PROJ-123 10001 --json '{
    "timeSpent": "3h",
    "comment": "Updated time estimate"
  }'

OUTPUT:
  On success: "Worklog 10001 updated successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Worklog not found: "API error (status 404)"
  - Invalid JSON: "Error parsing JSON: ..."
  - Missing JSON: "Error: --json flag is required"
  - PUT method disabled: "PUT method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueWorklogDeleteHelp() {
	fmt.Print(`jiracli issue worklog delete - Delete a worklog entry

USAGE:
  jiracli issue worklog delete <issue-key> <worklog-id> [flags]

DESCRIPTION:
  Deletes a worklog entry from an issue.

  Maps to: DELETE /rest/api/2/issue/{issueIdOrKey}/worklog/{id}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-deleteWorklog

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <worklog-id>   The worklog ID to delete

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue worklog delete PROJ-123 10001

OUTPUT:
  On success: "Worklog 10001 deleted successfully"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Worklog not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - DELETE method disabled: "DELETE method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"

WARNING:
  This action cannot be undone.
`)
}

// Watcher sub-action help
func printIssueWatcherListHelp() {
	fmt.Print(`jiracli issue watcher list - List all watchers of an issue

USAGE:
  jiracli issue watcher list <issue-key> [flags]

DESCRIPTION:
  Retrieves all watchers for a specific issue.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}/watchers
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getIssueWatchers

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue watcher list PROJ-123

OUTPUT:
  Returns JSON object with watchers:
  {
    "self": "https://jira.example.com/rest/api/2/issue/PROJ-123/watchers",
    "isWatching": true,
    "watchCount": 3,
    "watchers": [
      {"name": "john.doe", "displayName": "John Doe"},
      {"name": "jane.smith", "displayName": "Jane Smith"}
    ]
  }

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueWatcherAddHelp() {
	fmt.Print(`jiracli issue watcher add - Add a watcher to an issue

USAGE:
  jiracli issue watcher add <issue-key> <username> [flags]

DESCRIPTION:
  Adds a user as a watcher to an issue. The watcher will receive notifications.

  Maps to: POST /rest/api/2/issue/{issueIdOrKey}/watchers
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-addWatcher

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <username>     The username to add as watcher

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue watcher add PROJ-123 john.doe

OUTPUT:
  On success: "Watcher john.doe added to PROJ-123"

ERRORS:
  - Issue not found: "API error (status 404)"
  - User not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueWatcherRemoveHelp() {
	fmt.Print(`jiracli issue watcher remove - Remove a watcher from an issue

USAGE:
  jiracli issue watcher remove <issue-key> <username> [flags]

DESCRIPTION:
  Removes a user as a watcher from an issue.

  Maps to: DELETE /rest/api/2/issue/{issueIdOrKey}/watchers?username={username}
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-removeWatcher

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID
  <username>     The username to remove as watcher

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue watcher remove PROJ-123 john.doe

OUTPUT:
  On success: "Watcher john.doe removed from PROJ-123"

ERRORS:
  - Issue not found: "API error (status 404)"
  - User not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - DELETE method disabled: "DELETE method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

// Vote sub-action help
func printIssueVoteGetHelp() {
	fmt.Print(`jiracli issue vote get - Get vote information for an issue

USAGE:
  jiracli issue vote get <issue-key> [flags]

DESCRIPTION:
  Retrieves vote information for an issue, including vote count and whether
  the current user has voted.

  Maps to: GET /rest/api/2/issue/{issueIdOrKey}/votes
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-getVotes

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue vote get PROJ-123

OUTPUT:
  Returns JSON object with vote information:
  {
    "self": "https://jira.example.com/rest/api/2/issue/PROJ-123/votes",
    "votes": 5,
    "hasVoted": true
  }

ERRORS:
  - Issue not found: "API error (status 404)"
  - Permission denied: "API error (status 403)"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueVoteAddHelp() {
	fmt.Print(`jiracli issue vote add - Vote for an issue

USAGE:
  jiracli issue vote add <issue-key> [flags]

DESCRIPTION:
  Casts a vote for an issue.

  Maps to: POST /rest/api/2/issue/{issueIdOrKey}/votes
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-vote

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue vote add PROJ-123

OUTPUT:
  On success: "Vote added to PROJ-123"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Already voted: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}

func printIssueVoteRemoveHelp() {
	fmt.Print(`jiracli issue vote remove - Remove vote from an issue

USAGE:
  jiracli issue vote remove <issue-key> [flags]

DESCRIPTION:
  Removes a vote from an issue.

  Maps to: DELETE /rest/api/2/issue/{issueIdOrKey}/votes
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issue-unvote

ARGUMENTS:
  <issue-key>    The issue key (e.g., PROJ-123) or issue ID

FLAGS:
  --config <path>  Path to config file
                   Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli issue vote remove PROJ-123

OUTPUT:
  On success: "Vote removed from PROJ-123"

ERRORS:
  - Issue not found: "API error (status 404)"
  - Haven't voted: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - DELETE method disabled: "DELETE method is disabled in configuration"
  - Endpoint disabled: "issues endpoint is disabled in configuration"
`)
}
