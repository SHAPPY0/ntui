package utils

import (
	"os"
	"io"
	"log"
)

type Logger struct {
	LInfo 		*log.Logger
	LWarning 	*log.Logger
	LError 		*log.Logger
}

func NewLogger(logLevel string, file *os.File) *Logger {
	multiWriter := io.MultiWriter(file, os.Stdout)
	l := &Logger{
		LInfo:		log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
		LWarning:	log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		LError:		log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}
	return l
}

func (l *Logger) Log(kind, message string) {
	switch kind {
	case Info:
		l.LInfo.Println(message)
	case Warning:
		l.LWarning.Println(message)
	case Error:
		l.LError.Println(message)
	default:
		log.Println("Unsupported log type %s", kind)
	}
}

func (l *Logger) Info(message string) {
	l.LInfo.Println(message)
}

func (l *Logger) Warn(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Warning(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Error(message string) {
	l.LError.Println(message)
}