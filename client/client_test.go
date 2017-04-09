package client

import "testing"

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
