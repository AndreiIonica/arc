package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "scaffold",
		Short: "Quickly scafold new projects",
	}
}

func AddSubCommands(root *cobra.Command) {
	root.AddCommand(versionCmd)
	root.AddCommand(newProjectCmd)
}
