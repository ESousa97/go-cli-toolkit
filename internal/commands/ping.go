package commands

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/ESousa97/go-cli-toolkit/internal/config"
	"github.com/ESousa97/go-cli-toolkit/internal/ui"
	"github.com/spf13/cobra"
)

const (
	pingUse     = "ping [url...]"
	pingShort   = "Verifica se um ou mais hosts estão acessíveis"
	pingLong    = `Verifica se as URLs informadas estão acessíveis através de uma requisição HTTP GET concorrente. 
Se nenhuma URL for passada, utiliza a lista de 'hosts favoritos' do config.yaml.`
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
	RunE: func(cmd *cobra.Command, args []string) error {
		targets := args
		if len(targets) == 0 {
			targets = config.GetConfig().Hosts
		}

		if len(targets) == 0 {
			return fmt.Errorf("nenhum host informado e nenhum host favorito encontrado no config.yaml")
		}

		return runPing(targets)
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

	fmt.Printf("Iniciando ping em %d hosts...\n\n", len(urls))

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(ui.ColorGray)).
		Headers("HOST", "STATUS", "CODE", "DETAILS")

	var successes, failures int
	for res := range resultsChan {
		statusText := ui.StatusFailText
		codeStr := "---"
		details := res.message

		if res.online {
			successes++
			statusText = ui.StatusOkText
			codeStr = fmt.Sprintf("%d", res.status)
			details = "OK"
		} else {
			failures++
			if res.status != 0 {
				codeStr = fmt.Sprintf("%d", res.status)
			}
		}

		t.Row(res.url, statusText, codeStr, details)
	}

	fmt.Println(t.Render())

	fmt.Printf("\n--- Resumo ---\n")
	fmt.Printf("Sucessos: %d\n", successes)
	fmt.Printf("Falhas:   %d\n", failures)
	fmt.Printf("Total:    %d\n", len(urls))

	return nil
}

// pingHostConcurrently performs a single HTTP GET request to the targetURL
// and returns a [pingResult] with connectivity details.
func pingHostConcurrently(targetURL string) pingResult {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	// Garantir protocolo para evitar erros comuns de iniciante
	fullURL := targetURL
	if len(targetURL) < 4 || (targetURL[:4] != "http") {
		fullURL = "http://" + targetURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
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

