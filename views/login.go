package views

import (
	"desktop-app-template/auth"
	"log"

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

        user, err := auth.Login(username, password)
        if err != nil {
            log.Println("Failed to login:", err)
            dialog.ShowError(err, window)
        } else {
            log.Println("User logged in:", user.Username)
            dialog.ShowInformation("Login Successful", "Welcome, "+user.Username, window)
            showDashboard()
        }
    })

    registerButton := widget.NewButton("Register", func() {
        window.SetContent(RegisterView(window, showDashboard))
    })

    return container.NewVBox(
        usernameEntry,
        passwordEntry,
        loginButton,
        registerButton,
    )
}
