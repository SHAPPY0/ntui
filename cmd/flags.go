package cmd

import (
	"os"
)

var (
	AppName = "ntui"
	DefaultConfigDir = "." + AppName
	DefaultConfigFile = "config.json"
	userHomeDir, _ = os.UserHomeDir()
	DefaultRefreshRate = 5
	DefaultLogLevel = "info"
	DefaultConfigPath = userHomeDir + "/" + DefaultConfigDir + "/" + DefaultConfigFile
	DefaultRegion = ""
	DefaultNamespace = ""
	DefaultHomeDir = userHomeDir + "/" + DefaultConfigDir
	DefaultNomadHost = ""
	DefaultNomadToken = ""
	DefaultSkipVerify = false
)

type Flags struct {
	BuildNo			*string
	RefreshRate 	*int
	ConfigPath 		*string
	LogLevel 		*string
	Region 			*string
	Namespace 		*string
	HomeDir			*string
	NomadHost 		*string
	NomadToken		*string
	SkipVerify 		*bool
}

func NewFlags() *Flags {
	f := &Flags{
		RefreshRate:	IntPtr(DefaultRefreshRate),
		ConfigPath:		StrPtr(DefaultConfigPath),
		LogLevel:		StrPtr(DefaultLogLevel),
		Region:			StrPtr(DefaultRegion),
		Namespace:		StrPtr(DefaultNamespace),
		HomeDir:		StrPtr(DefaultHomeDir),
		NomadHost:		StrPtr(DefaultNomadHost),
		NomadToken:		StrPtr(DefaultNomadToken),
		SkipVerify:		BoolPtr(DefaultSkipVerify),
	}
	return f
}

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntPtr(i int) *int {
	return &i
}