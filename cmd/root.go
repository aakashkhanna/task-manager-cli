/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-manager-cli",
	Short: "A brief description of your application",
	Long:  `A CLI Task Manager is an ideal project to learn Go's practical features, including command-line parsing, file I/O, data persistence, and basic concurrency. Below are comprehensive requirements and suggestions, informed by best practices and real-world examples.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "d", false, "Do you just want to dry run the command.")
	viper.BindPFlag("dry-run", rootCmd.PersistentFlags().Lookup("dry-run"))
}
