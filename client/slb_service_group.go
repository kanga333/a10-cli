package client

import (
	"encoding/json"
	"fmt"
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
	Server    string  `json:"server"`
	Port      int     `json:"port"`
	Template  string  `json:"template"`
	Priority  int     `json:"priority"`
	Status    NumBool `json:"status"`
	StatsData NumBool `json:"stats_data"`
}

//ServiceGroupMemberCreate is a function to slb.service_group.member.create to a10.
func (c *Client) ServiceGroupMemberCreate(m *SGNameAndMember) error {
	c.logger.Printf("[INFO] start creating member: %s to sg: %s", m.Member.Server, m.Name)

	url, err := c.CreateSessionURL(memberCreate)
	if err != nil {
		return err
	}

	body, err := json.Marshal(m)
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

//ServiceGroupMemberDelete is a function to slb.service_group.member.delete to a10.
func (c *Client) ServiceGroupMemberDelete(m *SGNameAndMember) error {
	c.logger.Printf("[INFO] start deleting member: %s in sg: %s", m.Member.Server, m.Name)

	url, err := c.CreateSessionURL(memberDelete)
	if err != nil {
		return err
	}

	body, err := json.Marshal(m)
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

//ServiceGroupSearch is a function to slb.service_group.search to a10.
func (c *Client) ServiceGroupSearch(n string) (*ServiceGroup, error) {
	c.logger.Printf("[INFO] start serching sg: %s", n)

	url, err := c.CreateSessionURL(sgSearch)
	if err != nil {
		return nil, err
	}

	var name struct {
		Neme string `json:"name"`
	}
	name.Neme = n

	body, err := json.Marshal(name)
	if err != nil {
		return nil, err
	}

	resp, err := c.postJSON(url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsonBody struct {
		ServiceGroup ServiceGroup `json:"service_group"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}
	if &jsonBody == nil {
		return nil, fmt.Errorf("struct after JSON parsing is empty")
	}

	return &jsonBody.ServiceGroup, nil
}

//SGMemberSearch is a function to search for the specified member in ServiceGroup.
func (c *Client) SGMemberSearch(sg *ServiceGroup, server string) *Member {
	c.logger.Printf("[INFO] start serching member: %s in sg: %s", server, sg.Name)

	for _, m := range sg.MemberList {
		if m.Server == server {
			return &m
		}
	}
	return nil
}
