package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func makeJWT(aud any) string {
	payload, _ := json.Marshal(map[string]any{"aud": aud})
	enc := base64.RawURLEncoding.EncodeToString(payload)
	return "eyJhbGciOiJSUzI1NiJ9." + enc + ".sig"
}

func TestValidateGraphToken_acceptsGraphAudience(t *testing.T) {
	for _, aud := range []string{
		"https://graph.microsoft.com",
		"https://graph.microsoft.com/",
		"00000003-0000-0000-c000-000000000000",
	} {
		if err := ValidateGraphToken(makeJWT(aud)); err != nil {
			t.Fatalf("aud %q: %v", aud, err)
		}
	}
}

func TestValidateGraphToken_rejectsWrongAudience(t *testing.T) {
	err := ValidateGraphToken(makeJWT("https://api.spaces.skype.com"))
	if err == nil || !strings.Contains(err.Error(), "wrong token type") {
		t.Fatalf("expected wrong audience error, got %v", err)
	}
}
