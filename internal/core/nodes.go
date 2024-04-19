package core

import (
	// "fmt"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/models"
	// "github.com/shappy0/ntui/internal/widgets"
)

type Nodes struct {
	*views.Nodes
	App				*App
	Listener		*utils.Listener
	JobId			string
	SelectedValue	map[string]string
}

func NewNodes(app *App) *Nodes {
	n := &Nodes{
		Nodes:			views.NewNodes(),
		App:			app,
	}
	n.App.Layout.Body.AddPageX(n.GetTitle(), n, true, false)
	// n.SetOnSelectFn(n.OnRowSelected)
	n.Nodes.SetFocusFunc(n.OnFocus)
	n.Nodes.SetBlurFunc(n.OnBlur)
	return n
}

func (n *Nodes) OnFocus() {
	// n.App.Layout.Header.Menu.Add(widgets.RevertMenu, true)
	n.UpdateTable()
}

func (n *Nodes) OnBlur() {
	// n.App.Layout.Header.Menu.Remove(widgets.RevertMenu)
}

// func (n *Nodes) OnRowSelected(row, col int) {
// 	n.SelectedValue = n.GetSelectedItem()
// 	go func() {
// 		n.App.Layout.QueueUpdateDraw(func() {
// 			// v.App.Primitives.VersionDiff.Render(v.JobId, v.SelectedValue, v.Diffs[row - 1])
// 			n.App.Layout.OpenPage("versiondiff", true)
// 		})
// 	}()
// }

func (n *Nodes) UpdateTable() {
	Params := &models.NomadParams{
		Region:		n.App.Config.GetRegion(),
		Namespace:	n.App.Config.GetNamespace(),
	}
	nodes, err := n.App.NomadClient.NodeList(Params)
	if err != nil {
		n.App.Logger.Error("Error while getting Node List: " + err.Error())
	}
	n.UpdateTableData(nodes)
}
