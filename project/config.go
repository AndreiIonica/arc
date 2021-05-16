package project

import (
	"fmt"
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

type ConfigFile struct {
	Name string   `toml:"name"`
	Tag  string   `toml:"tag"`
	Lang []string `toml:"lang"`
	Repo bool     `toml:"repo"`
}

func (p *Project) writeConfig(path string) error {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not acces config file: %s \n\t", err.Error())
	}

	config := &ConfigFile{}
	err = toml.Unmarshal(bytes, config)

	if err != nil {
		return fmt.Errorf("malformed config file: %s", err.Error())
	}
	config.Name = p.Name
	config.Repo = p.Repo
	// Hardocing it for now, will add later
	// FIXME
	config.Tag = "Working"
	bytes, err = toml.Marshal(config)
	if err != nil {
		return fmt.Errorf("could not encode config file: %s \n\t", err.Error())
	}

	err = ioutil.WriteFile(path, bytes, 0775)

	if err != nil {
		return fmt.Errorf("could not write config file: %s \n\t", err.Error())
	}

	return nil
}
