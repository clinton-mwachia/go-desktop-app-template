package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Sidebar(window fyne.Window, showDashboard, showUsers, showTodos, showLogin func()) *fyne.Container {
	dashboardButton := widget.NewButton("Dashboard", func() {
		showDashboard()
	})

	usersButton := widget.NewButton("Users", func() {
		showUsers()
	})

	todosButton := widget.NewButton("Todos", func() {
		showTodos()
	})

	logoutButton := widget.NewButton("Logout", func() {
		showLogin()
	})

	sidebar := container.NewVBox(
		dashboardButton,
		usersButton,
		todosButton,
		// layout.NewSpacer(),
		logoutButton,
	)

	return sidebar
}
