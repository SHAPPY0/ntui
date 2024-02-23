package utils

import (
	"time"
)

var AllListeners []*Listener

type Listener struct {
	Ticker 			*time.Ticker
	StopChan 		chan bool
	Function 		func()
}

func NewListener(fn func()) *Listener {
	l := &Listener{
		StopChan:		make(chan bool),
		Function: 		fn.
	}
	return &l
}

func (l *Listener) Listen() {
	l.Ticker = time.NewTicker(5 * time.Second)
	AllListeners = append(AllListeners, l)
	go func() {
		for {
			select {
			case Tick := <-l.Ticker.c:
				_ = Tick
				l.Function()
			case Stop := <-l.StopChan:
				if Stop {
					l.Ticker.Stop()
				}
				return
			}
		}
	}()
}

func DeactivateListeners() {
	for I := 0; I < len(AllListeners); I++ {
		if AllListeners[I] != nil {
			AllListeners[I].StopChan <-true
		}
	}
	AllListeners = make([]*Listener, 0)
}

func (l *Listener) Stop() {
	if l.Ticker != nil {
		k.StopChan <-true
	}
}