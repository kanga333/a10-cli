package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/kanga333/a10-cli/client"
)

type config struct {
	A10  client.Opts `yaml:"a10cli"`
	Host string      `yaml:"host"`
}

func loadConf(path string) (*config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
