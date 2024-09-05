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
	darkModeIcon           *widget.Button
	settingsIcon           *widget.Button
	notifications          []models.Notification
)

// Struct to hold app settings
type AppSettings struct {
	IsDarkMode bool `json:"is_dark_mode"`
}

// Variable to track current theme mode
var isDarkMode bool = false

// Function to apply the theme based on the current mode
func applyTheme() {
	if isDarkMode {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		darkModeIcon.SetIcon(theme.VisibilityIcon())
	} else {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		darkModeIcon.SetIcon(theme.VisibilityOffIcon())
	}
}

// Function to toggle between light and dark mode

func toggleTheme() {
	isDarkMode = !isDarkMode
	applyTheme()
}

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
	var themeIcon = theme.VisibilityIcon()
	if isDarkMode {
		themeIcon = theme.VisibilityOffIcon()
	} else {
		themeIcon = theme.VisibilityIcon()
	}

	// dark mode icon
	darkModeIcon = widget.NewButtonWithIcon("", themeIcon, func() {
		toggleTheme()
	})

	// Set initial count
	updateNotificationCount(window)

	// Header container
	header := container.NewHBox(
		widget.NewLabel("Go Template"),
		layout.NewSpacer(),
		darkModeIcon,
		settingsIcon,
		notificationIcon,
		notificationCountLabel,
	)

	return header
}
