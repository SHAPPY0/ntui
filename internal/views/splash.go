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
	Version		string
}

var Logo = []string{
	` _  _  ____  __  __  ____`,
	`( \( )(_  _)(  )(  )(_  _)`,
	` )  (   )(   )(__)(  _)(_`,
	`(_)\_) (__) (______)(____)`,
}

func NewSplash(version string) *Splash {
	s := &Splash{
		Flex:	tview.NewFlex(),
		Title:	TitleSplash,
		Version:	version,
	}
	s.SetDirection(tview.FlexRow)
	LogoV := tview.NewTextView()
	LogoV.SetDynamicColors(true)
	LogoV.SetTextAlign(tview.AlignCenter)
	s.DrawLogo(LogoV)

	Version := tview.NewTextView()
	Version.SetDynamicColors(true)
	Version.SetTextAlign(tview.AlignCenter)

	s.DrawVersion(Version, s.Version)

	s.AddItem(LogoV, 10, 1, false)
	s.AddItem(Version, 1, 1, false)
	return s
}

func (s *Splash) GetTitle() string {
	return s.Title
}


func (s *Splash) DrawLogo(t *tview.TextView) {
	LogoV := strings.Join(Logo, fmt.Sprintf("\n[%s::b]", "#cccccc"))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
			strings.Repeat("\n", 2),
			"#cccccc",
			LogoV)
}

func (s *Splash) DrawVersion(t *tview.TextView, Version string) {
	fmt.Fprintf(t, "[%s::b]Version: [orange::b]%s", "#cccccc", Version)
}