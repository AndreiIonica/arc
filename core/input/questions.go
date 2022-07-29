package input

import (
	"scaffold/core/env"

	"github.com/AlecAivazis/survey/v2"
)

type UserAnswers struct {
	Name     string
	Template string
	Location string
}

// Use survey for interactive answers
func AskQuestions(e *env.UserEnv) (*UserAnswers, error) {
	// TODO: use flags AND questions for automation
	// survey question struct
	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the project name?",
				Default: e.DefaultProjectName,
			},
			Validate: survey.Required,
		},
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "Choose a tempalte:",
				Options: e.AvaialableTemplates,
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
