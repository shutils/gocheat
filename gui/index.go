package gui

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shutils/gocheat/common"
)

type indexPanel struct {
	entity *tview.Table
	parent *Gui
}

func getIndexPanel(g *Gui) *indexPanel {
	panel := &indexPanel{
		parent: g,
	}
	panel.initEntity(g)
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

func (i *indexPanel) initEntity(g *Gui) {
	i.entity = tview.NewTable()
	i.entity.SetTitle(i.getName()).SetBorder(true).SetTitleAlign(0)
	i.entity.SetSelectable(true, false)
	i.entity.SetSelectionChangedFunc(i.setScroll)
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
				if !i.entity.GetCell(row, column).NotSelectable {
					indexRow, err := strconv.Atoi(i.entity.GetCell(row, 1).Text)
					if err != nil {
						log.Fatalln(err)
					}
					tv.ScrollTo(indexRow, 0)
				}
			}
		}
	}
}

func (i *indexPanel) updatePanel(g *Gui) {
	cp := g.getPanelEntity("categories")
	cpt, ok := cp.(*tview.Table)
	if ok {
		i.entity.Clear()
		i.entity.SetCell(0, 0, &tview.TableCell{
			Text:          "headline",
			NotSelectable: true,
		})
		i.entity.SetCell(0, 1, &tview.TableCell{
			Text:          "index",
			NotSelectable: true,
		})
		row, colulmn := cpt.GetSelection()
		fn := cpt.GetCell(row, colulmn).Text
		fns := common.GetIndex(fn)
		for index, v := range fns {
			i.entity.SetCellSimple(1+index, 0, v.Text)
			i.entity.SetCellSimple(1+index, 1, strconv.Itoa(v.Index))
		}
	}
}
