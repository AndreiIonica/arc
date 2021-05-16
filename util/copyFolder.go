package util

import (
	"errors"
	"os/exec"
)

func CopyFolder(source, destination string) error {
	// TODO: right now using system call, will do it the proper way
	command := exec.Command("cp", "-r", source, destination)
	if err := command.Run(); err != nil {
		return errors.New("could not copy folder")
	}

	return nil
}
