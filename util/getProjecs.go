package util

import "os"

func GetFolders(path string) ([]string, error) {

	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}
	names := []string{}

	for _, file := range files {
		if file.IsDir() {
			names = append(names, file.Name())
		}

	}

	return names, nil

}
