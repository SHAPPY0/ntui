package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/hashicorp/nomad/api"
)

var TitleVersions = "versions"

type Versoins struct {
	*widgets.Table
	Title		string
	Headers		[]string
	Data		[]models.JobVersion
}

func NewVersions() *Versions {
	v := &Versions{
		Table:			widgets.NewTable(TitleVersions),
		Title:			TitleVersions,
		Headers:		[]string{"version", "stable", "submitted", "changes"},
	}
	v.Table.Headers = v.Headers
	v.Table.DrawHeader()
	return v
}

func (v *Versions) GetTitle() string {
	return v.Title
}

func (v *Versions) SetOnSelectFn(fn func(int, int)) {
	v.Table.SetOnSelectFn(fn)
}

func (v *Versions) SetOnTabPressFn(fn func()) {
	v.Table.SetOnTabPressFn(fn)
}

func FindChangesCount(taskGroupDiff []*api.TaskGroupDiff) int {
	Count := 0
	for I := 0; I < len(taskGroupDiff); I++ {
		if taskGroupDiff[I].Type == "Edited" {
			Count++
		}
	}
	return Count
}

func (v *Versions) UpdateTableData(jobId string, jobVersions []models.JobVersion, diffs []models.JobVersionDiff) {
	v.Data = jobVersions
	v.SetTableTitle(len(v.Data), jobId, "")
	v.Table.ClearTable()
	for I := 0; I <len(v.Data); I++ {
		RowTextColor := tcell.ColorWhite
		v.Table.DrawCell(I + 1, 0, "#" + utils.IntToStr(int(v.Data[I].Version)), RowTextColor)
		Stable := "No"
		if v.Data[I].Stable {
			Stable = "Yes"
		}
		v.Table.DrawCell(I + 1, 1, Stable, RowTextColor)
		v.Table.DrawCell(I + 1, 2, utils.DateTimeToStr(v.Data[I].SubmitTime), RowTextColor)
		ChangeCount := 0
		if len(diffs) > I {
			ChangeCount = FindChangesCount(diffs[I].TaskGroups)
		}
		Changes := "No Changes"
		if ChangeCount > 0 {
			Changes = utils.IntToStr(ChangeCount) + " Changes"
		}
		v.Table.DrawCell(I + 1, 3, Changes, RowTextColor)
	}
}