package project

import (
	"arc/util"
	"fmt"
)

type Project struct {
	Name     string
	Lang     []string
	Tag      string
	Repo     bool
	Location string
}

func (p *Project) execCommands(path string) error {
	err := util.RunCommands(path)
	if err != nil {
		//lint:ignore ST1005 we actually want this
		return fmt.Errorf("Error while executing commands: %s", err.Error())
	}
	return nil
}
