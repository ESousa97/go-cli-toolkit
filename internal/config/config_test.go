package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	// Prepare a temporary directory to act as a test environment
	tempDir := t.TempDir()

	// Change working directory for test scope so Viper finds config.yaml here.
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	// Reset Viper and globalConfig (useful for parallel tests)
	viper.Reset()
	globalConfig = Config{}

	// Test 1: File doesn't exist (should not return error, InitConfig ignores ConfigFileNotFoundError)
	err := InitConfig()
	assert.NoError(t, err, "Initialization without config.yaml should be handled without error")
	assert.Empty(t, GetConfig().Hosts, "Host list should be empty")

	// Test 2: Create valid config.yaml file
	configContent := []byte("hosts:\n  - test.com\n  - example.org")
	os.WriteFile(filepath.Join(tempDir, "config.yaml"), configContent, 0644)

	viper.Reset()
	globalConfig = Config{}

	err = InitConfig()
	assert.NoError(t, err, "Should not have error loading valid config")

	conf := GetConfig()
	assert.Len(t, conf.Hosts, 2)
	assert.Equal(t, "test.com", conf.Hosts[0])
	assert.Equal(t, "example.org", conf.Hosts[1])
}
