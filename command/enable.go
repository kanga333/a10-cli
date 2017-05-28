package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/config"
)

func CmdEnable(c *cli.Context) {
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

	host := conf.GetServerHost()

	server, err := a10.ServerSearch(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to server search: %s\n", err)
		os.Exit(1)
	}
	if server.Host == "" {
		fmt.Fprintf(os.Stderr, "server %s is not exist\n", host)
		os.Exit(0)
	}
	if server.Status == true {
		fmt.Fprintf(os.Stderr, "server %s status is already true\n", host)
		os.Exit(0)
	}
	server.Status = true
	err = a10.ServerUpdate(server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to server update: %s\n", err)
		os.Exit(1)
	}
	serverUpdated, err := a10.ServerSearch(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to server search after update: %s\n", err)
		os.Exit(1)
	}
	if serverUpdated.Status != true {
		fmt.Fprintf(os.Stderr, "server %s status update faild\n", host)
		os.Exit(1)
	}

}
