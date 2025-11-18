package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/ui"
	"github.com/vukan322/nuke-node_modules/internal/util"
)

var yesFlag = false

var nukeCmd = &cobra.Command{
	Use:   "nuke <path>",
	Short: "Delete node_modules directories",
	Long: "Nuke permanently deletes node_modules directories matching the criteria.\n\n" +
		"This command scans first, shows what will be deleted, then asks for confirmation\n" +
		"before proceeding (unless -y flag is used).",
	Example: "  # Delete old node_modules with confirmation\n" +
		"  nukenm nuke ~/Documents\n\n" +
		"  # Skip confirmation prompt\n" +
		"  nukenm nuke ~/Projects -y\n\n" +
		"  # Aggressive cleanup (0 days old)\n" +
		"  nukenm nuke ~/old-code --days 0 -y",
	Args: cobra.ExactArgs(1),
	RunE: runNuke,
}

func init() {
	nukeCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Skip confirmation prompt")
	rootCmd.AddCommand(nukeCmd)
}

func runNuke(cmd *cobra.Command, args []string) error {
	path := args[0]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s: %s\n", red("Error"), err)
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "%s: %s\n", red("Scan failed"), err)
		os.Exit(1)
	}

	if len(results.Folders) == 0 {
		fmt.Println(yellow("No node_modules folders found matching criteria"))
		os.Exit(2)
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
		if deleted.TotalCount > 0 {
			fmt.Printf("\n%s %d node_modules folder(s)\n", green("Deleted"), deleted.TotalCount)
			fmt.Printf("Total space: %s\n", green(util.FormatSize(deleted.TotalSize)))
		}
		fmt.Fprintf(os.Stderr, "\n%s %s\n", red("Error:"), err)
		if deleted.TotalCount > 0 {
			fmt.Fprintf(os.Stderr, "%s\n", yellow("Partial success: some folders were deleted successfully."))
		}
		os.Exit(1)
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
