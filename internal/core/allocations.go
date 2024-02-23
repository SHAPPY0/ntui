package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

type Allocations struct {
	*view.Allocations
	App				*App
	Listner			*utils.Listner
	SelectedValue	map[string]string
	TaskGroup		map[string]string
	Data			[]models.Allocations
}

func NewAllocations(app *App) *Allocations {
	a := &Allocations {
		Allocations:	views.NewAllocations(),
		App:			app,
	}
	a.App.Layout.Body.AddPageX(a.GetTitle(), a, true, false)
	a.SetOnSelectFn(a.OnRowSelected)
	a.Listner = utils.NewListner(a.Refresher)
	a.SetFocusFunc(a.OnFocus)
	a.SetBlurFunc(a.OnBlur)
	return a
}

func (a *Allocations) OnRowSelected(row, col int) {
	a.SelectedValue = a.GetSelectedItem()
	go func() {
		a.App.Layout.QueueUpdateDraw(func() {
			a.App.Layout.OpenPage("tasks", true)
		})
	}
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
	go a.Listner.Listen()
}

func (a *Allocations) OnBlur() {
	go a.Listner.Stop()
}

func (a *Allocations) Refresher() {
	a.UpdateTable(a.TaskGroup)
	a.App.Layout.Draw()
}