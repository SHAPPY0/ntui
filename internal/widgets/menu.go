package widgets

import (
	"fmt"
)

type Menu struct {
	*MapView
	Items		[]Item
}

type Item struct {
	Name 		string
	Description string
	Icon 		string
}

var (
	EnterMenu = Item{
		Name:		"enter",
		Icon:		"enter",
		Description: "Select Row",
	}
	UpArrowMenu = Item{
		Name:		"up_arrow",
		Icon:		"↑",
		Description: "Move UP",
	}
	DownArrowMenu = Item{
		Name:		"down_arrow",
		Icon:		"↓",
		Description: "Move Down",
	}
	EscMenu = Item{
		Name:		"esc",
		Icon:		"esc",
		Description: "Go Back",
	}
	RegionNMenu = Item{
		Name:		"region_namespace",
		Icon:		"ctrl+r",
		Description: "Show Region & Namespace",
	}
	LogMenu = Item{
		Name:		"log",
		Icon:		"l",
		Description: "Show Log",
	}
	StdoutMenu = Item{
		Name:		"stdout",
		Icon:		"o",
		Description: "STDOUT Logs",
	}
	StderrMenu = Item{
		Name:		"stderr",
		Icon:		"e",
		Description: "STDERR Logs",
	}
	VersionMenu = Item{
		Name:		"versions",
		Icon:		"v",
		Description: "Show Job Versions",
	}
	RevertMenu = Item{
		Name:		"revert_version",
		Icon:		"ctrl+v",
		Description: "Revert Version",
	}
	RestartTaskMenu = Item{
		Name:		"restart_task",
		Icon:		"ctrl+t",
		Description: "Restart Task",
	}
)

func NewMenu() *Menu {
	m := &Menu{
		MapView:	NewMapView(),
		Items:		make([]Item, 0),
	}
	return m
}

func MenuExist(menu *Menu, name string) bool {
	found := false
	for _, Item := range menu.Items {
		if Item.Name == name {
			found = true
		}
	}
	return found
}

func (m *Menu) Add(menu Item, refresh bool) *Menu {
	if MenuExist(m, menu.Name) {
		return m
	}
	m.Items = append(m.Items, menu)
	if refresh {
		m.Render()
	}
	return m
}

func (m *Menu) Render() {
	m.Clear()
	for _, Menu := range m.Items {
		Key := fmt.Sprintf("[%s]<%s>", "orange", Menu.Icon)
		Value := fmt.Sprintf("[%s]%s\n", "DimGray", Menu.Description)
		m.SetMapKeyValue(Key, Value)
	}
	m.DrawMapView()
}

func (m *Menu) Remove(menu Item) {
	for I, Menu := range m.Items {
		if Menu.Name == menu.Name {
			m.Items = append(m.Items[:I], m.Items[I + 1:]...)
		}
	}
	m.Render()
}