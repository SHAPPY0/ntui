package core

import (
	"github.com/shappy0/ntui/internal/widgets"
)

type Modal struct {
	*widgets.Modal
	App			*App
	Title 		string
	Buttons		[]string
	Data 		map[string]string
	ResponseFn  func(int, string)
}

func NewModal(app *App) *Modal {
	m := &Modal{
		Modal:		widgets.NewModal(),
		Buttons:	[]string{},
		App:		app,
	}
	m.App.Layout.Body.AddPageX("modal", m, true, false)
	return m
}

func (m *Modal) SetTitle(title string) {
	m.Title = title
	m.SetModalTitle(m.Title)
}

func (m *Modal) AddButtons(buttons []string) {
	m.Buttons = buttons
	m.ClearButtons()
	m.SetButtons(m.Buttons)
}

func (m *Modal) SetData(data map[string]string) {
	m.Data = data
}

func (m *Modal) SetResponseFunc(fn func(int, string)) {
	m.ResponseFn = fn
	m.SetDoneFn(m.ResponseFn)
}