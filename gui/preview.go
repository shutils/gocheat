package gui

import (
	"gocheat/common"
	"log"
	"path"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type preview struct {
	entity *tview.TextView
  parent *Gui
}

func getPreview(g *Gui) *preview {
	panel := &preview{
    parent: g,
  }
	panel.initEntity()
	return panel
}

func (p *preview) getName() (n string) {
	n = "preview"
	return
}

func (p *preview) getWidth() (w int) {
	w = 3
	return
}

func (p *preview) focus(g *Gui) {
	g.app.SetFocus(p.entity)
}

func (p *preview) initEntity() {
	p.entity = tview.NewTextView()
	p.entity.SetTitle(p.getName()).SetBorder(true).SetTitleAlign(0)
}

func (p *preview) setEntity(g *Gui) {
	p.initEntity()
	g.panels = append(g.panels, p)
	g.flex.AddItem(p.entity, 0, p.getWidth(), true)
}

func (p *preview) getEntity() tview.Primitive {
	return p.entity
}

func (p *preview) setKeybind(g *Gui) {
	g.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			g.focusPanel("index")
		case 'e':
			e := g.getPanelEntity("categories")
			t, ok := e.(*tview.Table)
			if !ok {
				log.Fatalln("It is not Table.")
			}
			r, c := t.GetSelection()
			category := t.GetCell(r, c).Text
			appDirName := common.GetAppDirName()
			g.app.Suspend(func() {
				common.EditFile(path.Join(appDirName, category))
			})
		}
		return event
	})
}
