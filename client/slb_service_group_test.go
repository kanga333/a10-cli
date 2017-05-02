package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const testSGMember = `
{
    "name": "T"
    ,"member": {
        "server": "T0684"
        ,"port": 2661
        ,"template": "T068"
        ,"priority": 11
        ,"status": 0
        ,"stats_data": 0
    }
}
`

const testSGMemberSubset = `
{
    "name": "T"
    ,"member": {
        "server": "T0684"
        ,"port": 2661
    }
}
`

const testServiceGroup = `
{
    "service_group": {
        "name": "VBX2842DH"
        ,"protocol": 2
        ,"lb_method": 2
        ,"health_monitor": "L8T0"
        ,"min_active_member": {
            "status": 0
            ,"number": 53
            ,"priority_set": 1
        }
        ,"backup_server_event_log_enable": 0
        ,"client_reset": 0
        ,"stats_data": 0
        ,"extended_stats": 0
        ,"member_list": [
            {
                "server": ""
                ,"port": 51084
                ,"template": ""
                ,"priority": 12
                ,"status": 0
                ,"stats_data": 0
            }
            ,{
                "server": "BT2DN60F"
                ,"port": 34144
                ,"template": "B"
                ,"priority": 9
                ,"status": 0
                ,"stats_data": 0
            }
        ]
    }
}
`

func TestServiceGroupMemberCreate(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.service_group.member.create" {
			t.Error("request QueryString should be method=slb.service_group.member.create but :", query["method"])
		}

		var sgMember SGNameAndMember
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&sgMember)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		var expectSgMember SGNameAndMember
		err = json.Unmarshal([]byte(testSGMember), &expectSgMember)
		if err != nil {
			t.Error("testServerData should be decoded as json")
		}

		if !reflect.DeepEqual(sgMember, expectSgMember) {
			t.Errorf("reqServer should be %v but %v", expectSgMember, sgMember)
		}

		respJSON := ""
		res.Header()["Content-Type"] = []string{"application/json"}
		fmt.Fprint(res, respJSON)
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	opts := Opts{
		Username: "admin",
		Password: "passwd",
		Target:   u.Host,
		Insecure: true,
	}

	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	var sgMember SGNameAndMember
	err = json.Unmarshal([]byte(testSGMember), &sgMember)
	if err != nil {
		t.Error("request body should be decoded as json")
	}

	err = client.ServiceGroupMemberCreate(&sgMember)

	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
}

func TestServiceGroupMemberDelete(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.service_group.member.delete" {
			t.Error("request QueryString should be method=slb.service_group.member.deletee but :", query["method"])
		}

		var sgMember SGNameAndMember
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&sgMember)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		var expectSgMember SGNameAndMember
		err = json.Unmarshal([]byte(testSGMemberSubset), &expectSgMember)
		if err != nil {
			t.Error("testServerData should be decoded as json")
		}

		if !reflect.DeepEqual(sgMember, expectSgMember) {
			t.Errorf("reqServer should be %v but %v", expectSgMember, sgMember)
		}

		respJSON := ""
		res.Header()["Content-Type"] = []string{"application/json"}
		fmt.Fprint(res, respJSON)
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	opts := Opts{
		Username: "admin",
		Password: "passwd",
		Target:   u.Host,
		Insecure: true,
	}

	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	var sgMember SGNameAndMember
	err = json.Unmarshal([]byte(testSGMemberSubset), &sgMember)
	if err != nil {
		t.Error("request body should be decoded as json")
	}

	err = client.ServiceGroupMemberDelete(&sgMember)

	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
}

func TestServiceGroupSearch(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.service_group.search" {
			t.Error("request QueryString should be method=slb.service_group.search but :", query["method"])
		}

		var jsonBody struct {
			Neme string `json:"name"`
		}
		err := json.NewDecoder(req.Body).Decode(&jsonBody)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		if jsonBody.Neme != "VBX2842DH" {
			t.Errorf("reqServer should be VBX2842DH but %v", jsonBody.Neme)
		}

		respJSON := testServiceGroup
		res.Header()["Content-Type"] = []string{"application/json"}
		fmt.Fprint(res, respJSON)
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	opts := Opts{
		Username: "admin",
		Password: "passwd",
		Target:   u.Host,
		Insecure: true,
	}

	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"
	sg, err := client.ServiceGroupSearch("VBX2842DH")
	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}

	var expectServiceGroup struct {
		ServiceGroup ServiceGroup `json:"service_group"`
	}
	err = json.Unmarshal([]byte(testServiceGroup), &expectServiceGroup)
	if err != nil {
		t.Error("testServerData should be decoded as json")
	}
	if reflect.DeepEqual(sg, expectServiceGroup.ServiceGroup) {
		t.Errorf("reqServer should be %v but %v", expectServiceGroup.ServiceGroup, sg)
	}

}
