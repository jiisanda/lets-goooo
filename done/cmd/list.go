// Package cmd /*
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks in the manager",
	Long: `List all the tasks in the Task manager... 
Usage example: ./bin/done list`,
	Run: func(cmd *cobra.Command, args []string) {
		listTask()
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

func listTask() {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 444)
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		return
	}
	defer file.Close()

	tasks, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error getting t from the file...")
		return
	}

	var taskRecords []Task
	for _, t := range tasks {
		id, _ := strconv.ParseInt(t[0], 10, 64)
		date := t[1]
		task := t[2]
		data := Task{
			ID:   id,
			date: date,
			task: task,
		}
		taskRecords = append(taskRecords, data)
	}

	for _, task := range taskRecords {
		fmt.Printf("%d \t %s \t %s\n", task.ID, task.date, task.task)
	}
}
