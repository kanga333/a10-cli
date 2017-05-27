package command

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"
)

func CmdLeave(c *cli.Context) {
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

	for _, v := range conf.ServiceGroups {
		sg, err := a10.ServiceGroupSearch(v.Name)
		if err != nil {
			log.Printf("[ERR] failed to search service group: %s", err)
			os.Exit(1)
		}
		if sg == nil {
			log.Printf("[INFO] service group: %v is not found", v.Name)
			break
		}
		m := a10.SGMemberSearch(sg, v.Member.Server)
		if m != nil {
			err = a10.ServiceGroupMemberDelete(&v)
			if err != nil {
				log.Printf("[ERR] failed to delete service group member: %s", err)
			}

		} else {
			log.Printf("[INFO] server: %s does not already exist in s: %s", v.Member.Server, v.Name)
		}
	}

	s, err := a10.ServerSearch(conf.Server.Host)
	if err != nil {
		log.Printf("[ERR] failed to search the server: %s", err)
		os.Exit(1)
	}
	if s != nil {
		fmt.Printf("Create %v.", conf.Server.Host)
		err = a10.ServerDelete(conf.Server.Host)
		if err != nil {
			log.Printf("[ERR] failed to delete server: %s", err)
			os.Exit(1)
		}
	} else {
		log.Printf("[INFO] server: %s does not already exist", conf.Server.Host)
	}

}
