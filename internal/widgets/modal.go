package widgets

import (
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/utils"
)

type Modal struct {
	*tview.Modal
	Title 		string
	Buttons 	[]string
}

func NewModal() *Modal {
	m := &Modal {
		Modal:	tview.NewModal(),
	}
	m.SetBackgroundColor(utils.ColorOrange)
	m.SetButtonBackgroundColor(utils.ColorOrange)
	m.SetBorder(false)
	return m
}

func (m *Modal) SetModalTitle(title string) {
	m.SetText(title)
}

func (m *Modal) SetButtons(buttons []string) {
	m.AddButtons(buttons)
}

func (m *Modal) SetDoneFn(fn func(int, string)) {
	m.SetDoneFn(fn)
}