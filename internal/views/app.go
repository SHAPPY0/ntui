package views

import (
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
}

func NewApp() *App {
	a := &App{
		Application:		tview.NewApplication(),
	}
	return a
}