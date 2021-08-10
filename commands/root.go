package commands

import (
	"fmt"
	"os"

	"github.com/r2d2-ai/ai-box-cli/api"
	"github.com/r2d2-ai/ai-box-cli/common"
	"github.com/r2d2-ai/ai-box-cli/util"
	"github.com/spf13/cobra"
)

const (
	VersionTpl = `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "cli version %s" .Version}}
`
)

var verbose bool

//Root command
var rootCmd = &cobra.Command{
	Use:   "AIflow [flags] [command]",
	Short: "AIflow cli",
	Long:  `AIflow command line interface for AIflow applications`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		preRun(cmd, args, verbose)
	},
}

func Initialize(version string) {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")

	if len(version) > 0 {
		rootCmd.Version = version // use version hardcoded by a "go generate" command
	} else {
		_, rootCmd.Version, _ = util.GetCLIInfo() // guess version from sources in $GOPATH/src
	}

	rootCmd.SetVersionTemplate(VersionTpl)

	//Get the list of commands from the registry of commands and add.
	commandList := common.GetPlugins()

	for _, command := range commandList {

		rootCmd.AddCommand(command)
	}
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func preRun(cmd *cobra.Command, args []string, verbose bool) {
	api.SetVerbose(verbose)
	common.SetVerbose(verbose)

	builtIn := cmd.Name() == "help" || cmd.Name() == "version"

	if len(os.Args) > 1 && !builtIn {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error determining working directory: %v\n", err)
			os.Exit(1)
		}
		appProject := api.NewAppProject(currentDir)

		err = appProject.Validate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error validating project: %v\n", err)
			os.Exit(1)
		}

		common.SetCurrentProject(appProject)
	}
}
