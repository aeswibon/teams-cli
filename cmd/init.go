package cmd

import (
	"fmt"
	"strings"

	"github.com/abhiudayg/teams-cli/internal/auth"
	"github.com/abhiudayg/teams-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	initToken            string
	initMSALClientID     string
	initMSALClientSecret string
	initMSALTenantID     string
	initUseMSAL          bool
	initUserID           string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration",
	Long: `Save MSAL credentials to ~/.teams-cli/config.toml.

Use MSAL for automatic token refresh:
  teams-cli init --use-msal \
    --msal-client-id "CLIENT_ID" \
    --msal-client-secret "SECRET" \
    --msal-tenant-id "TENANT_ID" \
    --user-id "YOUR_USER_ID"

Or manually provide a bearer token:
  teams-cli init --token "BEARER_TOKEN"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := &config.Config{API: "graph"}

		// MSAL configuration
		if initUseMSAL {
			if initMSALClientID == "" || initMSALClientSecret == "" || initMSALTenantID == "" {
				return fmt.Errorf("MSAL requires --msal-client-id, --msal-client-secret, and --msal-tenant-id")
			}

			if initUserID == "" {
				return fmt.Errorf("--user-id is required for MSAL app-only tokens")
			}

			cfg.UseMSAL = true
			cfg.MSALClientID = strings.TrimSpace(initMSALClientID)
			cfg.MSALClientSecret = strings.TrimSpace(initMSALClientSecret)
			cfg.MSALTenantID = strings.TrimSpace(initMSALTenantID)
			cfg.UserID = strings.TrimSpace(initUserID)

			// Test MSAL credentials by getting a token
			msalConfig := auth.MSALConfig{
				ClientID:     cfg.MSALClientID,
				ClientSecret: cfg.MSALClientSecret,
				TenantID:     cfg.MSALTenantID,
			}
			token, err := auth.GetMSALToken(msalConfig)
			if err != nil {
				return fmt.Errorf("MSAL authentication failed: %w", err)
			}

			// Validate the token
			if err := auth.ValidateGraphToken(token); err != nil {
				return fmt.Errorf("MSAL token validation failed: %w", err)
			}

			cfg.AccessToken = token

		} else {
			// Manual token configuration
			if initToken == "" {
				return fmt.Errorf("--token is required (or use --use-msal for automatic token refresh)")
			}

			cfg.AccessToken = strings.TrimSpace(initToken)

			if err := auth.ValidateGraphToken(cfg.AccessToken); err != nil {
				return err
			}
		}

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}

		fmt.Printf("\n✅ Configuration saved to %s\n", config.GetConfigPath())
		fmt.Printf("🔧 API backend: graph\n")
		if cfg.UseMSAL {
			fmt.Println("🔄 MSAL: Enabled (automatic token refresh)")
			fmt.Printf("👤 User ID: %s\n", cfg.UserID)
		}
		fmt.Println("🔒 File permissions: 0600 (user read/write only)")
		fmt.Println("\nRun 'teams-cli doctor' to verify authentication")

		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initToken, "token", "", "Bearer token (Authorization header)")
	initCmd.Flags().BoolVar(&initUseMSAL, "use-msal", false, "Use MSAL for automatic token refresh")
	initCmd.Flags().StringVar(&initMSALClientID, "msal-client-id", "", "MSAL Application (client) ID")
	initCmd.Flags().StringVar(&initMSALClientSecret, "msal-client-secret", "", "MSAL Client secret")
	initCmd.Flags().StringVar(&initMSALTenantID, "msal-tenant-id", "", "MSAL Directory (tenant) ID")
	initCmd.Flags().StringVar(&initUserID, "user-id", "", "User ID for app-only tokens (required for MSAL)")

	rootCmd.AddCommand(initCmd)
}
