package core

import (
	"fmt"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/widgets"
)

type Jobs struct {
	*views.Jobs
	App				*App
	Listener		*utils.Listener
	SelectedValue	map[string]string
}

func NewJobs(app *App) *Jobs {
	j := &Jobs{
		Jobs:	views.NewJobs(),
		App:	app,
	}
	j.App.Layout.Body.AddPageX(j.GetTitle(), j, true, false)
	j.Listener = utils.NewListener(j.Refresher)
	j.SetOnSelectFn(j.OnRowSelected)
	j.SetFocusFunc(j.OnFocus)
	j.SetBlurFunc(j.OnBlur)
	return j
}

func (j *Jobs) onFocus() {
	j.App.Layout.Header.Menu.Add(widgets.RegionMenu, true)
}

func (j *Jobs) OnBlur() {
	j.App.Layout.Header.Menu.Remove(widgets.RegionMenu)
}

func (j *Jobs) OnRowSelected(row, col int) {
	j.SelectedValue = j.GetSelectedItem()
	go func() {
		j.App.Layout.QueueUpdateDraw(func() {
			JobId := j.SelectedValue["name"]
			j.App.Primitives.TaskGroups.UpdateTable(JobId)
			j.App.Layout.OpenPage("taskgroups", true)
		})
	}()
}

func (j *Jobs) UpdateMenu() {
	j.App.Layout.Header.Menu.Add(widgets.RegionMenu, true)
}

func (j *Jobs) UpdateTable()  {
	if j.App.Config.GetRegion() != "" && j.App.Config.GetNamespace() != "" {
		Params := &models.NomadParams{
			Regions:	j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		j.UpdateMenu()
		Data, _ := j.App.NomadClient.Jobs(Params)
		j.UpdateTableData(Params, Data)
	} else {
		fmt.Println("No config values found")
	}
}

func (j *Jobs) Refresher() {
	j.UpdateTable()
	j.App.Layout.Draw()
}