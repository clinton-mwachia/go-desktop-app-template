package main

import (
	"desktop-app-template/utils"
	"desktop-app-template/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	utils.ConnectDB("mongodb://localhost:27017")

	application := app.New()
	window := application.NewWindow("Todo App")

	var content *fyne.Container

	dashboardView := views.DashboardView(window)
	usersView := views.UsersView(window)
	todosView := views.TodosView(window)

	content = dashboardView

	sidebar := container.NewVBox(
		widget.NewButton("Dashboard", func() {
			content.Objects = []fyne.CanvasObject{dashboardView}
			content.Refresh()
		}),
		widget.NewButton("Users", func() {
			content.Objects = []fyne.CanvasObject{usersView}
			content.Refresh()
		}),
		widget.NewButton("Todos", func() {
			content.Objects = []fyne.CanvasObject{todosView}
			content.Refresh()
		}),
	)

	mainContent := container.NewHSplit(sidebar, content)
	mainContent.Offset = 0.2

	loginView := views.LoginView(window, func() {
		window.SetContent(mainContent)
	})

	window.SetContent(loginView)
	window.ShowAndRun()
}
