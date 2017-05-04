package config

import (
	"io/ioutil"

	"github.com/ghodss/yaml"

	"github.com/kanga333/a10-cli/client"
)

// Config is a structure that expresses the setting required by a10
type Config struct {
	A10    client.Opts   `json:"a10cli"`
	Server client.Server `json:"server"`
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
