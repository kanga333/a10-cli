package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
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
		if strings.Join(query["method"], "") != "authenticate" {
			t.Error("request QueryString should be method=authenticate but :", query["method"])
		}

		var auth struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&auth)
		if err != nil {
			t.Error("request body should be decoded as json", err)
		}

		if auth.UserName != "admin" {
			t.Error("request body should have admin in the user column, but", auth.UserName)
		}
		if auth.Password != "passwd" {
			t.Error("request body should have passwd in the password column, but", auth.Password)
		}

		respJSON := `{"session_id": "FTNFPTD"}`
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

	err = client.Auth()
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if client.token != "FTNFPTD" {
		t.Error("clinet.token after auth() should be FTNFPTD but", client.token)
	}

}

func TestClose(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/services/rest/V2.1/" {
			t.Error("request URL should be /services/rest/V2.1/ but :", req.URL.Path)
		}

		if req.Method != "POST" {
			t.Error("request method should be POST but: ", req.Method)
		}

		query := req.URL.Query()
		if strings.Join(query["method"], "") != "close" {
			t.Error("request QueryString should be method=close but :", query["method"])
		}
		if strings.Join(query["session_id"], "") != "FTNFPTD" {
			t.Error("request QueryString should be session_id=FTNFPTD but :", query["method"])
		}

		respJSON := ``
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

	err = client.Close()
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	client.token = "FTNFPTD"

	err = client.Close()
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if client.token != "" {
		t.Error("clinet.token after close() should be empty but", client.token)
	}
}
