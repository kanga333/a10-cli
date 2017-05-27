package command

import (
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/kanga333/a10-cli/config"

	"github.com/codegangsta/cli"
)

func CmdDump(c *cli.Context) {
	conf, err := config.LoadConf(c.GlobalString("config"))
	if err != nil {
		log.Printf("[ERR] failed to read configuration file: %s", err)
		os.Exit(1)
	}

	a10, err := newAuthorizedClientwithFromConfig(conf)
	if err != nil {
		log.Printf("[ERR] failed to create authorized client: %s", err)
		os.Exit(1)
	}
	defer a10.Close()

	server, err := conf.GenerateServer()
	if err != nil {
		log.Printf("[ERR] failed to create server from config: %s", err)
		os.Exit(1)
	}

	s, err := a10.ServerSearch(server.Name)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
	b, err := yaml.Marshal(s)
	println(string(b))

	sgs, err := conf.GenerateSGNameAndMembers()
	if err != nil {
		log.Printf("[ERR] failed to create service groups from config: %s", err)
		os.Exit(1)
	}
	for _, v := range sgs {
		sg, err := a10.ServiceGroupSearch(v.Name)
		if err != nil {
			fmt.Printf("Unexpected error: %v", err)
			os.Exit(1)
		}
		if sg == nil {
			fmt.Printf("ServiceGroup %v is not found.", v.Name)
			os.Exit(1)
		}
		m := a10.SGMemberSearch(sg, v.Member.Server)
		by, err := yaml.Marshal(m)
		println(string(by))
	}
}
