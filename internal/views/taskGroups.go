package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var TitleTaskGroups = "taskgroups"

type TaskGroups struct {
	*widgets.Table
	Title 		string
	Headers		[]string
	Data		[]models.TaskGroups
}

func NewTaskGroups() *TaskGroups {
	tg := &TaskGroups{
		Table:		widgets.NewTable(TitleTaskGroups),
		Title:		TitleTaskGroups,
		Headers:	[]string{"name", "running", "completed", "failed", "lost", "queued", "starting", "unknown"},
	}
	tg.Table.Headers = tg.Headers
	tg.Table.DrawHeader()
	return tg
}

func (tg *TaskGroups) GetTitle() string {
	return tg.Title
}

func (tg *TaskGroups) SetOnSelectFn(fn func(int, int)) {
	tg.Table.SetOnSelectFn(fn)
}

func (tg *TaskGroups) SetOnTabPressFn(fn func()) {
	tg.Table.SetOnTabPressFn(fn)
}

func (tg *TaskGroups) UpdateTableData(jobId, region, namespace string, data []models.TaskGroups) {
	tg.Data = data
	tg.SetTableTitle(len(tg.Data), jobId, "")
	tg.Table.ClearTable()

	for I := 0; I < len(tg.Data); I++ {
		RowTextColor := tcell.ColorWhite
		tg.Table.DrawCell(I + 1, 0,  tg.Data[I].Name, RowTextColor)
		tg.Table.DrawCell(I + 1, 1,  utils.IntToStr(tg.Data[I].Running), RowTextColor)
		tg.Table.DrawCell(I + 1, 2,  utils.IntToStr(tg.Data[I].Complete), RowTextColor)
		tg.Table.DrawCell(I + 1, 3,  utils.IntToStr(tg.Data[I].Failed), RowTextColor)
		tg.Table.DrawCell(I + 1, 4,  utils.IntToStr(tg.Data[I].Lost), RowTextColor)
		tg.Table.DrawCell(I + 1, 5,  utils.IntToStr(tg.Data[I].Queued), RowTextColor)
		tg.Table.DrawCell(I + 1, 6,  utils.IntToStr(tg.Data[I].Starting), RowTextColor)
		tg.Table.DrawCell(I + 1, 7,  utils.IntToStr(tg.Data[I].Unknown), RowTextColor)
	}
}