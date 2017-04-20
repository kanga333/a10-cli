package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const testConfig = `
a10cli:
    username: admin
    password: password
    target: 127.0.0.1
    insecure: true
    proxy: ""
host: test-server
`

func TestLoadConft(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "config")
	if err != nil {
		t.Fatalf("shoud not raise error: %v", err)
	}
	defer os.Remove(file.Name())
	_, err = file.Write([]byte(testConfig))
	if err != nil {
		t.Fatalf("shoud not raise error: %v", err)
	}

	conf, err := loadConf(file.Name())
	if err != nil {
		t.Fatalf("shoud not raise error: %v", err)
	}

	if conf.A10.Username != "admin" {
		t.Error("conf.A10.Username should be admin but", conf.A10.Username)
	}
	if conf.A10.Insecure != true {
		t.Error("conf.A10.Insecure should be true but", conf.A10.Insecure)
	}

	if conf.Host != "test-server" {
		t.Error("conf.Host should be test-server but", conf.Host)
	}

}
