/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	services "task-manager-cli/services"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := services.ListTask()
		if err != nil {
			fmt.Println("Error listing task:", err)
			return
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Description", "Status", "Created At", "Updated At"})
		for i, task := range tasks {
			t.AppendRow(table.Row{
				i + 1,
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format("2006-01-02 15:04"),
				task.UpdatedAt.Format("2006-01-02 15:04"),
			})
		}
		t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
