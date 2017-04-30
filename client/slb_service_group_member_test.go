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
