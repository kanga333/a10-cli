package config

import (
	"io/ioutil"

	"github.com/ghodss/yaml"

	"fmt"

	"github.com/kanga333/a10-cli/client"
)

var (
	defaultServer = client.Server{
		GslbExternalAddress: "0.0.0.0",
		Weight:              1,
		HealthMonitor:       "(default)",
		Status:              false,
		ConnLimit:           8000000,
		ConnLimitLog:        true,
		ConnResume:          0,
		StatsData:           true,
		ExtendedStats:       false,
		SlowStart:           false,
		SpoofingCache:       false,
		Template:            "default",
		PortList:            []client.Port{},
	}
	defaultPort = client.Port{
		Protocol:      2,
		Status:        true,
		Weight:        1,
		NoSsl:         false,
		ConnLimit:     8000000,
		ConnLimitLog:  true,
		ConnResume:    1,
		Template:      "default",
		StatsData:     true,
		HealthMonitor: "(default)",
		ExtendedStats: false,
	}
	defaultSGMember = client.SGNameAndMember{
		Member: client.Member{
			Template:  "default",
			Priority:  1,
			Status:    true,
			StatsData: true,
		},
	}
)

// Config is a structure that expresses the setting required by a10
type Config struct {
	A10    client.Opts  `json:"a10ballancer"`
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Name                string             `json:"name"`
	Host                string             `json:"host"`
	GslbExternalAddress *string            `json:"gslb_external_address"`
	Weight              *int               `json:"weight"`
	HealthMonitor       *string            `json:"health_monitor"`
	ConnLimit           *int               `json:"conn_limit"`
	ConnLimitLog        *bool              `json:"conn_limit_log"`
	ConnResume          *int               `json:"conn_resume"`
	StatsData           *bool              `json:"stats_data"`
	ExtendedStats       *bool              `json:"extended_stats"`
	SlowStart           *bool              `json:"slow_start"`
	SpoofingCache       *bool              `json:"spoofing_cache"`
	Template            *string            `json:"template"`
	PortList            map[int]PortConfig `json:"port_list"`
}

type PortConfig struct {
	Protocol      *int    `json:"protocol"`
	Weight        *int    `json:"weight"`
	NoSsl         *bool   `json:"no_ssl"`
	ConnLimit     *int    `json:"conn_limit"`
	ConnLimitLog  *bool   `json:"conn_limit_log"`
	ConnResume    *int    `json:"conn_resume"`
	Template      *string `json:"template"`
	StatsData     *bool   `json:"stats_data"`
	HealthMonitor *string `json:"health_monitor"`
	ExtendedStats *bool   `json:"extended_stats"`
	SGName        string  `json:"sg_name"`
	SGTemplate    *string `json:"sg_template"`
	SGPriority    *int    `json:"sg_priority"`
	SGStatsData   *bool   `json:"ag_stats_data"`
}

// LoadConf reads the yaml setting from the specified path
func LoadConf(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) GetCliOpts() (*client.Opts, error) {
	if c.A10.Username == "" {
		return nil, validFail("username", c.A10.Username)
	}
	if c.A10.Password == "" {
		return nil, validFail("password", c.A10.Password)
	}
	return &c.A10, nil
}

func (c *Config) GetServer() (*client.Server, error) {
	if c.Server.Name == "" {
		return nil, validFail("name", c.Server.Name)
	}
	if c.Server.Host == "" {
		return nil, validFail("host", c.Server.Host)
	}
	server := defaultServer
	server.Name = c.Server.Name
	server.Host = c.Server.Host

	copySrtConf(&server.GslbExternalAddress, c.Server.GslbExternalAddress)
	copyIntConf(&server.Weight, c.Server.Weight)
	copySrtConf(&server.HealthMonitor, c.Server.HealthMonitor)
	copyIntConf(&server.ConnLimit, c.Server.ConnLimit)
	copyNumBoolConf(&server.ConnLimitLog, c.Server.ConnLimitLog)
	copyIntConf(&server.ConnResume, c.Server.ConnResume)
	copyNumBoolConf(&server.StatsData, c.Server.StatsData)
	copyNumBoolConf(&server.ExtendedStats, c.Server.ExtendedStats)
	copyNumBoolConf(&server.SlowStart, c.Server.SlowStart)
	copyNumBoolConf(&server.SpoofingCache, c.Server.SpoofingCache)
	copySrtConf(&server.Template, c.Server.Template)

	for num, conf := range c.Server.PortList {
		port := defaultPort
		port.PortNum = num
		copyIntConf(&port.Protocol, conf.Protocol)
		copyIntConf(&port.Weight, conf.Weight)
		copyNumBoolConf(&port.NoSsl, conf.NoSsl)
		copyIntConf(&port.ConnLimit, conf.ConnLimit)
		copyNumBoolConf(&port.ConnLimitLog, conf.ConnLimitLog)
		copyIntConf(&port.ConnResume, conf.ConnResume)
		copySrtConf(&port.Template, conf.Template)
		copyNumBoolConf(&port.StatsData, conf.StatsData)
		copySrtConf(&port.HealthMonitor, conf.HealthMonitor)
		copyNumBoolConf(&port.ExtendedStats, conf.ExtendedStats)

		server.PortList = append(server.PortList, port)
	}

	return &server, nil
}

func (c *Config) GetSGNameAndMembers() ([]client.SGNameAndMember, error) {
	if c.Server.Name == "" {
		return nil, validFail("name", c.Server.Name)
	}
	if len(c.Server.PortList) == 0 {
		return nil, validFail("port_list_num", "0")
	}
	var sgms []client.SGNameAndMember
	for num, conf := range c.Server.PortList {
		if conf.SGName == "" {
			return nil, validFail("sg_name", conf.SGName)
		}

		sgm := defaultSGMember
		sgm.Name = conf.SGName
		sgm.Member.Server = c.Server.Name
		sgm.Member.Port = num
		copySrtConf(&sgm.Member.Template, conf.SGTemplate)
		copyIntConf(&sgm.Member.Priority, conf.SGPriority)
		copyNumBoolConf(&sgm.Member.Status, conf.SGStatsData)

		sgms = append(sgms, sgm)

	}

	return sgms, nil
}

func validFail(n, v string) error {
	return fmt.Errorf("failed validation of the value: %s of the parameter: %s", v, n)
}

func copySrtConf(dest, src *string) {
	if src != nil {
		*dest = *src
	}
}

func copyIntConf(dest, src *int) {
	if src != nil {
		*dest = *src
	}
}

func copyNumBoolConf(dest *client.NumBool, src *bool) {
	if src != nil {
		*dest = client.NumBool(*src)
	}
}
