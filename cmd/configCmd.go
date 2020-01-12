package cmd

import (
	"fmt"

	h "tfutil/helpers"

	"github.com/spf13/cobra"
)

// Flag variables
var source string

func init() {
	rootCmd.AddCommand(cmdConfig)
	cmdConfig.AddCommand(scmdView)
	cmdConfig.PersistentFlags().StringVarP(&source, "source", "s", "", "Source of the terraform configurations to read from.")
}

var cmdConfig = &cobra.Command{
	Use:              "config",
	Short:            "Work with terraform configurations.",
	Long:             `Use 'config' command to work with terraform configuration files.`,
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: type `tfutil config -h` `tfutil config --help` for options.")
	},
}

// Subcommands section

var scmdView = &cobra.Command{
	Use:   "describe",
	Short: "Describes the terraform configuration.",
	Long:  `Use 'tfutil config describe' command to describe the terraform configuration from a .`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		h.RenderTable(h.FindResource(h.ReadSource(source)))
	},
}
