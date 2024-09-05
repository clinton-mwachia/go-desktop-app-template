package views

import (
	"desktop-app-template/models"
	"encoding/json"
	"log"
	"os"

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

const settingsFilePath = "settings.json"

// LoadSettings loads the app settings from a JSON file
func LoadSettings() (*AppSettings, error) {
	// Check if the settings file exists
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		// If it doesn't exist, return default settings
		return &AppSettings{IsDarkMode: false}, nil
	}

	// Read the settings file
	fileBytes, err := os.ReadFile(settingsFilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the AppSettings struct
	var settings AppSettings
	err = json.Unmarshal(fileBytes, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// SaveSettings saves the app settings to a JSON file
func SaveSettings(settings *AppSettings) error {
	// Marshal the settings struct into JSON format
	fileBytes, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	// Write the JSON data to the settings file
	err = os.WriteFile(settingsFilePath, fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
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

// Function to toggle between light and dark mode and save the settings
func toggleTheme() {
	isDarkMode = !isDarkMode
	applyTheme()

	// Save the current theme setting
	settings := &AppSettings{IsDarkMode: isDarkMode}
	err := SaveSettings(settings)
	if err != nil {
		log.Println("Error saving settings:", err)
	}
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
