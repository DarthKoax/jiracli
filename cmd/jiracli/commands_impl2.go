package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/darkkoax/jiracli/internal/api"
)

// Dashboard handlers
func handleDashboardList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	filter := parseFlag(args, "--filter")
	result, err := services.Dashboards.GetAll(context.Background(), startAt, maxResults, filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleDashboardGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: dashboard ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	dashboard, err := services.Dashboards.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(dashboard)
}

func handleDashboardCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var dashboard api.Dashboard
	if err := json.Unmarshal([]byte(jsonPayload), &dashboard); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Dashboards.Create(context.Background(), &dashboard)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleDashboardUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: dashboard ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var dashboard api.Dashboard
	if err := json.Unmarshal([]byte(jsonPayload), &dashboard); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Dashboards.Update(context.Background(), args[0], &dashboard)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleDashboardDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: dashboard ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Dashboards.Delete(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Dashboard %s deleted successfully\n", args[0])
}

func handleDashboardGadgets(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: dashboard ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	gadgets, err := services.Dashboards.GetGadgets(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(gadgets)
}

// Board handlers
func handleBoardList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	boardType := parseFlag(args, "--type")
	name := parseFlag(args, "--name")
	projectKey := parseFlag(args, "--project-key")
	result, err := services.Boards.GetAll(context.Background(), startAt, maxResults, boardType, name, projectKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleBoardGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	board, err := services.Boards.Get(context.Background(), boardID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(board)
}

func handleBoardCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var board api.Board
	if err := json.Unmarshal([]byte(jsonPayload), &board); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Boards.Create(context.Background(), &board)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleBoardDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err = services.Boards.Delete(context.Background(), boardID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Board %d deleted successfully\n", boardID)
}

func handleBoardConfig(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	config, err := services.Boards.GetConfig(context.Background(), boardID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(config)
}

func handleBoardIssues(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	jql := parseFlag(args, "--jql")
	fields := parseFlag(args, "--fields")
	var fieldList []string
	if fields != "" {
		fieldList = []string{fields}
	}
	result, err := services.Boards.GetIssues(context.Background(), boardID, startAt, maxResults, jql, fieldList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleBoardEpics(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	done := parseFlag(args, "--done") == "true"
	result, err := services.Boards.GetEpics(context.Background(), boardID, startAt, maxResults, done)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleBoardSprints(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	state := parseFlag(args, "--state")
	result, err := services.Boards.GetSprints(context.Background(), boardID, startAt, maxResults, state)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleBoardBacklog(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: board ID required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	jql := parseFlag(args, "--jql")
	fields := parseFlag(args, "--fields")
	var fieldList []string
	if fields != "" {
		fieldList = []string{fields}
	}
	result, err := services.Boards.GetBacklog(context.Background(), boardID, startAt, maxResults, jql, fieldList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

// Sprint handlers
func handleSprintGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	sprint, err := services.Sprints.Get(context.Background(), sprintID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(sprint)
}

func handleSprintCreate(args []string) {
	boardIDStr := parseFlag(args, "--board-id")
	if boardIDStr == "" {
		fmt.Fprintf(os.Stderr, "Error: --board-id required\n")
		os.Exit(1)
	}
	boardID, err := strconv.Atoi(boardIDStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid board ID\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var sprint api.Sprint
	if err := json.Unmarshal([]byte(jsonPayload), &sprint); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Sprints.Create(context.Background(), boardID, &sprint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleSprintUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var sprint api.Sprint
	if err := json.Unmarshal([]byte(jsonPayload), &sprint); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Sprints.Update(context.Background(), sprintID, &sprint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleSprintDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err = services.Sprints.Delete(context.Background(), sprintID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Sprint %d deleted successfully\n", sprintID)
}

func handleSprintIssues(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	jql := parseFlag(args, "--jql")
	fields := parseFlag(args, "--fields")
	var fieldList []string
	if fields != "" {
		fieldList = []string{fields}
	}
	result, err := services.Sprints.GetIssues(context.Background(), sprintID, startAt, maxResults, jql, fieldList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleSprintMove(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var payload struct {
		Issues []string `json:"issues"`
	}
	if err := json.Unmarshal([]byte(jsonPayload), &payload); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err = services.Sprints.MoveIssues(context.Background(), sprintID, payload.Issues)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Issues moved to sprint %d successfully\n", sprintID)
}

func handleSprintComplete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: sprint ID required\n")
		os.Exit(1)
	}
	sprintID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid sprint ID\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var payload struct {
		CompleteDate string `json:"completeDate"`
	}
	if err := json.Unmarshal([]byte(jsonPayload), &payload); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Sprints.Complete(context.Background(), sprintID, payload.CompleteDate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

// Field handlers
func handleFieldList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	fields, err := services.Fields.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(fields)
}

func handleFieldCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var field api.CustomField
	if err := json.Unmarshal([]byte(jsonPayload), &field); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Fields.CreateCustom(context.Background(), &field)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

// Component handlers
func handleComponentGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: component ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	component, err := services.Components.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(component)
}

func handleComponentCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var component api.Component
	if err := json.Unmarshal([]byte(jsonPayload), &component); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Components.Create(context.Background(), &component)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleComponentUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: component ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var component api.Component
	if err := json.Unmarshal([]byte(jsonPayload), &component); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Components.Update(context.Background(), args[0], &component)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Component %s updated successfully\n", args[0])
}

func handleComponentDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: component ID required\n")
		os.Exit(1)
	}
	replaceWith := parseFlag(args, "--replace-with")
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Components.Delete(context.Background(), args[0], replaceWith)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Component %s deleted successfully\n", args[0])
}

func handleComponentCount(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: component ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	count, err := services.Components.GetRelatedIssueCount(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(count)
}

// Version handlers
func handleVersionGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: version ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	version, err := services.Versions.Get(context.Background(), args[0], expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(version)
}

func handleVersionCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var version api.Version
	if err := json.Unmarshal([]byte(jsonPayload), &version); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Versions.Create(context.Background(), &version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleVersionUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: version ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var version api.Version
	if err := json.Unmarshal([]byte(jsonPayload), &version); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Versions.Update(context.Background(), args[0], &version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleVersionDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: version ID required\n")
		os.Exit(1)
	}
	moveFix := parseFlag(args, "--move-fix-issues-to")
	moveAffected := parseFlag(args, "--move-affected-issues-to")
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Versions.Delete(context.Background(), args[0], moveFix, moveAffected)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Version %s deleted successfully\n", args[0])
}

func handleVersionCount(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: version ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	count, err := services.Versions.GetRelatedIssueCounts(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(count)
}

func handleVersionMerge(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: version ID and target version ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Versions.Merge(context.Background(), args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Version %s merged into %s successfully\n", args[0], args[1])
}

// Priority handlers
func handlePriorityList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	priorities, err := services.Priorities.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(priorities)
}

func handlePriorityGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: priority ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	priority, err := services.Priorities.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(priority)
}

// Status handlers
func handleStatusList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	statuses, err := services.Statuses.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(statuses)
}

func handleStatusGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: status ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	status, err := services.Statuses.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(status)
}

func handleStatusCategories(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	categories, err := services.Statuses.GetCategories(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(categories)
}

func handleStatusCategory(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: category ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	category, err := services.Statuses.GetCategory(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(category)
}

// Resolution handlers
func handleResolutionList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	resolutions, err := services.Resolutions.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(resolutions)
}

func handleResolutionGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: resolution ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	resolution, err := services.Resolutions.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(resolution)
}

// Issue Type handlers
func handleIssueTypeList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	issueTypes, err := services.IssueTypes.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(issueTypes)
}

func handleIssueTypeGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue type ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	issueType, err := services.IssueTypes.Get(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(issueType)
}

func handleIssueTypeCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var issueType api.IssueType
	if err := json.Unmarshal([]byte(jsonPayload), &issueType); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.IssueTypes.Create(context.Background(), &issueType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleIssueTypeUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue type ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var issueType api.IssueType
	if err := json.Unmarshal([]byte(jsonPayload), &issueType); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.IssueTypes.Update(context.Background(), args[0], &issueType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Issue type %s updated successfully\n", args[0])
}

func handleIssueTypeDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue type ID required\n")
		os.Exit(1)
	}
	alternative := parseFlag(args, "--alternative-issue-type-id")
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.IssueTypes.Delete(context.Background(), args[0], alternative)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Issue type %s deleted successfully\n", args[0])
}

func handleIssueTypeAlternatives(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: issue type ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	alternatives, err := services.IssueTypes.GetAlternativeIssueTypes(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(alternatives)
}

// Permission handlers
func handlePermissionList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	permissions, err := services.Permissions.GetAll(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(permissions)
}

func handlePermissionMy(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	projectKey := parseFlag(args, "--project-key")
	projectID := parseFlag(args, "--project-id")
	issueKey := parseFlag(args, "--issue-key")
	issueID := parseFlag(args, "--issue-id")
	permID := parseFlag(args, "--permissions")
	permissions, err := services.Permissions.GetMyPermissions(context.Background(), projectKey, projectID, issueKey, issueID, permID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(permissions)
}

func handlePermissionScheme(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: action required (list, get, create, update, delete)\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	switch args[0] {
	case "list":
		startAt := parseFlagInt(args, "--start-at", 0)
		maxResults := parseFlagInt(args, "--max-results", 50)
		expand := parseFlag(args, "--expand")
		result, err := services.Permissions.GetAllSchemes(context.Background(), startAt, maxResults, expand)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		outputJSON(result)
	case "get":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: scheme ID required\n")
			os.Exit(1)
		}
		expand := parseFlag(args, "--expand")
		scheme, err := services.Permissions.GetScheme(context.Background(), args[1], expand)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		outputJSON(scheme)
	default:
		fmt.Fprintf(os.Stderr, "Unknown permission scheme action: %s\n", args[0])
		os.Exit(1)
	}
}

// Myself handlers
func handleMyselfGet(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	myself, err := services.Myself.Get(context.Background(), expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(myself)
}

func handleMyselfLocale(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	locale, err := services.Myself.GetLocale(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(locale)
}

// Screen handlers
func handleScreenList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	queryString := parseFlag(args, "--query-string")
	result, err := services.Screens.GetAll(context.Background(), startAt, maxResults, queryString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleScreenGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: screen ID required\n")
		os.Exit(1)
	}
	screenID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid screen ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	screen, err := services.Screens.Get(context.Background(), screenID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(screen)
}

func handleScreenCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var screen api.Screen
	if err := json.Unmarshal([]byte(jsonPayload), &screen); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Screens.Create(context.Background(), &screen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleScreenUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: screen ID required\n")
		os.Exit(1)
	}
	screenID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid screen ID\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var screen api.Screen
	if err := json.Unmarshal([]byte(jsonPayload), &screen); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err = services.Screens.Update(context.Background(), screenID, &screen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Screen %d updated successfully\n", screenID)
}

func handleScreenDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: screen ID required\n")
		os.Exit(1)
	}
	screenID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid screen ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err = services.Screens.Delete(context.Background(), screenID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Screen %d deleted successfully\n", screenID)
}

func handleScreenTabs(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: screen ID required\n")
		os.Exit(1)
	}
	screenID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid screen ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	tabs, err := services.Screens.GetTabs(context.Background(), screenID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(tabs)
}

// Reindex handlers
func handleReindexTrigger(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	reindexType := parseFlag(args, "--type")
	result, err := services.Reindex.Trigger(context.Background(), reindexType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleReindexStatus(args []string) {
	taskID := parseFlag(args, "--task-id")
	if taskID == "" {
		fmt.Fprintf(os.Stderr, "Error: --task-id required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Reindex.GetStatus(context.Background(), taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleReindexRequest(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: request ID required\n")
		os.Exit(1)
	}
	requestID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid request ID\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Reindex.GetRequestStatus(context.Background(), requestID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

// Application Properties handlers
func handleAppPropList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	key := parseFlag(args, "--key")
	permissionLevel := parseFlag(args, "--permission-level")
	props, err := services.ApplicationProperties.GetAll(context.Background(), key, permissionLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(props)
}

func handleAppPropSet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: property ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var payload struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal([]byte(jsonPayload), &payload); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.ApplicationProperties.Set(context.Background(), args[0], payload.Value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Property %s set successfully\n", args[0])
}

// Configuration handlers
func handleConfigGet(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	config, err := services.Configuration.Get(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(config)
}

// Avatar handlers
func handleAvatarSystem(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: entity type required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	avatars, err := services.Avatars.GetSystemAvatars(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(avatars)
}

func handleAvatarCustom(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: entity type and entity ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	avatars, err := services.Avatars.GetCustomAvatars(context.Background(), args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(avatars)
}

func handleAvatarDelete(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: entity type and avatar ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Avatars.DeleteCustomAvatar(context.Background(), args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Avatar %s deleted successfully\n", args[1])
}

func handleAvatarUpdate(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: entity type and avatar ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var avatar api.Avatar
	if err := json.Unmarshal([]byte(jsonPayload), &avatar); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Avatars.UpdateCustomAvatar(context.Background(), args[0], args[1], &avatar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Avatar %s updated successfully\n", args[1])
}
