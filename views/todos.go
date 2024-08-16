package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TodosView(window fyne.Window) *fyne.Container {
	label := widget.NewLabel("Manage Todos")
	return container.NewVBox(label)
}
