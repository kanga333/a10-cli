package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
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
	Username string `json:"username"`
	Password string `json:"password"`
	Target   string `json:"target"`
	Insecure bool   `json:"insecure"`
	Proxy    string `json:"proxy"`
}

// NewClient returns new a10.client.Client
func NewClient(opts Opts) (*Client, error) {
	log.Printf("[INFO] setting client to target %s", opts.Target)
	baseURL := scheme + opts.Target + baseURI

	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to parse base url: %s", baseURL)
	}

	client := &http.Client{}

	transport := &http.Transport{}
	tlsConfig := &tls.Config{}
	if opts.Insecure == true {
		log.Printf("[INFO] setting client insecure true")
		tlsConfig.InsecureSkipVerify = true
	}
	transport.TLSClientConfig = tlsConfig
	if opts.Proxy != "" {
		proxy, err := url.Parse(opts.Proxy)
		if err != nil {
			return nil, errors.Wrapf(err, "fail to parse proxy url: %s", opts.Proxy)
		}
		log.Printf("[INFO] setting client proxy %s", proxy)
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

//CreateSessionURL creates the URL of A10 API with session_id.
func (c *Client) CreateSessionURL(method string) (string, error) {
	if c.token == "" {
		return "", fmt.Errorf("Session is not authenticated")
	}

	v := make(url.Values)
	v.Add("method", method)
	v.Add("format", format)
	v.Add("session_id", c.token)
	u := c.baseURL.String() + "?" + v.Encode()

	return u, nil
}

//CreateURL creates the URL of A10 API.
func (c *Client) CreateURL(method string) string {
	v := make(url.Values)
	v.Add("method", method)
	v.Add("format", format)
	u := c.baseURL.String() + "?" + v.Encode()

	return u
}
