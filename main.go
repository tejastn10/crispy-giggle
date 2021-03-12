package main

import (
	"crispy-giggle/commands"
	"crispy-giggle/database"

	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	if err := database.Init(dbPath); err != nil {
		fmt.Println(err.Error())
	}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
