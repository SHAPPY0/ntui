package widgets

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	// "github.com/shappy0/ntui/internal/utils"
)

type Modal1 struct {
	*tview.Frame
	Content 	*tview.TextView
	Title 		string
	PageTitle 	string
	Buttons 	[]string
}

func NewModal1() *Modal1 {
	m1 := &Modal1{
		Content: 	tview.NewTextView(),
		Title:		"Output", // Used to show the title on modal
		PageTitle:	"CustomModal", //Used to open the modal
	}
	// m1.Content.SetText("Hello")
	m1.Content.SetDynamicColors(true)
	m1.Frame = tview.NewFrame(m1.Content)
	m1.SetBorders(1, 1, 0, 0, 2, 2)
	m1.SetBorder(true)
	m1.SetTitle(fmt.Sprintf(" %s ", m1.Title))
	return m1
}

func (m *Modal1) GetPageTitle() string {
	return m.PageTitle
}

func (m *Modal1) SetModalTitle(title string) {
	m.SetTitle(fmt.Sprintf(" %s ", title))
}

func (m *Modal1) SetDataX(data string) {
	m.Content.SetText(data)
}

func (m *Modal1) SetModalInputHandler(handler func(event *tcell.EventKey) *tcell.EventKey) {
	m.Frame.SetInputCapture(handler)
}