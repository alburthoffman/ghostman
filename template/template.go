package template

import (
	"github.com/newly12/ghostman/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var log = logging.GetLogger()

type Metadata struct {
	Usage string `yaml:"usage"`
}

type Headers struct {
	CT   string `yaml:"content-type,omitempty"`
	Auth string `yaml:"authorization,omitempty"`
}

type Attributes struct {
	Protc  string `yaml:"protocol"`
	Port   int    `ymal:"port"`
	Method string `yaml:"method,omitempty"`
	Url    string `yaml:"url"`
	Body   string `yaml:"body"`
}

type Defaults struct{}

type Template struct {
	Metadata
	Headers
	Attributes
	Defaults
}

func ParseTemplate(f string) *Template {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Errorf("failed to read config: %s", err)
	}

	t := Template{}
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Errorf("failed to parse config: %s", err)
	}

	return &t
}
