package cmd

import (
	"scaffold/core/env"
	"scaffold/core/input"
	"scaffold/core/template"

	"github.com/spf13/cobra"
)

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "scaffold a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		userEnv, err := env.GetUserEnv()
		if err != nil {
			return err
		}
		userAnswers, err := input.AskQuestions(userEnv)
		if err != nil {
			return err
		}
		err = template.Scaffold(userAnswers)
		return err
	},
}
