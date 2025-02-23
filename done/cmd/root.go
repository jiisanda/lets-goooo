package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "done",
	Short: "done a CLI Task Manager",
	Long:  "done is a CLI Task Manager that helps you to manage your tasks - Yup a TODO application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
