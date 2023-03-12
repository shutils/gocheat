package gui

import "github.com/rivo/tview"

type panel interface {
	getName() string
	getWidth() int
	initEntity(*Gui)
	getEntity() tview.Primitive
	focus(*Gui)
	setKeybind(*Gui)
	updatePanel(*Gui)
}
