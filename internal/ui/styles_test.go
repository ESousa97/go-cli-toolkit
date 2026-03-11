package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylesInitialization(t *testing.T) {
	// Como styles.go define apenas variáveis e constantes de estilo, 
	// um teste básico é garantir que eles foram renderizados/carregados sem panic.

	// Verifica se as cores não estão vazias internamente
	assert.NotNil(t, ColorPrimary)
	assert.NotNil(t, ColorSuccess)
	assert.NotNil(t, ColorError)
	
	// Verifica a renderização das constantes de status
	assert.Contains(t, StatusOkText, "ONLINE", "A constante de sucesso deve conter texto ONLINE")
	assert.Contains(t, StatusFailText, "OFFLINE", "A constante de erro deve conter texto OFFLINE")

	// Como é UI, a lógica aqui é mais para cobrir a sintaxe Go e evitar regressões
	// que quebrem a compilação por erros de tipo.
}
