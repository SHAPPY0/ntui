package core

import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
	"reflect"
	"github.com/spf13/cobra"
	"github.com/BurntSushi/toml"
	"github.com/akiyosi/tomlwriter"
)

var defaultConfigPath = ""

const (
	AppName = "ntui"
	ShortDesc = "A graphical CLI for your Hashicorp Nomad Cluster Management"
	LongDesc = "Ntui is a CLI to view and manage your Hashicorp Nomad Clusters."
	DefaultConfigDir = "." + AppName
	DefaultConfigFile = "config.toml"
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
	AppName			string 		`toml:"App_Name"`
	HomeDir 		string 		`toml:"Home_Dir"`
	LogLevel 		string		`toml:"Log_Level"`
	LogDir 			string		`toml:"Log_Dir"`
	RefreshRate 	int			`toml:"Refresh_Rate"`
	NomadBaseUrl	string		`toml:"Nomad_Server_Base_Url"`
	NomadHttpAuth	string		`toml:"Nomad_Http_Auth"`
	NomadToken		string		`toml:"Nomad_Token"`
	Region			string		`toml:"Nomad_Region"`
	Namespace		string		`toml:"Nomad_Namespace"`
	NomadCaCert 	string 		`toml:"Nomad_Cacert"`
	NomadCaPath		string 		`toml:"Nomad_Capath"`
	NomadClientCert string 		`toml:"Nomad_Client_Cert"`
	NomadClientKey 	string 		`toml:"Nomad_Client_Key"`
	NomadTlsServer	string 		`toml:"Nomad_Tls_Server"`
	NomadSkipVerify	bool		`toml:"Nomad_Skip_Verify"`
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
	if c.RefreshRate == 0 || (c.RefreshRate > 0 && flags.RefreshRate != 5) {
		c.RefreshRate = flags.RefreshRate
	}
	if flags.Host != "" {
		c.NomadBaseUrl = flags.Host
	}
	if c.LogLevel != "" {
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
	if flags.SkipVerify == true {
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
	defaultConfigPath = configPath 
	if _, err := os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("Invalid config file path")
		} else {
			return nil, fmt.Errorf("Unable to load config file.")
		}
	}
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		fmt.Println(err)
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

func (c *Config) UpdateConfigFile(key, newValue, oldValue string) {
	if defaultConfigPath != "" {
		file, _ := ioutil.ReadFile(defaultConfigPath)
		file, _ = tomlwriter.WriteValue(
			newValue, 
			file, 
			nil, 
			key, 
			oldValue,
		)
	}

}