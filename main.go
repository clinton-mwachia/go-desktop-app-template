package main

import (
	"desktop-app-template/utils"
	"desktop-app-template/views"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

type themeVariant struct {
	fyne.Theme

	variant fyne.ThemeVariant
}

func (f *themeVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func main() {
	application := app.New()
	window := application.NewWindow("Go desktop app template")
	// connect to DB
	utils.ConnectDB("mongodb://localhost:27017", window)

	// Placeholder for functions that need to reference each other
	var showDashboard, showUsers, showTodos, showLogs, showLogin func()

	// Load the settings on app startup
	settings, err := views.LoadSettings()
	if err != nil {
		dialog.ShowInformation("Loading settings", "Error loading settings: "+err.Error(), window)
	}

	if settings.IsDarkMode {
		fyne.CurrentApp().Settings().SetTheme(&themeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	} else {
		fyne.CurrentApp().Settings().SetTheme(&themeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
	}

	// Function to show the dashboard view
	showDashboard = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showLogs, showTodos, showLogin, utils.CurrentUserID)
		dashboard := views.DashboardView(window)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, dashboard))
	}

	// Function to show the users view
	showUsers = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showLogs, showTodos, showLogin, utils.CurrentUserID)
		users := views.UsersView(window)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, users))
	}

	// Function to show the todos view
	showTodos = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showLogs, showTodos, showLogin, utils.CurrentUserID)
		todos := views.TodosView(window, utils.CurrentUserID)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, todos))
	}

	// Function to show the users view
	showLogs = func() {
		sidebar := views.Sidebar(window, showDashboard, showUsers, showLogs, showTodos, showLogin, utils.CurrentUserID)
		logs := views.LogsView(window)
		window.SetContent(container.NewBorder(nil, nil, sidebar, nil, logs))
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
