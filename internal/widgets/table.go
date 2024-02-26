package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/utils"
)

type Table struct {
	*tview.Table
	Title 		string
	Data 		map[string]string
	Headers 	[]string
	OnSelectFn  func(int, int)
	OnTabPress  func()
}

func NewTable(title string) *Table {
	t := &Table{
		Table:		tview.NewTable(),
		Title:		title,
	}
	t.SetTableTitle(0, "", "")
	t.SetBorder(true)
	t.SetFixed(1, 1)
	t.SetSelectable(true, false)
	t.BindKeys()
	t.BindFunctions()
	return t
}

func (t *Table) SetOnSelectFn(fn func(int, int)) {
	t.OnSelectFn = fn
	t.SetSelectedFunc(fn)
}

func (t *Table) SetOnTabPressFn(fn func()) {
	t.OnTabPress = fn
}

func (t *Table) GetSelectedItem() map[string]string {
	Selected := make(map[string]string, 0)
	if t.GetRowCount() < 1 {
		return Selected
	}
	Row, _ := t.GetSelection()
	for Index, Name := range t.Headers {
		Value := t.GetCell(Row, Index).Text
		Selected[Name] = Value
	}
	return Selected
}

func (t *Table) SetTableTitle(count int, a, b string) {
	if a != "" && b == "" {
		t.SetTitle(fmt.Sprintf(" [::b][%s]%s(%s)[%d] ", utils.ColorTad7c5a, strings.ToUpper(t.Title), a, count))
	} else if a != "" && b != "" {
		t.SetTitle(fmt.Sprintf(" [::b][%s]%s(%s/%s)[%d] ", utils.ColorTad7c5a, strings.ToUpper(t.Title), a, b, count))
	} else {
		t.SetTitle(fmt.Sprintf(" [::b][%s]%s [%d] ", utils.ColorT70d5bf, strings.ToUpper(t.Title), count))
	}
}

func (t *Table) DrawHeader() {
	for I := 0; I < len(t.Headers); I++ {
		Header := fmt.Sprintf("[::b]%s", strings.ToUpper(t.Headers[I]))
		t.SetCell(0, I, 
			tview.NewTableCell(Header).
				SetExpansion(1).
				SetBackgroundColor(tcell.ColorGray).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
	}
}

func (t *Table) DrawCell(row, col int, value string, rowColor tcell.Color) {
	NewCell := tview.NewTableCell(value).
				SetExpansion(1).
				SetAlign(tview.AlignCenter)
	NewCell.SetTextColor(rowColor)
	t.SetCell(row, col, NewCell)
}

func (t *Table) DrawLeftCell(row, col int, value string, rowColor tcell.Color) {
	NewCell := tview.NewTableCell(value).
				SetExpansion(1).
				SetAlign(tview.AlignLeft)
	NewCell.SetTextColor(rowColor)
	t.SetCell(row, col, NewCell)
}

func (t *Table) BindKeys() {
	t.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch Key := event.Key(); Key {
		case utils.NtuiTabKey.Key:
			if t.OnTabPress != nil {
				t.OnTabPress()
			}
		}
		return event
	})
}

func (t *Table) BindFunctions() {
	t.SetFocusFunc(func() {
		t.SetFocusBorderColor(tcell.ColorOrange)
	})
	t.SetBlurFunc(func() {
		t.SetFocusBorderColor(tcell.ColorWhite)
	})
}

func (t *Table) SetFocusBorderColor(color tcell.Color) {
	if t.GetBorderColor() != color {
		t.SetBorderColor(color)
	}
}

func (t *Table) ClearTable() {
	t.Clear()
	t.DrawHeader()
}