package cli

import (
	"github.com/spf13/cobra"

	"arctic/cli/handlers"
)

var versionController = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Arctic",
	Long:  "All software has versions. This is Arctic's",
	Run:   handlers.HandleVersion,
}

var createController = &cobra.Command{
	Use:   "create",
	Short: "Scaffold a project",
	Run:   handlers.HandleCreation,
}
