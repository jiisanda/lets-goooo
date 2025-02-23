// Package cmd /*
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var taskID string

// scratchCmd represents the scratch command
var scratchCmd = &cobra.Command{
	Use:   "scratch",
	Short: "Scratch a task as done.",
	Long: `Scratch a task as done. this will delete the task from the task manager.
Usage example: ./bin/done scratch -t-id <task_id>`,
	Run: func(cmd *cobra.Command, args []string) {
		taskScratch(taskID)
	},
}

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleaning the task list.",
	Long: `Clean the task manager. will scratch all the tasks in task manager.
Usage example: ./bin/done clean`,
	Run: func(cmd *cobra.Command, args []string) {
		taskClean()
	},
}

func init() {
	rootCmd.AddCommand(scratchCmd)
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scratchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	scratchCmd.Flags().StringVarP(&taskID, "task-id", "t", "", "Task ID from the csv: Get task id from `done list`")
	err := addCmd.MarkFlagRequired("task-id")
	if err != nil {
		return
	}
}

func taskScratch(taskID string) {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error in opening file: tasks.csv")
		return
	}
	defer file.Close()

	tasks, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	idToDelete, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		fmt.Println("Invalid task id: ", taskID)
		return
	}

	var updatedTasks [][]string
	for _, task := range tasks {
		id, _ := strconv.ParseInt(task[0], 10, 64)
		if id != idToDelete {
			updatedTasks = append(updatedTasks, task)
		}
	}

	file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		fmt.Println("Error in opening file: tasks.csv")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range updatedTasks {
		if err := writer.Write(task); err != nil {
			fmt.Println("Error writing to file: ", err)
		}
	}

	fmt.Println("Task deleted: ", taskID)
}

func taskClean() {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error in opening file: tasks.csv")
		return
	}
	defer file.Close()

	tasks, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks to clean.")
		return
	}

	header := tasks[0]
	file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error in operating file: tasks.csv")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing to file: ", err)
	}

	fmt.Println("All tasks cleaned except the header.")
}
