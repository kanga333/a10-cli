# a10-cli

## Description
A10cli is a cli tool to the slb of acos 2.7.

## Usage

```console

a10-cli -h
NAME:
   a10-cli

USAGE:
   a10-cli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   kanga333

COMMANDS:
     join     Create a server in a10-slb and register the port in the specified service group
     leave    Delete the port registration from the service group and delete the server information from a10-slb
     status   Print the status of the server registered in a10-slb
     update   Update server setting to specified flag value
     disable  Disable load balancing status to server
     enable   Enable load balancing status to server
     dump     Dump the status of the specified server to JSON
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username value, -u value  Authentication user [$A10_USER]
   --password value, -p value  Authentication password [$A10_PASSWORD]
   --target value, -t value    Slb hostname or ip [$A10_TARGET]
   --config value, -c value    Location of config file [$A10_CONFIG]
   --help, -h                  show help
   --version, -v               print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/kanga333/a10-cli
```

## Contribution

1. Fork ([https://github.com/kanga333/a10-cli/fork](https://github.com/kanga333/a10-cli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[kanga333](https://github.com/kanga333)
