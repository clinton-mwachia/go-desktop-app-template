package views

import (
	"desktop-app-template/utils"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Header creates a common header for all views with a notification icon
func Header(window fyne.Window, userID primitive.ObjectID) *fyne.Container {
	notifications := utils.GetNotifications(userID, window)

	unreadCount := 0
	for _, notification := range notifications {
		if !notification.Read {
			unreadCount++
		}
	}

	notificationButton := widget.NewButtonWithIcon("", theme.MailComposeIcon(), func() {
		showNotificationsDialog(window, userID)
	})

	// Show unread notification count if there are any
	if unreadCount > 0 {
		notificationButton.SetText(strconv.Itoa(unreadCount))
	} else {
		notificationButton.SetText("")
	}

	// Create a header with the notification button
	header := container.NewHBox(
		widget.NewLabel("Go Template"),
		layout.NewSpacer(),
		notificationButton,
	)

	return header
}

// showNotificationsDialog displays a dialog with all notifications
func showNotificationsDialog(window fyne.Window, userID primitive.ObjectID) {
	notifications := utils.GetNotifications(userID, window)

	if len(notifications) == 0 {
		dialog.ShowInformation("Notifications", "No notifications available.", window)
		return
	}

	notificationList := widget.NewList(
		func() int { return len(notifications) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(notifications[i].Message)
		},
	)

	clearButton := widget.NewButton("Clear Notifications", func() {
		utils.ClearNotifications(userID, window)
		notificationList.Refresh()
		dialog.ShowInformation("Notifications", "All notifications cleared.", window)
	})

	dialog.ShowCustom("Notifications", "Close", container.NewBorder(notificationList, nil, nil, nil, clearButton), window)
}
