package views

import (
	"github.com/shappy0/ntui/internal/widgets"
)

type Log struct {
	*widgets.TextView
	Title 		string
	Log 		[]byte
}

var TitleLog = "log"

func NewLog() *Log {
	l := &Log{
		TextView:		widgets.NewTextView(TitleLog),
		Title:			TitleLog,
	}
	return l
}

func (l *Log) FollowX() {
	l.ScrollToEnd()
}

func (l *Log) GetTitle() string {
	return l.Title
}

func (l *Log) Render(log []byte) {
	l.Log = append(l.Log, log...)
	l.SetText(string(l.Log))
}

func (l *Log) ClearLogs() {
	l.Log = make([]byte, 0)
}