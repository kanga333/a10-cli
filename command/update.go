package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/config"
)

func CmdUpdate(c *cli.Context) {
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

	weight := c.Int("weight")
	if weight == 0 {
		fmt.Fprintln(os.Stderr, "weight 0 can not be set")
		os.Exit(1)
	}

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
	if server.Weight == weight {
		fmt.Fprintf(os.Stderr, "server %s weight is already %d\n", host, weight)
		os.Exit(0)
	}
	server.Weight = weight
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
	if serverUpdated.Weight != weight {
		fmt.Fprintf(os.Stderr, "server %s status update faild\n", host)
		os.Exit(1)
	}

}
