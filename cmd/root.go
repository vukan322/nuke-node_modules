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
	Use:     "nukenm",
	Short:   "Find and delete stale node_modules directories",
	Long:    `nukenm scans directories for node_modules folders and helps you reclaim disk space by deleting old, unused ones.`,
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
