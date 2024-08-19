package views

import (
	"desktop-app-template/models"
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UsersView(window fyne.Window) fyne.CanvasObject {
	var userList *widget.List
	var users []models.User

	updateUserList := func() {
		users = loadUsers(window)
		userList.Refresh()
	}

	// Header titles
	titleRow := container.NewHBox(
		widget.NewLabelWithStyle("Username", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Role", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Actions", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	userList = widget.NewList(
		func() int { return len(users) },
		func() fyne.CanvasObject {
			// Create widgets for each row
			usernameLabel := widget.NewLabel("")
			roleLabel := widget.NewLabel("")
			editButton := widget.NewButton("Edit", nil)
			deleteButton := widget.NewButton("Delete", nil)

			// Use container layout to arrange them
			row := container.NewHBox(
				usernameLabel,
				layout.NewSpacer(),
				roleLabel,
				layout.NewSpacer(),
				container.NewHBox(editButton, deleteButton),
			)
			return row
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			user := users[id]
			row := obj.(*fyne.Container)

			// Retrieve widgets based on their position in the layout
			usernameLabel := row.Objects[0].(*widget.Label)
			roleLabel := row.Objects[2].(*widget.Label)
			actionButtons := row.Objects[4].(*fyne.Container)
			editButton := actionButtons.Objects[0].(*widget.Button)
			deleteButton := actionButtons.Objects[1].(*widget.Button)

			// Set values and event handlers
			usernameLabel.SetText(user.Username)
			roleLabel.SetText(user.Role)

			editButton.OnTapped = func() {
				showUserForm(window, &user, updateUserList)
			}

			deleteButton.OnTapped = func() {
				dialog.ShowConfirm("Delete User", "Are you sure you want to delete this user?",
					func(ok bool) {
						if ok {
							utils.DeleteUser(user.ID, window)
							updateUserList()
						}
					}, window)
			}
		},
	)

	addUserButton := widget.NewButton("Add User", func() {
		showUserForm(window, nil, updateUserList)
	})

	// Initial load of users
	updateUserList()

	return container.NewBorder(addUserButton, nil, nil, nil, container.NewVBox(titleRow, userList))
}

// Function to load users from the database
func loadUsers(window fyne.Window) []models.User {
	return utils.GetAllUsers(window)
}

// Function to display the user form for adding or editing a user
func showUserForm(window fyne.Window, existing *models.User, onSubmit func()) {
	var user models.User
	isEdit := existing != nil
	if isEdit {
		user = *existing
	}

	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	username.SetText(user.Username)

	password := widget.NewEntry()
	password.SetPlaceHolder("Password")
	password.SetText(user.Password)

	role := widget.NewEntry()
	role.SetPlaceHolder("Role")
	role.SetText(user.Role)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: username},
			{Text: "Password", Widget: password},
			{Text: "Role", Widget: role},
		},
		OnSubmit: func() {
			user.Username = username.Text
			user.Password = password.Text
			user.Role = role.Text

			if isEdit {
				utils.UpdateUser(user, window)
			} else {
				user.ID = primitive.NewObjectID()
				utils.AddUser(user, window)
			}

			if onSubmit != nil {
				onSubmit()
			}
		},
	}

	formContainer := container.NewVBox(form)
	formContainer.Resize(fyne.NewSize(400, 250))

	dialog.ShowForm("User Form", "Save", "Cancel", form.Items, func(ok bool) {
		if ok {
			form.OnSubmit()
		}
	}, window)
}
