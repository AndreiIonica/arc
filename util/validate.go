package util

import (
	"errors"
	"regexp"
)

func ValidateName(raw interface{}) error {
	input, ok := raw.(string)

	if !ok {
		//lint:ignore ST1005 this can be capitalized because it gets printed directly
		return errors.New("Please provide a valid string")
	}

	// allow current folder
	if input == "." {
		return nil
	}

	// Names may only contain letters,numbers,underscores,-,#
	// TODO: rewrite regex or parse string manually
	matched, err := regexp.MatchString(`^([A-Za-z\-\_\d\# \-])+$`, input)
	if err != nil {
		//lint:ignore ST1005 this can be capitalized because it gets printed directly
		return errors.New("Error in regex")
	}

	if !matched {
		//lint:ignore ST1005 this can be capitalized because it gets printed directly
		return errors.New("Invalid name")
	}
	return nil
}
