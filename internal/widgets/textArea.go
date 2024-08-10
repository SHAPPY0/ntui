package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	// "github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/utils"
)

type TextArea struct {
	*tview.TextArea
	Title 	string
}

func NewTextArea(title string) *TextArea {
	ta := &TextArea{
		TextArea:	tview.NewTextArea(),
		Title:		title,
	}
	ta.SetTitleX(title)
	ta.WrapLines(true)
	ta.SetBorderX(true)
	return ta
}

func (ta *TextArea) SetTitleX(title string) {
	ta.SetTitle(fmt.Sprintf(" [::b][%s]%s ", utils.ColorT70d5bf, strings.ToUpper(title)))
}

func (ta *TextArea) SetPlaceholderX(ph string) {
	ta.SetPlaceholder(ph)
}

func (ta *TextArea) WrapLines(wrap bool) {
	ta.SetWrap(wrap)
}

func (ta *TextArea) SetBorderX(border bool) {
	ta.SetBorder(border)
}