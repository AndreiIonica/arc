package template

import (
	"path/filepath"
	"scaffold/core/input"
)

// Scaffold bootstraps a project
func Scaffold(ans *input.UserAnswers, templatePaths map[string]string) error {
	// TODO: platform independent implementation
	// TODO: replace project name in files
	// TODO: bootstrap.sh script

	dst, err := filepath.Abs(ans.Location)
	if err != nil {
		return err
	}

	err = copyFolder(templatePaths[ans.Template], dst)
	return err
}
