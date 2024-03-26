package core

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/widgets"
)

type Log struct {
	*views.Log
	App				*App
	SelectedAlloc	models.Allocations
	DefaultType		string
	Follow			bool
	StopLogChan		chan struct{}
}

const (
	LOG_STDOUT = "stdout"
	LOG_STDERR = "stderr"
)

func NewLog(app *App) *Log {
	l := &Log{
		Log:			views.NewLog(),
		App:			app,
		DefaultType:	LOG_STDOUT,
		Follow:			true,
	}
	l.App.Layout.Body.AddPageX(l.GetTitle(), l, true, false)
	l.SetFocusFunc(l.OnFocus)
	l.SetBlurFunc(l.OnBlur)
	return l
}

func (l *Log) BindInputCapture() {
	l.App.Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case utils.NtuiRuneKey.Key:
			switch event.Rune() {
			case 'e':
				if l.App.Layout.GetActivePage() == "log" {
					fmt.Println("E Pressed")
				}
				break
			}
		}
		return event
	})
}

func (l *Log) UpdateMenu() {
	l.App.Layout.Header.Menu.Remove(widgets.EnterMenu)
	l.App.Layout.Header.Menu.Remove(widgets.UpArrowMenu)
	l.App.Layout.Header.Menu.Remove(widgets.DownArrowMenu)
	if l.DefaultType == LOG_STDERR {
		l.App.Layout.Header.Menu.Remove(widgets.StderrMenu)
		l.App.Layout.Header.Menu.Add(widgets.StdoutMenu, true)
	} else {
		l.App.Layout.Header.Menu.Remove(widgets.StdoutMenu)
		l.App.Layout.Header.Menu.Add(widgets.StderrMenu, true)
	}
}

func (l *Log) OnFocus() {
	l.SelectedAlloc = l.App.Primitives.Tasks.SelectedValue
	l.UpdateMenu()
	allocName := l.SelectedAlloc.Name
	allocID := utils.GetID(l.SelectedAlloc.ID)
	l.SetTitleName(l.GetTitle() + "-" + l.DefaultType)
	l.SetTextVTitle(allocID, allocName)
	l.StopLogChan = make(chan struct{})
	l.FetchLog()
}

func (l *Log) FetchLog() {
	l.App.Alert.Loader(true)
	LogChan, ErrChan := l.App.NomadClient.Logs(
		l.SelectedAlloc.ID,
		l.SelectedAlloc.TaskName,
		l.DefaultType,
		"end",
		l.Follow,
		10000,
		l.StopLogChan,
	)
	if l.Follow {
		l.FollowX()
	}
	l.ClearLogs()
	l.App.Alert.Loader(false)
	go l.StartLogStream(LogChan, ErrChan)
}

func (l *Log) StartLogStream(logChan <-chan *api.StreamFrame, errChan <-chan error) {
	for {
		select {
		case Log := <-logChan:
			if Log == nil {
				return
			}
			l.Render(Log.Data)
			l.App.Layout.Draw()
		case Err := <-errChan:
			l.App.Alert.Error(Err.Error())
			return
		}
	}
}

func (l *Log) StopLogStream() {
	close(l.StopLogChan)
}

func (l *Log) SetLogType(logType string) {
	l.DefaultType = logType
}

func (l *Log) SetFollow(follow bool) {
	l.Follow = follow
}

func (l *Log) OnBlur() {
	l.StopLogStream()
	if l.DefaultType == LOG_STDOUT {
		l.App.Layout.Header.Menu.Remove(widgets.StderrMenu)
	} else {
		l.App.Layout.Header.Menu.Remove(widgets.StdoutMenu)
	}
}

func (l *Log) FetchStdOutLog() {
	l.DefaultType = LOG_STDOUT
	l.StopLogStream()
	l.OnFocus()
}

func (l *Log) FetchStdErrLog() {
	l.DefaultType = LOG_STDERR
	l.StopLogStream()
	l.OnFocus()
}