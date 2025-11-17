//go:build ignore

package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/vukan322/nuke-node_modules/cmd"
)

func main() {
	rootCmd := cmd.GetRootCmd()
	err := doc.GenManTree(rootCmd, &doc.GenManHeader{
		Title:   "NUKENM",
		Section: "1",
	}, "./docs/man")
	if err != nil {
		log.Fatal(err)
	}
}
