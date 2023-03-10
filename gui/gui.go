package gui

import (
	"github.com/rivo/tview"
)

type Gui struct {
	app    *tview.Application
	panels []panel
	flex   *tview.Flex
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
	g.flex = tview.NewFlex()
	cp := getCategories(g)
	cp.setEntity(g)
	cp.setKeybind(g)
	ip := getIndexPanel(g)
	ip.setEntity(g)
	pp := getPreview(g)
	pp.setEntity(g)
	cp.updatePreview(0, 0)
	g.app.SetRoot(g.flex, true)
}

func (g *Gui) focusPanel(name string) {
	for _, p := range g.panels {
		if p.getName() == name {
			p.focus(g)
			p.setKeybind(g)
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
