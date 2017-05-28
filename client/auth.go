package client

import "encoding/json"

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

// Auth is a function to authenticate to a10.
func (c *Client) Auth() error {
	c.logger.Printf("[INFO] start authenticate by user: %s", c.username)
	// If a value is set in token, it is already authenticated, so close the session first.
	if c.token != "" {
		c.logger.Println("[INFO] reauthentication as authenticated")
		c.Close()
	}

	url := c.CreateURL(auth)

	in := authInput{
		UserName: c.username,
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
	c.logger.Printf("[INFO] close the session with session_id : %s", c.token)
	if c.token == "" {
		c.logger.Println("[INFO] session already closed")
		return nil
	}

	url, err := c.CreateSessionURL(close)
	if err != nil {
		return err
	}

	resp, err := c.postJSON(url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.token = ""
	return nil
}
