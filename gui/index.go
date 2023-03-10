package gui

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type indexPanel struct {
	entity *tview.Table
	parent *Gui
}

func getIndexPanel(g *Gui) *indexPanel {
	panel := &indexPanel{
		parent: g,
	}
	panel.initEntity()
	return panel
}

func (i *indexPanel) getName() (n string) {
	n = "index"
	return
}

func (i *indexPanel) getWidth() (w int) {
	w = 1
	return
}

func (i *indexPanel) focus(g *Gui) {
	g.app.SetFocus(i.entity)
}

func (i *indexPanel) initEntity() {
	i.entity = tview.NewTable()
	i.entity.SetTitle(i.getName()).SetBorder(true).SetTitleAlign(0)
	i.entity.SetSelectable(true, false)
  i.entity.SetSelectionChangedFunc(i.setScroll)
}

func (i *indexPanel) setEntity(g *Gui) {
	i.initEntity()
	g.panels = append(g.panels, i)
	g.flex.AddItem(i.entity, 0, i.getWidth(), true)
}

func (i *indexPanel) getEntity() tview.Primitive {
	return i.entity
}

func (i *indexPanel) setKeybind(g *Gui) {
	g.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			g.focusPanel("categories")
		case 'l':
			g.focusPanel("preview")
		}
		return event
	})
}

func (i *indexPanel) setScroll(row int, column int) {
	for _, p := range i.parent.panels {
		if p.getName() == "preview" {
			tp := p.getEntity()
			tv, ok := tp.(*tview.TextView)
			if ok {
        indexRow, err := strconv.Atoi(i.entity.GetCell(row, 1).Text)
        if err != nil {
          log.Fatalln(err)
        }
				tv.ScrollTo(indexRow, 0)
			}
		}
	}
}
