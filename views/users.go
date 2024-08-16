package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UsersView(window fyne.Window) *fyne.Container {
	label := widget.NewLabel("Users Management")
	return container.NewVBox(label)
}
