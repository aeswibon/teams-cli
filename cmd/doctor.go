package cmd

import (
	"fmt"

	"github.com/aeswibon/teams-cli/internal/config"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Verify configuration, authentication, and API connectivity",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, source := config.GetToken(apiKey)
		cfg, _ := config.Load()
		if cfg == nil {
			cfg = &config.Config{}
		}

		result := map[string]any{
			"token_available": token != "",
			"auth_source":     source,
			"api":             "graph",
		}

		if token == "" {
			return doctorError(result, source, fmt.Errorf("authentication not configured"))
		}

		client, err := loadMessenger()
		if err != nil {
			return doctorError(result, source, err)
		}
		result["api"] = client.Backend()

		user, err := client.Doctor()
		if err != nil {
			result["message"] = "Authentication failed"
			return doctorError(result, source, err)
		}

		result["status"] = "ok"
		result["user"] = user
		result["message"] = "Successfully authenticated"

		if jsonOutput {
			return printJSON(result)
		}

		fmt.Println("✅ Status: ok")
		fmt.Printf("🔑 Token source: %s\n", source)
		fmt.Printf("🔧 API: graph\n")

		if appOnly, _ := user["appOnly"].(bool); appOnly {
			fmt.Println("🤖 Auth type: app-only (MSAL client credentials)")
			if cfg.UserID != "" {
				fmt.Printf("👤 User ID: %s\n", cfg.UserID)
			}
		} else {
			if name, _ := user["displayName"].(string); name != "" {
				fmt.Printf("👤 User: %s\n", name)
			}
			if mail, _ := user["mail"].(string); mail != "" {
				fmt.Printf("📧 Email: %s\n", mail)
			}
		}

		return nil
	},
}

func doctorError(result map[string]any, source string, err error) error {
	result["status"] = "error"
	result["error"] = err.Error()
	if jsonOutput {
		_ = printJSON(result)
		return err
	}
	fmt.Println("❌ Status: error")
	if source != "" {
		fmt.Printf("🔑 Token source: %s\n", source)
	}
	fmt.Printf("🔧 API: graph\n")
	fmt.Printf("⚠️  %v\n", err)
	return err
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
