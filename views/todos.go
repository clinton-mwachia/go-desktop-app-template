package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TodosView() *fyne.Container {
	todosContent := widget.NewLabel("Todos Content")
	return container.NewStack(todosContent)
}
