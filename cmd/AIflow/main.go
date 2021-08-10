package main

import (
	"fmt"
	"os"

	"github.com/r2d2-ai/aiflow-cli/commands"
	"github.com/r2d2-ai/aiflow-cli/util"
)

// Not set by default, will be filled by init() function in "./currentversion.go" file, if it exists.
// This latter file is generated with a "go generate" command.
var Version string = ""

//go:generate go run gen/version.go
func main() {

	if util.GetGoPath() == "" {
		_, _ = fmt.Fprintf(os.Stderr, "Error: GOPATH must be set before running AIflow cli\n")
		os.Exit(1)
	}

	//Initialize the commands
	_ = os.Setenv("GO111MODULE", "auto")
	commands.Initialize(Version)
	commands.Execute()
}
