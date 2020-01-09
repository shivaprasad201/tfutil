package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdModule)
}

var cmdModule = &cobra.Command{
	Use:   "module",
	Short: "Work with teraform modules.",
	Long: `Use 'module' command to work with terraform modules.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from tfutil module.")
		fmt.Println(args[0])
	},
}

// Subcommands section