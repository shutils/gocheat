package gui

import "github.com/rivo/tview"

type panel interface {
	getName() string
	getWidth() int
	initEntity()
	setEntity(*Gui)
  getEntity() tview.Primitive
	focus(*Gui)
	setKeybind(*Gui)
}
