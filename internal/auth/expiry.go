package auth

import (
	"encoding/json"
	"strings"
	"time"
)

type jwtClaims struct {
	Exp int64 `json:"exp"`
	Aud any   `json:"aud"`
}

func tokenExpiresAt(token string) time.Time {
	payload, err := decodeJWTPayload(strings.TrimSpace(token))
	if err != nil {
		return time.Time{}
	}
	var claims jwtClaims
	if json.Unmarshal(payload, &claims) != nil || claims.Exp == 0 {
		return time.Time{}
	}
	return time.Unix(claims.Exp, 0).UTC()
}

// TokenExpired reports whether the JWT exp is in the past (with 30s skew).
func TokenExpired(token string) bool {
	exp := tokenExpiresAt(token)
	if exp.IsZero() {
		return false
	}
	return time.Now().After(exp.Add(-30 * time.Second))
}
