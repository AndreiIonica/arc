package project

import (
	"arctic/util/command"
	"fmt"
)

func gitInit() error {
	c := command.ParseString("git init")
	err := c.Execute()
	if err != nil {
		//lint:ignore ST1005 we actually want this
		return fmt.Errorf("Error while running git init: %s", err.Error())
	}
	return nil
}
