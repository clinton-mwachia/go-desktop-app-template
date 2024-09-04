package views

import (
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Sidebar(window fyne.Window, showDashboard, showUsers, showLogs, showTodos, showLogin func(), userID primitive.ObjectID) *fyne.Container {
	isAdmin := utils.IsAdmin(userID, window)

	dashboardButton := widget.NewButton("Dashboard", func() {
		showDashboard()
	})

	var userButton *widget.Button
	var logsButton *widget.Button

	if isAdmin {
		userButton = widget.NewButton("Users", func() {
			showUsers()
		})
		logsButton = widget.NewButton("Logs", func() {
			showLogs()
		})
	} else {
		userButton = &widget.Button{}
		userButton.Hide()

		logsButton = &widget.Button{}
		logsButton.Hide()
	}

	todosButton := widget.NewButton("Todos", func() {
		showTodos()
	})

	logoutButton := widget.NewButton("Logout", func() {
		showLogin()
	})

	return container.NewVBox(
		dashboardButton,
		userButton,
		logsButton,
		todosButton,
		layout.NewSpacer(),
		logoutButton,
	)
}
