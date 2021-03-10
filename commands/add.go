package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmmd adds a new task
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Crispy-giggle adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your Task List. \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
