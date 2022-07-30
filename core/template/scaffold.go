package template

import (
	"os"
	"path/filepath"
	"scaffold/core/input"
)

// Scaffold bootstraps a project
func Scaffold(ans *input.UserAnswers) error {
	// TODO: platform independent implementation
	// TODO: replace project name in files
	// TODO: bootstrap.sh script

	dst, err := filepath.Abs(ans.Location)
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	templatePath := filepath.Join(home, ".scaffold-templates", ans.Template)
	err = copyFolder(templatePath, dst)
	return err
}
