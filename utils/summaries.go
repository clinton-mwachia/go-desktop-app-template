package utils

import (
	"fmt"
	"sort"
)

func DisplayDoneStatistics(doneData map[string]int) string {
	totalTodos := doneData["true"] + doneData["false"]
	donePercentage := (float64(doneData["true"]) / float64(totalTodos)) * 100
	notDonePercentage := (float64(doneData["false"]) / float64(totalTodos)) * 100

	return fmt.Sprintf(
		"Todo Completion Statistics:\n"+
			"- Completed: %d (%.2f%%)\n"+
			"- Not Completed: %d (%.2f%%)\n"+
			"Total Todos: %d",
		doneData["true"], donePercentage,
		doneData["false"], notDonePercentage,
		totalTodos,
	)
}

func TopFiveMonths(dateData map[string]int) string {
	// Create a slice to store the date and count pairs
	type monthData struct {
		Date  string
		Count int
	}
	months := []monthData{}

	// Populate the slice with data from the map
	for date, count := range dateData {
		months = append(months, monthData{Date: date, Count: count})
	}

	// Sort the slice in descending order by count
	sort.Slice(months, func(i, j int) bool {
		return months[i].Count > months[j].Count
	})

	// Prepare the top 5 months summary
	summary := "Top 5 Months with Most Todos Created:\n"
	for i := 0; i < 5 && i < len(months); i++ {
		summary += fmt.Sprintf("- %s: %d todos\n", months[i].Date, months[i].Count)
	}

	return summary
}

func CompareMonthlyTodoData(dateData map[string]int) string {
	months := []string{}
	for month := range dateData {
		months = append(months, month)
	}
	sort.Strings(months) // Ensure chronological order

	if len(months) < 2 {
		return "Comparison of Todo Creation:\nNot enough data for a meaningful comparison."
	}

	currentMonth := months[len(months)-1]
	previousMonth := months[len(months)-2]

	diff := dateData[currentMonth] - dateData[previousMonth]
	percentageChange := (float64(diff) / float64(dateData[previousMonth])) * 100

	return fmt.Sprintf(
		"Comparison of Todo Creation:\n"+
			"- This Month (%s): %d todos\n"+
			"- Last Month (%s): %d todos\n"+
			"Change: %d todos (%.2f%%)",
		currentMonth, dateData[currentMonth],
		previousMonth, dateData[previousMonth],
		diff, percentageChange,
	)
}

func MostProductiveMonth(dateData map[string]int) string {
	var mostProductive string
	maxTodos := 0
	for date, count := range dateData {
		if count > maxTodos {
			mostProductive = date
			maxTodos = count
		}
	}
	return fmt.Sprintf("Most Productive Month: %s \nwith %d todos", mostProductive, maxTodos)
}

func CalculateAverageTodosPerMonth(dateData map[string]int) string {
	totalMonths := len(dateData)
	if totalMonths == 0 {
		return "Average Todos Per Month: 0 (No data available)"
	}
	totalTodos := 0
	for _, count := range dateData {
		totalTodos += count
	}
	average := float64(totalTodos) / float64(totalMonths)
	return fmt.Sprintf("Average Todos Per Month: %.2f todos \n(across %d months)", average, totalMonths)
}

func CompletionRate(doneData map[string]int) string {
	totalTodos := doneData["true"] + doneData["false"]
	if totalTodos == 0 {
		return "Completion Rate: 0% (No todos found)"
	}
	completionPercentage := (float64(doneData["true"]) / float64(totalTodos)) * 100
	return fmt.Sprintf("Completion Rate: %.2f%% \n(%d out of %d todos completed)", completionPercentage, doneData["true"], totalTodos)
}
