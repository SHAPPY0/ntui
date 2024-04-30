package views

import (
	"github.com/shappy0/ntui/internal/widgets"
)

type Log struct {
	Container	*widgets.Flex
	LogView   	*widgets.TextView
	Title 		string
	Log 		[]byte
	Menus 		[]widgets.Item
}

var TitleLog = "log"

var LogMenu = []widgets.Item{
	widgets.StdoutMenu,
	widgets.StderrMenu,
	widgets.LogAutoScrollMenu,
}

func NewLog() *Log {
	l := &Log{
		Container:		widgets.NewFlex(),
		LogView:		widgets.NewTextView(TitleLog),
		Title:			TitleLog,
		Menus:			LogMenu,
	}
	l.Container.AddItemX(l.LogView, 0, 1, true)
	return l
}

func (l *Log) FollowX() {
	l.LogView.ScrollToEnd()
}

func (l *Log) GetTitle() string {
	return l.Title
}

func (l *Log) Render(log []byte) {
	l.Log = append(l.Log, log...)
	l.LogView.SetText(string(l.Log))
}

func (l *Log) ClearLogs() {
	l.Log = make([]byte, 0)
	l.Container.ClearFlex()
}