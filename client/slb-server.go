package client

import (
	"fmt"
)

const (
	search = "slb.server.search"
)

//Port represents slb.server.port object of A10.
type Port struct {
	Status       bool   `json:"user"`
	Weight       int    `json:"password"`
	NoSsl        bool   `json:"no_ssl"`
	ConnLimit    int    `json:"conn_limit"`
	ConnLimitLog bool   `json:"conn_limit_log"`
	ConnResume   int    `json:"conn_resume"`
	Template     string `json:"template"`
	StatsSata    bool   `json:"stats_data"`
	//This object is defined as Union - 1966932898,
	//and there is a possibility that follow_port may be entered instead
	HealthMonitor string `json:"health_monitor"`
	ExtendedStats bool   `json:"extended_stats"`
}

//Server represents slb.server object of A10.
type Server struct {
	Name                string `json:"name"`
	Host                string `json:"host"`
	GslbExternalAddress string `json:"gslb_external_address"`
	Weight              int    `json:"weight"`
	HealthMonitor       string `json:"health_monitor"`
	Status              bool   `json:"status"`
	ConnLimit           bool   `json:"conn_limit"`
	ConnLimitLog        bool   `json:"conn_limit_log"`
	ConnResume          bool   `json:"conn_resume"`
	StatsData           bool   `json:"stats_data"`
	ExtendedStats       string `json:"extended_stats"`
	SlowStart           bool   `json:"slow_start"`
	SpoofingCache       bool   `json:"spoofing_cache"`
	Template            string `json:"template"`
	PortList            []Port `json:"port_list"`
}

//Host represents slb.server.host object of A10.
//It is used by input of ServerSeach ().
//Normally it is a unique value that contains the ip address of host.
type Host struct {
	Host string `json:"host"`
}

// ServerSearch is a function to slb.server.search to a10
func (c *Client) ServerSearch(h Host) (*Server, error) {
	return nil, fmt.Errorf("Unimplemented")
}
