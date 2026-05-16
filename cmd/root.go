package cmd

import (
	"github.com/spf13/cobra"
)

var (
	jsonOutput bool
	apiKey     string
)

var rootCmd = &cobra.Command{
	Use:   "teams-cli",
	Short: "MS Teams CLI - interact with Microsoft Teams via Graph API",
	Long: `teams-cli is a command-line interface for Microsoft Teams.
Read messages and list chats via Microsoft Graph or the Teams Chat Service
(teams.cloud.microsoft / chatsvc) using tokens from your browser session.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Bearer token (overrides env/config)")
}
