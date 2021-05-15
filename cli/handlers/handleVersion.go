package handlers

import (
	"fmt"

	"github.com/spf13/cobra"
)

func HandleVersion(cmd *cobra.Command, args []string) {
	fmt.Println("Arctic Generator v0.1")
}
