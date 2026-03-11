package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	// Preparar um diretório temporário para atuar como ambiente de teste
	tempDir := t.TempDir()

	// Alterar o diretório de trabalho apenas para o escopo desse teste
	// para que o viper procure o config.yaml aqui.
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	// Resetar o viper e a globalConfig (útil se o teste rodar junto com outros)
	viper.Reset()
	globalConfig = Config{}

	// Teste 1: Arquivo não existe (não deve retornar erro, pois o InitConfig ignora ConfigFileNotFoundError)
	err := InitConfig()
	assert.NoError(t, err, "A inicialização sem config.yaml deve ser tratada sem erro")
	assert.Empty(t, GetConfig().Hosts, "A lista de hosts deveria estar vazia")

	// Teste 2: Criar arquivo config.yaml válido
	configContent := []byte("hosts:\n  - test.com\n  - example.org")
	os.WriteFile(filepath.Join(tempDir, "config.yaml"), configContent, 0644)

	viper.Reset()
	globalConfig = Config{}

	err = InitConfig()
	assert.NoError(t, err, "Não deveria haver erro ao carregar config válido")

	conf := GetConfig()
	assert.Len(t, conf.Hosts, 2)
	assert.Equal(t, "test.com", conf.Hosts[0])
	assert.Equal(t, "example.org", conf.Hosts[1])
}
