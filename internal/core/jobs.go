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
	j.Listener = utils.NewListener(app.Config.RefreshRate, j.Refresher)
	j.SetOnSelectFn(j.OnRowSelected)
	j.SetOnSelectionChanged(j.OnSelectionChanged)
	j.SetFocusFunc(j.OnFocus)
	j.SetBlurFunc(j.OnBlur)
	return j
}

func (j *Jobs) OnFocus() {
	j.App.Layout.Header.Menu.RenderMenu(j.Menus, true)
	go j.Listener.Listen()
}

func (j *Jobs) OnBlur() {
	j.App.Layout.Header.Menu.RemoveMenus(j.Menus)
	j.App.Layout.Header.Menu.Remove(widgets.StartJobMenu)
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
		j.OnSelectionChanged(1, 0)
	} else {
		msg := "No valid Region/Namespace found..."
		j.App.Logger.Error(msg)
		j.App.Alert.Error(msg)
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
	confirmModal.AddButtons([]string{"Cancel", "Stop"})
	confirmModal.SetResponseFunc(j.HandleStopModalResponse)
	j.App.Layout.OpenPage("modal", true)
}

func (j *Jobs) HandleStopModalResponse(index int, label string) {
	if index == 1 && label == "Stop" {
		params := &models.NomadParams{
			Region:		j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		j.App.Alert.Loader(true)
		if err := j.App.NomadClient.Deregister(j.SelectedValue["name"], false, params); err != nil {
			j.App.Alert.Loader(false)
			j.App.Alert.Error("Job stop request failed...")
			j.App.Logger.Errorf("Job stop request failed: %s", err.Error())
		} else {
			j.App.Alert.Loader(false)
		    msg := fmt.Sprintf("Job %s stopped successfully...", j.SelectedValue["name"])
		    j.App.Alert.Info(msg)
			j.App.Logger.Info(msg)
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
	confirmModal.AddButtons([]string{"Cancel", "Start"})
	confirmModal.SetResponseFunc(j.HandleStartModalResponse)
	j.App.Layout.OpenPage("modal", true)
}

func (j *Jobs) HandleStartModalResponse(index int, label string) {
	if index == 1 && label == "Start" {
		params := &models.NomadParams{
			Region:		j.App.Config.GetRegion(),
			Namespace:	j.App.Config.GetNamespace(),
		}
		j.App.Alert.Loader(true)
		if err := j.App.NomadClient.Register(j.SelectedValue["name"], params); err != nil {
			j.App.Alert.Loader(false)
			j.App.Alert.Error("Job start request failed...")
			j.App.Logger.Errorf("Job start request failed:: %s", err.Error())
		} else {
			j.App.Alert.Loader(false)
		    msg := fmt.Sprintf("Job %s started successfully...", j.SelectedValue["name"])
		    j.App.Alert.Info(msg)
			j.App.Logger.Info(msg)
		}
	}
	j.App.Layout.GoBack()
}

func (j *Jobs) GoToDefinitions() {
	j.SelectedValue = j.GetSelectedItem()
	j.App.Layout.OpenPage(j.App.Primitives.JobDefinition.GetTitle(), true)
}