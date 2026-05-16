package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	AccessToken string `toml:"access_token"`
	API         string `toml:"api"` // Always "graph"

	// MSAL credentials for automatic token refresh
	MSALClientID     string `toml:"msal_client_id"`
	MSALClientSecret string `toml:"msal_client_secret"`
	MSALTenantID     string `toml:"msal_tenant_id"`
	UseMSAL          bool   `toml:"use_msal"`

	// For app-only tokens (client credentials flow)
	UserID string `toml:"user_id"` // User ID for /users/{id}/chats endpoint
}

func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".teams-cli", "config.toml")
}

func Load() (*Config, error) {
	path := GetConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path := GetConfigPath()
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

// GetToken returns the access token from (in order):
// 1. Explicit apiKey parameter (if not empty)
// 2. TEAMS_CLI_TOKEN environment variable
// 3. Config file
func GetToken(apiKey string) (string, string) {
	if apiKey != "" {
		return apiKey, "flag"
	}

	if token := os.Getenv("TEAMS_CLI_TOKEN"); token != "" {
		return token, "env"
	}

	cfg, err := Load()
	if err == nil && cfg.AccessToken != "" {
		return cfg.AccessToken, "config"
	}

	return "", "missing"
}
