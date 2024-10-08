package views

import (
	"desktop-app-template/auth"
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func RegisterView(window fyne.Window, showDashboard func()) *fyne.Container {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	roleEntry := widget.NewEntry()
	roleEntry.SetPlaceHolder("Role (admin/user)")

	registerButton := widget.NewButton("Register", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		role := roleEntry.Text

		if username == "" || password == "" || role == "" {
			dialog.ShowInformation("User Register", "All fields are required", window)
		}

		err := auth.Register(username, password, role)
		if err != nil {
			dialog.ShowInformation("User Register", "Cannot Register account", window)
		} else {
			dialog.ShowInformation("Registration Successful", "Please login, "+username, window)
			window.SetContent(LoginView(window, showDashboard))
		}
	})

	loginButton := widget.NewButton("Login", func() {
		window.SetContent(LoginView(window, showDashboard))
	})

	form := container.NewVBox(
		usernameEntry,
		passwordEntry,
		roleEntry,
		registerButton,
		loginButton,
	)

	centeredForm := utils.NewFixedWidthCenter(form, 300)

	return container.NewCenter(centeredForm)
}
