package views

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
)

var TitleSplash = "splash"

type Splash struct {
	*tview.Flex
	Title 		string
}

var Logo = []string{
	` _  _  ____  __  __  ____`,
	`( \( )(_  _)(  )(  )(_  _)`,
	` )  (   )(   )(__)(  _)(_`,
	`(_)\_) (__) (______)(____)`,
}

func NewSplash() *Splash {
	s := &Splash{
		Flex:	tview.NewFlex(),
		Title:	TitleSplash,
	}
	s.SetDirection(tview.FlexRow)
	Logo := tview.NewTextView()
	Logo.SetDynamicColors(true)
	Logo.SetTextAlign(tview.AlignCenter)
	s.DrawLogo(Logo)

	Version := tview.NewTextView()
	Version.SetDynamicColors(true)
	Version.SetTextAlign(tview.AlignCenter)

	s.DrawVersion(Version, "0.1")

	s.AddItem(Log, 10, 1, false)
	s.AddItem(Versions, 1, 1, false)
	return s
}

func (s *Splash) GetTitle() string {
	return s.Title
}


func (s *Splash) DrawLogo(t *tview.TextView) {
	Logo := strings.Join(Logo, fmt.Sprintf("\n[%s::b]", "#cccccc"))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
			strings.Repeat("\n", 2),
			"#cccccc",
			logo)
}

func (s *Splash) DrawVersion(t *tview.TextView, Version string) {
	fmt.Fprintf(t, "[%s::b]Version: [orange::b]%s", "#cccccc", Version)
}