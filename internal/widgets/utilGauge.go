package widgets

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
)

type UtilGauge struct {
	*tvxwidgets.UtilModeGauge
	Label		string
}

func NewUtilGauge(label string) *UtilGauge {
	ug := &UtilGauge{
		UtilModeGauge:	tvxwidgets.NewUtilModeGauge(),
		Label:	label,
	}
	ug.SetLabel(ug.Label)
	ug.SetLabelColor(tcell.ColorLightSkyBlue)
	ug.SetRect(5, 4, 5, 3)
	ug.SetWarnPercentage(80)
	ug.SetCritPercentage(95)
	ug.SetBorder(false)
	return ug
}

func (ug *UtilGauge) SetValueX(value float64) {
	ug.SetValue(value)
}

func PrimitiveToGauge(primitive tview.Primitive) *UtilGauge{
	return primitive.(*UtilGauge)
}