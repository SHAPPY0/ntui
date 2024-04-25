package cmd

import (
	"fmt"
	"os"
	"reflect"
	"github.com/spf13/cobra"
	"github.com/BurntSushi/toml"
	"github.com/shappy0/ntui/internal/core"
)

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:	"version",
		Short:	"Show version/build info",
		Long:	"Show version/build info",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("Version:", Version)
			fmt.Println("Commit:", Commit)
		},
	}
	return &cmd
}

func configCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:	"config",
		Short:	"Show configurations",
		Long:	"Show configurations",
		Run:	func(cmd *cobra.Command, args []string) {
			ShowConfigs()
		},
	}
	return &cmd
}

func ShowConfigs() {
	var config  core.Config
	configPath := os.Getenv("CONFIG_PATH") 
	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = fmt.Sprintf("%s/%s/%s", homeDir, core.DefaultConfigDir, core.DefaultConfigFile)
	}
	if _, err := os.Stat(configPath); err != nil {
		panic("No configs found")
	}
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		fmt.Println(err)
	}
	cTypes := reflect.TypeOf(config)
	cValues := reflect.ValueOf(config)
	for i := 0; i < cTypes.NumField(); i++ {
		key := cTypes.Field(i)
		value := cValues.Field(i)
		var printValue interface{} 
		if value.Interface() != "" {
			printValue = value.Interface()
		}
		fmt.Printf("%s=%v\n", key.Name, printValue)
	}
}