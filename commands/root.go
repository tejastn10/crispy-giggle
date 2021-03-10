package commands

import "github.com/spf13/cobra"

// RootCmd confguration
var RootCmd = &cobra.Command{
	Use:   "crispy-giggle",
	Short: "Crispy-giggle a Cli Task Manager",
	Long:  `Crispy-giggle your personal companion to manage your daily tasks with a giggle`,
}
