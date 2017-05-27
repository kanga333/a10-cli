package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "A10_USER",
		Name:   "user",
		Value:  "",
		Usage:  "authentication user",
	},
	cli.StringFlag{
		EnvVar: "A10_PASSWORD",
		Name:   "password",
		Value:  "",
		Usage:  "authentication password",
	},
	cli.StringFlag{
		EnvVar: "A10_TARGET",
		Name:   "target",
		Value:  "",
		Usage:  "slb hostname or ip",
	},
	cli.StringFlag{
		EnvVar: "A10_CONFIG",
		Name:   "config",
		Value:  "",
		Usage:  "location of setting file",
	},
}

var Commands = []cli.Command{
	{
		Name:   "join",
		Usage:  "",
		Action: command.CmdJoin,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "leave",
		Usage:  "",
		Action: command.CmdLeave,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "status",
		Usage:  "",
		Action: command.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "",
				Name:   "server",
				Value:  "",
				Usage:  "server name",
			},
			cli.StringSliceFlag{
				EnvVar: "",
				Name:   "service-group",
				Value:  nil,
				Usage:  "service group name that can be specified more than once",
			},
		},
	},
	{
		Name:   "update",
		Usage:  "",
		Action: command.CmdUpdate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "disable",
		Usage:  "",
		Action: command.CmdDisable,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "enable",
		Usage:  "",
		Action: command.CmdEnable,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "dump",
		Usage:  "",
		Action: command.CmdDump,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "",
				Name:   "server",
				Value:  "",
				Usage:  "server name",
			},
			cli.StringSliceFlag{
				EnvVar: "",
				Name:   "service-group",
				Value:  nil,
				Usage:  "service group name that can be specified more than once",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
