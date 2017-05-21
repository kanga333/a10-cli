package command

import (
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"

	"github.com/codegangsta/cli"
)

func CmdExport(c *cli.Context) {
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
	b, err := yaml.Marshal(s)
	println(string(b))

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
		by, err := yaml.Marshal(m)
		println(string(by))
	}
}
