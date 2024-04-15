package core

import (
	"fmt"
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
		App:				app,
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
				v.App.Primitives.VersionDiff.Render(v.JobId, v.SelectedValue, v.Diffs[row - 1])
				v.App.Layout.OpenPage("versiondiff", true)
			}
		})
	}()
}

func (v *Versions) UpdateTable() {
	SelectedJob := v.App.Primitives.Jobs.SelectedValue
	if SelectedJob != nil {
		JobID := SelectedJob["name"]
		if JobID != "" {
			v.JobId = JobID
			Params := &models.NomadParams{
				Region:		v.App.Config.GetRegion(),
				Namespace:		v.App.Config.GetNamespace(),
			}
			JobVersions, Diffs, _ := v.App.NomadClient.Versions(v.JobId, Params)
			v.Diffs = Diffs
			v.UpdateTableData(v.JobId, JobVersions, Diffs)
		}
	}
}

func (v *Versions) ConfirmModal() {
	v.SelectedValue = v.GetSelectedItem()
	if v.SelectedValue["revertable"] == "Yes" {
		confirmModal := v.App.Primitives.Modal
		title := fmt.Sprintf("Are you sure to revert version %s%s?", v.JobId, v.SelectedValue["version"])
		confirmModal.SetTitle(title)
		confirmModal.SetData(v.SelectedValue)
		confirmModal.AddButtons([]string{"Revert", "Cancel"})
		confirmModal.SetResponseFunc(v.HandleButtonResponse)
		v.App.Layout.OpenPage("modal", true)
	}
}

func (v *Versions) HandleButtonResponse(index int, label string) {
	if index == 0 && label == "Revert" {
		Params := &models.NomadParams{
			Region:		v.App.Config.GetRegion(),
			Namespace:	v.App.Config.GetNamespace(),
		}
		v.App.Alert.Loader(true)
		verNum := utils.Split(v.SelectedValue["version"], "#")
		version := utils.IntToUint64(utils.StrToInt(verNum[1]))
		if err := v.App.NomadClient.Revert(v.JobId, version, Params); err != nil {
			v.App.Alert.Loader(false)
			v.App.Alert.Error(err.Error())
		} else {
			v.App.Alert.Loader(false)
		    msg := fmt.Sprintf("Version %s %s reverted successful...", v.JobId, v.SelectedValue["version"])
		    v.App.Alert.Info(msg)
		}
	}
	v.App.Layout.GoBack()
}