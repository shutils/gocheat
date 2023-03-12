package gui

import (
	"log"

	"github.com/shutils/gocheat/common"

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
	panel.initEntity(g)
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

func (c *categories) initEntity(g *Gui) {
	c.entity = tview.NewTable()
	c.entity.SetTitle(c.getName()).SetBorder(true).SetTitleAlign(0)
	c.entity.SetSelectable(true, false)
	c.entity.SetSelectionChangedFunc(c.updator)
	c.updatePanel(c.parent)
}

func (c *categories) getEntity() tview.Primitive {
	return c.entity
}

func (c *categories) setKeybind(g *Gui) {
	g.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'l':
			g.focusPanel("index")
		case 'n':
			c.openCreateFileModal()
			return tcell.NewEventKey(tcell.KeyBS, 'k', tcell.ModNone)
		}
		return event
	})
}

func (c *categories) updatePanel(g *Gui) {
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
			p.updatePanel(c.parent)
		}
	}
}

func (c *categories) updateIndex(row int, column int) {
	for _, p := range c.parent.panels {
		if p.getName() == "index" {
			p.updatePanel(c.parent)
		}
	}
}

func (c *categories) openCreateFileModal() {
	inputField := tview.NewInputField()
	inputField.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			c.parent.closeModal("createFileModal")
		case tcell.KeyEnter:
			if err := common.CreateFile(inputField.GetText()); err != nil {
				log.Fatalln(err)
			}
			c.parent.closeModal("createFileModal")
			c.parent.updateAllPanel()
		}
	})
	inputField.SetBorder(true).SetTitle("New category name")
	inputField.SetFieldWidth(30)
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}
	c.parent.openModal("createFileModal", modal(inputField, 0, 3), func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})
}
