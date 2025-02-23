// Package cmd /*
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var task string

type Task struct {
	ID   int64
	date string
	task string
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding a task to a task manager",
	Long: `Add a task to a task manager with -t command 
Usage example: ./bin/done add -t "task to add"
`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringVarP(&task, "task", "t", "", "Task to be added")
	err := addCmd.MarkFlagRequired("task")
	if err != nil {
		return
	}
}

func addTask(task string) {
	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id := time.Now().Unix()
	date := time.Now().Format("2006-01-02 15:04:05")

	record := []string{fmt.Sprintf("%d", id), date, task}
	if err := writer.Write(record); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("Task added: ", task)
}
