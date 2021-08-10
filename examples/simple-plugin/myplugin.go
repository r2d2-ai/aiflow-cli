package myplugin

import (
	"fmt"

	"github.com/r2d2-ai/aiflow-cli/common" // AIflow CLI support code
	"github.com/spf13/cobra"
)

func init() {
	common.RegisterPlugin(myCmd)
}

var myCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "says hello world",
	Long:  `This plugin command says hello world`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}
