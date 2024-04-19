package views

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
	// "github.com/hashicorp/nomad/api"
)

var TitleNodes = "nodes"

type Nodes struct {
	*widgets.Table
	Title		string
	Headers		[]string
	Data		[]models.Nodes
}

func NewNodes() *Nodes {
	n := &Nodes{
		Table:			widgets.NewTable(TitleNodes),
		Title:			TitleNodes,
		Headers:		[]string{"id", "name", "state", "address", "node pool", "datacenter", "version", "allocs"},
	}
	n.Table.Headers = n.Headers
	n.Table.DrawHeader()
	return n
}

func (n *Nodes) GetTitle() string {
	return n.Title
}

func (n *Nodes) SetOnSelectFn(fn func(int, int)) {
	n.Table.SetOnSelectFn(fn)
}

func (n *Nodes) SetOnTabPressFn(fn func()) {
	n.Table.SetOnTabPressFn(fn)
}

func (n *Nodes) UpdateTableData(data []models.Nodes) {
	n.Data = data
	n.SetTableTitle(len(n.Data), "", "")
	n.Table.ClearTable()
	for i := 0; i <len(n.Data); i++ {
		CellColor, Status := utils.ColorizeStatusCell(n.Data[i].Status)
		n.Table.DrawCell(i + 1, 0, utils.GetID(n.Data[i].ID), CellColor)
		n.Table.DrawCell(i + 1, 1, n.Data[i].Name, CellColor)
		n.Table.DrawCell(i + 1, 2, Status, CellColor)
		n.Table.DrawCell(i + 1, 3, n.Data[i].Address, CellColor)
		n.Table.DrawCell(i + 1, 4, n.Data[i].NodePool, CellColor)
		n.Table.DrawCell(i + 1, 5, n.Data[i].Datacenter, CellColor)
		n.Table.DrawCell(i + 1, 6, n.Data[i].Version, CellColor)
		// n.Table.DrawCell(i + 1, 7, "n.Data[i].Volumes", CellColor)
		n.Table.DrawCell(i + 1, 7, utils.IntToStr(n.Data[i].AllocsCount), CellColor)
	}
}