package widgets

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

type MapView struct {
	*tview.Table
	Keys		[]string
	Values		[]string
	Size 		int
}


func NewMapView() *MapView {
	mv := &MapView{
		Table:	tview.NewTable(),
	}
	mv.SetBorder(false)
	return mv
}

func (mv *MapView) SetMapKeys(keys []string) {
	mv.Keys = keys
	mv.Size += 1 
}

func (mv *MapView) SetMapValues(values []string) {
	mv.Values = values
}

func (mv *MapView) SetMapKeyValue(key, value string) {
	mv.Keys = append(mv.Keys, key)
	mv.Values = append(mv.Values, value)
	mv.Size += 1
}

func (mv *MapView) DrawMapView() {
	for I := 0; I < len(mv.Keys); I++ {
		mv.Table.SetCell(I, 0, mv.KeyCell(mv.Keys[I]))
		mv.Table.SetCell(I, 1, mv.ValueCell(mv.Values[I]))
	}
}

func (mv *MapView) KeyCell(key string) *tview.TableCell {
	KeyCell := tview.NewTableCell(key)
	KeyCell.SetAlign(tview.AlignLeft)
	KeyCell.SetTextColor(tcell.ColorOrange)
	return KeyCell
}

func (mv *MapView) ValueCell(value string) *tview.TableCell {
	ValueCell := tview.NewTableCell(value)
	ValueCell.SetExpansion(0)
	ValueCell.SetAlign(tview.AlignLeft)
	return ValueCell
}

func (mv *MapView) Clear() {
	mv.Table.Clear()
	mv.Keys = make([]string, 0)
	mv.Values = make([]string, 0)
	mv.Size = 0
}