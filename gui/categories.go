package gui

import (
	"github.com/shutils/gocheat/common"
	"path/filepath"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type categories struct {
	entity *tview.Table
	parent *Gui
}

func getCategories(g *Gui) *categories {
	panel := &categories{
		parent: g,
	}
	panel.initEntity()
	return panel
}

func (c *categories) getName() (n string) {
	n = "categories"
	return
}

func (c *categories) getWidth() (w int) {
	w = 1
	return
}

func (c *categories) focus(g *Gui) {
	g.app.SetFocus(c.entity)
}

func (c *categories) initEntity() {
	c.entity = tview.NewTable()
	c.loadFiles()
	c.entity.SetTitle(c.getName()).SetBorder(true).SetTitleAlign(0)
	c.entity.SetSelectable(true, false)
	c.entity.SetSelectionChangedFunc(c.updator)
}

func (c *categories) setEntity(g *Gui) {
	c.initEntity()
	g.panels = append(g.panels, c)
	g.flex.AddItem(c.entity, 0, c.getWidth(), true)
}

func (c *categories) getEntity() tview.Primitive {
	return c.entity
}

func (c *categories) setKeybind(g *Gui) {
	g.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'l':
			g.focusPanel("index")
		}
		return event
	})
}

func (c *categories) loadFiles() {
	appDirName := common.GetAppDirName()
	fs := common.GetFileNames(appDirName)
	for i, v := range fs {
		c.entity.SetCell(i, 0, tview.NewTableCell(v))
	}
}

func (c *categories) updator(row int, column int) {
	c.updatePreview(row, column)
	c.updateIndex(row, column)
}

func (c *categories) updatePreview(row int, column int) {
	for _, p := range c.parent.panels {
		if p.getName() == "preview" {
			tp := p.getEntity()
			tv, ok := tp.(*tview.TextView)
			if ok {
				text := common.GetText(filepath.Join(common.GetAppDirName(), c.entity.GetCell(row, column).Text))
				tv.SetText(text)
			}
		}
	}
}

func (c *categories) updateIndex(row int, column int) {
	for _, p := range c.parent.panels {
		if p.getName() == "index" {
			tp := p.getEntity()
			t, ok := tp.(*tview.Table)
			if ok {
        t.Clear()
				indexes := common.GetIndex(filepath.Join(c.entity.GetCell(row, column).Text))
				for i, v := range indexes {
					t.SetCell(i, 0, tview.NewTableCell(v.Text))
					t.SetCell(i, 1, tview.NewTableCell(strconv.Itoa(v.Index)))
				}
			}
		}
	}
}
