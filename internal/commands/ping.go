package commands

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ESousa97/go-cli-toolkit/internal/config"
	"github.com/ESousa97/go-cli-toolkit/internal/ui"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

const (
	pingUse   = "ping [url...]"
	pingShort = "Checks if one or more hosts are accessible"
	pingLong  = `Checks if the provided URLs are accessible via concurrent HTTP GET requests. 
If no URLs are provided, it uses the 'favorite hosts' list from the configuration file.`
	pingTimeout = 5 * time.Second
)

type pingResult struct {
	url     string
	online  bool
	status  int
	message string
}

// pingCmd represents the ping command
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
			return fmt.Errorf("no host provided and no favorite hosts found in config repository")
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

	// Close channel when all routines finish
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	fmt.Printf("Starting ping for %d hosts...\n\n", len(urls))

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

	// Ensure protocol to avoid common beginner errors
	fullURL := targetURL
	if len(targetURL) < 4 || (targetURL[:4] != "http") {
		fullURL = "http://" + targetURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return pingResult{url: targetURL, online: false, message: "error creating request"}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return pingResult{url: targetURL, online: false, message: "host unreachable"}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return pingResult{url: targetURL, online: true, status: resp.StatusCode}
	}

	return pingResult{url: targetURL, online: false, status: resp.StatusCode, message: "invalid status code"}
}
