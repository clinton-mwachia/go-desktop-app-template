package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DashboardView() *fyne.Container {
	dashboardContent := widget.NewLabel("Dashboard Content")
	return container.NewStack(dashboardContent)
}
