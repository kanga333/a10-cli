package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const testServerData = `
{
    "server": {
        "name": "8"
        ,"host": "8X4"
        ,"gslb_external_address": "106.184.78.162"
        ,"weight": 96
        ,"health_monitor": "8X4D204D"
        ,"status": 0
        ,"conn_limit": 3795251
        ,"conn_limit_log": 0
        ,"conn_resume": 160051
        ,"stats_data": 0
        ,"extended_stats": 0
        ,"slow_start": 0
        ,"spoofing_cache": 0
        ,"template": "8X"
        ,"port_list": [
            {
                "port_num": 32777
                ,"protocol": 2
                ,"status": 0
                ,"weight": 99
                ,"no_ssl": 0
                ,"conn_limit": 2614761
                ,"conn_limit_log": 0
                ,"conn_resume": 219590
                ,"template": "B2RL"
                ,"stats_data": 0
                ,"health_monitor": "(default)"
                ,"extended_stats": 0
            }
            ,{
                "port_num": 1239
                ,"protocol": 2
                ,"status": 0
                ,"weight": 40
                ,"no_ssl": 0
                ,"conn_limit": 6620575
                ,"conn_limit_log": 0
                ,"conn_resume": 70525
                ,"template": ""
                ,"stats_data": 0
                ,"health_monitor": "(default)"
                ,"extended_stats": 0
            }
        ]
    }
}
`

func TestServerSearch(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}

		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}

		query := req.URL.Query()
		if strings.Join(query["format"], "") != "json" {
			t.Error("request QueryString should be format=json but :", query["json"])
		}
		if strings.Join(query["method"], "") != "slb.server.search" {
			t.Error("request QueryString should be method=slb.server.search but :", query["method"])
		}
		if strings.Join(query["session_id"], "") != "FTNFPTD" {
			t.Error("request QueryString should be session_id=FTNFPTD but :", query["method"])
		}

		body, _ := ioutil.ReadAll(req.Body)

		var h Host

		err := json.Unmarshal(body, &h)
		if err != nil {
			t.Error("request body should be decoded as json", string(body))
		}

		if h.Host != "8X4" {
			t.Error("request body should have 8X4 in the name column, but", h.Host)
		}

		respJSON := testServerData
		res.Header()["Content-Type"] = []string{"application/json"}
		fmt.Fprint(res, respJSON)

	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	opts := Opts{
		user:     "admin",
		password: "passwd",
		target:   u.Host,
		insecure: true,
	}

	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	h := Host{
		Host: "8x4",
	}

	s, err := client.ServerSearch(h)
	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
	if s.Name != "8" {
		t.Error("s.Name after ServerSearch() should be 8 but", s.Name)
	}
	if s.Weight != 96 {
		t.Error("s.Weight after ServerSearch() should be 96 but", s.Weight)
	}
	if s.Status != false {
		t.Error("s.Status after ServerSearch() should be false but", s.Status)
	}
	if len(s.PortList) != 2 {
		t.Error("s.Port length after ServerSearch() should be 2 but", len(s.PortList))
	}
	p := s.PortList[0]

	if p.Status != false {
		t.Error("p.Status after ServerSearch() should be false but", p.Status)
	}

	if p.Weight != 99 {
		t.Error("p.Weight after ServerSearch() should be 99 but", p.Weight)
	}
	if p.Template != "B2RL" {
		t.Error("p.Template after ServerSearch() should be B2RL but", p.Template)
	}

}
