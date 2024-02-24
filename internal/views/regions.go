package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

var TitleRegions = "regions"

type Regions struct {
	*widgets.Table
	Headers		[]string
	Data		[]models.Regions
}

func NewRegions() *Regions {
	reg := &Regions{
		Table:	widgets.NewTable(TitleRegions),
		Headers:	[]string{"id", "region"},
	}
	reg.Table.Headers = reg.Headers
	reg.Table.DrawHeader()
	return reg
}

func (reg *Regions) SetOnSelectFn(fn func(int, int)) {
	reg.Table.SetOnSelectFn(fn)
}

func (reg *Regions) SetOnTabPressFn(fn func()) {
	reg.Table.SetOnTabPressFn(fn)
}

func (reg *Regions) UpdateTableData(data []models.Regions) {
	reg.Data = data
	reg.SetTableTitle(len(reg.Data), "", "")
	RowTextColor := tcell.ColorWhite
	for I := 0; I < len(reg.Data); I++ {
		reg.Table.DrawCell(I + 1, 0, reg.Data[I].Id, RowTextColor)
		reg.Table.DrawCell(I + 1, 1, reg.Data[I].Name, RowTextColor)
	}
}