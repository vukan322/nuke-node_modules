package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/ui"
)

var (
	yesFlag = false
	redNuke = color.New(color.FgRed).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
)

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
		return fmt.Errorf("%s: %s", redNuke("Error"), err)
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
		return fmt.Errorf("%s: %w", redNuke("Scan failed"), err)
	}

	if len(results.Folders) == 0 {
		fmt.Println(yellow("No node_modules folders found matching criteria"))
		return nil
	}

	ui.PrintResults(results, false)

	if !yesFlag {
		if !confirmDeletion() {
			fmt.Println(yellow("Operation cancelled"))
			return nil
		}
	}

	sp = spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	sp.Suffix = " Deleting..."
	if !verbose {
		sp.Start()
	}

	deleted, err := s.Delete(results)

	if !verbose {
		sp.Stop()
	}

	if err != nil {
		return fmt.Errorf("%s: %w", redNuke("Deletion failed"), err)
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
