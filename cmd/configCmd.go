package cmd

import (
	"fmt"

	h "tfutil/helpers"

	"github.com/spf13/cobra"
)

// Flag variables
var source string
var blockKey string

func init() {
	rootCmd.AddCommand(cmdConfig)
	cmdConfig.AddCommand(scmdDescribe, scmdGenerate, scmdInit)
	cmdConfig.PersistentFlags().StringVarP(&source, "source", "s", "", "Source of the terraform configurations to read from.")
	scmdDescribe.PersistentFlags().StringVarP(&blockKey, "block", "b", "", "Terraform configuration blocks to describe.")
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

var scmdDescribe = &cobra.Command{
	Use:   "describe",
	Short: "Describes the terraform configuration.",
	Long:  `Use 'tfutil config describe' command to describe the terraform configuration from a source or a path.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		d := h.FindResource(h.ReadSource(source))
		switch blockKey {
		case "all":
			for k := range d {
				h.MakeTable(h.GetTableData(d, k))
			}
	
		default:
			h.MakeTable(h.GetTableData(d, blockKey))
		}
	},
}

var scmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generates the terraform configuration.",
	Long:  `Use 'tfutil config generate' command to generate the terraform configuration from a source module.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// h.RenderTable(h.FindResource(h.ReadSource(source)))
	},
}

var scmdInit = &cobra.Command{
	Use:   "init",
	Short: "Initialises a set of empty terraform files to put configurations in.",
	Long:  `Use 'tfutil config init' command to create terrraform files.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// h.RenderTable(h.FindResource(h.ReadSource(source)))
	},
}
