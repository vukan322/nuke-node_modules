package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	verbose       bool
	days          int
	includeHidden bool
	version       = "dev"

	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

var rootCmd = &cobra.Command{
	Use:   "nukenm",
	Short: "Find and delete stale node_modules directories",
	Long: "nukenm is a fast CLI tool to find and delete stale node_modules directories,\n" +
		"freeing up disk space from JavaScript projects you're not actively using.\n\n" +
		"By default, it only targets folders not modified in the last 14 days and\n" +
		"skips hidden directories to protect system tools like .nvm and .npm.",
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed output during scan")
	rootCmd.PersistentFlags().IntVar(&days, "days", 14, "Only process node_modules not modified in N days")
	rootCmd.PersistentFlags().BoolVar(&includeHidden, "include-hidden", false, "Include hidden directories (starting with .)")
}
