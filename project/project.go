package project

import (
	"arctic/util"
	"arctic/util/command"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Project struct {
	Name     string
	Lang     string
	Tag      string
	Repo     bool
	Location string
}

func (p *Project) CreatePoject() error {
	// REFACTOR: this seems iffy, especially the file acces stuff,
	//           will try to use go for that
	projectLocation := fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))

	src := fmt.Sprintf("%s/%s", projectLocation, p.Lang)
	dest := fmt.Sprintf("./%s", p.Location)

	err := util.CopyFolder(src, dest)
	if err != nil {
		s := fmt.Sprintf("Error while copying template: %s", err.Error())
		return errors.New(s)
	}

	// Go into the project folder in order to execute commands
	os.Chdir(dest)

	current, _ := os.Getwd()
	if p.Repo {
		c := command.ParseString("git init")
		err = c.Execute()
		if err != nil {
			s := fmt.Sprintf("Error while running git init: %s", err.Error())
			return errors.New(s)
		}
	}

	err = util.RunCommands(filepath.Join(current, "commands.txt"))
	if err != nil {
		s := fmt.Sprintf("Error while executing commands: %s", err.Error())
		return errors.New(s)
	}

	return nil
}
