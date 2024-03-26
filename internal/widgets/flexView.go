package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/utils"
)

type Flex struct {
	*tview.Flex
}

func NewFlex() *Flex {
	f := &Flex{
		Flex:	tview.NewFlex(),
	}
	f.SetBorder(true)
	return f
}

func (f *Flex) AddItemX(primitive tview.Primitive, fixedSize, proportion int, focus bool) {
	f.AddItem(primitive, fixedSize, proportion, focus)
}

func (f *Flex) SetTitleX(title, a string) {
	if a != "" {
		f.SetTitle(fmt.Sprintf(" [::b][%s]%s(%s) ", utils.ColorT70d5bf, strings.ToUpper(title), a))
	} else {
		f.SetTitle(fmt.Sprintf(" [::b][%s]%s ", utils.ColorT70d5bf, strings.ToUpper(title)))
	}
}

func (f *Flex) ClearFlex() {
	f.Clear()
}