package client

import (
	"encoding/json"
	"fmt"
	"log"
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

// ServerSearch is a function to slb.server.search to a10
func (c *Client) ServerSearch(h string) (*Server, error) {
	log.Println("Start server search.")

	url, err := c.CreateSessionURL(search)
	if err != nil {
		log.Println("Error in creating session url.")
		return nil, err
	}

	var input struct {
		Host string `json:"host"`
	}
	input.Host = h

	body, err := json.Marshal(input)
	if err != nil {
		log.Println("Error in creating serverSearch request.")
		return nil, err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in serverSearch request.")
		return nil, err
	}
	defer resp.Body.Close()
	var jsonBody struct {
		Server Server `json:"server"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		log.Println("Error in parsing serverSearch request response.")
		return nil, err
	}
	if &jsonBody == nil {
		return nil, fmt.Errorf("Struct after JSON parsing is empty")
	}

	return &jsonBody.Server, nil
}

// ServerCreate is a function to slb.server.create to a10
func (c *Client) ServerCreate(s *Server) error {
	log.Println("Start server create.")

	url, err := c.CreateSessionURL(create)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

	body, err := json.Marshal(s)
	if err != nil {
		log.Println("Error in creating server create request.")
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in server create request.")
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ServerDelete is a function to slb.server.delete to a10
func (c *Client) ServerDelete(h string) error {
	log.Println("Start server delete.")

	url, err := c.CreateSessionURL(delete)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

	var input struct {
		Host string `json:"host"`
	}
	input.Host = h

	body, err := json.Marshal(&input)
	if err != nil {
		log.Println("Error in creating server create request.")
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in server delete request.")
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ServerUpdate is a function to slb.server.update to a10
func (c *Client) ServerUpdate(s *Server) error {
	log.Println("Start server update.")

	url, err := c.CreateSessionURL(update)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

	body, err := json.Marshal(s)
	if err != nil {
		log.Println("Error in creating server update request.")
		return err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in server update request.")
		return err
	}
	defer resp.Body.Close()

	return nil
}
