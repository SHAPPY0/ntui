package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
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
	reg.SetOnTabPressFn(reg.OnTabPress)
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
	reg.App.Alert.Loader(true)
	go func() {
		metadata := models.Metadata{
			Host:	reg.App.Config.NomadBaseUrl,
			Region:	reg.App.Config.Region,
			Namespace:	reg.App.Config.Namespace,
		}
		reg.App.Layout.Header.SetMetadata(metadata)
		reg.App.Layout.QueueUpdateDraw(func() {
			reg.App.Primitives.Namespaces.UpdateTable()
			reg.App.Alert.Loader(false)
			reg.OnTabPress()
		})
	}()
}

func (reg *Regions) UpdateTable() {
	Data, _ := reg.App.NomadClient.Regions()
	reg.UpdateTableData(Data)
}