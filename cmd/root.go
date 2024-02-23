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
		Use:	"ntui",
		Short:	"Short Desc of ntui",
		Long:	"Long Desc of ntui",
		RunE:	Init,
	}
)

func Run() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting ntui")
	if Err := rootCmd.Execute(); Err != nil {
		fmt.Println(Err)
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
	App.Runx()
	return nil
}