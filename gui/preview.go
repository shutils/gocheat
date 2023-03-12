package gui

import (
	"github.com/shutils/gocheat/common"
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
	panel.initEntity(g)
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

func (p *preview) initEntity(g *Gui) {
	p.entity = tview.NewTextView()
	p.entity.SetTitle(p.getName()).SetBorder(true).SetTitleAlign(0)
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
      g.updateAllPanel()
		}
		return event
	})
}

func (p *preview) updatePanel(g *Gui) {
  cp := g.getPanelEntity("categories")
  cpt, ok := cp.(*tview.Table)
  if ok {
    row, colulmn := cpt.GetSelection()
    fn := cpt.GetCell(row, colulmn).Text
    t := common.GetText(fn)
    p.entity.SetText(t)
  }
}
