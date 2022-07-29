package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type UserEnv struct {
	DefaultProjectName  string
	AvaialableTemplates []string
}

func GetUserEnv() (*UserEnv, error) {
	// TODO: XDG compliance
	workdingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	templatesDir := filepath.Join(home, ".scaffold-templates")

	return getUserEnv(workdingDir, templatesDir)
}

// Internal implementation for the exported function.
// The exported function is a thin wrapper around this for testability
func getUserEnv(workingDir string, templatesDir string) (*UserEnv, error) {
	e := &UserEnv{}

	workingDir = filepath.Base(workingDir)
	e.DefaultProjectName = workingDir
	entries, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			e.AvaialableTemplates = append(e.AvaialableTemplates, entry.Name())
		}
	}

	return e, nil
}
