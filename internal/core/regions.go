package core

import (
	"github.com/shappy0/ntui/internal/views"
)

type Regions struct {
	*views.Regions
	App				*App
	SelectedValue	map[string]string
}

func NewRegions(app *App) *Regions {
	reg := &Regions{
		Regions:	views.NewRegions(),
		App:		app,
	}
	reg.App.Layout.Body.AddPageX(reg.GetTitle(), reg, true, false)
	reg.UpdateTable()
	reg.SetOnSelectFn(reg.OnRowSelected)
	reg.SetOnTabPress(reg.OnTabPress)
	return reg
}

func (reg *Regions) OnTabPress() {
	if reg.App.Primitives.Regions.HasFocus() {
		if reg.App.Primitives.Namespaces.GetRowCount() > 1 {
			reg.App.Layout.ChangeFocusX(reg.App.Primitives.Namespaces)
		}
	}
}

func (reg *Regions) OnRowSelected(row, col int) {
	reg.SelectedValue = reg.GetSelectedItem()
	reg.App.Config.SetRegion(reg.SelectedValue["region"])
	go func() {
		reg.App.Layout.QueueUpdateDraw(func() {
			reg.App.Primitives.Namespaces.UpdateTable()
			reg.OnTabPress()
		})
	}()
}

func (reg *Regions) UpdateTable() {
	Data, _ := reg.App.NomadClient.Regions()
	reg.UpdateTableData(Data)
}