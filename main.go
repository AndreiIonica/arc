package main

import (
	"fmt"
	"os"
	"scaffold/cmd"
)

func main() {
	root := cmd.NewRootCmd()
	cmd.AddSubCommands(root)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
