package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/darkkoax/jiracli/internal/api"
)

func handleIssue(args []string) {
	if len(args) == 0 {
		printIssueHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			// Help for specific action: jiracli issue help <action>
			printIssueActionHelp(args[1])
		} else {
			// Help for issue command: jiracli issue help
			printIssueHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help: jiracli issue <action> help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printIssueActionHelp(action)
		return
	}

	switch action {
	case "get":
		handleIssueGet(actionArgs)
	case "create":
		handleIssueCreate(actionArgs)
	case "update":
		handleIssueUpdate(actionArgs)
	case "delete":
		handleIssueDelete(actionArgs)
	case "transitions":
		handleIssueTransitions(actionArgs)
	case "transition":
		handleIssueTransition(actionArgs)
	case "comment":
		handleIssueComment(actionArgs)
	case "worklog":
		handleIssueWorklog(actionArgs)
	case "watcher":
		handleIssueWatcher(actionArgs)
	case "vote":
		handleIssueVote(actionArgs)
	case "assign":
		handleIssueAssign(actionArgs)
	default:
		fmt.Fprintf(os.Stderr, "Unknown issue action: %s\n\n", action)
		printIssueHelp()
		os.Exit(1)
	}
}

func printIssueActionHelp(action string) {
	switch action {
	case "get":
		printIssueGetHelp()
	case "create":
		printIssueCreateHelp()
	case "update":
		printIssueUpdateHelp()
	case "delete":
		printIssueDeleteHelp()
	case "transitions":
		printIssueTransitionsHelp()
	case "transition":
		printIssueTransitionHelp()
	case "comment":
		printIssueCommentHelp()
	case "worklog":
		printIssueWorklogHelp()
	case "watcher":
		printIssueWatcherHelp()
	case "vote":
		printIssueVoteHelp()
	case "assign":
		printIssueAssignHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown issue action: %s\n\n", action)
		printIssueHelp()
		os.Exit(1)
	}
}

func handleIssueGet(args []string) {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueGetHelp()
		return
	}

	issueKey := args[0]
	var fields, expand []string
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--fields":
			if i+1 < len(args) {
				fields = splitComma(args[i+1])
				i++
			}
		case "--expand":
			if i+1 < len(args) {
				expand = splitComma(args[i+1])
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	issue, err := services.Issues.Get(ctx, issueKey, fields, expand)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(issue)
}

func handleIssueCreate(args []string) {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueCreateHelp()
		return
	}

	var jsonPayload string
	var configPath string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json":
			if i+1 < len(args) {
				jsonPayload = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag is required\n")
		os.Exit(1)
	}

	var issue api.Issue
	if err := json.Unmarshal([]byte(jsonPayload), &issue); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	result, err := services.Issues.Create(ctx, &issue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(result)
}

func handleIssueUpdate(args []string) {
	if len(args) < 2 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueUpdateHelp()
		return
	}

	issueKey := args[0]
	var jsonPayload string
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--json":
			if i+1 < len(args) {
				jsonPayload = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag is required\n")
		os.Exit(1)
	}

	var issue api.Issue
	if err := json.Unmarshal([]byte(jsonPayload), &issue); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.Update(ctx, issueKey, &issue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Issue %s updated successfully\n", issueKey)
}

func handleIssueDelete(args []string) {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueDeleteHelp()
		return
	}

	issueKey := args[0]
	var deleteSubtasks bool
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--delete-subtasks":
			deleteSubtasks = true
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.Delete(ctx, issueKey, deleteSubtasks)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Issue %s deleted successfully\n", issueKey)
}

func handleIssueTransitions(args []string) {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueTransitionsHelp()
		return
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	transitions, err := services.Issues.GetTransitions(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(transitions)
}

func handleIssueTransition(args []string) {
	if len(args) < 2 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueTransitionHelp()
		return
	}

	issueKey := args[0]
	transitionID := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.DoTransition(ctx, issueKey, transitionID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Issue %s transitioned successfully\n", issueKey)
}

func handleIssueComment(args []string) {
	if len(args) == 0 {
		printIssueCommentHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			// Help for specific action: jiracli issue comment help <action>
			printIssueCommentActionHelp(args[1])
		} else {
			// Help for comment command: jiracli issue comment help
			printIssueCommentHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help: jiracli issue comment <action> help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printIssueCommentActionHelp(action)
		return
	}

	switch action {
	case "list":
		handleIssueCommentList(actionArgs)
	case "add":
		handleIssueCommentAdd(actionArgs)
	case "update":
		handleIssueCommentUpdate(actionArgs)
	case "delete":
		handleIssueCommentDelete(actionArgs)
	default:
		fmt.Fprintf(os.Stderr, "Unknown comment action: %s\n", action)
		os.Exit(1)
	}
}

func printIssueCommentActionHelp(action string) {
	switch action {
	case "list":
		printIssueCommentListHelp()
	case "add":
		printIssueCommentAddHelp()
	case "update":
		printIssueCommentUpdateHelp()
	case "delete":
		printIssueCommentDeleteHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown comment action: %s\n\n", action)
		printIssueCommentHelp()
		os.Exit(1)
	}
}

func handleIssueCommentList(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	comments, err := services.Issues.GetComments(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(comments)
}

func handleIssueCommentAdd(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and --body required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var body, configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 < len(args) {
				body = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if body == "" {
		fmt.Fprintf(os.Stderr, "Error: --body flag is required\n")
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	comment := &api.Comment{Body: body}
	result, err := services.Issues.AddComment(ctx, issueKey, comment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(result)
}

func handleIssueCommentUpdate(args []string) {
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Error: issue key, comment ID, and --body required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	commentID := args[1]
	var body, configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 < len(args) {
				body = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if body == "" {
		fmt.Fprintf(os.Stderr, "Error: --body flag is required\n")
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	comment := &api.Comment{Body: body}
	err := services.Issues.UpdateComment(ctx, issueKey, commentID, comment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comment %s updated successfully\n", commentID)
}

func handleIssueCommentDelete(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and comment ID required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	commentID := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.DeleteComment(ctx, issueKey, commentID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comment %s deleted successfully\n", commentID)
}

func handleIssueWorklog(args []string) {
	if len(args) == 0 {
		printIssueWorklogHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			printIssueWorklogActionHelp(args[1])
		} else {
			printIssueWorklogHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printIssueWorklogActionHelp(action)
		return
	}

	switch action {
	case "list":
		handleIssueWorklogList(actionArgs)
	case "add":
		handleIssueWorklogAdd(actionArgs)
	case "update":
		handleIssueWorklogUpdate(actionArgs)
	case "delete":
		handleIssueWorklogDelete(actionArgs)
	default:
		fmt.Fprintf(os.Stderr, "Unknown worklog action: %s\n", action)
		os.Exit(1)
	}
}

func printIssueWorklogActionHelp(action string) {
	switch action {
	case "list":
		printIssueWorklogListHelp()
	case "add":
		printIssueWorklogAddHelp()
	case "update":
		printIssueWorklogUpdateHelp()
	case "delete":
		printIssueWorklogDeleteHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown worklog action: %s\n\n", action)
		printIssueWorklogHelp()
		os.Exit(1)
	}
}

func handleIssueWorklogList(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	worklogs, err := services.Issues.GetWorklogs(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(worklogs)
}

func handleIssueWorklogAdd(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var jsonPayload, configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--json":
			if i+1 < len(args) {
				jsonPayload = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag is required\n")
		os.Exit(1)
	}

	var worklog api.Worklog
	if err := json.Unmarshal([]byte(jsonPayload), &worklog); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	result, err := services.Issues.AddWorklog(ctx, issueKey, &worklog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(result)
}

func handleIssueWorklogUpdate(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and worklog ID required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	worklogID := args[1]
	var jsonPayload, configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--json":
			if i+1 < len(args) {
				jsonPayload = args[i+1]
				i++
			}
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag is required\n")
		os.Exit(1)
	}

	var worklog api.Worklog
	if err := json.Unmarshal([]byte(jsonPayload), &worklog); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.UpdateWorklog(ctx, issueKey, worklogID, &worklog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Worklog %s updated successfully\n", worklogID)
}

func handleIssueWorklogDelete(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and worklog ID required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	worklogID := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.DeleteWorklog(ctx, issueKey, worklogID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Worklog %s deleted successfully\n", worklogID)
}

func handleIssueWatcher(args []string) {
	if len(args) == 0 {
		printIssueWatcherHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			printIssueWatcherActionHelp(args[1])
		} else {
			printIssueWatcherHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printIssueWatcherActionHelp(action)
		return
	}

	switch action {
	case "list":
		handleIssueWatcherList(actionArgs)
	case "add":
		handleIssueWatcherAdd(actionArgs)
	case "remove":
		handleIssueWatcherRemove(actionArgs)
	default:
		fmt.Fprintf(os.Stderr, "Unknown watcher action: %s\n", action)
		os.Exit(1)
	}
}

func printIssueWatcherActionHelp(action string) {
	switch action {
	case "list":
		printIssueWatcherListHelp()
	case "add":
		printIssueWatcherAddHelp()
	case "remove":
		printIssueWatcherRemoveHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown watcher action: %s\n\n", action)
		printIssueWatcherHelp()
		os.Exit(1)
	}
}

func handleIssueWatcherList(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	watchers, err := services.Issues.GetWatchers(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(watchers)
}

func handleIssueWatcherAdd(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and username required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	username := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.AddWatcher(ctx, issueKey, username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Watcher %s added to %s\n", username, issueKey)
}

func handleIssueWatcherRemove(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: issue key and username required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	username := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.RemoveWatcher(ctx, issueKey, username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Watcher %s removed from %s\n", username, issueKey)
}

func handleIssueVote(args []string) {
	if len(args) == 0 {
		printIssueVoteHelp()
		return
	}

	// Check if first arg is "help"
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		if len(args) > 1 {
			printIssueVoteActionHelp(args[1])
		} else {
			printIssueVoteHelp()
		}
		return
	}

	action := args[0]
	actionArgs := args[1:]

	// Check if action is followed by help
	if len(actionArgs) > 0 && (actionArgs[0] == "help" || actionArgs[0] == "--help" || actionArgs[0] == "-h") {
		printIssueVoteActionHelp(action)
		return
	}

	switch action {
	case "get":
		handleIssueVoteGet(actionArgs)
	case "add":
		handleIssueVoteAdd(actionArgs)
	case "remove":
		handleIssueVoteRemove(actionArgs)
	default:
		fmt.Fprintf(os.Stderr, "Unknown vote action: %s\n", action)
		os.Exit(1)
	}
}

func printIssueVoteActionHelp(action string) {
	switch action {
	case "get":
		printIssueVoteGetHelp()
	case "add":
		printIssueVoteAddHelp()
	case "remove":
		printIssueVoteRemoveHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown vote action: %s\n\n", action)
		printIssueVoteHelp()
		os.Exit(1)
	}
}

func handleIssueVoteGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	vote, err := services.Issues.GetVotes(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	outputJSON(vote)
}

func handleIssueVoteAdd(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.AddVote(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Vote added to %s\n", issueKey)
}

func handleIssueVoteRemove(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue key required\n")
		os.Exit(1)
	}

	issueKey := args[0]
	var configPath string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.RemoveVote(ctx, issueKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Vote removed from %s\n", issueKey)
}

func handleIssueAssign(args []string) {
	if len(args) < 2 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printIssueAssignHelp()
		return
	}

	issueKey := args[0]
	assignee := args[1]
	var configPath string

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	services := getServices(configPath)
	ctx := context.Background()

	err := services.Issues.Assign(ctx, issueKey, assignee)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Issue %s assigned to %s\n", issueKey, assignee)
}

// Helper functions
func splitComma(s string) []string {
	if s == "" {
		return nil
	}
	result := []string{}
	for _, part := range split(s, ",") {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func split(s, sep string) []string {
	result := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i = start - 1
		}
	}
	result = append(result, s[start:])
	return result
}
