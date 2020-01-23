package cmd

import (

	h "tfutil/helpers"

	"github.com/spf13/cobra"
)

// Flag variables
var initPath string

func init() {
	rootCmd.AddCommand(cmdInit)
	cmdInit.PersistentFlags().StringVarP(&initPath, "path", "p", "", "Path to initialise terraform files in.")
}

var cmdInit = &cobra.Command{
	Use:              "init",
	Short:            "Initialise terraform configurations.",
	Long:             `
Use 'init' command to initialise terraform configuration files.
These files can be used to put terraform configurations for usage as a config or a module.`,
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		h.InitFiles(initPath)
	},
}

// Subcommands section
