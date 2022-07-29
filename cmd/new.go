package cmd

import (
	"github.com/spf13/cobra"
)

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "scaffold a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
