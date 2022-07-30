package template

import (
	"io/ioutil"
	"path/filepath"
)

// Returns a map with the keys being the template name and the value being the path
func LoadTemplates(path string) (map[string]string, error) {
	paths := make(map[string]string)
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			paths[entry.Name()] = filepath.Join(path, entry.Name())
		}
	}
	return paths, nil
}

// Returns a slice of the template names from the map returned by LoadTemplates
func GetTemplateNames(paths map[string]string) []string {
	keys := make([]string, len(paths))

	i := 0
	for k := range paths {
		keys[i] = k
		i++
	}
	return keys
}
