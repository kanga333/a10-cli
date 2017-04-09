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

		body, _ := ioutil.ReadAll(req.Body)

		var jsonBody struct {
			User     string `json:"user"`
			Password string `json:"password"`
		}

		err := json.Unmarshal(body, &jsonBody)
		if err != nil {
			t.Error("request body should be decoded as json", string(body))
		}

		if jsonBody.User != "admin" {
			t.Error("request body should have admin in the user column, but", jsonBody.User)
		}
		if jsonBody.Password != "passwd" {
			t.Error("request body should have passwd in the password column, but", jsonBody.Password)
		}

		respJSON := `{"session_id": "FTNFPTD"}`
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

	err = client.Auth()
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if client.token != "FTNFPTD" {
		t.Error("clinet.token after auth() should be FTNFPTD but", client.token)
	}

}
