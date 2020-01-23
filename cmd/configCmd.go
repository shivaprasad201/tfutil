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
	cmdConfig.AddCommand(scmdList)
	scmdList.PersistentFlags().StringVarP(&source, "source", "s", "", "Source of the terraform configurations to read from.")
	scmdList.PersistentFlags().StringVarP(&blockKey, "block", "b", "", "Terraform configuration blocks to describe.")
}

var cmdConfig = &cobra.Command{
	Use:              "config",
	Short:            "Work with terraform configurations.",
	Long:             `
Use 'config' command to work with terraform configuration files.`,
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: type `tfutil config -h` `tfutil config --help` for options.")
	},
}

// Subcommands section

var scmdList = &cobra.Command{
	Use:   "list",
	Short: "Lists blocks in the terraform configuration.",
	Long:  `
Use 'tfutil config list' command to list the blocks in terraform configuration from a source or a path.`,
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

