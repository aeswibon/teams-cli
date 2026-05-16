package cmd

import (
	"fmt"

	"github.com/aeswibon/teams-cli/internal/api"
	"github.com/aeswibon/teams-cli/internal/config"
)

func loadMessenger() (api.Messenger, error) {
	token, _ := config.GetToken(apiKey)
	if token == "" {
		return nil, fmt.Errorf("no authentication token found — run 'teams-cli init' or set TEAMS_CLI_TOKEN")
	}

	cfg, err := config.Load()
	if err != nil {
		cfg = &config.Config{}
	}

	return api.NewMessenger(cfg, token)
}
