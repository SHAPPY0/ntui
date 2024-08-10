package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

type Tasks struct {
	*views.Tasks
	App				*App
	Listener		*utils.Listener
	SelectedValue	models.Allocations
}

func NewTasks(app *App) *Tasks {
	t := &Tasks{
		Tasks:		views.NewTasks(),
		App:		app,
	}
	t.App.Layout.Body.AddPageX(t.GetTitle(), t, true, false)
	t.Listener = utils.NewListener(app.Config.RefreshRate, t.Refresher)
	t.Tasks.SetFocusFunc(t.OnFocus)
	t.Tasks.SetBlurFunc(t.OnBlur)
	return t
}

func (t *Tasks) OnFocus() {
	selectedAlloc := t.App.Primitives.Allocations.SelectedValue
	t.SelectedValue, _ = t.App.Primitives.Allocations.GetAllocationData(selectedAlloc["id"])
	t.App.Layout.Header.Menu.RenderMenu(t.Menus, true)
	t.App.Layout.Header.Menu.RemoveMenus(t.RemoveMenus)
	t.DrawView(t.SelectedValue)
	go t.Listener.Listen()
}

func (t *Tasks) OnBlur() {
	t.InfoView.ClearFlex()
	t.UsageView.Clear()
	t.DetailsView.Clear()
	t.Tasks.Clear()
	t.App.Layout.Header.Menu.RenderMenu(t.RemoveMenus, true)
	t.App.Layout.Header.Menu.RemoveMenus(t.Menus)
	go t.Listener.Stop()
}

func (t *Tasks) RefreshUsageStats() {
	taskGroup := t.App.Primitives.Allocations.TaskGroup
	t.App.Primitives.Allocations.UpdateTable(taskGroup)
	selectedAlloc := t.App.Primitives.Allocations.SelectedValue
	t.SelectedValue, _ = t.App.Primitives.Allocations.GetAllocationData(selectedAlloc["id"])
	t.SetUsageData(t.SelectedValue.Tasks.Resources)
	t.EventsView(t.SelectedValue.Events)
}

func (t *Tasks) Refresher() {
	t.RefreshUsageStats()
	t.App.Layout.Draw()
}