package views

import (
	"desktop-app-template/auth"
	"desktop-app-template/models"
	"desktop-app-template/utils"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			fmt.Println(err)
			dialog.ShowError(err, window)
		} else {
			log.Println("User logged in:", user.Username)
			parsedTime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			myLog := models.Log{
				ID:        primitive.NewObjectID(),
				Timestamp: parsedTime,
				Details:   user.Username + " logged in",
				Status:    "SUCCESS",
			}
			utils.AddLog(myLog, window)
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
