package utils

import (
	"time"
	"github.com/shappy0/ntui/internal/models"
)

const (
	DefaultFlashDelay	=	3 * time.Second
)

type Alert struct {
	Message 	models.AlertMessage
	Duration	time.Duration
	AlertChan 	chan models.AlertMessage
}

func NewAlert() *Alert {
	a := &Alert{
		Duration:		DefaultFlashDelay,
		AlertChan:		make(models.AlertChan, 3),
	}
	return a
}

func (a *Alert) Channel() models.AlertChan {
	return a.AlertChan
}

func (a *Alert) Info(Msg string) {
	a.SendMessage(Info, Msg)
}

func (a *Alert) Warning(Msg string) {
	a.SendMessage(Warning, Msg)
}

func (a *Alert) Error(Msg string) {
	a.SendMessage(Error, Msg)
}

func (a *Alert) SendMessage(Type, Msg string) {
	a.Message = models.AlertMessage{Type: Type, Text: Msg}
	a.AlertChan <-a.Message
	go a.Hide()
}

func (a *Alert) Hide() {
	for {
		select{
		case <-time.After(a.Duration):
			a.AlertChan <-models.AlertMessage{}
			return
		}
	}
}
