package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/ui"
)

var yesFlag bool

var nukeCmd = &cobra.Command{
	Use:   "nuke <path>",
	Short: "Delete node_modules directories",
	Args:  cobra.ExactArgs(1),
	RunE:  runNuke,
}

func init() {
	nukeCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Skip confirmation prompt")
	rootCmd.AddCommand(nukeCmd)
}

func runNuke(cmd *cobra.Command, args []string) error {
	path := args[0]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}

	s := scanner.New(path, days, verbose, includeHidden)
	results, err := s.Scan()
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	if len(results.Folders) == 0 {
		fmt.Println("No node_modules folders found matching criteria")
		return nil
	}

	ui.PrintResults(results, false)

	if !yesFlag {
		if !confirmDeletion() {
			fmt.Println("Operation cancelled")
			return nil
		}
	}

	fmt.Println("\nDeleting...")
	deleted, err := s.Delete(results)
	if err != nil {
		return fmt.Errorf("deletion failed: %w", err)
	}

	ui.PrintResults(deleted, true)
	return nil
}

func confirmDeletion() bool {
	fmt.Print("\nProceed with deletion? (y/N): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}
