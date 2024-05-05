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
	multiWriter := io.MultiWriter(file) //file, os.Stdout
	multiWriterErr := io.MultiWriter(file) //file, os.Stderr
	l := &Logger{
		LInfo:		log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
		LWarning:	log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime),
		LError:		log.New(multiWriterErr, "ERROR: ", log.Ldate|log.Ltime),
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

func (l *Logger) Infof(format string, v ...any) {
	l.LInfo.Printf(format, v...)
}

func (l *Logger) Warn(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.LWarning.Printf(format, v...)
}

func (l *Logger) Warning(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.LWarning.Printf(format, v...)
}

func (l *Logger) Error(message string) {
	l.LError.Println(message)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.LError.Printf(format, v...)
}