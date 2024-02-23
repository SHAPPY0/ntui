package core

import (
	"github.com/shappy0/ntui/internal/views"
)

type Splash struct {
	*views.Splash
	App			*App
}

func NewSplash(app *App) *Splash {
	s := &Splash{
		Splash:		views.NewSplash(),
		App:		app,
	}
	s.App.Layout.Body.AddPageX(s.GetTitle(), s, true, false)
	return s
}