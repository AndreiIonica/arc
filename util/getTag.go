package util

import "strings"

func GetTag(path string, tagChoices []string) string {

	for _, tag := range tagChoices {
		if strings.Contains(path, tag) {
			return tag
		}
	}

	return "None"
}
