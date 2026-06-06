package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		auth      string
		wantKey   string
		wantError bool
	}{
		{"ApiKey my-secret-key", "my-secret-key", false},
		{"", "", true},
		{"Bearer token", "", true},
		{"ApiKey", "", true},
		{"apikey my-key", "", true},
	}

	for _, tt := range tests {
		headers := http.Header{}
		if tt.auth != "" {
			headers.Set("Authorization", tt.auth)
		}

		key, err := GetAPIKey(headers)

		if key != tt.wantKey {
			t.Errorf("auth=%q: got key %q, want %q", tt.auth, key, tt.wantKey)
		}
		if (err != nil) != tt.wantError {
			t.Errorf("auth=%q: got error %v, want error %t", tt.auth, err, tt.wantError)
		}
	}
}
