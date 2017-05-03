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

PortStatus
{{range .PortList}}
PortNum:	{{.PortNum}}({{.Status | boolToState}})
{{end}}
`

func CmdStatus(c *cli.Context) {
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
	printStatus(s)
}

func printStatus(s *client.Server) {
	tmpl, err := template.New("status").
		Funcs(template.FuncMap{"boolToState": boolToState}).
		Parse(templ)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
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
