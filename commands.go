package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/command"
)

// GlobalFlags is a flag used for the entire a10-cli.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "A10_USER",
		Name:   "username,u",
		Value:  "",
		Usage:  "Authentication user",
	},
	cli.StringFlag{
		EnvVar: "A10_PASSWORD",
		Name:   "password,p",
		Value:  "",
		Usage:  "Authentication password",
	},
	cli.StringFlag{
		EnvVar: "A10_TARGET",
		Name:   "target,t",
		Value:  "",
		Usage:  "Slb hostname or ip",
	},
	cli.StringFlag{
		EnvVar: "A10_CONFIG",
		Name:   "config,c",
		Value:  "",
		Usage:  "Location of config file",
	},
}

// Commands stores information on subcommands used with a10-cli.
var Commands = []cli.Command{
	{
		Name:   "join",
		Usage:  "Create a server in a10-slb and register the port in the specified service group",
		Action: command.CmdJoin,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "leave",
		Usage:  "Delete the port registration from the service group and delete the server information from a10-slb",
		Action: command.CmdLeave,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "status",
		Usage:  "Print the status of the server registered in a10-slb",
		Action: command.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "",
				Name:   "server",
				Value:  "",
				Usage:  "Server name to print status",
			},
		},
	},
	{
		Name:   "update",
		Usage:  "Update server setting to specified flag value",
		Action: command.CmdUpdate,
		Flags: []cli.Flag{
			cli.IntFlag{
				EnvVar: "",
				Name:   "weight,w",
				Value:  1,
				Usage:  "Server weight",
			},
		},
	},
	{
		Name:   "disable",
		Usage:  "Disable load balancing status to server",
		Action: command.CmdDisable,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "enable",
		Usage:  "Enable load balancing status to server",
		Action: command.CmdEnable,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "dump",
		Usage:  "Dump the status of the specified server to JSON",
		Action: command.CmdDump,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "",
				Name:   "server",
				Value:  "",
				Usage:  "Server name to dump",
			},
			cli.StringSliceFlag{
				EnvVar: "",
				Name:   "service-group",
				Value:  nil,
				Usage:  "Service group name to dump (can be set more than once)",
			},
		},
	},
}

// CommandNotFound displays an error when calling an unregistered subcommand.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
