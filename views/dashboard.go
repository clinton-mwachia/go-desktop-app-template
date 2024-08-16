package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DashboardView(window fyne.Window) *fyne.Container {
	label := widget.NewLabel("Welcome to the Dashboard")
	return container.NewVBox(label)
}
