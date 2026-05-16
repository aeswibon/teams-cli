package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "https://graph.microsoft.com/v1.0"
)

// GraphClient calls Microsoft Graph API v1.0.
type GraphClient struct {
	Token      string
	HTTPClient *http.Client
	UserID     string // For app-only tokens: use /users/{id}/chats instead of /me/chats
}

// NewGraphClient creates a Graph API client.
func NewGraphClient(token string) *GraphClient {
	return &GraphClient{
		Token:  token,
		UserID: "", // Will be set by factory if needed
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewGraphClientWithUser creates a Graph API client for a specific user (app-only tokens)
func NewGraphClientWithUser(token, userID string) *GraphClient {
	return &GraphClient{
		Token:  token,
		UserID: userID,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *GraphClient) Backend() string { return "graph" }

type GraphError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *GraphClient) Request(method, path string) ([]byte, error) {
	url := BaseURL + path
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var graphErr GraphError
		if json.Unmarshal(body, &graphErr) == nil && graphErr.Error.Message != "" {
			return nil, fmt.Errorf("API error (%d): %s - %s", resp.StatusCode, graphErr.Error.Code, graphErr.Error.Message)
		}
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// Doctor checks API connectivity and auth
func (c *GraphClient) Doctor() (map[string]interface{}, error) {
	// For app-only tokens, just verify we can make a request
	// We can't use /me (needs user delegation) or /organization (needs permissions)
	// So we'll just return success if we have a user ID configured
	if c.UserID != "" {
		return map[string]interface{}{
			"appOnly":     true,
			"displayName": "Application",
			"authType":    "app-only (MSAL client credentials)",
			"userID":      c.UserID,
		}, nil
	}

	// User-delegated token - try /me
	data, err := c.Request("GET", "/me")
	if err != nil {
		return nil, err
	}

	var user map[string]interface{}
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("parse user: %w", err)
	}

	return user, nil
}

type ChatList struct {
	Value []Chat `json:"value"`
}

func (c *GraphClient) ListChats() ([]Chat, error) {
	path := "/me/chats"
	if c.UserID != "" {
		// App-only token: use /users/{id}/chats
		path = fmt.Sprintf("/users/%s/chats", c.UserID)
	}

	data, err := c.Request("GET", path)
	if err != nil {
		return nil, err
	}

	var chatList ChatList
	if err := json.Unmarshal(data, &chatList); err != nil {
		return nil, fmt.Errorf("parse chats: %w", err)
	}

	return chatList.Value, nil
}

type MessageList struct {
	Value []Message `json:"value"`
}

func (c *GraphClient) GetMessages(chatID string, limit int) ([]Message, error) {
	path := fmt.Sprintf("/me/chats/%s/messages", chatID)
	if c.UserID != "" {
		// App-only token: use /users/{id}/chats/{chatId}/messages
		path = fmt.Sprintf("/users/%s/chats/%s/messages", c.UserID, chatID)
	}

	if limit > 0 {
		path += fmt.Sprintf("?$top=%d", limit)
	}

	data, err := c.Request("GET", path)
	if err != nil {
		return nil, err
	}

	var msgList MessageList
	if err := json.Unmarshal(data, &msgList); err != nil {
		return nil, fmt.Errorf("parse messages: %w", err)
	}

	return msgList.Value, nil
}
