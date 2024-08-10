package widgets

import (
	// "fmt"
	"github.com/rivo/tview"
)

type Pages struct {
	*tview.Pages
	ActivePage 		string
	History 		[]string
}

func NewPages() *Pages {
	p := &Pages{
		Pages:	tview.NewPages(),
	}
	return p
}

func (p *Pages) GetActivePage() string {
	return p.ActivePage
}

func (p *Pages) AddPageX(name string, item tview.Primitive, resize, visible bool) {
	if visible {
		p.ActivePage = name
	}
	p.AddPage(name, item, resize, visible)
}

func (p *Pages) AddHistory(name string) {
	if name != "" {
		p.History = append(p.History, name)
	}
}

func (p *Pages) OpenPageX(name string, addHistory bool) {
	if addHistory && p.ActivePage != "" {
		p.AddHistory(p.ActivePage)
	}
	if name == "main" && len(p.History) > 0 {
		p.FlushHistory()
	}
	p.SwitchToPage(name)
	p.ActivePage = name
}

func (p *Pages) OpenPage1(name string, addHistory bool) {
	p.ShowPage(name)
}

func (p *Pages) ShowPageX(name string) {
	p.AddHistory(name)
	p.ShowPage(name)
	p.ActivePage = name
}

func (p *Pages) GoBack() {
	if len(p.History) > 0 {
		PrevPage := p.HistoryPop()
		p.OpenPageX(PrevPage, false)
	}
}

func (p *Pages) HistoryPop() string {
	LastIndex := len(p.History) - 1
	PageName := p.History[LastIndex]
	p.History = p.History[:LastIndex]
	return PageName
}

func (p *Pages) FlushHistory() {
	for i := range(len(p.History)) {
		_ = i
		p.HistoryPop()
	}
}