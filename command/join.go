package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"
)

func CmdJoin(c *cli.Context) {
	conf, err := config.LoadConf(c.GlobalString("config"))
	if err != nil {
		log.Printf("[ERR] failed to read configuration file: %s", err)
		os.Exit(1)
	}
	a10, err := client.NewClient(conf.A10)
	if err != nil {
		log.Printf("[ERR] failed to create client: %s", err)
		os.Exit(1)
	}
	err = a10.Auth()
	if err != nil {
		log.Printf("[ERR] failed on authentication: %s", err)
		os.Exit(1)
	}
	defer a10.Close()

	s, err := a10.ServerSearch(conf.Server.Host)
	if err != nil {
		log.Printf("[ERR] failed to search the server: %s", err)
		os.Exit(1)
	}
	if s.Host != "" {
		log.Printf("[INFO] server: %s is already exist", s.Host)
	} else {
		err = a10.ServerCreate(&conf.Server)
		if err != nil {
			log.Printf("[ERR] failed to create server: %s", err)
			os.Exit(1)
		}
	}

	for _, v := range conf.ServiceGroups {
		sg, err := a10.ServiceGroupSearch(v.Name)
		if err != nil {
			log.Printf("[ERR] failed to search service group: %s", err)
			os.Exit(1)
		}
		if sg == nil {
			log.Printf("[ERR] service group: %s is not exist", v.Name)
			os.Exit(1)
		}
		m := a10.SGMemberSearch(sg, v.Member.Server)
		if m != nil {
			log.Printf("[INFO] server: %s is already exist in sg: %s", s.Host, sg.Name)
		} else {
			err = a10.ServiceGroupMemberCreate(&v)
			if err != nil {
				log.Printf("[ERR] failed to service group create: %s", err)
			}
		}
	}
}
