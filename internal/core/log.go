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
	l.App.Layout.Body.AddPageX(l.GetTitle(), l.LogView, true, false)
	l.LogView.SetFocusFunc(l.OnFocus)
	l.LogView.SetBlurFunc(l.OnBlur)
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
	l.App.Layout.Header.Menu.RemoveMenus([]widgets.Item{
		widgets.EnterMenu, 
		widgets.UpArrowMenu, 
		widgets.DownArrowMenu,
	})
	if l.DefaultType == LOG_STDERR {
		l.App.Layout.Header.Menu.Remove(l.Menus[1])
		l.App.Layout.Header.Menu.Add(l.Menus[0], true)
	} else {
		l.App.Layout.Header.Menu.Remove(l.Menus[0])
		l.App.Layout.Header.Menu.Add(l.Menus[1], true)
	}
	for _, m := range l.Menus[2:] {
		l.App.Layout.Header.Menu.Add(m, true)
	}
	
}

func (l *Log) OnFocus() {
	l.SelectedAlloc = l.App.Primitives.Tasks.SelectedValue
	l.UpdateMenu()
	allocName := l.SelectedAlloc.Name
	allocID := utils.GetID(l.SelectedAlloc.ID)
	l.LogView.SetTitleName(l.GetTitle() + "-" + l.DefaultType)
	l.LogView.SetTextVTitle(allocID, allocName)
	l.StopLogChan = make(chan struct{})
	l.FetchLog()
	
}

func (l *Log) FetchLog() {
	l.App.Alert.Loader(true)
	l.ClearLogs()
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
		case _ = <-l.StopLogChan:
			// l.ClearLogs()
			return
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
	l.FollowX()
}

func (l *Log) OnBlur() {
	l.StopLogStream()
	if l.DefaultType == LOG_STDOUT {
		l.App.Layout.Header.Menu.Remove(l.Menus[1])
	} else {
		l.App.Layout.Header.Menu.Remove(l.Menus[0])
	}
	l.App.Layout.Header.Menu.RemoveMenus(l.Menus[2:])
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

func (l *Log) ShowFullScreen(fs bool) {
	l.Container.FullScreen(fs)
}