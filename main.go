package main

import (
	"os"

	"github.com/vukan322/nuke-node_modules/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
