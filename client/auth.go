package client

import (
	"encoding/json"
	"log"
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
	log.Println("Start authentication.")
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
		log.Println("Error in creating authentication request.")
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in aythentication request.")
		return err
	}
	defer resp.Body.Close()

	var out authOutput
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		log.Println("Error in parsing authentication request response.")
		return err
	}

	c.token = out.SessionID
	log.Println("Authentication is complete.")
	return nil
}

// Close is a function to session.close to a10
func (c *Client) Close() error {
	log.Println("Start closing session.")
	if c.token == "" {
		log.Println("Session already closed.")
		return nil
	}

	parm := make(url.Values)
	parm.Add("method", close)
	parm.Add("format", format)
	parm.Add("session_id", c.token)

	url := c.baseURL.String() + "?" + parm.Encode()

	resp, err := c.postJSON(url, nil)
	if err != nil {
		log.Println("Error in session close request.")
		return err
	}
	defer resp.Body.Close()

	c.token = ""
	log.Println("Closing the session is complete.")
	return nil
}
