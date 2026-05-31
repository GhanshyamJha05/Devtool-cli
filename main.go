package main

import (
	"github.com/GhanshyamJha05/devtool-cli/cmd"
)

func main() {
	// In clean CLI architecture, main.go is kept completely bare.
	// Its only job is to delegate execution to the root command.
	cmd.Execute()
}
