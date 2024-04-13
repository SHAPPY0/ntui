package cmd

import (
	"fmt"
	"os"
	// "bytes"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
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
	var (
		Configs 		core.Config
		// PrettyConfig	bytes.Buffer
	)
	ConfigPath := os.Getenv("CONFIG_PATH") 
	if ConfigPath == "" {
		HomeDir, _ := os.UserHomeDir()
		ConfigPath = fmt.Sprintf("%s/%s/%s", HomeDir, core.DefaultConfigDir, core.DefaultConfigFile)
	}
	ConfigJson, Err := ioutil.ReadFile(ConfigPath)
	if Err != nil {
		panic(Err)
	}
	if Err = json.Unmarshal([]byte(ConfigJson), &Configs); Err != nil {
		panic(Err)
	}

	CTypes := reflect.TypeOf(Configs)
	CValues := reflect.ValueOf(Configs)
	for I := 0; I < CTypes.NumField(); I++ {
		key := CTypes.Field(I)
		value := CValues.Field(I)
		var printValue interface{} 
		if value.Interface() != "" {
			printValue = value.Interface()
		}
		fmt.Printf("%s = %v\n", key.Name, printValue)
	}

	// ConfigJson, Err = json.Marshal(Configs)
	// if Err = json.Indent(&PrettyConfig, ConfigJson, "", "\t"); Err != nil {
	// 	panic(Err)
	// }
	// fmt.Println(string(PrettyConfig.Bytes()))
	
}