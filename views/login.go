package views

import (
	"desktop-app-template/auth"
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func LoginView(window fyne.Window, showDashboard func()) *fyne.Container {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if username == "" || password == "" {
			dialog.ShowInformation("User Register", "All fields are required", window)
		}

		user, err := auth.Login(username, password)
		if err != nil {
			utils.Logger(username+" wrong password/username", "ERROR", window)
			dialog.ShowInformation("User Login", "Wrong password/username ", window)
		} else {
			detail := user.Username + " Logged in"
			utils.Logger(detail, "SUCCESS", window)
			utils.CurrentUserID = user.ID
			dialog.ShowInformation("Login Successful", "Welcome, "+user.Username, window)
			showDashboard()
		}
	})

	registerButton := widget.NewButton("Register", func() {
		window.SetContent(RegisterView(window, showDashboard))
	})

	form := container.NewVBox(
		usernameEntry,
		passwordEntry,
		loginButton,
		registerButton,
	)

	centeredForm := utils.NewFixedWidthCenter(form, 300) // Set width to 300

	return container.NewCenter(centeredForm)

}
