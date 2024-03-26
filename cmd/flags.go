package cmd

const (
	DefaultRefreshRate = 5
	DefaultLogLevel = "info"
	DefaultConfigPath = ""
	DefaultRegion = ""
	DefaultNamespace = ""
)

type Flags struct {
	Version 		*string
	BuildNo			*string
	RefreshRate 	*int
	ConfigPath 		*string
	LogLevel 		*string
	Region 			*string
	Namespace 		*string
}

func NewFlags() *Flags {
	f := &Flags{
		Version:		StrPtr(""),
		RefreshRate:	IntPtr(DefaultRefreshRate),
		ConfigPath:		StrPtr(DefaultConfigPath),
		LogLevel:		StrPtr(DefaultLogLevel),
		Region:			StrPtr(DefaultRegion),
		Namespace:		StrPtr(DefaultNamespace),
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