package client

import (
	"net/http"
	"net/url"
)

const (
	scheme  = "https://"
	baseURI = "/services/rest/V2.1/"
	format  = "json"
)

// Client is a10-cli　client
type Client struct {
	baseURL    *url.URL
	user       string
	password   string
	httpClient *http.Client
}

// Opts is an option used to generate a10.client.Client
type Opts struct {
	user     string
	password string
	target   string
}

// NewClient returns new a10.client.Client
func NewClient(opts Opts) (*Client, error) {
	baseURL := scheme + opts.target + baseURI

	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &Client{url, opts.user, opts.password, client}, nil
}
