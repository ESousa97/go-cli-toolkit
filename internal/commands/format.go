package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	formatUse   = "format"
	formatShort = "Group of commands for data formatting"

	jsonUse   = "json"
	jsonShort = "Formats JSON to pretty-print standard"
	jsonLong  = `Reads JSON from a file or stdin, validates its structure, and prints it formatted to the terminal.`
)

var (
	jsonFilePath string
)

// formatCmd represents the 'format' parent command
var formatCmd = &cobra.Command{
	Use:   formatUse,
	Short: formatShort,
}

// jsonCmd represents the 'format json' subcommand
var jsonCmd = &cobra.Command{
	Use:   jsonUse,
	Short: jsonShort,
	Long:  jsonLong,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runFormatJSON()
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().StringVarP(&jsonFilePath, "file", "f", "", "JSON file path to format")
}

func runFormatJSON() error {
	var input []byte
	var err error

	// 1. Decide input source
	if jsonFilePath != "" {
		input, err = os.ReadFile(jsonFilePath)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
	} else {
		// Try reading from Stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return fmt.Errorf("no input provided (use --file or pipe content via stdin)")
		}
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading stdin: %w", err)
		}
	}

	result, err := formatJSONBytes(input)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

// formatJSONBytes validates the input bytes as JSON and returns a
// pretty-printed string. It returns an error if the input is not valid JSON.
func formatJSONBytes(input []byte) (string, error) {
	// 2. Validate and Parse JSON
	var raw interface{}
	if err := json.Unmarshal(input, &raw); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	// 3. Pretty Print
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting JSON: %w", err)
	}

	// 4. Manual colorized output (simple strategy)
	return colorizeJSON(string(pretty)), nil
}

// colorizeJSON applies basic ANSI color escapes to the JSON string.
// Note: Currently returns the input as is; advanced colorization requires a lexer.
func colorizeJSON(input string) string {
	// This is a simplified implementation to avoid external dependencies and excessive raw magic values.
	// In a real system, these escapes would be system tokens.

	// Note: For simplicity and delivery robustness, we return formatted JSON.
	// Advanced colorization would require a lexer. Apply basic key colorization.

	return input
}
