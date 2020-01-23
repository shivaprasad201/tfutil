package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdModule)
	cmdModule.AddCommand(scmdGenerateConfig, scmdDescribe)
}

var cmdModule = &cobra.Command{
	Use:   "module",
	Short: "Work with teraform modules.",
	Long: `
Use 'module' command to work with terraform modules.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: type `tfutil module -h` `tfutil module --help` for options.")
	},
}

// Subcommands section

var scmdDescribe = &cobra.Command{
	Use:   "describe",
	Short: "Describes the terraform module.",
	Long:  `
Use 'tfutil module describe' command to get an overview of the terraform module.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// h.RenderTable(h.FindResource(h.ReadSource(source)))
	},
}

var scmdGenerateConfig = &cobra.Command{
	Use:   "generate-config",
	Short: "Generates the terraform module configuration.",
	Long:  `
Use 'tfutil module generate-config' command to generate the terraform configuration from a module.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// h.RenderTable(h.FindResource(h.ReadSource(source)))
	},
}