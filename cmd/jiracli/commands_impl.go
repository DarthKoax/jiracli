package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/darkkoax/jiracli/internal/api"
	"github.com/darkkoax/jiracli/internal/client"
	"github.com/darkkoax/jiracli/internal/config"
)

// Helper function to get services
func getServices(configPath string) *api.Services {
	if configPath == "" {
		configPath = getDefaultConfigPath()
	}
	cfg, err := loadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
	c, err := createClient(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}
	return api.NewServices(c)
}

func loadConfig(configPath string) (*config.Config, error) {
	return config.Load(configPath)
}

func createClient(cfg *config.Config) (*client.Client, error) {
	return client.New(cfg)
}

func getDefaultConfigPath() string {
	if configDir := os.Getenv("JIRA_CONFIG_DIR"); configDir != "" {
		return configDir + "/config.toml"
	}
	home, _ := os.UserHomeDir()
	return home + "/.config/darthkoax/jiracli/config.toml"
}

func outputJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func parseFlag(args []string, flag string) string {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == flag {
			return args[i+1]
		}
	}
	return ""
}

func parseFlagInt(args []string, flag string, defaultVal int) int {
	val := parseFlag(args, flag)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

// Project handlers
func handleProjectList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	projects, err := services.Projects.GetAll(context.Background(), expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(projects)
}

func handleProjectGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	projectKey := args[0]
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	project, err := services.Projects.Get(context.Background(), projectKey, expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(project)
}

func handleProjectCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var project api.Project
	if err := json.Unmarshal([]byte(jsonPayload), &project); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Projects.Create(context.Background(), &project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleProjectUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var project api.Project
	if err := json.Unmarshal([]byte(jsonPayload), &project); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Projects.Update(context.Background(), args[0], &project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %s updated successfully\n", args[0])
}

func handleProjectDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Projects.Delete(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %s deleted successfully\n", args[0])
}

func handleProjectRoles(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	roles, err := services.Projects.GetRoles(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(roles)
}

func handleProjectComponents(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	components, err := services.Projects.GetComponents(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(components)
}

func handleProjectVersions(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: project key required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	versions, err := services.Projects.GetVersions(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(versions)
}

// User handlers
func handleUserGet(args []string) {
	username := parseFlag(args, "--username")
	accountID := parseFlag(args, "--account-id")
	if username == "" && accountID == "" {
		fmt.Fprintf(os.Stderr, "Error: --username or --account-id required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	var user *api.User
	var err error
	if username != "" {
		user, err = services.Users.Get(context.Background(), username, expandList)
	} else {
		user, err = services.Users.GetByAccountID(context.Background(), accountID)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(user)
}

func handleUserSearch(args []string) {
	username := parseFlag(args, "--username")
	if username == "" {
		fmt.Fprintf(os.Stderr, "Error: --username required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	users, err := services.Users.Search(context.Background(), username, startAt, maxResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(users)
}

func handleUserCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var user api.User
	if err := json.Unmarshal([]byte(jsonPayload), &user); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Users.Create(context.Background(), &user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleUserDelete(args []string) {
	username := parseFlag(args, "--username")
	if username == "" {
		fmt.Fprintf(os.Stderr, "Error: --username required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Users.Delete(context.Background(), username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("User %s deleted successfully\n", username)
}

func handleUserColumns(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: action required (get, set, reset)\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	switch args[0] {
	case "get":
		columns, err := services.Users.GetColumns(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		outputJSON(columns)
	case "set":
		jsonPayload := parseFlag(args, "--json")
		if jsonPayload == "" {
			fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
			os.Exit(1)
		}
		var columns []map[string]interface{}
		if err := json.Unmarshal([]byte(jsonPayload), &columns); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
			os.Exit(1)
		}
		err := services.Users.SetColumns(context.Background(), columns)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("User columns updated successfully")
	case "reset":
		err := services.Users.ResetColumns(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("User columns reset successfully")
	default:
		fmt.Fprintf(os.Stderr, "Unknown columns action: %s\n", args[0])
		os.Exit(1)
	}
}

// Group handlers
func handleGroupGet(args []string) {
	groupname := parseFlag(args, "--groupname")
	if groupname == "" {
		fmt.Fprintf(os.Stderr, "Error: --groupname required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	group, err := services.Groups.Get(context.Background(), groupname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(group)
}

func handleGroupCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var payload map[string]string
	if err := json.Unmarshal([]byte(jsonPayload), &payload); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Groups.Create(context.Background(), payload["name"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleGroupDelete(args []string) {
	groupname := parseFlag(args, "--groupname")
	if groupname == "" {
		fmt.Fprintf(os.Stderr, "Error: --groupname required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Groups.Delete(context.Background(), groupname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group %s deleted successfully\n", groupname)
}

func handleGroupMembers(args []string) {
	groupname := parseFlag(args, "--groupname")
	if groupname == "" {
		fmt.Fprintf(os.Stderr, "Error: --groupname required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	members, err := services.Groups.GetMembers(context.Background(), groupname, startAt, maxResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(members)
}

func handleGroupSearch(args []string) {
	query := parseFlag(args, "--query")
	if query == "" {
		fmt.Fprintf(os.Stderr, "Error: --query required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	maxResults := parseFlagInt(args, "--max-results", 50)
	groups, err := services.Groups.Search(context.Background(), query, 0, maxResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(groups)
}

// Search handlers
func handleSearchGet(args []string) {
	jql := parseFlag(args, "--jql")
	if jql == "" {
		fmt.Fprintf(os.Stderr, "Error: --jql required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	fields := parseFlag(args, "--fields")
	expand := parseFlag(args, "--expand")
	var fieldList, expandList []string
	if fields != "" {
		fieldList = []string{fields}
	}
	if expand != "" {
		expandList = []string{expand}
	}
	result, err := services.Search.SearchGet(context.Background(), jql, startAt, maxResults, fieldList, expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

// Filter handlers
func handleFilterGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: filter ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	expand := parseFlag(args, "--expand")
	var expandList []string
	if expand != "" {
		expandList = []string{expand}
	}
	filter, err := services.Filters.Get(context.Background(), args[0], expandList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(filter)
}

func handleFilterCreate(args []string) {
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var filter api.Filter
	if err := json.Unmarshal([]byte(jsonPayload), &filter); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Filters.Create(context.Background(), &filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleFilterUpdate(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: filter ID required\n")
		os.Exit(1)
	}
	jsonPayload := parseFlag(args, "--json")
	if jsonPayload == "" {
		fmt.Fprintf(os.Stderr, "Error: --json flag required\n")
		os.Exit(1)
	}
	var filter api.Filter
	if err := json.Unmarshal([]byte(jsonPayload), &filter); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	result, err := services.Filters.Update(context.Background(), args[0], &filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(result)
}

func handleFilterDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: filter ID required\n")
		os.Exit(1)
	}
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	err := services.Filters.Delete(context.Background(), args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Filter %s deleted successfully\n", args[0])
}

func handleFilterList(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	filters, err := services.Filters.GetFavourite(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(filters)
}

func handleFilterMy(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	filters, err := services.Filters.GetMy(context.Background(), startAt, maxResults, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(filters)
}

func handleFilterSearch(args []string) {
	configPath := parseFlag(args, "--config")
	services := getServices(configPath)
	filterName := parseFlag(args, "--filter-name")
	owner := parseFlag(args, "--owner")
	groupname := parseFlag(args, "--groupname")
	startAt := parseFlagInt(args, "--start-at", 0)
	maxResults := parseFlagInt(args, "--max-results", 50)
	filters, err := services.Filters.Search(context.Background(), filterName, owner, groupname, startAt, maxResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	outputJSON(filters)
}
