package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Graph accepts these audience values for /v1.0 API calls.
var graphAudiences = map[string]bool{
	"https://graph.microsoft.com":          true,
	"https://graph.microsoft.com/":         true,
	"00000003-0000-0000-c000-000000000000": true,
}

type jwtPayload struct {
	Aud   any      `json:"aud"`
	Scp   string   `json:"scp"`
	Roles []string `json:"roles"`
	Typ   string   `json:"typ"`
}

// ValidateGraphToken checks that a JWT is intended for Microsoft Graph.
func ValidateGraphToken(token string) error {
	token = strings.TrimSpace(token)
	if !strings.HasPrefix(token, "eyJ") {
		return fmt.Errorf("token should start with eyJ (did you copy the full value after 'Bearer '?)")
	}

	payload, err := decodeJWTPayload(token)
	if err != nil {
		return err
	}

	var p jwtPayload
	if err := json.Unmarshal(payload, &p); err != nil {
		return fmt.Errorf("parse token payload: %w", err)
	}

	if strings.EqualFold(p.Typ, "refresh") {
		return fmt.Errorf("this looks like a refresh token, not a Graph access token — pick a different request")
	}

	auds := normalizeAudiences(p.Aud)
	if len(auds) == 0 {
		return fmt.Errorf("token has no audience (aud) claim")
	}

	for _, aud := range auds {
		if graphAudiences[aud] {
			return nil
		}
	}

	return fmt.Errorf(
		"wrong token type: audience is %q (expected Microsoft Graph)\n\n"+
			"For teams.cloud.microsoft use chatsvc instead:\n"+
			"  teams-cli init --api chatsvc --token BEARER --skype-token SKYPE\n"+
			"See README.md for copying both tokens from a chatsvc request",
		strings.Join(auds, ", "),
	)
}

func normalizeAudiences(aud any) []string {
	switch v := aud.(type) {
	case string:
		return []string{v}
	case []any:
		out := make([]string, 0, len(v))
		for _, item := range v {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	default:
		return nil
	}
}

func decodeJWTPayload(token string) ([]byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid JWT format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		// Some tokens use standard base64 padding
		payload, err = base64.URLEncoding.DecodeString(parts[1])
		if err != nil {
			return nil, fmt.Errorf("decode token payload: %w", err)
		}
	}
	return payload, nil
}

// Audience returns the aud claim for display (doctor hints).
func Audience(token string) string {
	payload, err := decodeJWTPayload(strings.TrimSpace(token))
	if err != nil {
		return ""
	}
	var p jwtPayload
	if json.Unmarshal(payload, &p) != nil {
		return ""
	}
	auds := normalizeAudiences(p.Aud)
	return strings.Join(auds, ", ")
}
