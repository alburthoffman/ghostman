package config

import (
	"github.com/newly12/ghostman/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var log = logging.GetLogger()

type Runtime struct {
	TmplDir   string `yaml:"templateDir,omitempty"`
	MaxConcur int    `yaml:"maxConcurrent,omitempty"`
}

type Defaults struct {
	Rcconf string
	Runtime
}

func GetDefaultsFromConfig(f string) *Defaults {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Errorf("failed to read config: %s", err)
	}

	defaults := Defaults{}
	err = yaml.Unmarshal(yamlFile, &defaults)
	if err != nil {
		log.Errorf("failed to parse config: %s", err)
	}

	return &defaults
}
