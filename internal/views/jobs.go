package views

import (
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var TitleJobs = "jobs"

type Jobs struct {
	*widgets.Table
	Headers 		[]string
	Data 			[]models.Jobs
}

func NewJobs() *Jobs {
	j := &Jobs{
		Table: 		widgets.NewTable(TitleJobs),
		Headers:	[]string{"name", "status", "type", "groups", "priority"},
	}

	j.Table.Headers = j.Headers
	j.Table.DrawHeader()
	return j
}

func (j *Jobs) SetOnSelectFn(fn func(int, int)) {
	j.Table.SetOnSelectFn(fn)
}

func (j *Jobs) SetOnTabPress(fn func()) {
	j.Table.SetOnTabPress(fn)
}

func (j *Jobs) UpdateTableData(params *models.NomadParams, data []models.Jobs) {
	j.Data = data
	j.SetTableTitle(len(j.Data), params.Region, params.Namespace)
	for I := 0; I < len(j.Data); I++ {
		CellColor, Status := utils.ColorizeStatusCell(j.Data[I].Status)
		j.Table.DrawCell(I + 1, 0, j.Data[I].Name, CellColor)
		j.Table.DrawCell(I + 1, 1, Status, CellColor)
		j.Table.DrawCell(I + 1, 2, j.Data[I].Type, CellColor)
		j.Table.DrawCell(I + 1, 3, utils.IntToStr(j.Data[I].StatusSummary.Total), CellColor)
		j.Table.DrawCell(I + 1, 4, utils.IntToStr(j.Data[I].Priority), CellColor)
	}
}