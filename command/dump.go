package command

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kanga333/a10-cli/config"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/client"
)

func CmdDump(c *cli.Context) {
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

	var dumpServerName string
	var dumpSGNames []string
	if c.String("server") == "" {
		if c.StringSlice("service-group") != nil {
			fmt.Fprintln(os.Stderr, "service-group must be used with server option")
			os.Exit(1)
		}
		dumpServerName = conf.GetServerName()
		dumpSGNames = conf.GetServiceGroupName()
	} else {
		dumpServerName = c.String("server")
		dumpSGNames = c.StringSlice("service-group")
	}

	server, err := a10.ServerSearchByName(dumpServerName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to server search: %s", err)
		os.Exit(1)
	}

	var SGMembers []client.SGNameAndMember
	for _, sgName := range dumpSGNames {
		sg, err := a10.ServiceGroupSearch(sgName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to service group search: %s", err)
			os.Exit(1)
		}
		if sg == nil {
			fmt.Fprintf(os.Stderr, "service group :%s is not found", sgName)
			os.Exit(1)
		}
		m := a10.SGMemberSearch(sg, dumpServerName)
		if m == nil {
			fmt.Fprintf(os.Stderr, "server %s is not member inservice group :%s", dumpServerName, sgName)
			os.Exit(1)
		}
		var sgm = client.SGNameAndMember{
			Name:   sgName,
			Member: *m,
		}
		SGMembers = append(SGMembers, sgm)
	}

	byteServer, err := json.Marshal(server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to server marshal :%s", err)
		os.Exit(1)
	}
	fmt.Println(string(byteServer))
	for _, sgm := range SGMembers {
		byteSGM, err := json.Marshal(sgm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to server marshal :%s", err)
			os.Exit(1)
		}
		fmt.Println(string(byteSGM))
	}
}
