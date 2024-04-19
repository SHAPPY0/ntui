package core

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
	"github.com/spf13/cobra"
)

const (
	AppName = "ntui"
	ShortDesc = "A graphical CLI for your Hashicorp Nomad Cluster Management"
	LongDesc = "Ntui is a CLI to view and manage your Hashicorp Nomad Clusters."
	DefaultConfigDir = "." + AppName
	DefaultConfigFile = "config.json"
	DefaultLogDir = "logs"
	DefaultLogFile = "/ntui.log"
)

var EnvVars = map[string]string{
	"NOMAD_ADDR": "NomadBaseUrl",
	"NOMAD_SKIP": "NomadSkipVerify",
	"NOMAD_TOKEN": "NomadToken",
}

type Flags struct {
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
	LogDir 			string		`json:"Log_Dir"`
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

func SetConfigValues(c *Config, flags Flags) {
	if c.HomeDir == "" {
		c.HomeDir = flags.HomeDir
	}
	if c.RefreshRate == 0 {
		c.RefreshRate = flags.RefreshRate
	}
	if flags.Host != "" {
		c.NomadBaseUrl = flags.Host
	}
	if c.LogLevel == "" {
		c.LogLevel = flags.LogLevel
	}
	if c.HomeDir != "" && c.LogDir == "" {
		c.LogDir = c.HomeDir + "/" + DefaultLogDir 
	}
	if flags.Region != "" {
		c.Region = flags.Region
	}
	if flags.Namespace != "" {
		c.Namespace = flags.Namespace
	}
	if flags.Token != "" {
		c.NomadToken = flags.Token
	}
	if flags.SkipVerify != false {
		c.NomadSkipVerify = flags.SkipVerify
	}
}

func (c *Config) Load(flags Flags) (*Config, error) {
	config := c
	configPath := os.Getenv("CONFIG_PATH")
	if flags.ConfigPath != "" {
		configPath = flags.ConfigPath
	}
	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = fmt.Sprintf("%s/%s/%s", homeDir, DefaultConfigDir, DefaultConfigFile)
	}
	if configPath == "" {
		panic(fmt.Errorf("No config file found"))
	}

	configJson, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	if err = json.Unmarshal([]byte(configJson), &config); err != nil {
		return config, err
	}

	SetConfigValues(config, flags)
	for key, value := range EnvVars {
		os.Setenv(key, GetValue(config, value))
	}

	return config, nil
}

func (f *Flags) ParseCommand(cmd *cobra.Command) *Flags {
	customConfigPath, err := cmd.Flags().GetString("config-path")
	if err != nil {
		panic(err)
	}
	f.ConfigPath = customConfigPath

	refreshRate, err := cmd.Flags().GetInt("refresh")
	if err != nil {
		panic(err)
	}
	f.RefreshRate = refreshRate

	homeDir, err := cmd.Flags().GetString("home-dir")
	if err != nil {
		panic(err)
	}
	f.HomeDir = homeDir

	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		panic(err)
	}
	f.LogLevel = logLevel

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		panic(err)
	}
	f.Host = host
	
	namespace, err := cmd.Flags().GetString("namespace")
	if err != nil {
		panic(err)
	}
	f.Namespace = namespace

	region, err := cmd.Flags().GetString("region")
	if err != nil {
		panic(err)
	}
	f.Region = region

	skipVerify, err := cmd.Flags().GetBool("skip-verify")
	if err != nil {
		panic(err)
	}
	f.SkipVerify = skipVerify

	token, err := cmd.Flags().GetString("token")
	if err != nil {
		panic(err)
	}
	f.Token = token

	return f
}