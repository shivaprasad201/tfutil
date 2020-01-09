package cmd

import (
	"fmt"
	"os"

	h "tfutil/helpers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tfutil",
	Short: "Tfutil",
	Long:  `A simple utility to help working with Terraform.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: \n  tfutil --help or tfutil -h for help")
	},
}

//Execute handles root command execution
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		h.HandleError(err, "Error: Unable to execute the root command.")
		os.Exit(1)
	}
}
