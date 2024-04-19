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
	j.SetOnSelectionChanged(j.OnSelectionChanged)
	j.SetFocusFunc(j.OnFocus)
	j.SetBlurFunc(j.OnBlur)
	return j
}

func (j *Jobs) OnFocus() {
	j.App.Layout.Header.Menu.RenderMenu(j.Menus)
	go j.Listener.Listen()
}

func (j *Jobs) OnBlur() {
	j.App.Layout.Header.Menu.RemoveMenus(j.Menus)
	go j.Listener.Stop()
}

func (j *Jobs) OnRowSelected(row, col int) {
	j.SelectedValue = j.GetSelectedItem()
	j.App.Alert.Loader(true)
	go func() {
		j.App.Layout.QueueUpdateDraw(func() {
			JobId := j.SelectedValue["name"]
			j.App.Primitives.TaskGroups.UpdateTable(JobId)
			j.App.Alert.Loader(false)
			j.App.Layout.OpenPage("taskgroups", true)
		})
	}()
}

func (j *Jobs) OnSelectionChanged(row, col int) {
	selectedRow := j.GetSelectedItem()
	if selectedRow["status"] == "Dead" {
		j.App.Layout.Header.Menu.Replace(widgets.StopJobMenu, widgets.StartJobMenu)
	} else {
		j.App.Layout.Header.Menu.Replace(widgets.StartJobMenu, widgets.StopJobMenu)
	}
}

func (j *Jobs) UpdateMenu() {
	j.App.Layout.Header.Menu.Add(widgets.RegionNMenu, true)
}

func (j *Jobs) UpdateTable()  {
	if j.App.Config.GetRegion() != "" && j.App.Config.GetNamespace() != "" {
		Params := &models.NomadParams{
			Region:	j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		// j.UpdateMenu()
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

func (j *Jobs) StopModal() {
	j.SelectedValue = j.GetSelectedItem()
	confirmModal := j.App.Primitives.Modal
	title := fmt.Sprintf("Are you sure to stop job %s?", j.SelectedValue["name"])
	confirmModal.SetTitle(title)
	confirmModal.SetData(j.SelectedValue)
	confirmModal.AddButtons([]string{"Stop", "Cancel"})
	confirmModal.SetResponseFunc(j.HandleStopModalResponse)
	j.App.Layout.OpenPage("modal", true)
}

func (j *Jobs) HandleStopModalResponse(index int, label string) {
	if index == 0 && label == "Stop" {
		params := &models.NomadParams{
			Region:		j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		j.App.Alert.Loader(true)
		if err := j.App.NomadClient.Deregister(j.SelectedValue["name"], false, params); err != nil {
			j.App.Alert.Loader(false)
			j.App.Alert.Error(err.Error())
		} else {
			j.App.Alert.Loader(false)
		    msg := fmt.Sprintf("Job %s stopped successfully...", j.SelectedValue["name"])
		    j.App.Alert.Info(msg)
		}
	}
	j.App.Layout.GoBack()
}

func (j *Jobs) StartModal() {
	j.SelectedValue = j.GetSelectedItem()
	confirmModal := j.App.Primitives.Modal
	title := fmt.Sprintf("Are you sure to start job %s?", j.SelectedValue["name"])
	confirmModal.SetTitle(title)
	confirmModal.SetData(j.SelectedValue)
	confirmModal.AddButtons([]string{"Start", "Cancel"})
	confirmModal.SetResponseFunc(j.HandleStartpModalResponse)
	j.App.Layout.OpenPage("modal", true)
}

func (j *Jobs) HandleStartpModalResponse(index int, label string) {
	if index == 0 && label == "Start" {
		params := &models.NomadParams{
			Region:		j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		j.App.Alert.Loader(true)
		if err := j.App.NomadClient.Register(j.SelectedValue["name"], params); err != nil {
			j.App.Alert.Loader(false)
			j.App.Alert.Error(err.Error())
		} else {
			j.App.Alert.Loader(false)
		    msg := fmt.Sprintf("Job %s started successfully...", j.SelectedValue["name"])
		    j.App.Alert.Info(msg)
		}
	}
	j.App.Layout.GoBack()
}