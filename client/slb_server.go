package client

import (
	"encoding/json"
	"fmt"
)

const (
	search = "slb.server.search"
	create = "slb.server.create"
	delete = "slb.server.delete"
	update = "slb.server.update"
)

//Port represents slb.server.port object of A10.
type Port struct {
	PortNum      int     `json:"port_num"`
	Protocol     int     `json:"protocol"`
	Status       NumBool `json:"status"`
	Weight       int     `json:"weight"`
	NoSsl        NumBool `json:"no_ssl"`
	ConnLimit    int     `json:"conn_limit"`
	ConnLimitLog NumBool `json:"conn_limit_log"`
	ConnResume   int     `json:"conn_resume"`
	Template     string  `json:"template"`
	StatsData    NumBool `json:"stats_data"`
	//This object is defined as Union - 1966932898,
	//and there is a possibility that follow_port may be entered instead
	HealthMonitor string  `json:"health_monitor"`
	ExtendedStats NumBool `json:"extended_stats"`
}

//Server represents slb.server object of A10.
type Server struct {
	Name                string  `json:"name"`
	Host                string  `json:"host"`
	GslbExternalAddress string  `json:"gslb_external_address"`
	Weight              int     `json:"weight"`
	HealthMonitor       string  `json:"health_monitor"`
	Status              NumBool `json:"status"`
	ConnLimit           int     `json:"conn_limit"`
	ConnLimitLog        NumBool `json:"conn_limit_log"`
	ConnResume          int     `json:"conn_resume"`
	StatsData           NumBool `json:"stats_data"`
	ExtendedStats       NumBool `json:"extended_stats"`
	SlowStart           NumBool `json:"slow_start"`
	SpoofingCache       NumBool `json:"spoofing_cache"`
	Template            string  `json:"template"`
	PortList            []Port  `json:"port_list"`
}

// ServerSearch is a function to slb.server.search to a10.
func (c *Client) ServerSearch(h string) (*Server, error) {
	c.logger.Printf("[INFO] start searching server: %s", h)
	url, err := c.CreateSessionURL(search)
	if err != nil {
		return nil, err
	}

	var input struct {
		Host string `json:"host"`
	}
	input.Host = h

	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jsonBody struct {
		Server Server `json:"server"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}
	if &jsonBody == nil {
		return nil, fmt.Errorf("struct after JSON parsing is empty")
	}

	return &jsonBody.Server, nil
}

// ServerSearchByName is a function to slb.server.search to a10.
func (c *Client) ServerSearchByName(n string) (*Server, error) {
	c.logger.Printf("[INFO] start searching server by name: %s", n)
	url, err := c.CreateSessionURL(search)
	if err != nil {
		return nil, err
	}

	var input struct {
		Name string `json:"name"`
	}
	input.Name = n

	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jsonBody struct {
		Server Server `json:"server"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}
	if &jsonBody == nil {
		return nil, fmt.Errorf("struct after JSON parsing is empty")
	}

	return &jsonBody.Server, nil
}

// ServerCreate is a function to slb.server.create to a10
func (c *Client) ServerCreate(s *Server) error {
	c.logger.Printf("[INFO] start creating name: %s , host: %s", s.Name, s.Host)
	url, err := c.CreateSessionURL(create)
	if err != nil {
		return err
	}

	body, err := json.Marshal(s)
	if err != nil {
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ServerDelete is a function to slb.server.delete to a10
func (c *Client) ServerDelete(h string) error {
	c.logger.Printf("[INFO] start deleting server: %s", h)
	url, err := c.CreateSessionURL(delete)
	if err != nil {
		return err
	}

	var input struct {
		Host string `json:"host"`
	}
	input.Host = h

	body, err := json.Marshal(&input)
	if err != nil {
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ServerUpdate is a function to slb.server.update to a10
func (c *Client) ServerUpdate(s *Server) error {
	c.logger.Printf("[INFO] start updating server: %s , host: %s", s.Name, s.Host)
	url, err := c.CreateSessionURL(update)
	if err != nil {
		return err
	}

	body, err := json.Marshal(s)
	if err != nil {
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
