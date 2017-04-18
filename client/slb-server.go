package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

const (
	search = "slb.server.search"
)

//Port represents slb.server.port object of A10.
type Port struct {
	PortNum      int     `json:"port_num"`
	Protocol     int     `json:"protocol"`
	Status       a10Bool `json:"user"`
	Weight       int     `json:"weight"`
	NoSsl        a10Bool `json:"no_ssl"`
	ConnLimit    int     `json:"conn_limit"`
	ConnLimitLog a10Bool `json:"conn_limit_log"`
	ConnResume   int     `json:"conn_resume"`
	Template     string  `json:"template"`
	StatsSata    a10Bool `json:"stats_data"`
	//This object is defined as Union - 1966932898,
	//and there is a possibility that follow_port may be entered instead
	HealthMonitor string  `json:"health_monitor"`
	ExtendedStats a10Bool `json:"extended_stats"`
}

//Server represents slb.server object of A10.
type Server struct {
	Name                string  `json:"name"`
	Host                string  `json:"host"`
	GslbExternalAddress string  `json:"gslb_external_address"`
	Weight              int     `json:"weight"`
	HealthMonitor       string  `json:"health_monitor"`
	Status              a10Bool `json:"status"`
	ConnLimit           int     `json:"conn_limit"`
	ConnLimitLog        a10Bool `json:"conn_limit_log"`
	ConnResume          int     `json:"conn_resume"`
	StatsData           a10Bool `json:"stats_data"`
	ExtendedStats       a10Bool `json:"extended_stats"`
	SlowStart           a10Bool `json:"slow_start"`
	SpoofingCache       a10Bool `json:"spoofing_cache"`
	Template            string  `json:"template"`
	PortList            []Port  `json:"port_list"`
}

//Host represents slb.server.host object of A10.
//It is used by input of ServerSeach ().
//Normally it is a unique value that contains the ip address of host.
type Host struct {
	Host string `json:"host"`
}

// ServerSearch is a function to slb.server.search to a10
func (c *Client) ServerSearch(h Host) (*Server, error) {
	log.Println("Start server search.")
	if c.token == "" {
		return nil, fmt.Errorf("Session is not authenticated")
	}

	parm := make(url.Values)
	parm.Add("method", search)
	parm.Add("format", format)
	parm.Add("session_id", c.token)

	url := c.baseURL.String() + "?" + parm.Encode()

	body, err := json.Marshal(h)
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
