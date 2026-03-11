package commands

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

const (
	pingUse     = "ping [url...]"
	pingShort   = "Verifica se um ou mais hosts estão acessíveis"
	pingLong    = `Verifica se as URLs informadas estão acessíveis através de uma requisição HTTP GET concorrente.`
	pingTimeout = 5 * time.Second
)

type pingResult struct {
	url     string
	online  bool
	status  int
	message string
}

// pingCmd representa o comando ping
var pingCmd = &cobra.Command{
	Use:   pingUse,
	Short: pingShort,
	Long:  pingLong,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPing(args)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

func runPing(urls []string) error {
	var wg sync.WaitGroup
	resultsChan := make(chan pingResult, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(targetURL string) {
			defer wg.Done()
			resultsChan <- pingHostConcurrently(targetURL)
		}(url)
	}

	// Fecha o channel quando todos terminarem
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var successes, failures int
	fmt.Printf("Iniciando ping em %d hosts...\n\n", len(urls))

	for res := range resultsChan {
		if res.online {
			successes++
			fmt.Printf("[OK]   %s (Status: %d)\n", res.url, res.status)
		} else {
			failures++
			fmt.Printf("[ERRO] %s (%s)\n", res.url, res.message)
		}
	}

	fmt.Printf("\n--- Resumo ---\n")
	fmt.Printf("Sucessos: %d\n", successes)
	fmt.Printf("Falhas:   %d\n", failures)
	fmt.Printf("Total:    %d\n", len(urls))

	return nil
}

// pingHostConcurrently realiza o ping e retorna um objeto de resultado
func pingHostConcurrently(targetURL string) pingResult {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return pingResult{url: targetURL, online: false, message: "erro ao criar requisição"}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return pingResult{url: targetURL, online: false, message: "host inacessível"}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return pingResult{url: targetURL, online: true, status: resp.StatusCode}
	}

	return pingResult{url: targetURL, online: false, status: resp.StatusCode, message: "status code inválido"}
}
