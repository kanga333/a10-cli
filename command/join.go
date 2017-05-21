package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"
)

func CmdJoin(c *cli.Context) {
	conf, err := config.LoadConf(c.GlobalString("config"))
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
	a10, err := client.NewClient(conf.A10)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
	err = a10.Auth()
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
	defer a10.Close()

	s, err := a10.ServerSearch(conf.Server.Host)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
	if s.Host != "" {
		fmt.Printf("Server %v is already exist.", conf.Server.Host)
	} else {
		fmt.Printf("Create %v.", conf.Server.Host)
		err = a10.ServerCreate(&conf.Server)
		if err != nil {
			fmt.Printf("Unexpected error: %v", err)
			os.Exit(1)
		}
	}

	for _, v := range conf.ServiceGroups {
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
		if m != nil {
			fmt.Printf("Server %v is already exist.", v.Member.Server)
		} else {
			err = a10.ServiceGroupMemberCreate(&v)
			if err != nil {
				fmt.Printf("Unexpected error: %v", err)
			}
		}
	}
}
