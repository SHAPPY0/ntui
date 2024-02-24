package views

import (
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var TitleAllocations = "allocations"

type Allocations struct {
	*widgets.Table
	Title 			string
	Headers 		[]string
	Data			[]models.Allocations
}

func NewAllocations() *Allocations {
	a := &Allocations{
		Table:		widgets.NewTable(TitleAllocations),
		Title:		TitleAllocations,
		Headers:	[]string{"id", "name", "status", "version", "created", "modified", "client", "volumn", "cpu", "memory"},
	}
	a.Table.Headers = a.Headers
	a.Table.DrawHeader()
	return a
}

func (a *Allocations) GetTitle() string {
	return a.Title
}

func (a *Allocations) SetOnSelectFn(fn func(int, int)) {
	a.Table.SetOnSelectFn(fn)
}

func (a *Allocations) SetOnTabPressFn(fn func()) {
	a.Table.SetOnTabPressFn(fn)
}

func (a *Allocations) UpdateTableData(data []models.Allocations) {
	a.Data = data
	TgName := ""
	if len(data) > 0{
		TgName = data[0].TaskName
	}
	a.SetTableTitle(len(a.Data), TgName, "")
	a.Table.ClearTable()
	for I := 0; I < len(a.Data); I++ {
		CellColor, Status := utils.ColorizeStatusCell(a.Data[I].Status)
		a.Table.DrawCell(I + 1, 0, utils.GetID(a.Data[I].ID), CellColor)
		a.Table.DrawCell(I + 1, 1, a.Data[I].Name, CellColor)
		a.Table.DrawCell(I + 1, 2, Status, CellColor)
		a.Table.DrawCell(I + 1, 3, utils.IntToStr(a.Data[I].Version), CellColor)
		a.Table.DrawCell(I + 1, 4, utils.DateTimeToStr(a.Data[I].Created), CellColor)
		a.Table.DrawCell(I + 1, 5, utils.DateTimeToStr(a.Data[I].Modified), CellColor)
		a.Table.DrawCell(I + 1, 6, utils.GetID(a.Data[I].Client), CellColor)
		a.Table.DrawCell(I + 1, 7, a.Data[I].Volumn, CellColor)
		CpuStat := utils.IntToStr(a.Data[I].CpuUsage) + "MHz/" + utils.IntToStr(a.Data[I].Cpu) + "MHz"
		if a.Data[I].CpuUsage >= a.Data[I].Cpu {
			CpuStat = utils.SetCellTextColor(CpuStat, "red")
		}
		a.Table.DrawCell(I + 1, 8, CpuStat, CellColor)
		MemoryStat := utils.IntToStr(utils.FormatMemoryUsage(a.Data[I].MemoryUsage)) + "MiB/" + utils.IntToStr(a.Data[I].Memory) + "MiB"
		if utils.FormatMemoryUsage(a.Data[I].MemoryUsage) >= a.Data[I].Memory {
			MemoryStat = utils.SetCellTextColor(MemoryStat, "red")
		}
		a.Table.DrawCell(I + 1, 9, MemoryStat, CellColor)
 	}
}