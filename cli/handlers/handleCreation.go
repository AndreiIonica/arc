package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func validateName(raw interface{}) error {
	input, ok := raw.(string)

	if !ok {
		return errors.New("please provide a valid string")
	}
	if input == "." {
		return nil
	}

	// Names may only contain letters,numbers,underscores,-,#
	matched, err := regexp.MatchString(`^([A-Za-z\-\_\d\# \-])+$`, input)
	if err != nil {
		return errors.New("error in regex")
	}

	if !matched {
		return errors.New("invalid name")
	}
	return nil
}

func getProjects(projectLocation string) []string {

	files, err := ioutil.ReadDir(projectLocation)

	if err != nil {
		return nil
	}
	names := []string{}

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names

}

var projectLocation = fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Project name:"},
		Validate: validateName,
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Project type:",
			Options: getProjects(projectLocation),
		},
	},
	{
		Name:     "folder",
		Prompt:   &survey.Input{Message: "Where to create:"},
		Validate: validateName,
	},
}

func HandleCreation(cmd *cobra.Command, args []string) {
	// TODO: refactor this(regarding filenames)
	answers := struct {
		Name   string
		Type   string
		Folder string
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO: right now using system call, will do it the proper way
	initialPath := fmt.Sprintf("%v/%v", projectLocation, answers.Type)
	resultingPath := fmt.Sprintf("./%v", answers.Folder)

	command := exec.Command("cp", "-r", initialPath, resultingPath)
	if err := command.Run(); err != nil {
		fmt.Println("Error in copyinng folder")
	}

	os.Chdir(resultingPath)

	current, _ := os.Getwd()
	err = execute(filepath.Join(current, "commands.txt"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
