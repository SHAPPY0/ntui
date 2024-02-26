package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/utils"
)

type TextView struct {
	*tview.TextView
	Title		string
}

func NewTextView(title string) *TextView {
	tv := &TextView{
		TextView:	tview.NewTextView(),
		Title:		title,
	}
	tv.SetBorder(true)
	tv.SetScrollable(true)
	tv.SetTextVTitle("", "")
	return tv
}

func (tv *TextView) SetTitleName(name string) {
	tv.Title = name
}

func (tv *TextView) SetTextVTitle(a, b string) {
	if a != "" && b != "" {
		tv.SetTitle(fmt.Sprintf(" [::b][%s]%s (%s/%s) ", utils.ColorTad7c5a, strings.ToUpper(tv.Title), a, b))
	} else {
		tv.SetTitle(fmt.Sprintf(" [::b][%s]%s ", utils.ColorTad7c5a, strings.ToUpper(tv.Title)))
	}
}

func (tv *TextView) SetTextX(text string) {
	tv.SetText(text)
}

func (tv *TextView) ClearX() {
	tv.Clear()
}