package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/config"
)

func CmdLeave(c *cli.Context) {
	conf, err := config.LoadConf(c.GlobalString("config"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read configuration file: %s\n", err)
		os.Exit(1)
	}

	a10, err := newAuthorizedClientwithFromConfig(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create authorized client: %s\n", err)
		os.Exit(1)
	}
	defer a10.Close()

	sgs, err := conf.GenerateSGNameAndMembers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create service groups from config: %s\n", err)
		os.Exit(1)
	}

	for _, v := range sgs {
		sg, err := a10.ServiceGroupSearch(v.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to search service group: %s\n", err)
			os.Exit(1)
		}
		if sg == nil {
			break
		}
		m := a10.SGMemberSearch(sg, v.Member.Server)
		if m != nil {
			err = a10.ServiceGroupMemberDelete(&v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to delete service group member: %s\n", err)
			}

		}
	}

	server, err := conf.GenerateServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create server from config: %s\n", err)
		os.Exit(1)
	}

	s, err := a10.ServerSearch(server.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to search the server: %s\n", err)
		os.Exit(1)
	}
	if s.Host != "" {
		err = a10.ServerDelete(server.Host)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to delete server: %s\n", err)
			os.Exit(1)
		}
	}

}
