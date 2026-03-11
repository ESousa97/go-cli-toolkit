package commands

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

const (
	pingUse     = "ping [url]"
	pingShort   = "Verifica se um host está acessível"
	pingLong    = `Verifica se a URL informada está acessível através de uma requisição HTTP GET.`
	pingTimeout = 5 * time.Second
)

// pingCmd representa o comando ping
var pingCmd = &cobra.Command{
	Use:   pingUse,
	Short: pingShort,
	Long:  pingLong,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		return pingHost(url)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

// pingHost isola a lógica de negócio do comando de console em si
func pingHost(targetURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição para %s: %w", targetURL, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("host %s está OFFLINE ou inacessível: %w", targetURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Printf("O host %s está ONLINE (Status: %d)\n", targetURL, resp.StatusCode)
	} else {
		fmt.Printf("O host %s retornou erro (Status: %d)\n", targetURL, resp.StatusCode)
	}

	return nil
}
