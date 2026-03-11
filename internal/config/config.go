package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config defines the toolkit configuration structure mapped from yaml.
type Config struct {
	Hosts []string `mapstructure:"hosts"`
}

var globalConfig Config

// InitConfig initializes Viper to discover and load the configuration file.
// It looks for config.yaml in the current directory or $HOME/.toolkit.
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

// GetConfig returns a copy of the globally loaded [Config].
func GetConfig() Config {
	return globalConfig
}
