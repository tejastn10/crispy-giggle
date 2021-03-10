package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd remvoes the task from the task list
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Crispy-giggle removes your task from your task list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
