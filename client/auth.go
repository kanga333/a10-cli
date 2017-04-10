package client

import (
	"encoding/json"
	"net/url"
)

const (
	auth  = "authenticate"
	close = "close"
)

type authInput struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type authOutput struct {
	SessionID string `json:"session_id"`
}

// Auth is a function to authenticate to a10
func (c *Client) Auth() error {
	parm := make(url.Values)
	parm.Add("method", auth)
	parm.Add("format", format)

	url := c.baseURL.String() + "?" + parm.Encode()

	in := authInput{
		User:     c.user,
		Password: c.password,
	}

	body, err := json.Marshal(in)
	if err != nil {
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var out authOutput
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return err
	}

	c.token = out.SessionID
	return nil
}

// Close is a function to session.close to a10
func (c *Client) Close() error {
	if c.token == "" {
		return nil
	}

	parm := make(url.Values)
	parm.Add("method", close)
	parm.Add("format", format)
	parm.Add("session_id", c.token)

	url := c.baseURL.String() + "?" + parm.Encode()

	resp, err := c.postJSON(url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.token = ""
	return nil
}