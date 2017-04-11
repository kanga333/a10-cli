package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	opts := Opts{
		user:     "admin",
		password: "passwd",
		target:   "127.0.0.1",
	}
	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	if client.user != "admin" {
		t.Error("should be admin but :", client.user)
	}

	expect := "https://127.0.0.1/services/rest/V2.1/"
	actual := client.baseURL.String()
	if actual != expect {
		t.Errorf("BaseURL should be: \n%s\nbut:\n%s", expect, actual)
	}
}

func TestProxy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "I'm Proxy")
	}))
	defer ts.Close()

	opts := Opts{
		user:     "admin",
		password: "passwd",
		target:   "127.0.0.1",
		proxy:    ts.URL,
	}

	client, err := NewClient(opts)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	resp, err := client.httpClient.Get("http://192.0.2.0")
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if string(b) != "I'm Proxy" {
		t.Error("response body should be I'm Proxy but", b)
	}
}
