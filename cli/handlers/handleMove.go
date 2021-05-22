package handlers

import (
	"arctic/project"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func HandleMove(cmd *cobra.Command, args []string) {
	fmt.Println("Starting to move")
	current, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting working dir")
		return
	}

	proj, err := project.ReadConfig(path.Join(current, ".arctic.toml"))
	if err != nil {
		fmt.Printf("Error reading config: %s", err.Error())
		return
	}

	fmt.Println(proj)
	fmt.Println(current)

}
