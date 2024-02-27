package template

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"scaffold/core/input"
)

// Scaffold bootstraps a project
func Scaffold(ans *input.UserAnswers, templatePaths map[string]string) error {
	// TODO: platform independent implementation

	dest, err := filepath.Abs(ans.Location)
	if err != nil {
		return err
	}

	err = copyFolder(templatePaths[ans.Template], dest)

	if err != nil {
		return err
	}

	if runtime.GOOS == "linux" {
		fmt.Println("Detected linux, will run shell scripts using BASH")
		scriptPath := filepath.Join(dest, "scaffold-bootstrap.sh")

		if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
			// No script exists
			fmt.Println("No bootscrap script detected. SKIPPING...")
			return nil
		}

		// Run the command and pipe the output in real-time
		bootstrap := exec.Command("/bin/sh", scriptPath)
		bootstrap.Stdout = os.Stdout
		bootstrap.Stdin = os.Stdin
		bootstrap.Stderr = os.Stderr

		err = bootstrap.Run()

		return err
	}
	return nil
}
