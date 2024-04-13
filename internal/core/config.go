package core

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
)

const (
	AppName = "ntui"
	ShortDesc = "A graphical CLI for your Hashicorp Nomad Cluster Management"
	LongDesc = "Ntui is a CLI to view and manage your Hashicorp Nomad Clusters."
	DefaultConfigDir = "." + AppName
	DefaultConfigFile = "config.json"
)

var EnvVars = map[string]string{
	"NOMAD_ADDR": "NomadBaseUrl",
	"NOMAD_SKIP": "NomadSkipVerify",
	"NOMAD_TOKEN": "NomadToken",
}

type CmdParams struct {
	ConfigPath		string
	RefreshRate		int
	HomeDir			string
	LogLevel		string
	Region			string
	Namespace 		string
	Host 			string
	Token			string
	SkipVerify		bool
}

type Config struct {
	AppName			string 		`json:"App_Name"`
	HomeDir 		string 		`json:"Home_Dir"`
	LogLevel 		string		`json:"Log_Level"`
	RefreshRate 	int			`json:"Refresh_Rate"`
	NomadBaseUrl	string		`json:"Nomad_Server_Base_Url"`
	NomadHttpAuth	string		`json:"Nomad_Http_Auth"`
	NomadToken		string		`json:"Nomad_Token"`
	Region			string		`json:"Nomad_Region"`
	Namespace		string		`json:"Nomad_Namespace"`
	NomadCaCert 	string 		`json:"Nomad_Cacert"`
	NomadCaPath		string 		`json:"Nomad_Capath"`
	NomadClientCert string 		`json:"Nomad_Client_Cert"`
	NomadClientKey 	string 		`json:"Nomad_Client_Key"`
	NomadTlsServer	string 		`json:"Nomad_Tls_Server"`
	NomadSkipVerify	bool		`json:"Nomad_Skip_Verify"`
}

func NewConfig() *Config {
	c := &Config{
		AppName:	AppName,
	}
	return c
}

func (c *Config) SetRegion(region string) {
	c.Region = region
}

func (c *Config) GetRegion() string {
	return c.Region
}

func (c *Config) SetNamespace(ns string) {
	c.Namespace = ns
}

func (c *Config) GetNamespace() string {
	return c.Namespace
}

func (c *Config) IsRegionInConfig() bool {
	return c.Region != "" && c.Namespace != ""
}

func GetValue(c *Config, key string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(key)
	return string(f.String())
}

func SetConfigValues(c *Config, params CmdParams) {
	if params.HomeDir != "" {
		c.HomeDir = params.HomeDir
	} else {
		HomeDir, _ := os.UserHomeDir()
		c.HomeDir = HomeDir
	}
	if params.RefreshRate > 0 {
		c.RefreshRate = params.RefreshRate
	}
	if params.Host != "" {
		c.NomadBaseUrl = params.Host
	}
	if params.LogLevel != "" {
		c.LogLevel = params.LogLevel
	}
	if params.Region != "" {
		c.Region = params.Region
	}
	if params.Namespace != "" {
		c.Namespace = params.Namespace
	}
	if params.Token != "" {
		c.NomadToken = params.Token
	}
	if params.SkipVerify != false {
		c.NomadSkipVerify = params.SkipVerify
	}
}

func (c *Config) Load(params CmdParams) (*Config, error) {
	Config := c
	ConfigPath := os.Getenv("CONFIG_PATH")
	if params.ConfigPath != "" {
		ConfigPath = params.ConfigPath
	}
	if ConfigPath == "" {
		HomeDir, _ := os.UserHomeDir()
		ConfigPath = fmt.Sprintf("%s/%s/%s", HomeDir, DefaultConfigDir, DefaultConfigFile)
	}
	if ConfigPath == "" {
		panic(fmt.Errorf("No config file found"))
	}

	ConfigJson, Err := ioutil.ReadFile(ConfigPath)
	if Err != nil {
		return Config, Err
	}
	if Err = json.Unmarshal([]byte(ConfigJson), &Config); Err != nil {
		return Config, Err
	}

	SetConfigValues(Config, params)

	for Key, Value := range EnvVars {
		os.Setenv(Key, GetValue(Config, Value))
	}

	return Config, nil
}