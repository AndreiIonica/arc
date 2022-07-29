package cmd

import (
	"fmt"
	"scaffold/core/env"
	"scaffold/core/input"

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
		fmt.Println(userAnswers)
		return nil
	},
}
