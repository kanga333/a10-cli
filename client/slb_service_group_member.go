package client

import (
	"encoding/json"
	"log"
)

const (
	memberCreate = "slb.service_group.member.create"
	memberDelete = "slb.service_group.member.delete"
)

//SGNameAndMember represents slb.service_group.member io object of A10.
type SGNameAndMember struct {
	Name   string `json:"name"`
	Member Member `json:"member"`
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
