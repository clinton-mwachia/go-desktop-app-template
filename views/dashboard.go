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
	doneData, dateData := utils.FetchTodoDataForCharts(userID, window)

	// Create statistics boxes
	totalBox := createStatisticsBox("Total Todos", strconv.Itoa(totalTodos))
	completedBox := createStatisticsBox("Completed Todos", strconv.Itoa(completedTodos))
	pendingBox := createStatisticsBox("Pending Todos", strconv.Itoa(pendingTodos))

	// summaries
	doneStats := utils.DisplayDoneStatistics(doneData)
	dateStats := utils.TopFiveMonths(dateData)
	comparisonStats := utils.CompareMonthlyTodoData(dateData)
	mostProductiveStat := utils.MostProductiveMonth(dateData)
	avgTodosPerMonth := utils.CalculateAverageTodosPerMonth(dateData)
	completionRate := utils.CompletionRate(doneData)

	// Layout for the statistics boxes
	statsContainer := container.New(layout.NewGridLayout(3),
		totalBox,
		completedBox,
		pendingBox,
	)

	summariesContainer := container.New(layout.NewGridLayout(3),
		widget.NewLabelWithStyle(doneStats, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(dateStats, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(comparisonStats, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	anotherContainer := container.New(layout.NewGridLayout(3),
		widget.NewLabelWithStyle(mostProductiveStat, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(avgTodosPerMonth, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(completionRate, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	return container.NewBorder(statsContainer, nil, nil, nil, container.NewVBox(summariesContainer, anotherContainer))
}
