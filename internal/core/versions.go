package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/widgets"
)

type Versions struct {
	*views.Versions
	App				*App
	Listener		*utils.Listener
	JobId			string
	SelectedValue	map[string]string
	Diffs 			[]models.JobVersionDiff
}

func NewVersions(app *App) *Versions {
	v := &Versions{
		Versions:			views.NewVersions(),
		App:				app
	}
	v.App.Layout.Body.AddPageX(v.GetTitle(), v, true, false)
	v.SetOnSelectFn(v.OnRowSelected)
	v.Versions.SetFocusFunc(v.OnFocus)
	v.Versions.SetBlurFunc(v.OnBlur)
	return v
}

func (v *Versions) OnFocus() {
	v.App.Layout.Header.Menu.Add(widgets.RevertMenu, true)
	v.UpdateTable()
}

func (v *Versions) OnBlur() {
	v.App.Layout.Header.Menu.Remove(widgets.RevertMenu)
}

func (v *Versions) OnRowSelected(row, col int) {
	v.SelectedValue = v.GetSelectedItem()
	go func() {
		v.App.Layout.QueueUpdateDraw(func() {
			if v.SelectedValue["changes"] != "No Changes" {
				v.App.Primitives.VersionDiff.Render(v.JobId, v.SelectedValue, v.Diff[row - 1])
				v.App.Layout.OpenPage("versionDiff", true)
			}
		})
	}()
}

func (v *Versios) UpdateTable() {
	SelectedJob := v.App.Primitives.Jobs.SelectedValue
	if SelectedJob != nil {
		JobID := SelectedJob["name"]
		if JobID != "" {
			v.JobId = JobID
			Params := &models.NomadParams{
				Regions:		v.App.Config.GetRegion(),
				Namespace:		v.App.Config.GetNamespace(),
			}
			JobVersions, Diff, _ := v.App.NomadClient.Versions(v.JobId, Params)
			v.Diffs = Diffs
			v.UpdateTableData(v.JobId, JobVersions, Diffs)
		}
	}
}