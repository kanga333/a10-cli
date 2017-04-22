package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/kanga333/a10-cli/client"
)

// Config is a structure that expresses the setting required by a10
type Config struct {
	A10  client.Opts `yaml:"a10cli"`
	Host string      `yaml:"host"`
}

// LoadConf reads the yaml setting from the specified path
func LoadConf(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
