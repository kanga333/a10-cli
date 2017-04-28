package client

import (
	"encoding/json"
	"log"
)

const (
	auth  = "authenticate"
	close = "close"
)

type authInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type authOutput struct {
	SessionID string `json:"session_id"`
}

// Auth is a function to authenticate to a10
func (c *Client) Auth() error {
	log.Println("Start authentication.")

	if c.token != "" {
		log.Println("Currently authentication has already been completed.")
		log.Println("Close the session and reauthenticate.")
		err := c.Close()
		if err != nil {
			log.Println("Closing the session failed but processing continues.", err)
		}
	}

	url := c.CreateURL(auth)

	in := authInput{
		UserName: c.username,
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

	url, err := c.CreateSessionURL(close)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

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
