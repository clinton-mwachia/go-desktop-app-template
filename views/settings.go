package views

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// settings
func showSettings(window fyne.Window) {
	// Load image from the assets directory
	imagePath := filepath.Join("assets", "profile.jpg")
	imageFile := canvas.NewImageFromFile(imagePath)
	imageFile.FillMode = canvas.ImageFillContain
	imageFile.SetMinSize(fyne.NewSize(200, 200))

	appSettings := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Theme:"),
			widget.NewLabel("My theme"),
		),
		// Add your theme selection widget here
		container.NewHBox(
			widget.NewLabel("Font:"),
			widget.NewLabel("My Font"),
		),
		// Add your font size adjustment widget here
	)

	content := container.NewHBox(
		container.NewVBox(
			imageFile,
			widget.NewButton("Update User Details", func() {
				showUpdateUserDetailsDialog(window)
			}),
			widget.NewButton("Change Password", func() {
				showChangePasswordDialog(window)
			}),
		),
		appSettings,
	)

	// Center the content and set the size
	centeredContent := container.NewCenter(content)
	containerWithWidth := container.New(layout.NewStackLayout(), centeredContent)
	containerWithWidth.Resize(fyne.NewSize(800, 500)) // Adjust the width as needed

	dialog.ShowCustom("Settings", "Close", content, window)
}

// update user details
func showUpdateUserDetailsDialog(window fyne.Window) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: usernameEntry},
			{Text: "Email", Widget: emailEntry},
		},
		OnSubmit: func() {
			info := fmt.Sprintf(
				"Username: %s\nEmail: %s:\n",
				usernameEntry.Text,
				emailEntry.Text,
			)
			dialog.ShowInformation("Profile Information", info, window)
		},
	}

	dialog.ShowForm("Update User Details", "Save", "Cancel", form.Items, nil, window)
}

// update password
func showChangePasswordDialog(window fyne.Window) {
	oldPassword := widget.NewPasswordEntry()
	oldPassword.SetPlaceHolder("Old Password")

	newPassword := widget.NewPasswordEntry()
	newPassword.SetPlaceHolder("New Password")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Old Password", Widget: oldPassword},
			{Text: "New Password", Widget: newPassword},
		},
		OnSubmit: func() {
			info := fmt.Sprintf(
				"oldPassword: %s\nnewPassword: %s:\n",
				oldPassword.Text,
				newPassword.Text,
			)
			dialog.ShowInformation("Password", info, window)
		},
	}

	dialog.ShowForm("Change Password", "Save", "Cancel", form.Items, nil, window)
}
