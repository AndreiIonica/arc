package cmd

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var versionController = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Arctic",
	Long:  `All software has versions. This is Arctic's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Arctic Generator v0.1")
	},
}

var createController = &cobra.Command{
	Use:   "create",
	Short: "Scaffold a project",
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			// ^([A-Za-z\-\_\d\# \-])+$
			matched, err := regexp.MatchString(`^([A-Za-z\-\_\d\# \-])+$`, input)
			if err != nil {
				return errors.New("Error in regex...")
			}

			if !matched {
				return errors.New("Invalid name...")
			}
			return nil

		}

		prompt := promptui.Prompt{
			Label:    "Name",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
	},
}
