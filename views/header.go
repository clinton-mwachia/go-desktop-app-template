package views

import (
	"desktop-app-template/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	notificationCountLabel *widget.Label
	notificationIcon       *widget.Button
	settingsIcon           *widget.Button
	notifications          []models.Notification
)

func Header(window fyne.Window) *fyne.Container {
	// Notification icon button with initial count
	notificationCountLabel = widget.NewLabel("0")
	notificationIcon = widget.NewButtonWithIcon("", theme.MailComposeIcon(), func() {
		showNotifications(window)
	})

	// settings icon
	settingsIcon = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		showSettings(window)
	})

	// Set initial count
	updateNotificationCount(window)

	// Header container
	header := container.NewHBox(
		widget.NewLabel("Go Template"),
		layout.NewSpacer(),
		settingsIcon,
		notificationIcon,
		notificationCountLabel,
	)

	return header
}
