package commands

import (
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

// Execute adiciona todos os comandos filhos ao root command e configura flags de forma apropriada.
// Isso é chamado pela função main.main(). Esta função precisa ser executada apenas uma vez pelo rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Variáveis e Flags globais podem ser inicializadas aqui
}
