package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config define a estrutura de configuração do toolkit
type Config struct {
	Hosts []string `mapstructure:"hosts"`
}

var globalConfig Config

// InitConfig inicializa o Viper para carregar as configurações
func InitConfig() error {
	viper.SetConfigName("config") // nome do arquivo (sem extensão)
	viper.SetConfigType("yaml")   // ou viper.SetConfigType("yml")

	// Caminhos de busca
	viper.AddConfigPath(".") // Diretório atual

	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(home, ".toolkit"))
	}

	// Tenta ler o arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Arquivo não encontrado, podemos ignorar ou definir padrões
			return nil
		}
		return fmt.Errorf("erro ao ler config.yaml: %w", err)
	}

	// Mapeia para a struct
	if err := viper.Unmarshal(&globalConfig); err != nil {
		return fmt.Errorf("erro ao mapear configuração: %w", err)
	}

	return nil
}

// GetConfig retorna a configuração global carregada
func GetConfig() Config {
	return globalConfig
}
