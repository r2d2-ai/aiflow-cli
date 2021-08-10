package commands

import (
	"fmt"
	"os"

	"github.com/r2d2-ai/ai-box-cli/api"
	"github.com/spf13/cobra"
)

var AIflowJsonPath string
var coreVersion string

func init() {
	CreateCmd.Flags().StringVarP(&AIflowJsonPath, "file", "f", "", "specify a AIflow.json to create project from")
	CreateCmd.Flags().StringVarP(&coreVersion, "cv", "", "", "specify core library version (ex. master)")
	rootCmd.AddCommand(CreateCmd)
}

var CreateCmd = &cobra.Command{
	Use:              "create [flags] [appName]",
	Short:            "create a AIflow application project",
	Long:             `Creates a AIflow application project.`,
	Args:             cobra.RangeArgs(0, 1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {

		api.SetVerbose(verbose)
		appName := ""
		if len(args) > 0 {
			appName = args[0]
		}

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error determining working directory: %v\n", err)
			os.Exit(1)
		}
		_, err = api.CreateProject(currentDir, appName, AIflowJsonPath, coreVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
			os.Exit(1)
		}
	},
}
