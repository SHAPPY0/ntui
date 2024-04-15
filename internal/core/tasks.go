package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/widgets"
)

type Tasks struct {
	*views.Tasks
	App				*App
	SelectedValue	models.Allocations
}

func NewTasks(app *App) *Tasks {
	t := &Tasks{
		Tasks:		views.NewTasks(),
		App:		app,
	}
	t.App.Layout.Body.AddPageX(t.GetTitle(), t, true, false)
	t.Tasks.SetFocusFunc(t.OnFocus)
	t.Tasks.SetBlurFunc(t.OnBlur)
	return t
}

func (t *Tasks) OnFocus() {
	allocsList := t.App.Primitives.Allocations.Data
	selectedAlloc := t.App.Primitives.Allocations.SelectedValue
	for _, alloc := range allocsList {
		if utils.GetID(alloc.ID) == selectedAlloc["id"] {
			t.SelectedValue = alloc
		}
	}
	t.App.Layout.Header.Menu.Remove(widgets.EnterMenu)
	t.App.Layout.Header.Menu.Remove(widgets.UpArrowMenu)
	t.App.Layout.Header.Menu.Remove(widgets.DownArrowMenu)
	t.App.Layout.Header.Menu.Add(widgets.RestartTaskMenu, false)
	t.App.Layout.Header.Menu.Add(widgets.LogMenu, true)
	t.DrawView(t.SelectedValue)
}

func (t *Tasks) OnBlur() {
	t.InfoView.ClearFlex()
	t.UsageView.Clear()
	t.DetailsView.Clear()
	t.Tasks.Clear()
	t.App.Layout.Header.Menu.Remove(widgets.LogMenu)
	t.App.Layout.Header.Menu.Remove(widgets.RestartTaskMenu)
	t.App.Layout.Header.Menu.Add(widgets.EnterMenu, false)
	t.App.Layout.Header.Menu.Add(widgets.UpArrowMenu, false)
	t.App.Layout.Header.Menu.Add(widgets.DownArrowMenu, true)
}