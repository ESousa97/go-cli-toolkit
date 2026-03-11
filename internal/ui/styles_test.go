package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylesInitialization(t *testing.T) {
	// styles.go defines variables and style constants,
	// a basic test ensures they are loaded/rendered without panic.

	// Verifica se as cores não estão vazias internamente
	assert.NotNil(t, ColorPrimary)
	assert.NotNil(t, ColorSuccess)
	assert.NotNil(t, ColorError)

	// Verifica a renderização das constantes de status
	assert.Contains(t, StatusOkText, "ONLINE", "Success constant should contain ONLINE text")
	assert.Contains(t, StatusFailText, "OFFLINE", "Error constant should contain OFFLINE text")

	// Como é UI, a lógica aqui é mais para cobrir a sintaxe Go e evitar regressões
	// que quebrem a compilação por erros de tipo.
}
