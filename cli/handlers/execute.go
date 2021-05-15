package handlers

import (
	"arctic/cli/handlers/command"

	"bufio"
	"errors"
	"fmt"
	"os"
)

func execute(filepath string) error {
	fmt.Println("Starting to execute")
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("error while reading file")
	}
	defer file.Close()

	in := bufio.NewScanner(file)

	for in.Scan() {
		cmd := command.ParseString(in.Text())
		err = cmd.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
