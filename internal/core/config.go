package core

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
)

var EnvVars = map[string]string{
	"NOMAD_ADDR": "NomadBaseUrl",
	"NOMAD_SKIP": "NomadSkipVerify",
	"NOMAD_TOKEN": "NomadToken",
}

type Config struct {
	AppName			string 		`json:"AppName"`
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
		AppName:	"ntui-app",
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

func (c *Config) Load() (*Config, error) {
	ConfigPath := "/home/ipp-dev/.ntui/config.json"
	Config := c
	ConfigJson, Err := ioutil.ReadFile(ConfigPath)
	if Err != nil {
		return Config, Err
	}
	if Err = json.Unmarshal([]byte(ConfigJson), &Config); Err != nil {
		return Config, Err
	}
	for Key, Value := range EnvVars {
		os.Setenv(Key, GetValue(Config, Value))
	}
	return Config, nil
}