package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		wantKey     string
		wantErr     bool
	}{
		{
			name:        "valid api key",
			headerValue: "ApiKey abc123",
			wantKey:     "abc123",
			wantErr:     false,
		},
		{
			name:        "missing authorization header",
			headerValue: "",
			wantKey:     "",
			wantErr:     true,
		},
		{
			name:        "wrong auth scheme",
			headerValue: "Bearer abc123",
			wantKey:     "",
			wantErr:     true,
		},
		{
			name:        "malformed header missing key",
			headerValue: "ApiKey",
			wantKey:     "",
			wantErr:     true,
		},
		{
			name:        "extra spaces still works",
			headerValue: "ApiKey abc123 extra",
			wantKey:     "abc123",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.headerValue != "" {
				headers.Set("Authorization", tt.headerValue)
			}

			gotKey, err := GetAPIKey(headers)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error=%v, got err=%v", tt.wantErr, err)
			}

			if gotKey != tt.wantKey {
				t.Fatalf("expected key=%q, got %q", tt.wantKey, gotKey)
			}
		})
	}
}
