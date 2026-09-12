package main

import (
	"context"
	"fmt"
	"os"

	"github.com/darkkoax/jiracli/internal/api"
	"github.com/darkkoax/jiracli/internal/client"
	"github.com/darkkoax/jiracli/internal/config"
)

func handleConnect(args []string) {
	if len(args) > 0 && (args[0] == "help" || args[0] == "--help" || args[0] == "-h") {
		printConnectHelp()
		return
	}

	var configPath string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--config":
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
		}
	}

	if configPath == "" {
		configPath = config.DefaultConfigPath()
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	c, err := client.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	services := api.NewServices(c)
	ctx := context.Background()

	myself, err := services.Myself.Get(ctx, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Jira: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to Jira DC as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
	fmt.Printf("Jira instance: %s\n", cfg.Jira.BaseURL)
}

func printConnectHelp() {
	fmt.Print(`jiracli connect - Connect to Jira and verify authentication

USAGE:
  jiracli connect [--config <path>]

DESCRIPTION:
  Connects to the Jira Data Center instance using the configuration file and
  verifies authentication by fetching the current user's information.

  This command is useful for testing your configuration before using other
  commands. It will fail if:
  - The config file doesn't exist or is invalid
  - The base_url is unreachable
  - The API token is invalid or expired
  - Custom CA certificate is invalid (if specified)

FLAGS:
  --config <path>    Path to config file
                     Default: ~/.config/jiracli/config.toml

EXAMPLES:
  jiracli connect                              # Use default config location
  jiracli connect --config /path/to/config     # Use custom config file

OUTPUT:
  On success:
    Connected to Jira DC as: John Doe (john.doe@example.com)
    Jira instance: https://jira.example.com

  On failure:
    Error connecting to Jira: <error message>
    Exit code: 1

COMMON ERRORS:
  - "reading config file: no such file or directory"
    → Run 'jiracli init' to create a config file

  - "jira.base_url is required"
    → Edit config.toml and set base_url

  - "jira.api_token is required"
    → Edit config.toml and set api_token

  - "API error (status 401)"
    → API token is invalid or expired. Generate a new token in Jira.

  - "x509: certificate signed by unknown authority"
    → Set custom_ca_cert in config.toml to your CA certificate bundle

TROUBLESHOOTING:
  1. Verify config file exists: jiracli init
  2. Check base_url is correct and accessible
  3. Verify API token has appropriate permissions
  4. If using custom CA, ensure cert path is correct
  5. Check network connectivity to Jira instance
`)
}
