package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Info = "info"
	Warning = "warning"
	Error = "error"
)

func SetLogLevels() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func log(kind, message string) {
	switch kind {
	case Info:
		log.Info().Msg(message)
	case Warning:
		log.Warn().Msg(message)
	case Error:
		log.Error().Msg(message)
	default:
		fmt.Println("Invalid Log Type %s", kind)
	}
}