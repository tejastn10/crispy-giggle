package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd shows the task list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Crispy-giggle shows your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
