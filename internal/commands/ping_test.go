package commands

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHostConcurrently(t *testing.T) {
	// Cria servidor de teste que retorna HTTP 200 OK
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer serverOK.Close()

	// Cria servidor de teste que retorna HTTP 500 Internal Server Error
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
			name:       "Host Erro 500",
			targetURL:  serverErr.URL,
			wantOnline: false,
			wantStatus: 500,
		},
		{
			name:       "Host Inexistente / Timeout Replicado",
			targetURL:  "http://localhost:59999", // Porta local normalmente fechada, gera connection refused rápido
			wantOnline: false,
			wantStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pingHostConcurrently(tt.targetURL)
			
			assert.Equal(t, tt.targetURL, res.url, "A URL de resposta deve ser a mesma.")
			assert.Equal(t, tt.wantOnline, res.online, "O status online retornado divergiu.")
			assert.Equal(t, tt.wantStatus, res.status, "O Http StatusCode retornado divergiu.")
		})
	}
}
