package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
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
	baseURL  *url.URL
	user     string
	password string

	httpClient *http.Client

	token string
}

// Opts is an option used to generate a10.client.Client
type Opts struct {
	user     string
	password string
	target   string
	insecure bool
	proxy    string
}

// NewClient returns new a10.client.Client
func NewClient(opts Opts) (*Client, error) {
	baseURL := scheme + opts.target + baseURI

	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	transport := &http.Transport{}
	tlsConfig := &tls.Config{}
	if opts.insecure == true {
		tlsConfig.InsecureSkipVerify = true
	}
	transport.TLSClientConfig = tlsConfig
	if opts.proxy != "" {
		proxy, err := url.Parse(opts.proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxy)
	}
	client.Transport = transport

	return &Client{
		baseURL:    url,
		user:       opts.user,
		password:   opts.password,
		httpClient: client,
	}, nil
}

func (c *Client) postJSON(path string, body []byte) (*http.Response, error) {
	bodyReader := bytes.NewReader(body)
	log.Printf("Send post request to address %v", path)
	resp, err := c.httpClient.Post(path, "application/json", bodyReader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("http request error, response code:%v", resp.StatusCode)
	}
	return resp, nil
}
