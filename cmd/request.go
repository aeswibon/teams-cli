package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aeswibon/teams-cli/internal/api"
	"github.com/aeswibon/teams-cli/internal/config"
	"github.com/spf13/cobra"
)

var requestMethod string

var requestCmd = &cobra.Command{
	Use:   "request <path>",
	Short: "Make raw Microsoft Graph API request",
	Long:  `Graph API only. Path is relative to https://graph.microsoft.com/v1.0/`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _ := config.Load()
		if cfg == nil {
			cfg = &config.Config{}
		}
		token, _ := config.GetToken(apiKey)
		if token == "" {
			return fmt.Errorf("no authentication token found")
		}
		if cfg.ResolveBackend(token) != config.APIGraph {
			return fmt.Errorf("request requires graph API; use chats and messages for chatsvc")
		}

		client := api.NewGraphClient(token)
		body, err := client.Request(requestMethod, args[0])
		if err != nil {
			return err
		}

		if jsonOutput {
			var data any
			if json.Unmarshal(body, &data) == nil {
				return printJSON(data)
			}
		}
		fmt.Print(string(body))
		if len(body) > 0 && body[len(body)-1] != '\n' {
			fmt.Println()
		}
		return nil
	},
}

func init() {
	requestCmd.Flags().StringVar(&requestMethod, "method", http.MethodGet, "HTTP method")
	rootCmd.AddCommand(requestCmd)
}
