package command

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/config"
)

func CmdLeave(c *cli.Context) {
	conf, err := config.LoadConf(c.GlobalString("config"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read configuration file: %s", err)
		os.Exit(1)
	}

	a10, err := newAuthorizedClientwithFromConfig(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create authorized client: %s", err)
		os.Exit(1)
	}
	defer a10.Close()

	sgs, err := conf.GenerateSGNameAndMembers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create service groups from config: %s", err)
		os.Exit(1)
	}

	for _, v := range sgs {
		sg, err := a10.ServiceGroupSearch(v.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to search service group: %s", err)
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
				fmt.Fprintf(os.Stderr, "failed to delete service group member: %s", err)
			}

		} else {
			log.Printf("[INFO] server: %s does not already exist in s: %s", v.Member.Server, v.Name)
		}
	}

	server, err := conf.GenerateServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create server from config: %s", err)
		os.Exit(1)
	}

	s, err := a10.ServerSearch(server.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to search the server: %s", err)
		os.Exit(1)
	}
	if s != nil {
		fmt.Printf("Create %v.", server.Host)
		err = a10.ServerDelete(server.Host)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to delete server: %s", err)
			os.Exit(1)
		}
	} else {
		log.Printf("[INFO] server: %s does not already exist", server.Host)
	}

}
