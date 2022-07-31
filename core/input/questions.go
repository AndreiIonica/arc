package input

import (
	"github.com/AlecAivazis/survey/v2"
)

type UserAnswers struct {
	ProjectName string
	Template    string
	Location    string
}

// Use survey for interactive answers
func AskQuestions(defaultProjectName string, templates []string) (*UserAnswers, error) {
	// TODO: use flags AND questions for automation
	// survey question struct
	var qs = []*survey.Question{
		{
			Name: "projectName",
			Prompt: &survey.Input{
				Message: "What is the project name?",
				Default: defaultProjectName,
			},
			Validate: survey.Required,
		},
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "Choose a tempalte:",
				Options: templates,
			},
			Validate: survey.Required,
		},
		{
			Name: "location",
			Prompt: &survey.Input{
				Message: "Where to create project:",
				Default: ".",
			},
			Validate: survey.Required,
		},
	}
	answers := &UserAnswers{}
	err := survey.Ask(qs, answers)
	if err != nil {
		return nil, err
	}
	return answers, err
}
