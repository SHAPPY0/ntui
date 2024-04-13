package cmd

const (
	DefaultRefreshRate = 5
	DefaultLogLevel = "info"
	DefaultConfigPath = ""
	DefaultRegion = ""
	DefaultNamespace = ""
	DefaultHomeDir = ""
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
		NomadToken:		StrPtr(DefaultNomadHost),
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