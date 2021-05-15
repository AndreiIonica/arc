package command

import (
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Program string
	Argv    []string
}

func ParseString(str string) Command {
	c := Command{}
	arguments := strings.Split(str, " ")
	c.Program = arguments[0]
	c.Argv = arguments[1:]

	return c
}

func (cmd *Command) Execute() error {

	executor := exec.Command(cmd.Program, cmd.Argv...)
	executor.Stdout = os.Stdout

	err := executor.Run()

	if err != nil {
		return err
	}

	return nil
}
