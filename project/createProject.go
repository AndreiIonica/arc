package project

import (
	"arc/util"
	"fmt"
	"os"
	"path/filepath"
)

func (p *Project) CreatePoject() error {
	// REFACTOR: this seems iffy, especially the file acces stuff,
	//           will try to use golang for that
	projectLocation := fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))

	// index zero because we can have multiple languages
	src := fmt.Sprintf("%s/%s", projectLocation, p.Lang[0])
	dest := fmt.Sprintf("./%s", p.Location)

	err := util.CopyFolder(src, dest)
	if err != nil {
		//lint:ignore ST1005 we actually want this
		// return error because continuing doesn't make sense because there is no folder
		return fmt.Errorf("Error while copying template: %s\n\t", err.Error())
	}

	// go into the project folder in order to execute commands
	os.Chdir(dest)
	// from here on any error will pe printed but wont stop de function

	current, _ := os.Getwd()
	err = p.WriteConfig(filepath.Join(current, ".arc.toml"))
	if err != nil {
		fmt.Printf("\tError while writing config: %s\n\t", err.Error())
	}

	if p.Repo {
		err = gitInit()
		if err != nil {
			return err
		}
	}

	err = p.execCommands(filepath.Join(current, "commands.txt"))
	if err != nil {
		fmt.Printf("\tError while %s\n\t", err.Error())
	}

	return nil
}
