package util

import "io/ioutil"

func GetProjects(projectLocation string) []string {

	files, err := ioutil.ReadDir(projectLocation)

	if err != nil {
		return nil
	}
	names := []string{}

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names

}
