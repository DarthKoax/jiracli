package main

import (
	"fmt"
	"os"
	"strings"
)

// Generic command handler pattern for all commands
func handleCommandWithHelp(commandName string, args []string, printHelp func(), actionHandlers map[string]func([]string), printActionHelp func(string)) {
	if len(args) == 0 {
		printHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			// Help for specific action: jiracli <command> help <action>
			printActionHelp(args[1])
		} else {
			// Help for command: jiracli <command> help
			printHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help: jiracli <command> <action> help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printActionHelp(action)
		return
	}

	if handler, ok := actionHandlers[action]; ok {
		handler(actionArgs)
	} else {
		fmt.Fprintf(os.Stderr, "Unknown %s action: %s\n\n", commandName, action)
		printHelp()
		os.Exit(1)
	}
}

// Stub handlers for all commands - to be implemented
func handleProject(args []string) {
	actionHandlers := map[string]func([]string){
		"list":       handleProjectList,
		"get":        handleProjectGet,
		"create":     handleProjectCreate,
		"update":     handleProjectUpdate,
		"delete":     handleProjectDelete,
		"roles":      handleProjectRoles,
		"components": handleProjectComponents,
		"versions":   handleProjectVersions,
	}
	handleCommandWithHelp("project", args, printProjectHelp, actionHandlers, printProjectActionHelp)
}

func printProjectActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown project action: %s\n\n", action)
		printProjectHelp()
		os.Exit(1)
	}
}

func handleUser(args []string) {
	actionHandlers := map[string]func([]string){
		"get":     handleUserGet,
		"search":  handleUserSearch,
		"create":  handleUserCreate,
		"delete":  handleUserDelete,
		"columns": handleUserColumns,
	}
	handleCommandWithHelp("user", args, printUserHelp, actionHandlers, printUserActionHelp)
}

func printUserActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown user action: %s\n\n", action)
		printUserHelp()
		os.Exit(1)
	}
}

func handleGroup(args []string) {
	actionHandlers := map[string]func([]string){
		"get":     handleGroupGet,
		"create":  handleGroupCreate,
		"delete":  handleGroupDelete,
		"members": handleGroupMembers,
		"search":  handleGroupSearch,
	}
	handleCommandWithHelp("group", args, printGroupHelp, actionHandlers, printGroupActionHelp)
}

func printGroupActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown group action: %s\n\n", action)
		printGroupHelp()
		os.Exit(1)
	}
}

func handleSearch(args []string) {
	if len(args) == 0 || (len(args) > 0 && strings.HasPrefix(args[0], "--")) {
		handleSearchGet(args)
		return
	}
	actionHandlers := map[string]func([]string){
		"get": handleSearchGet,
	}
	handleCommandWithHelp("search", args, printSearchHelp, actionHandlers, printSearchActionHelp)
}

func printSearchActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown search action: %s\n\n", action)
		printSearchHelp()
		os.Exit(1)
	}
}

func handleFilter(args []string) {
	actionHandlers := map[string]func([]string){
		"get":    handleFilterGet,
		"create": handleFilterCreate,
		"update": handleFilterUpdate,
		"delete": handleFilterDelete,
		"list":   handleFilterList,
		"my":     handleFilterMy,
		"search": handleFilterSearch,
	}
	handleCommandWithHelp("filter", args, printFilterHelp, actionHandlers, printFilterActionHelp)
}

func printFilterActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown filter action: %s\n\n", action)
		printFilterHelp()
		os.Exit(1)
	}
}

func handleDashboard(args []string) {
	actionHandlers := map[string]func([]string){
		"list":    handleDashboardList,
		"get":     handleDashboardGet,
		"create":  handleDashboardCreate,
		"update":  handleDashboardUpdate,
		"delete":  handleDashboardDelete,
		"gadgets": handleDashboardGadgets,
	}
	handleCommandWithHelp("dashboard", args, printDashboardHelp, actionHandlers, printDashboardActionHelp)
}

func printDashboardActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown dashboard action: %s\n\n", action)
		printDashboardHelp()
		os.Exit(1)
	}
}

func handleBoard(args []string) {
	actionHandlers := map[string]func([]string){
		"list":    handleBoardList,
		"get":     handleBoardGet,
		"create":  handleBoardCreate,
		"delete":  handleBoardDelete,
		"config":  handleBoardConfig,
		"issues":  handleBoardIssues,
		"epics":   handleBoardEpics,
		"sprints": handleBoardSprints,
		"backlog": handleBoardBacklog,
	}
	handleCommandWithHelp("board", args, printBoardHelp, actionHandlers, printBoardActionHelp)
}

func printBoardActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown board action: %s\n\n", action)
		printBoardHelp()
		os.Exit(1)
	}
}

func handleSprint(args []string) {
	actionHandlers := map[string]func([]string){
		"get":      handleSprintGet,
		"create":   handleSprintCreate,
		"update":   handleSprintUpdate,
		"delete":   handleSprintDelete,
		"issues":   handleSprintIssues,
		"move":     handleSprintMove,
		"complete": handleSprintComplete,
	}
	handleCommandWithHelp("sprint", args, printSprintHelp, actionHandlers, printSprintActionHelp)
}

func printSprintActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown sprint action: %s\n\n", action)
		printSprintHelp()
		os.Exit(1)
	}
}

func handleField(args []string) {
	actionHandlers := map[string]func([]string){
		"list":   handleFieldList,
		"create": handleFieldCreate,
	}
	handleCommandWithHelp("field", args, printFieldHelp, actionHandlers, printFieldActionHelp)
}

func printFieldActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown field action: %s\n\n", action)
		printFieldHelp()
		os.Exit(1)
	}
}

func handleComponent(args []string) {
	actionHandlers := map[string]func([]string){
		"get":    handleComponentGet,
		"create": handleComponentCreate,
		"update": handleComponentUpdate,
		"delete": handleComponentDelete,
		"count":  handleComponentCount,
	}
	handleCommandWithHelp("component", args, printComponentHelp, actionHandlers, printComponentActionHelp)
}

func printComponentActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown component action: %s\n\n", action)
		printComponentHelp()
		os.Exit(1)
	}
}

func handleVersionResource(args []string) {
	actionHandlers := map[string]func([]string){
		"get":    handleVersionGet,
		"create": handleVersionCreate,
		"update": handleVersionUpdate,
		"delete": handleVersionDelete,
		"count":  handleVersionCount,
		"merge":  handleVersionMerge,
	}
	handleCommandWithHelp("version-resource", args, printVersionResourceHelp, actionHandlers, printVersionResourceActionHelp)
}

func printVersionResourceActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown version-resource action: %s\n\n", action)
		printVersionResourceHelp()
		os.Exit(1)
	}
}

func handlePriority(args []string) {
	actionHandlers := map[string]func([]string){
		"list": handlePriorityList,
		"get":  handlePriorityGet,
	}
	handleCommandWithHelp("priority", args, printPriorityHelp, actionHandlers, printPriorityActionHelp)
}

func printPriorityActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown priority action: %s\n\n", action)
		printPriorityHelp()
		os.Exit(1)
	}
}

func handleStatus(args []string) {
	actionHandlers := map[string]func([]string){
		"list":       handleStatusList,
		"get":        handleStatusGet,
		"categories": handleStatusCategories,
		"category":   handleStatusCategory,
	}
	handleCommandWithHelp("status", args, printStatusHelp, actionHandlers, printStatusActionHelp)
}

func printStatusActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown status action: %s\n\n", action)
		printStatusHelp()
		os.Exit(1)
	}
}

func handleResolution(args []string) {
	actionHandlers := map[string]func([]string){
		"list": handleResolutionList,
		"get":  handleResolutionGet,
	}
	handleCommandWithHelp("resolution", args, printResolutionHelp, actionHandlers, printResolutionActionHelp)
}

func printResolutionActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown resolution action: %s\n\n", action)
		printResolutionHelp()
		os.Exit(1)
	}
}

func handleIssueType(args []string) {
	actionHandlers := map[string]func([]string){
		"list":         handleIssueTypeList,
		"get":          handleIssueTypeGet,
		"create":       handleIssueTypeCreate,
		"update":       handleIssueTypeUpdate,
		"delete":       handleIssueTypeDelete,
		"alternatives": handleIssueTypeAlternatives,
	}
	handleCommandWithHelp("issuetype", args, printIssueTypeHelp, actionHandlers, printIssueTypeActionHelp)
}

func printIssueTypeActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown issuetype action: %s\n\n", action)
		printIssueTypeHelp()
		os.Exit(1)
	}
}

func handlePermission(args []string) {
	actionHandlers := map[string]func([]string){
		"list":   handlePermissionList,
		"my":     handlePermissionMy,
		"scheme": handlePermissionScheme,
	}
	handleCommandWithHelp("permission", args, printPermissionHelp, actionHandlers, printPermissionActionHelp)
}

func printPermissionActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown permission action: %s\n\n", action)
		printPermissionHelp()
		os.Exit(1)
	}
}

func handleMyself(args []string) {
	if len(args) == 0 || (len(args) > 0 && strings.HasPrefix(args[0], "--")) {
		handleMyselfGet(args)
		return
	}
	actionHandlers := map[string]func([]string){
		"locale": handleMyselfLocale,
	}
	handleCommandWithHelp("myself", args, printMyselfHelp, actionHandlers, printMyselfActionHelp)
}

func printMyselfActionHelp(action string) {
	switch action {
	case "locale":
		printMyselfLocaleHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown myself action: %s\n\n", action)
		printMyselfHelp()
		os.Exit(1)
	}
}

func handleScreen(args []string) {
	actionHandlers := map[string]func([]string){
		"list":   handleScreenList,
		"get":    handleScreenGet,
		"create": handleScreenCreate,
		"update": handleScreenUpdate,
		"delete": handleScreenDelete,
		"tabs":   handleScreenTabs,
	}
	handleCommandWithHelp("screen", args, printScreenHelp, actionHandlers, printScreenActionHelp)
}

func printScreenActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown screen action: %s\n\n", action)
		printScreenHelp()
		os.Exit(1)
	}
}

func handleReindex(args []string) {
	actionHandlers := map[string]func([]string){
		"trigger": handleReindexTrigger,
		"status":  handleReindexStatus,
		"request": handleReindexRequest,
	}
	handleCommandWithHelp("reindex", args, printReindexHelp, actionHandlers, printReindexActionHelp)
}

func printReindexActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown reindex action: %s\n\n", action)
		printReindexHelp()
		os.Exit(1)
	}
}

func handleAppProp(args []string) {
	actionHandlers := map[string]func([]string){
		"list": handleAppPropList,
		"set":  handleAppPropSet,
	}
	handleCommandWithHelp("appprop", args, printAppPropHelp, actionHandlers, printAppPropActionHelp)
}

func printAppPropActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown appprop action: %s\n\n", action)
		printAppPropHelp()
		os.Exit(1)
	}
}

func handleConfig(args []string) {
	actionHandlers := map[string]func([]string){
		"get": handleConfigGet,
	}
	handleCommandWithHelp("config", args, printConfigHelp, actionHandlers, printConfigActionHelp)
}

func printConfigActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown config action: %s\n\n", action)
		printConfigHelp()
		os.Exit(1)
	}
}

func handleAvatar(args []string) {
	actionHandlers := map[string]func([]string){
		"system": handleAvatarSystem,
		"custom": handleAvatarCustom,
		"delete": handleAvatarDelete,
		"update": handleAvatarUpdate,
	}
	handleCommandWithHelp("avatar", args, printAvatarHelp, actionHandlers, printAvatarActionHelp)
}

func printAvatarActionHelp(action string) {
	switch action {
	default:
		fmt.Fprintf(os.Stderr, "Unknown avatar action: %s\n\n", action)
		printAvatarHelp()
		os.Exit(1)
	}
}
