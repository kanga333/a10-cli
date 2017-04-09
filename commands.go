package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_USER",
		Name:   "user",
		Value:  "",
		Usage:  "",
	},
	cli.StringFlag{
		EnvVar: "ENV_PASSWORD",
		Name:   "password",
		Value:  "",
		Usage:  "",
	},
	cli.StringFlag{
		EnvVar: "ENV_TARGET",
		Name:   "target",
		Value:  "",
		Usage:  "",
	},
	cli.StringFlag{
		EnvVar: "ENV_CONFIG",
		Name:   "config",
		Value:  "",
		Usage:  "",
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
		Flags:  []cli.Flag{},
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
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
