package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/ui"
)

var scanCmd = &cobra.Command{
	Use:   "scan <path>",
	Short: "Scan for node_modules directories (dry run)",
	Args:  cobra.ExactArgs(1),
	RunE:  runScan,
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func runScan(cmd *cobra.Command, args []string) error {
	path := args[0]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", red("Error"), err)
	}

	s := scanner.New(path, days, verbose, includeHidden)

	sp := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	sp.Suffix = " Scanning for node_modules..."
	if !verbose {
		sp.Start()
	}

	results, err := s.Scan()

	if !verbose {
		sp.Stop()
	}

	if err != nil {
		return fmt.Errorf("%s: %w", red("Scan failed"), err)
	}

	ui.PrintResults(results, false)
	return nil
}
