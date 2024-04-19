package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
	"github.com/shappy0/ntui/internal/core"
	"github.com/shappy0/ntui/internal/utils"
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
	rootCmd.AddCommand(
		versionCmd(),
		configCmd(),
	)
	InitFlags()
	if Err := rootCmd.Execute(); Err != nil {
		// log.Error().Msg(Err.Error())
	}
}

func Init(cmd *cobra.Command, args []string) error {
	var flags core.Flags
	flags.ParseCommand(cmd)
	
	Config, Err := core.NewConfig().Load(flags)
	if Err != nil {
		return fmt.Errorf("Unable to load config file")
	}
	
	if err := utils.EnsureDirPath(Config.LogDir + core.DefaultLogFile, utils.DefaultDirMod); err != nil {
		return err
	}

	logFile, err := os.OpenFile(Config.LogDir + core.DefaultLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, utils.DefaultFileMod)
	if err != nil {
		return err
	}

	defer func() {
		if logFile != nil {
			_ = logFile.Close()
		}
	}()
	
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v.\n", err)
		}
	}()

	var logger = utils.NewLogger(Config.LogLevel, logFile)

	app, err := core.NewApp(Config, logger)
	if err != nil {
		return err
	}

	if err := app.Init(); err != nil {
		logger.Error(err.Error())
		return err
	}
	
	logger.Info("Starting ntui ...")
	
	app.RunX()
	
	return nil
}

func InitFlags() {
	flags := NewFlags()

	rootCmd.Flags().StringVarP(
		flags.LogLevel,
		"log-level", "l",
		DefaultLogLevel,
		"Set Log Level (info, warning, error)",
	)
	rootCmd.Flags().IntVarP(
		flags.RefreshRate,
		"refresh", "r",
		DefaultRefreshRate,
		"Set the default refresh rate as integer[sec]",
	)
	rootCmd.Flags().StringVarP(
		flags.ConfigPath,
		"config-path", "c",
		DefaultConfigPath,
		"Set the config file path",
	)
	rootCmd.Flags().StringVarP(
		flags.HomeDir,
		"home-dir", "",
		DefaultHomeDir,
		"Set the home dir of ntui app",
	)
	rootCmd.Flags().StringVarP(
		flags.Region,
		"region", "",
		DefaultRegion,
		"Set default Nomad Region",
	)
	rootCmd.Flags().StringVarP(
		flags.Namespace,
		"namespace", "n",
		DefaultNamespace,
		"Set default Nomad Namespace",
	)
	rootCmd.Flags().StringVarP(
		flags.NomadHost,
		"host", "",
		DefaultNomadHost,
		"Set Nomad Host",
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