package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
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
		tv.SetTitle(fmt.Sprintf(" [::b][orange]%s (%s/%s) ", strings.ToUpper(tv.Title), a, b))
	} else {
		tv.SetTitle(fmt.Sprintf(" [::b][orage]%s ", strings.ToUpper(tv.Title)))
	}
}

func (tv *TextView) SetTextX(text string) {
	tv.SetText(text)
}

func (tv *TextView) ClearX() {
	tv.Clear()
}