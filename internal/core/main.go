package core

import (
	"github.com/rivo/tview"
)

type Main struct {
	*tview.Flex
	App			*App
}

func NewMain(app *App) *Main {
	m := &Main{
		App:		app,
		Flex:		tview.NewFlex(),
	}
	m.App.Layout.Body.AddPageX("main", m, true, false)
	m.SetDirection(tview.FlexRow)
	m.AddItem(m.App.Primitives.Regions, 0,  1, true)
	m.AddItem(m.App.Primitives.Namespace, 0, 1, false)
	return m
}