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
	username string
	password string

	httpClient *http.Client

	token string
}

// Opts is an option used to generate a10.client.Client
type Opts struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Target   string `yaml:"target"`
	Insecure bool   `yaml:"insecure"`
	Proxy    string `yaml:"proxy"`
}

// NewClient returns new a10.client.Client
func NewClient(opts Opts) (*Client, error) {
	baseURL := scheme + opts.Target + baseURI

	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	transport := &http.Transport{}
	tlsConfig := &tls.Config{}
	if opts.Insecure == true {
		tlsConfig.InsecureSkipVerify = true
	}
	transport.TLSClientConfig = tlsConfig
	if opts.Proxy != "" {
		proxy, err := url.Parse(opts.Proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxy)
	}
	client.Transport = transport

	return &Client{
		baseURL:    url,
		username:   opts.Username,
		password:   opts.Password,
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

//AuthenticatedRequest is a function of the request to the authenticated A10 API.
func (c *Client) AuthenticatedRequest(method string, body []byte) (*http.Response, error) {
	if c.token == "" {
		return nil, fmt.Errorf("Session is not authenticated")
	}

	parm := make(url.Values)
	parm.Add("method", method)
	parm.Add("format", format)
	parm.Add("session_id", c.token)
	url := c.baseURL.String() + "?" + parm.Encode()

	return c.postJSON(url, body)
}

//Request is a function of the request to the A10 API.
func (c *Client) Request(method string, body []byte) (*http.Response, error) {

	parm := make(url.Values)
	parm.Add("method", method)
	parm.Add("format", format)
	url := c.baseURL.String() + "?" + parm.Encode()

	return c.postJSON(url, body)
}
