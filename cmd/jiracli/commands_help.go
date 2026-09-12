package main

import "fmt"

func printProjectHelp() {
	fmt.Print(`jiracli project - Manage projects

USAGE:
  jiracli project <action> [flags]

ACTIONS:
  list        List all projects
  get         Get project details
  create      Create a new project
  update      Update a project
  delete      Delete a project
  roles       Get project roles
  components  Get project components
  versions    Get project versions

EXAMPLES:
  jiracli project list
  jiracli project get PROJ
  jiracli project create --json '{"key":"PROJ","name":"My Project"}'

For more information about an action, run:
  jiracli project help <action>

JIRA API:
  Maps to: /rest/api/2/project endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/project
`)
}

func printUserHelp() {
	fmt.Print(`jiracli user - Manage users

USAGE:
  jiracli user <action> [flags]

ACTIONS:
  get       Get user details
  search    Search for users
  create    Create a new user (admin)
  delete    Delete a user (admin)
  columns   Manage user columns

EXAMPLES:
  jiracli user get --username john.doe
  jiracli user search --query john
  jiracli user create --json '{"name":"new.user","emailAddress":"new@example.com"}'

For more information about an action, run:
  jiracli user help <action>

JIRA API:
  Maps to: /rest/api/2/user endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/user
`)
}

func printGroupHelp() {
	fmt.Print(`jiracli group - Manage groups

USAGE:
  jiracli group <action> [flags]

ACTIONS:
  get        Get group details
  create     Create a new group (admin)
  delete     Delete a group (admin)
  members    Get group members
  search     Search for groups

EXAMPLES:
  jiracli group get --groupname developers
  jiracli group create --json '{"name":"new-group"}'
  jiracli group members --groupname developers

For more information about an action, run:
  jiracli group help <action>

JIRA API:
  Maps to: /rest/api/2/group endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/group
`)
}

func printSearchHelp() {
	fmt.Print(`jiracli search - Search for issues using JQL

USAGE:
  jiracli search --jql <jql-query> [flags]

DESCRIPTION:
  Searches for issues using Jira Query Language (JQL). Returns matching
  issues with specified fields.

  Maps to: POST /rest/api/2/search
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/search-search

FLAGS:
  --jql <query>         JQL query string (required)
                        Example: "project = PROJ AND status = Open"

  --start-at <num>      Index of first result (default: 0)
                        Use for pagination

  --max-results <num>   Maximum number of results (default: 50)
                        Use for pagination

  --fields <list>       Comma-separated list of fields to return
                        Example: --fields "summary,status,assignee"

  --expand <list>       Comma-separated list of properties to expand
                        Options: schema, names, operations, editmeta, changelog

  --config <path>       Path to config file
                        Default: ~/.config/darthkoax/jiracli/config.toml

EXAMPLES:
  # Simple search
  jiracli search --jql "project = PROJ"

  # Search with status filter
  jiracli search --jql "project = PROJ AND status = Open"

  # Search with pagination
  jiracli search --jql "project = PROJ" --start-at 0 --max-results 20

  # Search with specific fields
  jiracli search --jql "assignee = john.doe" --fields "summary,status,priority"

  # Complex JQL query
  jiracli search --jql "project = PROJ AND status != Closed AND priority in (Highest, High) ORDER BY created DESC"

OUTPUT:
  Returns JSON object with search results:
  {
    "startAt": 0,
    "maxResults": 50,
    "total": 123,
    "issues": [
      {
        "id": "10001",
        "key": "PROJ-123",
        "fields": {
          "summary": "Issue summary",
          "status": {"name": "Open"},
          ...
        }
      },
      ...
    ]
  }

ERRORS:
  - Invalid JQL: "API error (status 400)"
  - Permission denied: "API error (status 403)"
  - POST method disabled: "POST method is disabled in configuration"
  - Endpoint disabled: "search endpoint is disabled in configuration"

JQL REFERENCE:
  Common JQL operators:
  - = (equals)
  - != (not equals)
  - > < >= <= (comparison)
  - IN (in list)
  - NOT IN (not in list)
  - ~ (contains)
  - IS / IS NOT (null checks)

  Common JQL fields:
  - project, status, assignee, reporter, priority, issuetype
  - created, updated, due, resolution
  - labels, components, versions, fixVersions
  - summary, description, comment

  See Jira documentation for full JQL reference.
`)
}

func printFilterHelp() {
	fmt.Print(`jiracli filter - Manage filters (saved searches)

USAGE:
  jiracli filter <action> [flags]

ACTIONS:
  get          Get filter details
  create       Create a new filter
  update       Update a filter
  delete       Delete a filter
  list         List favourite filters
  my           List my filters
  search       Search for filters

EXAMPLES:
  jiracli filter get 10000
  jiracli filter create --json '{"name":"My Filter","jql":"project=PROJ"}'
  jiracli filter list

For more information about an action, run:
  jiracli filter help <action>

JIRA API:
  Maps to: /rest/api/2/filter endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/filter
`)
}

func printDashboardHelp() {
	fmt.Print(`jiracli dashboard - Manage dashboards

USAGE:
  jiracli dashboard <action> [flags]

ACTIONS:
  list      List all dashboards
  get       Get dashboard details
  create    Create a new dashboard
  update    Update a dashboard
  delete    Delete a dashboard
  gadgets   Manage dashboard gadgets

EXAMPLES:
  jiracli dashboard list
  jiracli dashboard get 10000
  jiracli dashboard create --json '{"name":"My Dashboard"}'

For more information about an action, run:
  jiracli dashboard help <action>

JIRA API:
  Maps to: /rest/api/2/dashboard endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/dashboard
`)
}

func printBoardHelp() {
	fmt.Print(`jiracli board - Manage agile boards

USAGE:
  jiracli board <action> [flags]

ACTIONS:
  list        List all boards
  get         Get board details
  create      Create a new board
  delete      Delete a board
  config      Get board configuration
  issues      Get board issues
  epics       Get board epics
  sprints     Get board sprints
  backlog     Get board backlog

EXAMPLES:
  jiracli board list
  jiracli board get 1
  jiracli board issues 1 --jql "sprint in openSprints()"

For more information about an action, run:
  jiracli board help <action>

JIRA API:
  Maps to: /rest/agile/1.0/board endpoints
  https://docs.atlassian.com/jira-software/REST/8.20.0/#agile/1.0/board
`)
}

func printSprintHelp() {
	fmt.Print(`jiracli sprint - Manage sprints

USAGE:
  jiracli sprint <action> [flags]

ACTIONS:
  get         Get sprint details
  create      Create a new sprint
  update      Update a sprint
  delete      Delete a sprint
  issues      Get sprint issues
  move        Move issues to sprint
  complete    Complete a sprint

EXAMPLES:
  jiracli sprint get 1
  jiracli sprint create --json '{"name":"Sprint 1","originBoardId":1}'
  jiracli sprint issues 1

For more information about an action, run:
  jiracli sprint help <action>

JIRA API:
  Maps to: /rest/agile/1.0/sprint endpoints
  https://docs.atlassian.com/jira-software/REST/8.20.0/#agile/1.0/sprint
`)
}

func printFieldHelp() {
	fmt.Print(`jiracli field - Manage issue fields

USAGE:
  jiracli field <action> [flags]

ACTIONS:
  list      List all fields
  create    Create a custom field

EXAMPLES:
  jiracli field list
  jiracli field create --json '{"name":"Custom Field","type":"textfield"}'

For more information about an action, run:
  jiracli field help <action>

JIRA API:
  Maps to: /rest/api/2/field endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/field
`)
}

func printComponentHelp() {
	fmt.Print(`jiracli component - Manage project components

USAGE:
  jiracli component <action> [flags]

ACTIONS:
  get       Get component details
  create    Create a new component
  update    Update a component
  delete    Delete a component
  count     Get related issue count

EXAMPLES:
  jiracli component get 10000
  jiracli component create --json '{"name":"Backend","project":"PROJ"}'

For more information about an action, run:
  jiracli component help <action>

JIRA API:
  Maps to: /rest/api/2/component endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/component
`)
}

func printVersionResourceHelp() {
	fmt.Print(`jiracli version-resource - Manage project versions

USAGE:
  jiracli version-resource <action> [flags]

ACTIONS:
  get       Get version details
  create    Create a new version
  update    Update a version
  delete    Delete a version
  count     Get related issue counts
  merge     Merge versions

EXAMPLES:
  jiracli version-resource get 10000
  jiracli version-resource create --json '{"name":"1.0","project":"PROJ"}'

For more information about an action, run:
  jiracli version-resource help <action>

JIRA API:
  Maps to: /rest/api/2/version endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/version
`)
}

func printPriorityHelp() {
	fmt.Print(`jiracli priority - View priorities

USAGE:
  jiracli priority <action> [flags]

ACTIONS:
  list      List all priorities
  get       Get priority details

EXAMPLES:
  jiracli priority list
  jiracli priority get 1

For more information about an action, run:
  jiracli priority help <action>

JIRA API:
  Maps to: /rest/api/2/priority endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/priority
`)
}

func printStatusHelp() {
	fmt.Print(`jiracli status - View statuses

USAGE:
  jiracli status <action> [flags]

ACTIONS:
  list          List all statuses
  get           Get status details
  categories    List status categories
  category      Get status category details

EXAMPLES:
  jiracli status list
  jiracli status get 1
  jiracli status categories

For more information about an action, run:
  jiracli status help <action>

JIRA API:
  Maps to: /rest/api/2/status endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/status
`)
}

func printResolutionHelp() {
	fmt.Print(`jiracli resolution - View resolutions

USAGE:
  jiracli resolution <action> [flags]

ACTIONS:
  list      List all resolutions
  get       Get resolution details

EXAMPLES:
  jiracli resolution list
  jiracli resolution get 1

For more information about an action, run:
  jiracli resolution help <action>

JIRA API:
  Maps to: /rest/api/2/resolution endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/resolution
`)
}

func printIssueTypeHelp() {
	fmt.Print(`jiracli issuetype - Manage issue types

USAGE:
  jiracli issuetype <action> [flags]

ACTIONS:
  list          List all issue types
  get           Get issue type details
  create        Create a new issue type
  update        Update an issue type
  delete        Delete an issue type
  alternatives  Get alternative issue types

EXAMPLES:
  jiracli issuetype list
  jiracli issuetype get 1
  jiracli issuetype create --json '{"name":"Epic","description":"Large story"}'

For more information about an action, run:
  jiracli issuetype help <action>

JIRA API:
  Maps to: /rest/api/2/issuetype endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/issuetype
`)
}

func printPermissionHelp() {
	fmt.Print(`jiracli permission - Manage permissions

USAGE:
  jiracli permission <action> [flags]

ACTIONS:
  list      List all permissions
  my        Get my permissions
  scheme    Manage permission schemes

EXAMPLES:
  jiracli permission list
  jiracli permission my
  jiracli permission scheme list

For more information about an action, run:
  jiracli permission help <action>

JIRA API:
  Maps to: /rest/api/2/permissions endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/permissions
`)
}

func printMyselfHelp() {
	fmt.Print(`jiracli myself - Get current user information

USAGE:
  jiracli myself [flags]

DESCRIPTION:
  Retrieves information about the currently authenticated user.

  Maps to: GET /rest/api/2/myself
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/myself-getUser

FLAGS:
  --expand <list>    Comma-separated list of properties to expand
                     Options: groups, applicationRoles
                     Example: --expand "groups,applicationRoles"

  --config <path>    Path to config file
                     Default: ~/.config/darthkoax/jiracli/config.toml

EXAMPLES:
  jiracli myself
  jiracli myself --expand "groups"

OUTPUT:
  Returns JSON object with user information:
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

ERRORS:
  - Not authenticated: "API error (status 401)"
  - Endpoint disabled: "myself endpoint is disabled in configuration"
`)
}

func printScreenHelp() {
	fmt.Print(`jiracli screen - Manage screens

USAGE:
  jiracli screen <action> [flags]

ACTIONS:
  list      List all screens
  get       Get screen details
  create    Create a new screen (admin)
  update    Update a screen (admin)
  delete    Delete a screen (admin)
  tabs      Manage screen tabs

EXAMPLES:
  jiracli screen list
  jiracli screen get 1
  jiracli screen tabs 1

For more information about an action, run:
  jiracli screen help <action>

JIRA API:
  Maps to: /rest/api/2/screens endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/screens
`)
}

func printReindexHelp() {
	fmt.Print(`jiracli reindex - Trigger and monitor reindex (admin)

USAGE:
  jiracli reindex <action> [flags]

DESCRIPTION:
  Triggers a reindex of the Jira instance and monitors progress.
  Requires admin_mode = true in config.

ACTIONS:
  trigger     Trigger a reindex
  status      Get reindex status
  request     Get reindex request status

EXAMPLES:
  jiracli reindex trigger --type BACKGROUND
  jiracli reindex status --task-id 123

For more information about an action, run:
  jiracli reindex help <action>

JIRA API:
  Maps to: /rest/api/2/reindex endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/reindex

REQUIRES:
  admin_mode = true in config.toml
`)
}

func printAppPropHelp() {
	fmt.Print(`jiracli appprop - Manage application properties (admin)

USAGE:
  jiracli appprop <action> [flags]

DESCRIPTION:
  Manages Jira application properties.
  Requires admin_mode = true in config.

ACTIONS:
  list      List all properties
  set       Set a property value

EXAMPLES:
  jiracli appprop list
  jiracli appprop set jira.title --value "My Jira"

For more information about an action, run:
  jiracli appprop help <action>

JIRA API:
  Maps to: /rest/api/2/application-properties endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/application-properties

REQUIRES:
  admin_mode = true in config.toml
`)
}

func printConfigHelp() {
	fmt.Print(`jiracli config - View global configuration (admin)

USAGE:
  jiracli config <action> [flags]

DESCRIPTION:
  Retrieves Jira global configuration settings.
  Requires admin_mode = true in config.

ACTIONS:
  get       Get global configuration

EXAMPLES:
  jiracli config get

For more information about an action, run:
  jiracli config help <action>

JIRA API:
  Maps to: /rest/api/2/configuration endpoint
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/configuration

REQUIRES:
  admin_mode = true in config.toml
`)
}

func printAvatarHelp() {
	fmt.Print(`jiracli avatar - Manage avatars

USAGE:
  jiracli avatar <action> [flags]

ACTIONS:
  system      Get system avatars
  custom      Get custom avatars
  delete      Delete a custom avatar
  update      Update a custom avatar

EXAMPLES:
  jiracli avatar system project
  jiracli avatar custom project 10000

For more information about an action, run:
  jiracli avatar help <action>

JIRA API:
  Maps to: /rest/api/2/avatar endpoints
  https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/avatar
`)
}

func printMyselfLocaleHelp() {
	fmt.Print(`jiracli myself locale - Get current user locale

USAGE:
  jiracli myself locale [flags]

DESCRIPTION:
  Retrieves the locale setting for the currently authenticated user.

  Maps to: GET /rest/api/2/myself/locale
  Jira API: https://docs.atlassian.com/software/jira/docs/api/REST/8.20.0/#api/2/myself-getLocale

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/darthkoax/jiracli/config.toml

EXAMPLES:
  jiracli myself locale

OUTPUT:
  Returns JSON object with locale information:
  {
    "locale": "en_US"
  }

ERRORS:
  - Not authenticated: "API error (status 401)"
  - Endpoint disabled: "myself endpoint is disabled in configuration"
`)
}
