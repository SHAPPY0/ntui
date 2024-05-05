package core

import (
	"github.com/rivo/tview"
)

var titleMain = "main"

type Main struct {
	*tview.Flex
	App			*App
	Title		string
}

func NewMain(app *App) *Main {
	m := &Main{
		App:		app,
		Flex:		tview.NewFlex(),
		Title:		titleMain,
	}
	m.App.Layout.Body.AddPageX(m.Title, m, true, false)
	m.SetDirection(tview.FlexRow)
	m.AddItem(m.App.Primitives.Regions, 0,  1, true)
	m.AddItem(m.App.Primitives.Namespaces, 0, 1, false)
	return m
}

func (m *Main) GetTitle() string {
	return m.Title
}