package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdConfig)
	cmdConfig.AddCommand(scmdView)
}

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Work with terraform configurations.",
	Long: `Use 'config' command to work with terraform configuration files.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from tfutil config.")
		fmt.Println(args[0])
	},
}

// Subcommands section

var scmdView = &cobra.Command{
	Use:   "view",
	Short: "prints out the config file",
	Long: `Use 'config' command to work with terraform configuration files.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from tfutil config view.")
		fmt.Println(args[0])
	},
}