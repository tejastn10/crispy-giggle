package main

import (
	"crispy-giggle/commands"
	"fmt"
	"os"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
