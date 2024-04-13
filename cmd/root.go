package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shappy0/ntui/internal/core"
)

var (
	Version, Commit = "v0.0", "NA"
	rootCmd = &cobra.Command{
		Use:	core.AppName,
		Short:	core.ShortDesc,
		Long:	core.LongDesc,
		RunE:	Init,
	}
)

func Run() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	rootCmd.AddCommand(
		versionCmd(),
		configCmd(),
	)
	InitFlags()
	if Err := rootCmd.Execute(); Err != nil {
		// log.Error().Msg(Err.Error())
	}
}

func ParseCmdParams(cmd *cobra.Command, params *core.CmdParams) *core.CmdParams {
	CustomConfigPath, Err := cmd.Flags().GetString("config-path")
	if Err != nil {
		panic(Err)
	}
	params.ConfigPath = CustomConfigPath

	RefreshRate, Err := cmd.Flags().GetInt("refresh")
	if Err != nil {
		panic(Err)
	}
	params.RefreshRate = RefreshRate

	HomeDir, Err := cmd.Flags().GetString("home-dir")
	if Err != nil {
		panic(Err)
	}
	params.HomeDir = HomeDir

	LogLevel, Err := cmd.Flags().GetString("log-level")
	if Err != nil {
		panic(Err)
	}
	params.LogLevel = LogLevel

	Namespace, Err := cmd.Flags().GetString("namespace")
	if Err != nil {
		panic(Err)
	}
	params.Namespace = Namespace

	Region, Err := cmd.Flags().GetString("region")
	if Err != nil {
		panic(Err)
	}
	params.Region = Region

	SkipVerify, Err := cmd.Flags().GetBool("skip-verify")
	if Err != nil {
		panic(Err)
	}
	params.SkipVerify = SkipVerify

	Token, Err := cmd.Flags().GetString("token")
	if Err != nil {
		panic(Err)
	}
	params.Token = Token

	return params
}

func Init(cmd *cobra.Command, args []string) error {
	var Params core.CmdParams
	ParseCmdParams(cmd, &Params)
	Config, Err := core.NewConfig().Load(Params)
	if Err != nil {
		return fmt.Errorf("Unable to load config file")
	}	
	App, Err := core.NewApp(Config)
	if Err != nil {
		return Err
	}
	if Err := App.Init(); Err != nil {
		return Err
	}
	log.Info().Msg("Starting NTUI")
	App.RunX()
	return nil
}

func InitFlags() {
	flags := NewFlags()
	rootCmd.Flags().StringVarP(
		flags.LogLevel,
		"log-level", "l",
		DefaultLogLevel,
		"Specify Log Level",
	)
	rootCmd.Flags().IntVarP(
		flags.RefreshRate,
		"refresh", "r",
		DefaultRefreshRate,
		"Specify the default refresh rate as integer[sec]",
	)
	rootCmd.Flags().StringVarP(
		flags.ConfigPath,
		"config-path", "c",
		"",
		"Specify the config file path",
	)
	rootCmd.Flags().StringVarP(
		flags.HomeDir,
		"home-dir", "",
		DefaultHomeDir,
		"Specify the home dir of ntui app",
	)
	rootCmd.Flags().StringVarP(
		flags.Region,
		"region", "",
		DefaultRegion,
		"Specify default Nomad Region",
	)
	rootCmd.Flags().StringVarP(
		flags.Namespace,
		"namespace", "n",
		DefaultNamespace,
		"Specify default Nomad Namespace",
	)
	rootCmd.Flags().StringVarP(
		flags.NomadHost,
		"host", "",
		DefaultNomadHost,
		"Specify Nomad Host",
	)
	rootCmd.Flags().StringVarP(
		flags.NomadToken,
		"token", "t",
		DefaultNomadToken,
		"Specify Nomad Token",
	)
	rootCmd.Flags().BoolVarP(
		flags.SkipVerify,
		"skip-verify", "",
		DefaultSkipVerify,
		"Specify skip nomad ssl verify",
	)
	rootCmd.Flags()
}