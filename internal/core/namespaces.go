package core

import (
	"github.com/shappy0/ntui/internal/views"
)

type Namespaces struct {
	*views.Namespaces
	App				*App
	SelectedValue	map[string]string
}

func NewNamespaces(app *App) *Namespaces {
	n := &Namespaces{
		Namespaces:		views.NewNamespaces(),
		App:			app,
	}
	n.App.Layout.Body.AddPageX(n.GetTitle(), n, true, false)
	n.SetOnSelectFn(n.OnRowSelected)
	n.SetOnTabPress(n.OnTabPress)
	return n
}

func (n *Namespaces) OnTabPress() {
	if n.App.Primitives.Namespaces.HasFocus() {
		if n.App.Primitives.Regions.GetRowCount() > 1 {
			n.App.Layout.ChangeFocusX(n.App.Primitives.Regions)
		}
	}
}

func (n *Namespaces) OnRowSelected(row, col int) {
	n.SelectedValue = n.GetSelectedItem()
	n.App.Config.SetNamespace(n.SelectedValue["name"])
	n.App.Alert.Loader(true)
	go func() {
		n.App.Layout.QueueUpdateDraw(func() {
			n.App.Primitives.Jobs.UpdateTable()
			n.App.Alert.Loader(false)
			n.App.Layout.OpenPage("jobs", true)
		})
	}()
}

func (n *Namespaces) UpdateTable() {
	Data, _ := n.App.NomadClient.Namespaces()
	Region := n.App.Config.GetRegion()
	n.UpdateTableData(Data, Region)
}