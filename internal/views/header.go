package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/widgets"
)

type Header struct {
	*tview.Flex
	Logo			*tview.TextView
	Menu 			*widgets.Menu
	Metadata 		*widgets.TextView
}

var Logo = []string{
	` _  _  ____  __  __  ____`,
	`( \( )(_  _)(  )(  )(_  _)`,
	` )  (   )(   )(__)(  _)(_`,
	`(_)\_) (__) (______)(____)`,
}

func NewHeader() *Header {
	h := &Header{
		Felx:		tview.NewFlex(),
		Logo:		tview.NewTextView(),
		Menu:		widgets.NewMenu(),
		Metadata:	tview.NewTextView(),
	}
	h.RenderLogo()
	h.RenderMenu()
	h.AddItem(h.Logo, 0, 1, false)
	h.AddItem(h.Menu.MapView, 0, 1, false)
	h.AddItem(h.Metadata, 0, 1, false)
	return h
}

func (h *Header) RenderLogo() error {
	h.Logo.SetDynamicColors(true)
	for I, S := range Logo {
		fmt.Fprintf(h.Logo, "[%s::b]%s", "", s)
		if I + 1 < len(Logo) {
			fmt.Fprintf(h.Logo, "\n")
		}
	}
	return nil
}

func (h *Header) RenderMenu() error {
	h.Menu.Add(widgets.EnterMenu, false).
		Add(widgets.UpArrowMenu, false).
		Add(widgets.DownArrowMenu, false).
		Add(widgets.EscMenu, false)
	h.Menu.Render()
	return nil
}