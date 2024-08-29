package views

import (
	"desktop-app-template/utils"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// createStatisticsBox creates a statistics display box
func createStatisticsBox(title, value string) fyne.CanvasObject {
	// Create a border
	border := canvas.NewRectangle(color.Gray{})
	border.StrokeWidth = 2

	return container.NewVBox(
		widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(value, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

}

func DashboardView(window fyne.Window) *fyne.Container {
	userID := utils.CurrentUserID

	// Fetch statistics and data from the database
	totalTodos, completedTodos, pendingTodos := utils.FetchTodoStatistics(userID, window)
	//statusData, creationData := utils.FetchTodoDataForCharts(userID, window)

	// Create statistics boxes
	totalBox := createStatisticsBox("Total Todos", strconv.Itoa(totalTodos))
	completedBox := createStatisticsBox("Completed Todos", strconv.Itoa(completedTodos))
	pendingBox := createStatisticsBox("Pending Todos", strconv.Itoa(pendingTodos))

	// Layout for the statistics boxes
	statsContainer := container.New(layout.NewGridLayout(3),
		totalBox,
		completedBox,
		pendingBox,
	)

	return container.NewStack(statsContainer)
}
