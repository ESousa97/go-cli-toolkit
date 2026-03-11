package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatJSONBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    string
		wantErr bool
	}{
		{
			name:    "JSON válido simples",
			input:   []byte(`{"chave":"valor","ativo":true}`),
			want:    "{\n  \"ativo\": true,\n  \"chave\": \"valor\"\n}",
			wantErr: false,
		},
		{
			name:    "JSON com array",
			input:   []byte(`[1,2,3]`),
			want:    "[\n  1,\n  2,\n  3\n]",
			wantErr: false,
		},
		{
			name:    "JSON inválido",
			input:   []byte(`{"chave":valor}`),
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatJSONBytes(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
