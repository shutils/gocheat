package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Gui struct {
	app          *tview.Application
	panels       []panel
	pages        *tview.Pages
	currentPanel string
}

func New() *Gui {
	gui := &Gui{
		app: tview.NewApplication(),
	}
	gui.initApp()
	return gui
}

func (g *Gui) Run() {
	g.app.Run()
}

func (g *Gui) initApp() {
	g.panels = append(g.panels, getCategories(g))
	g.panels = append(g.panels, getIndexPanel(g))
	g.panels = append(g.panels, getPreview(g))
	g.pages = tview.NewPages()
	g.initPanel()
	g.focusPanel("categories")
	g.app.SetRoot(g.pages, true).SetFocus(g.pages)
	g.updateAllPanel()
}

func (g *Gui) initPanel() {
	flex := tview.NewFlex()
	for _, p := range g.panels {
		flex.AddItem(p.getEntity(), 0, p.getWidth(), true)
	}
	g.pages.AddPage("main", flex, true, true)
}

func (g *Gui) updateAllPanel() {
	for _, p := range g.panels {
		p.updatePanel(g)
	}
}

func (g *Gui) focusPanel(name string) {
	for _, p := range g.panels {
		if p.getName() == name {
			p.focus(g)
			p.setKeybind(g)
			g.currentPanel = p.getName()
		}
	}
}

func (g *Gui) getPanelEntity(name string) (e tview.Primitive) {
	for _, p := range g.panels {
		if p.getName() == name {
			e = p.getEntity()
		}
	}
	return
}

func (g *Gui) openModal(name string, modal tview.Primitive, keybind func(event *tcell.EventKey) *tcell.EventKey) {
	g.pages.AddPage(name, modal, true, true)
	g.app.SetInputCapture(keybind)
}

func (g *Gui) closeModal(name string) {
	g.pages.HidePage(name)
	g.pages.RemovePage(name)
	g.focusPanel(g.currentPanel)
}
