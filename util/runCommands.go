package util

import (
	"arctic/util/command"

	"bufio"
	"errors"
	"fmt"
	"os"
)

func RunCommands(filepath string) error {
	fmt.Println("Starting to execute")
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("error while reading file")
	}
	defer file.Close()

	in := bufio.NewScanner(file)

	for in.Scan() {
		text := in.Text()
		cmd := command.ParseString(text)
		err = cmd.Execute()
		if err != nil {
			fmt.Printf(`Error while executing "%s": %s`, text, err.Error())
		}
	}

	return nil
}
