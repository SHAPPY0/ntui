package core

import (
	"fmt"
	"strings"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
	// "github.com/shappy0/ntui/internal/widgets"
)

type Allocations struct {
	*views.Allocations
	App				*App
	Listener		*utils.Listener
	SelectedValue	map[string]string
	TaskGroup		map[string]string
	SelectedAlloc   models.Allocations
	Data			[]models.Allocations
}

func NewAllocations(app *App) *Allocations {
	a := &Allocations {
		Allocations:	views.NewAllocations(),
		App:			app,
	}
	a.App.Layout.Body.AddPageX(a.GetTitle(), a, true, false)
	a.SetOnSelectFn(a.OnRowSelected)
	a.Listener = utils.NewListener(app.Config.RefreshRate, a.Refresher)
	a.SetFocusFunc(a.OnFocus)
	a.SetBlurFunc(a.OnBlur)
	return a
}

func (a *Allocations) OnRowSelected(row, col int) {
	a.SetSelectedRow()
	a.App.Alert.Loader(true)
	go func() {
		a.App.Layout.QueueUpdateDraw(func() {
			a.App.Layout.OpenPage(a.App.Primitives.Tasks.GetTitle(), true)
			a.App.Alert.Loader(false)
		})
	}()
}

func (a *Allocations) Render(data map[string]string) {
	a.UpdateTable(data)
}

func (a *Allocations) UpdateTable(data map[string]string) {
	Params := &models.NomadParams{
		Region:		a.App.Config.GetRegion(),
		Namespace:	a.App.Config.GetNamespace(),
	}
	a.TaskGroup = data
	Data, _ := a.App.NomadClient.Allocations(a.TaskGroup["name"], Params)
	a.Data = Data
	a.UpdateTableData(Data)
}

func (a *Allocations) OnFocus() {
	a.App.Layout.Header.Menu.RenderMenu(a.Menus, true)
	go a.Listener.Listen()
}

func (a *Allocations) SetSelectedRow() {
	a.SelectedValue = a.GetSelectedItem()
	a.SelectedAlloc, _ = a.GetAllocationData(a.SelectedValue["id"])
}

func (a *Allocations) OnBlur() {
	a.App.Layout.Header.Menu.RemoveMenus(a.Menus)
	go a.Listener.Stop()
}

func (a *Allocations) Refresher() {
	a.UpdateTable(a.TaskGroup)
	a.App.Layout.Draw()
}

func (a *Allocations) GetAllocationData(id string) (models.Allocations, bool) {
	if id != "" {
		for i := 0; i < len(a.Data); i++ {
			if strings.HasPrefix(a.Data[i].ID, id) {
				return a.Data[i], true
			}
		}
	}
	return models.Allocations{}, false
}

func (a *Allocations) HandleButtonResponse(index int, label string) {
	if index == 1 && label == "Restart" {
		Allocations, Ok := a.GetAllocationData(a.SelectedValue["id"])
		if Ok {
			a.App.Alert.Loader(true)
			Params := &models.NomadParams{
				Region:		a.App.Config.GetRegion(),
				Namespace:	a.App.Config.GetNamespace(),
			}
			if Err := a.App.NomadClient.Restart(Allocations.ID, Allocations.TaskName, Params); Err != nil {
				a.App.Alert.Loader(false)
				a.App.Logger.Errorf("Error: While restarting %s task: %s", a.SelectedValue["name"], Err.Error())
				a.App.Alert.Error("Restart request failed...")
			} else {
				a.App.Alert.Loader(false)
				Msg := fmt.Sprintf("Task %s/%s restarted successfully...", a.SelectedValue["id"], a.SelectedValue["name"])
				a.App.Logger.Info(Msg)
				a.App.Alert.Info(Msg)
			}
			a.App.Layout.GoBack()
		} else {
			a.App.Alert.Error("Restart request failed...")
		}
	} else {
		a.App.Layout.GoBack()
	}
}

func (a *Allocations) InitRestartModal() {
	a.SelectedValue = a.GetSelectedItem()
	ConfirmModal := a.App.Primitives.Modal
	Title := fmt.Sprintf("Are you sure to restart %s/%s?", a.SelectedValue["id"], a.SelectedValue["name"])
	ConfirmModal.SetTitle(Title)
	ConfirmModal.SetData(a.SelectedValue)
	ConfirmModal.AddButtons([]string{"Cancel", "Restart"})
	ConfirmModal.SetResponseFunc(a.HandleButtonResponse)
	a.App.Layout.OpenPage("modal", true)
}