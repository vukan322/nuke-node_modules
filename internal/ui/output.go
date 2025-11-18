package ui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/util"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func PrintResults(result *scanner.ScanResult, isDeleted bool, quiet bool) {
	if result.TotalCount == 0 {
		if !quiet {
			fmt.Println(yellow("No node_modules folders found"))
		}
		return
	}

	if quiet {
		fmt.Printf("%d folders, %s\n", result.TotalCount, util.FormatSize(result.TotalSize))
		return
	}

	action := "Found"
	colorFunc := cyan
	if isDeleted {
		action = "Deleted"
		colorFunc = green
	}

	fmt.Printf("\n%s %s\n", colorFunc(action), colorFunc(fmt.Sprintf("%d node_modules folder(s)", result.TotalCount)))
	fmt.Printf("Total space: %s\n", colorFunc(util.FormatSize(result.TotalSize)))

	if !isDeleted {
		fmt.Printf("\nFolders:\n")
		for _, folder := range result.Folders {
			fmt.Printf("  %s (%s)\n", folder.Path, util.FormatSize(folder.Size))
		}
	}
}
