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
		if strings.Join(query["method"], "") != "slb.server.search" {
			t.Error("request QueryString should be method=slb.server.search but :", query["method"])
		}

		var host struct {
			Host string `json:"Host"`
		}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&host)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		if host.Host != "8X4" {
			t.Error("request body should have 8X4 in the name column, but", host.Host)
		}

		respJSON := testServerData
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

	client, err := NewClient(&opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	server, err := client.ServerSearch("8X4")

	var expectServer struct {
		Server Server `json:"server"`
	}
	err = json.Unmarshal([]byte(testServerData), &expectServer)
	if err != nil {
		t.Error("testServerData should be decoded as json")
	}

	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}

	if !reflect.DeepEqual(server, &expectServer.Server) {
		t.Errorf("server should be %v but %v", server, &expectServer.Server)
	}

}

func TestServerCreate(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.server.create" {
			t.Error("request QueryString should be method=slb.server.create but :", query["method"])
		}

		var reqServer Server
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqServer)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		var expectServer struct {
			Server Server `json:"server"`
		}
		err = json.Unmarshal([]byte(testServerData), &expectServer)
		if err != nil {
			t.Error("testServerData should be decoded as json")
		}

		if !reflect.DeepEqual(reqServer, expectServer.Server) {
			t.Errorf("reqServer should be %v but %v", reqServer, expectServer.Server)
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

	client, err := NewClient(&opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	var server struct {
		Server Server `json:"server"`
	}
	err = json.Unmarshal([]byte(testServerData), &server)
	if err != nil {
		t.Error("request body should be decoded as json")
	}

	err = client.ServerCreate(&server.Server)

	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
}

func TestServerDelete(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.server.delete" {
			t.Error("request QueryString should be method=slb.server.delete but :", query["method"])
		}

		var host struct {
			Host string `json:"Host"`
		}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&host)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		if host.Host != "8X4" {
			t.Error("request body should have 8X4 in the name column, but", host.Host)
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

	client, err := NewClient(&opts)
	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	err = client.ServerDelete("8X4")
	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
}

func TestServerUpdate(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}
		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}
		query := req.URL.Query()
		if strings.Join(query["method"], "") != "slb.server.update" {
			t.Error("request QueryString should be method=slb.server.update but :", query["method"])
		}

		var reqServer Server
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqServer)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		var expectServer struct {
			Server Server `json:"server"`
		}
		err = json.Unmarshal([]byte(testServerData), &expectServer)
		if err != nil {
			t.Error("testServerData should be decoded as json")
		}

		if !reflect.DeepEqual(reqServer, expectServer.Server) {
			t.Errorf("reqServer should be %v but %v", reqServer, expectServer.Server)
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

	client, err := NewClient(&opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	client.token = "FTNFPTD"

	var server struct {
		Server Server `json:"server"`
	}
	err = json.Unmarshal([]byte(testServerData), &server)
	if err != nil {
		t.Error("request body should be decoded as json")
	}

	err = client.ServerUpdate(&server.Server)

	if err != nil {
		t.Fatalf("should not raise error: %v", err)
	}
}
