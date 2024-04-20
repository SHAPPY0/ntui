package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

type Header struct {
	*tview.Flex
	Logo			*tview.TextView
	Menu 			*widgets.Menu
	Metadata 		*widgets.MapView
}

// var Logo = []string{
// 	` _  _  ____  __  __  ____`,
// 	`( \( )(_  _)(  )(  )(_  _)`,
// 	` )  (   )(   )(__)(  _)(_`,
// 	`(_)\_) (__) (______)(____)`,
// }

func NewHeader() *Header {
	h := &Header{
		Flex:		tview.NewFlex(),
		Logo:		tview.NewTextView(),
		Menu:		widgets.NewMenu(),
		Metadata:	widgets.NewMapView(),
	}
	h.RenderLogo()
	h.RenderMenu(make([]widgets.Item, 0))
	h.AddItem(h.Logo, 0, 1, false)
	h.AddItem(h.Menu.Grid1, 0, 1, false)
	h.AddItem(h.Menu.Grid2, 0, 1, false)
	h.AddItem(h.Menu.Grid3, 0, 1, false)
	h.AddItem(h.Metadata, 0, 1, false)
	return h
}

func (h *Header) RenderLogo() error {
	h.Logo.SetDynamicColors(true)
	for I, S := range Logo {
		fmt.Fprintf(h.Logo, "[%s::b]%s", "", S)
		if I + 1 < len(Logo) {
			fmt.Fprintf(h.Logo, "\n")
		}
	}
	return nil
}

func (h *Header) RenderMenu(menus []widgets.Item) error {
	h.Menu.RenderGlobalMenus()
	h.Menu.RenderMenu(menus)
	return nil
}

func (h *Header) SetMetadata(metadata models.Metadata) {
	h.Metadata.Clear()
	hostKey := fmt.Sprintf("[%s]%s:", "cadetblue", "Host")
	hostValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.Host != "" {
		hostValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.Host)
	}
	h.Metadata.SetMapKeyValue(hostKey, hostValue)

	regionKey := fmt.Sprintf("[%s]%s:", "cadetblue", "Region")
	regionValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.Region != "" {
		regionValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.Region)
	}
	h.Metadata.SetMapKeyValue(regionKey, regionValue)

	namespaceKey := fmt.Sprintf("[%s]%s:", "cadetblue", "Namespace")
	namespaceValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.Namespace != "" {
		namespaceValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.Namespace)
	}
	h.Metadata.SetMapKeyValue(namespaceKey, namespaceValue)

	h.Metadata.DrawMapView()
}