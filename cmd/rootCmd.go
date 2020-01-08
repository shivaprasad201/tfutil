package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdBuild, cmdGenConfig)
}

var cmdBuild = &cobra.Command{
	Use:   "build [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		BuildImage(".", "test")
	},
}

var cmdGenConfig = &cobra.Command{
	Use:   "generate-config",
	Short: "Generates a tf config from a module. ",
	Long: `Usage:
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from tfutil")
	},
}

var rootCmd = &cobra.Command{
	Use:   "tfutil",
	Short: "Tfutil",
	Long:  `A simple utility to help working with Terraform configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: \n  tfutil --help or tfutil -h for help")
	},
}

//Execute handles root command execution
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
