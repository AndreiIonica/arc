package handlers

import (
	"arctic/util"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	projectLocation string
	projects        []string
	qs              []*survey.Question
)

func init() {
	// This function is ran first in this file
	// so i can set config vars here
	projectLocation = fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))
	projects = util.GetProjects(projectLocation)

	qs = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Project name:"},
			Validate: util.ValidateName,
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Project type:",
				Options: projects,
			},
		},
		{
			Name:     "folder",
			Prompt:   &survey.Input{Message: "Where to create:"},
			Validate: util.ValidateName,
		},
	}
}

// Type of response
type Answer struct {
	Name   string
	Type   string
	Folder string
}

func HandleCreation(cmd *cobra.Command, args []string) {
	answers := Answer{}

	err := survey.Ask(qs, &answers)

	if err != nil {
		fmt.Printf("Error while asking questions: %s", err.Error())
		return
	}

	// REFACTOR: this filename stuff
	src := fmt.Sprintf("%s/%s", projectLocation, answers.Type)
	dest := fmt.Sprintf("./%s", answers.Folder)

	err = util.CopyFolder(src, dest)
	if err != nil {
		fmt.Printf("Error while copying template: %s", err.Error())
		return
	}

	// Go into the project folder in order to execute commands
	os.Chdir(dest)

	current, _ := os.Getwd()

	err = util.RunCommands(filepath.Join(current, "commands.txt"))
	if err != nil {
		fmt.Printf("Error while executing commands: %s", err.Error())
	}
}
