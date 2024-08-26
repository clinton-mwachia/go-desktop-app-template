package views

import (
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Sidebar(window fyne.Window, showDashboard, showUsers, showTodos, showLogin func(), userID primitive.ObjectID) *fyne.Container {
	isAdmin := utils.IsAdmin(userID, window)

	dashboardButton := widget.NewButton("Dashboard", func() {
		showDashboard()
	})

	var userButton *widget.Button
	if isAdmin {
		userButton = widget.NewButton("Users", func() {
			showUsers()
		})
	} else {
		userButton = &widget.Button{}
		userButton.Hide()
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
		todosButton,
		// layout.NewSpacer(),
		logoutButton,
	)
}
