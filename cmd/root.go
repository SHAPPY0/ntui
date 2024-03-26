package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shappy0/ntui/internal/core"
)

var (
	rootCmd = &cobra.Command{
		Use:	core.AppName,
		Short:	core.ShortDesc,
		Long:	core.LongDesc,
		RunE:	Init,
	}
)

func Run() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting ntui")
	InitFlags()
	if Err := rootCmd.Execute(); Err != nil {
		fmt.Println(Err)
		log.Error().Msg(Err.Error())
	}
}

func Init(cmd *cobra.Command, args []string) error {
	Config, Err := core.NewConfig().Load()
	if Err != nil {
		return fmt.Errorf("While loading config file")
	}	
	App, Err := core.NewApp(Config)
	if Err != nil {
		return Err
	}
	if Err := App.Init(); Err != nil {
		return Err
	}
	App.RunX()
	return nil
}

func InitFlags() {
	flags := NewFlags()
	rootCmd.Flags().StringVarP(
		flags.Version,
		"version", "v",
		"",
		"Get the version of ntui app",
	)
	rootCmd.Flags().InitVarP(
		flags.RefreshRate,
		"refresh", "r",
		DefaultRefreshRate,
		"Specify the default refresh rate as an integer(sec)",
	)
	rootCmd.Flags().StringVarP(
		flags.ConfigPath,
		"config_path", "c",
		"Specify the config file path",
	)
	rootCmd.Flags()
}