package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UsersView() *fyne.Container {
	usersContent := widget.NewLabel("Users Content")
	return container.NewStack(usersContent)
}
