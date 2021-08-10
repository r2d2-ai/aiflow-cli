package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/r2d2-ai/ai-box-cli/api"
	"github.com/r2d2-ai/ai-box-cli/common"
	"github.com/r2d2-ai/ai-box-cli/util"
	"github.com/spf13/cobra"
)

var updateAll bool

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&updateAll, "all", "", false, "update all contributions")
}

const (
	fJsonFile = "AIflow.json"
)

var updateCmd = &cobra.Command{
	Use:   "update [flags] <contribution|dependency>",
	Short: "update a project contribution/dependency",
	Long:  `Updates a contribution or dependency in the project`,
	Run: func(cmd *cobra.Command, args []string) {

		updatePackage(common.CurrentProject(), args, updateAll)

	},
}

func updatePackage(project common.AppProject, args []string, all bool) {

	if !all {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Contribution not specified")
			os.Exit(1)
		}
		err := api.UpdatePkg(project, args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating contribution/dependency: %v\n", err)
			os.Exit(1)
		}

	} else {
		//Get all imports
		imports, err := util.GetAppImports(filepath.Join(project.Dir(), fJsonFile), project.DepManager(), true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating all contributions: %v\n", err)
			os.Exit(1)
		}
		//Update each package in imports
		for _, imp := range imports.GetAllImports() {

			err = api.UpdatePkg(project, imp.GoGetImportPath())

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error updating contribution/dependency: %v\n", err)
				os.Exit(1)
			}
		}
	}

}
