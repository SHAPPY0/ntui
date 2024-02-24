package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

var TitleNamespace = "namespaces"

type Namespaces struct {
	*widgets.Table
	Headers		[]string
	Data		[]models.Namespaces
}

func NewNamespaces() *Namespaces {
	n := &Namespaces {
		Table:		widgets.NewTable(TitleNamespace),
		Headers:	[]string{"id", "name", "description"},
	}
	n.Table.Headers = n.Headers
	n.Table.DrawHeader()
	return n
}

func (n *Namespaces) SetOnSelectFn(fn func(int, int)) {
	n.Table.SetOnSelectFn(fn)
}

func (n *Namespaces) SetOnTabPress(fn func()) {
	n.Table.SetOnTabPressFn(fn)
}

func (n *Namespaces) UpdateTableData(data []models.Namespaces, region string) {
	n.Data = data
	n.SetTableTitle(len(n.Data), region, "")
	RowTextColor := tcell.ColorWhite
	for I := 0; I < len(n.Data); I++ {
		n.Table.DrawCell(I + 1, 0, n.Data[I].Id, RowTextColor)
		n.Table.DrawCell(I + 1, 1, n.Data[I].Name, RowTextColor)
		n.Table.DrawCell(I + 1, 2, n.Data[I].Description, RowTextColor)
	}
}