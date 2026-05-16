package api

import (
	"fmt"
	"strings"

	"github.com/aeswibon/teams-cli/internal/auth"
	"github.com/aeswibon/teams-cli/internal/config"
)

// NewMessenger returns a Graph API client.
func NewMessenger(cfg *config.Config, bearerOverride string) (Messenger, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	bearer := strings.TrimSpace(bearerOverride)

	// If MSAL is enabled, get a fresh token
	if bearer == "" && cfg.UseMSAL && cfg.MSALClientID != "" {
		msalConfig := auth.MSALConfig{
			ClientID:     cfg.MSALClientID,
			ClientSecret: cfg.MSALClientSecret,
			TenantID:     cfg.MSALTenantID,
		}
		token, err := auth.GetMSALToken(msalConfig)
		if err != nil {
			return nil, fmt.Errorf("get MSAL token: %w", err)
		}
		bearer = token
	}

	if bearer == "" {
		bearer = strings.TrimSpace(cfg.AccessToken)
	}
	if bearer == "" {
		return nil, fmt.Errorf("no access token — run teams-cli init")
	}

	// Validate Graph token
	if err := auth.ValidateGraphToken(bearer); err != nil {
		return nil, err
	}

	// Check if we need to use app-only mode with user ID
	if cfg.UserID != "" {
		return NewGraphClientWithUser(bearer, cfg.UserID), nil
	}
	return NewGraphClient(bearer), nil
}
