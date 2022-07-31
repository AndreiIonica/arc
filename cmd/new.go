package cmd

import (
	"os"
	"path/filepath"

	"scaffold/core/input"
	"scaffold/core/template"

	"github.com/spf13/cobra"
)

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "scaffold a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDir, err := os.Getwd()
		if err != nil {
			return err
		}
		defaultProjectName := filepath.Base(workingDir)

		userHome, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		templatePaths, err := template.LoadTemplates(filepath.Join(userHome, ".scaffold-templates"))
		if err != nil {
			return err
		}

		userAnswers, err := input.AskQuestions(defaultProjectName, template.GetTemplateNames(templatePaths))
		if err != nil {
			return err
		}

		err = template.Scaffold(userAnswers, templatePaths)
		return err
	},
}
