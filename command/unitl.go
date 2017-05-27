package command

import (
	"github.com/kanga333/a10-cli/client"
	"github.com/kanga333/a10-cli/config"
)

func newAuthorizedClientwithFromConfig(c *config.Config) (*client.Client, error) {
	ops, err := c.GetCliOpts()
	if err != nil {
		return nil, err
	}
	a10, err := client.NewClient(ops)
	if err != nil {
		return nil, err
	}
	err = a10.Auth()
	if err != nil {
		return nil, err
	}
}
