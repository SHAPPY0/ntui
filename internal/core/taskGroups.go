package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/widgets"
)

type TaskGroups struct {
	*views.TaskGroups
	App				*App
	Listener		*utils.Listener
	JobId 			string
	SelectedValue	map[string]string
}

func NewTaskGroups(app *App) *TaskGroups {
	tg := &TaskGroups{
		TaskGroups:		views.NewTaskGroups(),
		App:			app,
	}
	tg.App.Layout.Body.AddPageX(tg.GetTitle(), tg, true, false)
	tg.Listener = utils.NewListener(tg.Refresher)
	tg.SetOnSelectFn(tg.OnRowSelected)
	tg.TaskGroups.SetBlurFunc(tg.OnBlur)
	tg.TaskGroups.SetFocusFunc(tg.OnFocus)
	return tg
}

func (tg *TaskGroups) OnFocus() {
	tg.App.Layout.Header.Menu.Add(widgets.VersionMenu,  true)
	go tg.Listener.Listen()
}

func (tg *TaskGroups) OnBlur() {
	tg.App.Layout.Header.Menu.Remove(widgets.VersionMenu)
	go tg.Listener.Stop()
}

func (tg *TaskGroups) OnRowSelected(row, col int) {
	tg.SelectedValue = tg.GetSelectedItem()
	go func() {
		tg.App.Layout.QueueUpdateDraw(func() {
			tg.SelectedValue["jobId"] = tg.JobId
			tg.App.Primitives.Allocations.Render(tg.SelectedValue)
			tg.Aapp.Layout.OpenPage("allocations", true)
		})
	}()
}

func (tg *TaskGroups) UpdateTable(jobId string) {
	Region := tg.App.Config.GetRegion()
	Namespace := tg.App.Config.GetNamespace()
	tg.JobId = jobId
	Data, _ := tg.App.NomadClient.TaskGroups(jobId, Region, Namespace)
	tg.UpdateTableData(jobId, Regions, Namespace, Data)
}

func (tg *TaskGroups) Refresher() {
	tg.UpdateTable(tg.JobId)
	tg.App.Layout.Draw()
}

