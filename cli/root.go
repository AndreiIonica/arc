package cli

import (
	"github.com/spf13/cobra"
)

var (
	cli = &cobra.Command{
		Use:   "arc",
		Short: "Smart project generator",
		Long: `arc is a project generator.
This application is a tool to generate the needed files
to quickly create an application.`,
	}
)

func init() {
	cli.AddCommand(versionController)
	cli.AddCommand(createController)
	cli.AddCommand(moveController)
}

// Starts the cli
func Execute() error {
	return cli.Execute()
}
