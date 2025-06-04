/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	models "task-manager-cli/models"
	services "task-manager-cli/services"

	"github.com/spf13/cobra"
)

// markStatusCmd represents the markStatus command
var markStatusCmd = &cobra.Command{
	Use:   "mark-status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		status, err := models.ParseStatus(args[1])
		if err != nil {
			fmt.Println("Error parsing status:", err)
			return
		}
		id, err := services.MarkTask(taskId, status)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Printf("Task updated successfully with id: %s!\n", id)
	},
}

func init() {
	rootCmd.AddCommand(markStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
