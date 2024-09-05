package views

import (
	"desktop-app-template/auth"
	"desktop-app-template/models"
	"desktop-app-template/utils"
	"errors"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/bcrypt"
)

// showSettings displays the settings view with user details and update options
func showSettings(window fyne.Window) {
	var user models.User

	loadUser := func() {
		userID := utils.CurrentUserID
		user = utils.GetUserByID(userID, window)
	}
	loadUser()
	// Load image from the assets directory
	imagePath := filepath.Join("assets", "profile.jpg")
	imageFile := canvas.NewImageFromFile(imagePath)
	imageFile.FillMode = canvas.ImageFillContain
	imageFile.SetMinSize(fyne.NewSize(200, 200))

	// User details section (Right Side)
	userDetailsContainer := container.NewVBox()
	userDetailsContainer.Resize(fyne.NewSize(300, 200))

	// refresh user details in the UI
	refreshUserDetails := func() {
		loadUser()

		userDetailsContainer.Objects = []fyne.CanvasObject{
			widget.NewLabelWithStyle("Username: "+user.Username, fyne.TextAlignLeading, fyne.TextStyle{}),
			widget.NewLabelWithStyle("Role: "+user.Role, fyne.TextAlignLeading, fyne.TextStyle{}),
		}
		userDetailsContainer.Refresh()
	}

	refreshUserDetails()

	// Variable to track current theme mode

	content := container.NewHBox(
		imageFile,
		container.NewVBox(
			userDetailsContainer,
			container.NewGridWithColumns(3,
				widget.NewButton("Update Details", func() {
					showUpdateUserDetailsDialog(window, user, refreshUserDetails)
				}),
				widget.NewButton("Change Password", func() {
					showChangePasswordDialog(window, user)
				}),
			),
		),
	)

	// Center the content
	centeredContent := container.NewCenter(content)
	containerWithWidth := container.New(layout.NewStackLayout(), centeredContent)
	containerWithWidth.Resize(fyne.NewSize(800, 500))

	dialog.ShowCustom("Settings", "Close", content, window)
}

// showUpdateUserDetailsDialog displays a dialog for updating user details
func showUpdateUserDetailsDialog(window fyne.Window, user models.User, updateUserDetailsInView func()) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter Username")
	usernameEntry.SetText(user.Username)

	// Create a select widget for the role field
	roles := []string{"admin", "user"}
	roleSelect := widget.NewSelect(roles, func(selectedRole string) {
		user.Role = selectedRole
	})
	roleSelect.SetSelected(user.Role)

	formItems := []*widget.FormItem{
		{Text: "Username", Widget: usernameEntry},
		{Text: "Role", Widget: roleSelect},
	}

	form := utils.NewFixedWidthCenter(container.NewVBox(widget.NewForm(formItems...)), 300)

	dialog.ShowCustomConfirm("Update User Details", "Save", "Cancel", container.NewCenter(form), func(ok bool) {
		if !ok {
			return
		}

		user.Username = usernameEntry.Text
		user.Role = roleSelect.Selected

		// Update user details in the database
		utils.UpdateUser(user, window)

		// Update user details in the view
		updateUserDetailsInView()

		dialog.ShowInformation("Success", "User details updated successfully.", window)
	}, window)
}

// showChangePasswordDialog displays a dialog for changing the user's password
func showChangePasswordDialog(window fyne.Window, user models.User) {
	currentPasswordEntry := widget.NewPasswordEntry()
	currentPasswordEntry.SetPlaceHolder("Enter Current Password")

	newPasswordEntry := widget.NewPasswordEntry()
	newPasswordEntry.SetPlaceHolder("Enter New Password")

	confirmPasswordEntry := widget.NewPasswordEntry()
	confirmPasswordEntry.SetPlaceHolder("Confirm New Password")

	formItems := []*widget.FormItem{
		{Text: "Current Password", Widget: currentPasswordEntry},
		{Text: "New Password", Widget: newPasswordEntry},
		{Text: "Confirm Password", Widget: confirmPasswordEntry},
	}

	form := utils.NewFixedWidthCenter(container.NewVBox(widget.NewForm(formItems...)), 400)

	dialog.ShowCustomConfirm("Change Password", "Save", "Cancel", container.NewCenter(form), func(ok bool) {
		if !ok {
			return
		}

		oldPassword := currentPasswordEntry.Text
		newPassword := newPasswordEntry.Text
		confirmPassword := confirmPasswordEntry.Text

		// Verify that old password matches the stored password hash.
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
			dialog.ShowError(errors.New("old password is incorrect"), window)
			return
		}

		// Check if the new password is the same as the old password.
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newPassword)); err == nil {
			dialog.ShowError(errors.New("new password cannot be the same as the old password"), window)
			return
		}

		if newPassword != confirmPassword {
			dialog.ShowError(errors.New("new password and confirm password do not match"), window)
			return
		}

		// Update password in the database
		err := auth.UpdateUserPassword(user.ID, newPassword, window)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Success", "Password changed successfully. Please log in again.", window)

	}, window)
}
