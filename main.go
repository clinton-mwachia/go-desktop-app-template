package main

import (
	"desktop-app-template/utils"
	"desktop-app-template/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	utils.ConnectDB("mongodb://localhost:27017")

	application := app.New()
	window := application.NewWindow("Go desktop app template")

	// Placeholder for functions that need to reference each other
	var showDashboard, showUsers, showTodos, showLogin func()

	// Function to show the dashboard view
	showDashboard = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showTodos, showLogin, utils.CurrentUserID)
		dashboard := views.DashboardView()
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, dashboard))
	}

	// Function to show the users view
	showUsers = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showTodos, showLogin, utils.CurrentUserID)
		users := views.UsersView(window)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, users))
	}

	// Function to show the todos view
	showTodos = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showTodos, showLogin, utils.CurrentUserID)
		todos := views.TodosView(window, utils.CurrentUserID)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, todos))
	}

	// Function to show the login view
	showLogin = func() {
		window.SetContent(views.LoginView(window, showDashboard))
	}

	// Initial view when the application starts
	showLogin()
	window.Resize(fyne.NewSize(400, 300))
	window.CenterOnScreen()
	window.ShowAndRun()
}
