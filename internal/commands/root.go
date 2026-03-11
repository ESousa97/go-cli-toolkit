package commands

import (
	"github.com/ESousa97/go-cli-toolkit/internal/config"
	"github.com/spf13/cobra"
)

const (
	toolkitUse   = "toolkit"
	toolkitShort = "Toolkit é uma CLI utilitária"
	toolkitLong  = `Toolkit é uma Interface de Linha de Comando (CLI) utilitária em Go.
Use os subcomandos para iterar sobre as funcionalidades disponíveis.`
)

// rootCmd representa o comando base quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:           toolkitUse,
	Short:         toolkitShort,
	Long:          toolkitLong,
	SilenceErrors: true, // O main.go já trata e loga o erro
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		if err := config.InitConfig(); err != nil {
			// Em uma CLI, erros de configuração podem ser reportados mas nem sempre impedem a execução
			// Dependendo da criticidade, poderíamos usar panic ou exit.
		}
	})
}

