package handlers

import (
	"arctic/project"
	"arctic/util/command"
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	qsMove []*survey.Question
)

func init() {
	qsMove = []*survey.Question{
		{
			Name: "tag",
			Prompt: &survey.Select{
				Message: "New Tag:",
				Options: tagChoices,
			},
		},
	}
}

type answerMove struct {
	Tag string
}

func HandleMove(cmd *cobra.Command, args []string) {
	fmt.Println("Starting to move")
	current, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting working dir")
		return
	}

	// discarding project obj, will use it when i add central store
	_, err = project.ReadConfig(path.Join(current, ".arctic.toml"))
	if err != nil {
		fmt.Printf("Error reading config: %s", err.Error())
		return
	}

	answers := answerMove{}

	err = survey.Ask(qsMove, &answers)

	if err != nil {
		fmt.Printf("Error while asking questions: %s", err.Error())
		return
	}

	s := fmt.Sprintf("mv %s %s/dev/%s", current, homeLocation, answers.Tag)

	move := command.ParseString(s)
	fmt.Println(s)
	err = move.Execute()

	if err != nil {
		fmt.Printf("Error moving folder")
		return
	}

}
