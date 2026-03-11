package commands

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHostConcurrently(t *testing.T) {
	// Create test server returning HTTP 200 OK
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer serverOK.Close()

	// Create test server returning HTTP 500 Internal Server Error
	serverErr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer serverErr.Close()

	tests := []struct {
		name       string
		targetURL  string
		wantOnline bool
		wantStatus int
	}{
		{
			name:       "Host Online (200 OK)",
			targetURL:  serverOK.URL,
			wantOnline: true,
			wantStatus: 200,
		},
		{
			name:       "Host Error 500",
			targetURL:  serverErr.URL,
			wantOnline: false,
			wantStatus: 500,
		},
		{
			name:       "Non-existent Host / Replicated Timeout",
			targetURL:  "http://localhost:59999", // Local port normally closed, generates connection refused quickly
			wantOnline: false,
			wantStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pingHostConcurrently(tt.targetURL)

			assert.Equal(t, tt.targetURL, res.url, "Response URL must match.")
			assert.Equal(t, tt.wantOnline, res.online, "Returned online status mismatch.")
			assert.Equal(t, tt.wantStatus, res.status, "Returned Http StatusCode mismatch.")
		})
	}
}
