package commands

import (
	"crispy-giggle/database"

	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmmd adds a new task
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list ➕",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := database.CreateTask(task)
		if err != nil {
			fmt.Println("❌ Something went Wrong!", err)
			return
		}
		fmt.Printf("Added \"%s\" to your Task List ✅\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
