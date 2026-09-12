package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printRootHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help", "--help", "-h":
		if len(os.Args) > 2 {
			printCommandHelp(os.Args[2])
		} else {
			printRootHelp()
		}
	case "version", "--version", "-v":
		printVersion()
	case "init":
		handleInit(os.Args[2:])
	case "connect":
		handleConnect(os.Args[2:])
	case "issue":
		handleIssue(os.Args[2:])
	case "project":
		handleProject(os.Args[2:])
	case "user":
		handleUser(os.Args[2:])
	case "group":
		handleGroup(os.Args[2:])
	case "search":
		handleSearch(os.Args[2:])
	case "filter":
		handleFilter(os.Args[2:])
	case "dashboard":
		handleDashboard(os.Args[2:])
	case "board":
		handleBoard(os.Args[2:])
	case "sprint":
		handleSprint(os.Args[2:])
	case "field":
		handleField(os.Args[2:])
	case "component":
		handleComponent(os.Args[2:])
	case "version-resource":
		handleVersionResource(os.Args[2:])
	case "priority":
		handlePriority(os.Args[2:])
	case "status":
		handleStatus(os.Args[2:])
	case "resolution":
		handleResolution(os.Args[2:])
	case "issuetype":
		handleIssueType(os.Args[2:])
	case "permission":
		handlePermission(os.Args[2:])
	case "myself":
		handleMyself(os.Args[2:])
	case "screen":
		handleScreen(os.Args[2:])
	case "reindex":
		handleReindex(os.Args[2:])
	case "appprop":
		handleAppProp(os.Args[2:])
	case "config":
		handleConfig(os.Args[2:])
	case "avatar":
		handleAvatar(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printRootHelp()
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("jiracli version %s\n", version)
	fmt.Printf("GitHub: https://github.com/darkkoax/jiracli\n")
}

func printRootHelp() {
	fmt.Printf(`jiracli %s - Jira Data Center / Server CLI

USAGE:
  jiracli <command> [flags]
  jiracli <command> <action> [flags]
  jiracli help <command>

COMMANDS:
  Setup & Connection:
    init              Create a default configuration file
    connect           Connect to Jira and verify authentication

  Issue Management:
    issue             Manage issues (get, create, update, delete, comment, worklog, etc.)
    search            Search for issues using JQL
    filter            Manage filters (saved searches)

  Project Management:
    project           Manage projects
    component         Manage project components
    version-resource  Manage project versions

  Agile/Scrum:
    board             Manage agile boards
    sprint            Manage sprints

  User & Group Management:
    user              Manage users
    group             Manage groups

  Configuration:
    field             Manage issue fields
    issuetype         Manage issue types
    priority          View priorities
    status            View statuses
    resolution        View resolutions
    screen            Manage screens
    dashboard         Manage dashboards
    permission        Manage permissions

  System:
    myself            Get current user information
    reindex           Trigger and monitor reindex (admin)
    appprop           Manage application properties (admin)
    config            View global configuration (admin)
    avatar            Manage avatars

  Other:
    version           Print version information
    help              Show help for a command

EXAMPLES:
  jiracli init                                    # Initialize config
  jiracli connect                                 # Test connection
  jiracli issue get PROJ-123                      # Get issue details
  jiracli issue create --json '{"fields":{...}}'  # Create issue
  jiracli search --jql "project=PROJ"             # Search issues
  jiracli help issue                              # Get help for issue command

For more information about a command, run:
  jiracli help <command>
`, version)
}

func printCommandHelp(command string) {
	switch command {
	case "init":
		printInitHelp()
	case "connect":
		printConnectHelp()
	case "issue":
		printIssueHelp()
	case "project":
		printProjectHelp()
	case "user":
		printUserHelp()
	case "group":
		printGroupHelp()
	case "search":
		printSearchHelp()
	case "filter":
		printFilterHelp()
	case "dashboard":
		printDashboardHelp()
	case "board":
		printBoardHelp()
	case "sprint":
		printSprintHelp()
	case "field":
		printFieldHelp()
	case "component":
		printComponentHelp()
	case "version-resource":
		printVersionResourceHelp()
	case "priority":
		printPriorityHelp()
	case "status":
		printStatusHelp()
	case "resolution":
		printResolutionHelp()
	case "issuetype":
		printIssueTypeHelp()
	case "permission":
		printPermissionHelp()
	case "myself":
		printMyselfHelp()
	case "screen":
		printScreenHelp()
	case "reindex":
		printReindexHelp()
	case "appprop":
		printAppPropHelp()
	case "config":
		printConfigHelp()
	case "avatar":
		printAvatarHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printRootHelp()
		os.Exit(1)
	}
}
