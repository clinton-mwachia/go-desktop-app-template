package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// settings
func showSettings(window fyne.Window) {
	dialog.ShowInformation("title", "testing", window)
}
