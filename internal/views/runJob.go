package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var titleRunJob = "run job"

type RunJob struct {
	*tview.Grid
	TextAreaX	*widgets.TextArea
	HelpMenu	*tview.TextView
	Title		string
	Data		[]models.Namespaces
}

var (
	PasteMenu = fmt.Sprintf(" [%s]<ctrl+v> [%s]Paste", utils.ColorTOrange, utils.ColorTWhite)
	DryRunMenu = fmt.Sprintf(" [%s]<ctrl+d> [%s]Dry Run", utils.ColorTOrange, utils.ColorTWhite)
	MoveUpMenu = fmt.Sprintf(" [%s]<↑> [%s]Scroll Up", utils.ColorTOrange, utils.ColorTWhite)
	MoveDownMenu = fmt.Sprintf(" [%s]<↓> [%s]Scroll Down", utils.ColorTOrange, utils.ColorTWhite)
	// SubmitJobMenu = fmt.Sprintf(", [%s]<ctrl+s> [%s]Submit", utils.ColorTOrange, utils.ColorTWhite)
)

func NewRunJob() *RunJob {
	rj := &RunJob {
		Grid:		tview.NewGrid(),
		Title:		titleRunJob,
		TextAreaX:	widgets.NewTextArea(titleRunJob),
		HelpMenu:	tview.NewTextView(),
	}
	rj.Render()
	return rj
}

func (rj *RunJob) GetTitle() string {
	return rj.Title
}

func (rj *RunJob) GetTextAreaText() string {
	return rj.TextAreaX.GetText()
}

func (rj *RunJob) SetContent(data string) {
	rj.TextAreaX.SetText(data, true)
}

func (rj *RunJob) Clear() {
	rj.TextAreaX.SetText("", false)
}

func (rj *RunJob) Render() {
	helpMenu := PasteMenu + MoveUpMenu + MoveDownMenu + DryRunMenu
	rj.HelpMenu.SetDynamicColors(true).SetText(helpMenu)
	rj.AddItem(rj.TextAreaX, 0, 0, 50, 2, 0, 0, true)
	rj.AddItem(rj.HelpMenu, 50, 0, 1, 1, 0, 0, false)
	rj.TextAreaX.SetPlaceholderX("Enter Job HCL...")
}

func (rj *RunJob) UpdateHelpMenu(menus string) {
	rj.HelpMenu.SetText(menus)
}


