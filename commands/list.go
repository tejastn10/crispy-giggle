package commands

import (
	"crispy-giggle/database"

	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd shows the task list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View your task list 🧐",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.GetTasks()
		if err != nil {
			fmt.Println("❌ Something went Wrong!", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks pending  🥳  ")
			return
		}
		fmt.Println("You have the Following tasks pending to be done:  ⏲️  ")
		for i, task := range tasks {
			fmt.Printf("%d. %s \n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
