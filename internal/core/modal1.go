package core

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
)

type CustomModal struct {
	*tview.Grid
	Modal1 		*widgets.Modal1
	App			*App
	Title 		string
	Buttons		[]string
	Data 		string
	ResponseFn  func(int, string)
}

func NewCustomModal(app *App) *CustomModal {
	cm := &CustomModal{
		Grid:		tview.NewGrid(),
		Modal1:		widgets.NewModal1(),
		Buttons:	[]string{},

		App:		app,
	}
	// cm.SetColumns(0, 64, 0)
	// cm.SetRows(0, 22, 0)
	cm.SetColumns(0, 50, 0)
	cm.SetRows(0, 15, 0)
	cm.AddItem(cm.Modal1, 1, 1, 1, 1, 0, 0, true)
	cm.App.Layout.Body.AddPageX(cm.GetPageTitle(), cm, true, false)
	return cm
}

func (cm *CustomModal) GetPageTitle() string {
	return cm.Modal1.GetPageTitle()
}

func (cm *CustomModal) SetTitle(title string) {
	cm.Modal1.SetModalTitle(title)
}

// func (m *Modal) AddButtons(buttons []string) {
// 	m.Buttons = buttons
// 	m.ClearButtons()
// 	m.SetButtons(m.Buttons)
// }

func (cm *CustomModal) SetData(data string) {
	cm.Data = data
	cm.Modal1.SetDataX(cm.Data)
}

// func (m *Modal) SetResponseFunc(fn func(int, string)) {
// 	m.ResponseFn = fn
// 	m.SetDoneFunc(m.ResponseFn)
// }

func (cm *CustomModal) SetInputHandler(handler func(event *tcell.EventKey) *tcell.EventKey) {
	cm.Modal1.SetModalInputHandler(handler)
}