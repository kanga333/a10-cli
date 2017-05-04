package client

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	memberCreate = "slb.service_group.member.create"
	memberDelete = "slb.service_group.member.delete"
	sgSearch     = "slb.service_group.search"
)

//SGNameAndMember represents slb.service_group.member io object of A10.
type SGNameAndMember struct {
	Name   string `json:"name"`
	Member Member `json:"member"`
}

//ServiceGroup represents slb.service_group object of A10.
type ServiceGroup struct {
	Name                       string          `json:"name"`
	Protocol                   int             `json:"protocol"`
	LbMethod                   int             `json:"lb_method"`
	HealthMonitor              string          `json:"health_monitor"`
	MinActiveMember            MinActiveMember `json:"min_active_member"`
	BackupServerEventLogEnable int             `json:"backup_server_event_log_enable"`
	ClientReset                int             `json:"client_reset"`
	StatsData                  int             `json:"stats_data"`
	ExtendedStats              int             `json:"extended_stats"`
	MemberList                 []Member        `json:"member_list"`
}

//MinActiveMember represents slb.service_group.min_active_member object of A10.
type MinActiveMember struct {
	Status      int `json:"status"`
	Number      int `json:"number"`
	PrioritySet int `json:"priority_set"`
}

//Member represents slb.service_group.member member object of A10.
type Member struct {
	Server    string   `json:"server"`
	Port      int      `json:"port"`
	Template  *string  `json:"template,omitempty"`
	Priority  *int     `json:"priority,omitempty"`
	Status    *NumBool `json:"status,omitempty"`
	StatsData *NumBool `json:"stats_data,omitempty"`
}

//ServiceGroupMemberCreate is a function to slb.service_group.member.create to a10.
func (c *Client) ServiceGroupMemberCreate(m *SGNameAndMember) error {
	log.Println("Start member create.")

	url, err := c.CreateSessionURL(memberCreate)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

	body, err := json.Marshal(m)
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

//ServiceGroupMemberDelete is a function to slb.service_group.member.delete to a10.
func (c *Client) ServiceGroupMemberDelete(m *SGNameAndMember) error {
	log.Println("Start member delete.")

	url, err := c.CreateSessionURL(memberDelete)
	if err != nil {
		log.Println("Error in creating session url.")
		return err
	}

	body, err := json.Marshal(m)
	if err != nil {
		log.Println("Error in creating server delete request.")
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

//ServiceGroupSearch is a function to slb.service_group.search to a10.
func (c *Client) ServiceGroupSearch(n string) (*ServiceGroup, error) {
	log.Println("Start sg search.")

	url, err := c.CreateSessionURL(sgSearch)
	if err != nil {
		log.Println("Error in creating session url.")
		return nil, err
	}

	var name struct {
		Neme string `json:"name"`
	}
	name.Neme = n

	body, err := json.Marshal(name)
	if err != nil {
		log.Println("Error in creating server delete request.")
		return nil, err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		log.Println("Error in server create request.")
		return nil, err
	}
	defer resp.Body.Close()
	var jsonBody struct {
		ServiceGroup ServiceGroup `json:"service_group"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		log.Println("Error in parsing serverSearch request response.")
		return nil, err
	}
	if &jsonBody == nil {
		return nil, fmt.Errorf("Struct after JSON parsing is empty")
	}

	return &jsonBody.ServiceGroup, nil
}

//SGMemberSearch is a function to search for the specified member in ServiceGroup.
func (c *Client) SGMemberSearch(sg *ServiceGroup, server string) *Member {
	log.Println("Start SGmember search.")

	for _, m := range sg.MemberList {
		if m.Server == server {
			return &m
		}
	}
	return nil
}
