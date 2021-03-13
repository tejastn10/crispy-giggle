package commands

import (
	"crispy-giggle/database"

	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd remvoes the task from the task list
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Removes your task from your task list 📝",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("❌ Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := database.GetTasks()
		if err != nil {
			fmt.Println("❌ Something wemt Wrong!", err)
			return
		}
		for _, id := range ids {
			if (id <= 0) || (id > len(tasks)) {
				fmt.Println("❗ Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := database.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("❌ Failed to mark \"%d\" as completed. Error: %s \n", id, err)
			} else {
				fmt.Printf("✅ Marked \"%d\" as complete\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
