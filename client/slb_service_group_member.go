package client

import (
	"encoding/json"
	"log"
)

const (
	memberCreate = "slb.service_group.member.create"
)

//SGNameAndMember represents slb.service_group.member io object of A10.
type SGNameAndMember struct {
	Name   string `json:"name"`
	Member Member `json:"member"`
}

//Member represents slb.service_group.member member object of A10.
type Member struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	Template  string `json:"template"`
	Priority  int    `json:"priority"`
	Status    int    `json:"status"`
	StatsData int    `json:"stats_data"`
}

//ServiceGroupMemberCreate s a function to slb.service_group.member.create to a10.
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
