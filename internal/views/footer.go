package views

import (
	"fmt"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

type Footer struct {
	*widgets.Flex
	// Grid1		*widgets.TextView
	Grid2		*widgets.TextView
	// Grid3		*widgets.TextView
}

func NewFooter() *Footer {
	f := &Footer{
		Flex:	widgets.NewFlex(),
		// Grid1:	widgets.NewTextView(""),
		Grid2:	widgets.NewTextView(""),
		// Grid3:	widgets.NewTextView(""),
	}
	f.SetBorder(false)

	//Footer Grid 1
	// f.Grid1.SetBorder(false)
	
	//Footer Grid 2
	f.Grid2.SetBorder(false)
	f.Grid2.SetDynamicColors(true)
	f.Grid2.SetTextAlignX("AlignCenter")

	//Footer Grid 3
	// f.Grid3.SetBorder(false)

	// f.AddItemX(f.Grid1, 0, 1, false)
	f.AddItemX(f.Grid2, 0, 1, false)
	// f.AddItemX(f.Grid3, 0, 1, false)
	return f
}

func (f *Footer) SetAlert(message models.AlertMessage) {
	if message.Text == "" {
		f.Grid2.ClearX()
		return
	}
	Color := utils.ColorTGreen
	if message.Type == utils.Warning {
		Color = utils.ColorTOrange
	} else if message.Type == utils.Error {
		Color = utils.ColorTRed
	} else if message.Type == utils.Loader {
		Color = utils.ColorTad7c5a
	}
	fmt.Fprintf(f.Grid2, "[%s]%s", Color, message.Text)
}