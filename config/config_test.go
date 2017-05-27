package config

import (
	"io/ioutil"
	"os"
	"testing"
)

const testConfig = `
a10ballancer:
  username: admin
  password: password
  target: 127.0.0.1
  insecure: true
  proxy: ""
server:
  name: "8"
  host: 8X4
  gslb_external_address: 106.184.78.162
  weight: 96
  health_monitor: 8X4D204D
  conn_limit: 3795251
  conn_limit_log: false
  conn_resume: 160051
  stats_data: false
  extended_stats: false
  slow_start: false
  spoofing_cache: false
  template: 8X
  port_list:
    32777:
      protocol: 2
      weight: 99
      no_ssl: true
      conn_limit: 2614761
      conn_limit_log: false
      conn_resume: 219590
      template: B2RL
      stats_data: false
      health_monitor: (default)
      extended_stats: false
      sg_name: VBX2842DH
      sg_template: "B"
      sg_priority: 9
      ag_stats_data: false
    1239:
      sg_name: 3
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

	conf, err := LoadConf(file.Name())
	if err != nil {
		t.Fatalf("shoud not raise error: %v", err)
	}

	if conf.A10.Username != "admin" {
		t.Error("conf.A10.Username should be admin but", conf.A10.Username)
	}
	if conf.A10.Insecure != true {
		t.Error("conf.A10.Insecure should be true but", conf.A10.Insecure)
	}

	if conf.Server.Name != "8" {
		t.Error("conf.Server.Name should be 8 but", conf.Server)
	}

	if *conf.Server.Weight != 96 {
		t.Error("conf.Server.Weight should be 96 but", *conf.Server.Weight)
	}

	if len(conf.Server.PortList) != 2 {
		t.Error("port list should be 2 but", len(conf.Server.PortList))
	}

	if *conf.Server.PortList[32777].NoSsl != true {
		t.Error("*conf.Server.PortList[32777].NoSsl should be true but", *conf.Server.PortList[32777].NoSsl)
	}

	if conf.Server.PortList[32777].SGName != "VBX2842DH" {
		t.Error("conf.Server.PortList[32777].SGName should be VBX2842DH but", conf.Server.PortList[32777].SGName)
	}
}
