package handlers

import (
	"arctic/project"
	"arctic/util"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	projectLocation string
	projectChoices  []string
	qs              []*survey.Question
)

func init() {
	// This function is ran first in this file
	// so i can set config vars here
	projectLocation = fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))
	projectChoices = util.GetProjects(projectLocation)

	qs = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Project name:"},
			Validate: util.ValidateName,
		},
		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Project language:",
				Options: projectChoices,
			},
		},
		{
			Name:     "folder",
			Prompt:   &survey.Input{Message: "Where to create:"},
			Validate: util.ValidateName,
		},
		{
			Name: "repo",
			Prompt: &survey.Select{
				Message: "Create git repo?",
				Options: []string{
					"Yes",
					"No",
				},
			},
		},
	}
}

// Type of response
type Answer struct {
	Name   string
	Lang   string `survey:"language"`
	Folder string
	Repo   string
}

func HandleCreation(cmd *cobra.Command, args []string) {
	answers := Answer{}

	err := survey.Ask(qs, &answers)

	if err != nil {
		fmt.Printf("Error while asking questions: %s", err.Error())
		return
	}
	project := project.Project{
		Name:     answers.Name,
		Lang:     answers.Lang,
		Repo:     answers.Repo == "Yes",
		Location: answers.Folder,
	}

	err = project.CreatePoject()
	if err != nil {
		fmt.Printf("Error while creating project:\n\t %s", err.Error())
	}

}
