package ui

import (
	"fmt"

	"github.com/vukan322/nuke-node_modules/internal/scanner"
	"github.com/vukan322/nuke-node_modules/internal/util"
)

func PrintResults(result *scanner.ScanResult, isDeleted bool) {
	if result.TotalCount == 0 {
		fmt.Println("No node_modules folders found")
		return
	}

	action := "Found"
	if isDeleted {
		action = "Deleted"
	}

	fmt.Printf("\n%s %d node_modules folder(s)\n", action, result.TotalCount)
	fmt.Printf("Total space: %s\n", util.FormatSize(result.TotalSize))

	if !isDeleted {
		fmt.Printf("\nFolders:\n")
		for _, folder := range result.Folders {
			fmt.Printf("  %s (%s)\n", folder.Path, util.FormatSize(folder.Size))
		}
	}
}
