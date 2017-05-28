package command

import (
	"html/template"
	"os"

	"fmt"

	"github.com/codegangsta/cli"
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"
)

const templ = `ServerStatus: ({{.Status | boolToState}})
Name:	{{.Name}}
Host:	{{.Host}}
Weight:	{{.Weight}}

PortStatus:
{{range .PortList}}
PortNum:	{{.PortNum}}({{.Status | boolToState}})
{{end}}
`

func CmdStatus(c *cli.Context) {
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

	s, err := a10.ServerSearch(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to search the server: %s\n", err)
		os.Exit(1)
	}
	if s.Host == "" {
		fmt.Fprintf(os.Stderr, "server %s is not exist\n", host)
		os.Exit(0)
	}

	printStatus(s)
}

func printStatus(s *client.Server) {
	tmpl, err := template.New("status").
		Funcs(template.FuncMap{"boolToState": boolToState}).
		Parse(templ)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create template: %s\n", err)
		os.Exit(1)
	}
	tmpl.Execute(os.Stdout, s)
}

func boolToState(b client.NumBool) string {
	if b {
		return "active"
	}
	return "inactive"
}
