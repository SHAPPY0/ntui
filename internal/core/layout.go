package core

import (
	"time"
	"github.com/rivo/tview"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

type Layout struct {
	*views.App
	Pages			*tview.Pages
	MainLayout		*tview.Flex
	Splash			*views.Splash
	Header			*views.Header
	Body			*widgets.Pages
	Footer			*views.Footer
}

func NewLayout(version string) *Layout {
	l := &Layout{
		App:		views.NewApp(),
		MainLayout:	tview.NewFlex(),
		Splash:		views.NewSplash(version),
		Header:		views.NewHeader(),
		Body:		widgets.NewPages(),
		Footer:		views.NewFooter(),
	}
	l.MainLayout.SetDirection(tview.FlexRow).
				AddItem(l.Header, 5, 1, false).
				AddItem(l.Body, 0, 1, true).
				AddItem(l.Footer, 1, 1, false)

	l.SetRoot(l.Splash, true)
	return l
}

func (l *Layout) Run(app *App) error {
	metadata := models.Metadata{
		Host:		app.Config.NomadBaseUrl,
		Namespace: 	app.Config.Namespace,
		Region:		app.Config.Region,
	}
	l.Header.SetMetadata(metadata)
	go func() {
		<- time.After(1 * time.Second)
		l.QueueUpdateDraw(func() {
			l.SetRoot(l.MainLayout, true)
			l.SetFocus(l.MainLayout)
			if app.Config.IsRegionInConfig() {
				app.Primitives.Jobs.UpdateTable()
				l.OpenPage("jobs", true)
			} else {
				l.OpenPage("main", true)
			}
		})
	}()
	if Err := l.SetFocus(l.Splash).Run(); Err != nil {
		return Err
	}
	return nil
}

func (l *Layout) ChangeFocusX(p tview.Primitive) {
	l.SetFocus(p)
}

func (l *Layout) OpenPage(name string, addHistory bool) {
	l.Body.OpenPageX(name, addHistory)
}

func (l *Layout) GetActivePage() string {
	return l.Body.GetActivePage()
}

func (l *Layout) GoBack() {
	l.Body.GoBack()
}