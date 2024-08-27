package views

import (
	"desktop-app-template/models"
	"desktop-app-template/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	pageSize = 15 // Number of todos per page
)

func TodosView(window fyne.Window, userID primitive.ObjectID) fyne.CanvasObject {
	var todoList *widget.List
	var todos []models.Todo
	var currentPage int = 1
	var totalTodos int64 = 0
	var pageLabel *widget.Label
	var prevButton, nextButton *widget.Button
	var searchResults []models.Todo
	var searchEntry *widget.Entry
	var noResultsLabel *widget.Label

	header := Header(window)

	// Load todos for the specified page
	loadTodos := func(page int) {
		// Check if search is active
		if searchEntry.Text != "" {
			// Use filtered todos when a search query is active
			todos = searchResults
			totalTodos = int64(len(todos))
		} else {
			// Use all todos for normal pagination
			todos = utils.GetTodosPaginated(page, pageSize, userID, window)
			totalTodos = utils.CountTodos(userID, window)
		}

		todoList.Refresh()

		// Enable or disable pagination buttons based on the current page and total pages
		totalPages := int(math.Ceil(float64(totalTodos) / float64(pageSize)))

		// Update page label
		pageLabel.SetText(fmt.Sprintf("Page %d of %d", currentPage, totalPages))

		prevButton.Disable()
		nextButton.Disable()
		if currentPage > 1 {
			prevButton.Enable()
		}
		if currentPage < totalPages {
			nextButton.Enable()
		}
	}

	// Update visibility of no results label
	updateNoResultsLabel := func() {
		if len(todos) == 0 {
			noResultsLabel.Show()
		} else {
			noResultsLabel.Hide()
		}
	}

	updateTodoList := func() {
		loadTodos(currentPage)
		updateNoResultsLabel()
	}

	// Header Row with Titles
	titleRow := container.NewHBox(
		widget.NewLabelWithStyle("Title", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Content", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Actions", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	// Create the todos list
	todoList = widget.NewList(
		func() int {
			return len(todos)
		},
		func() fyne.CanvasObject {
			titleLabel := widget.NewLabel("")
			contentLabel := widget.NewLabel("")
			editButton := widget.NewButton("Edit", nil)
			deleteButton := widget.NewButton("Delete", nil)

			row := container.NewHBox(
				titleLabel,
				layout.NewSpacer(),
				contentLabel,
				layout.NewSpacer(),
				container.NewHBox(editButton, deleteButton),
			)
			return row
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			todo := todos[id]
			row := obj.(*fyne.Container)

			titleLabel := row.Objects[0].(*widget.Label)
			contentLabel := row.Objects[2].(*widget.Label)
			actionButtons := row.Objects[4].(*fyne.Container)
			editButton := actionButtons.Objects[0].(*widget.Button)
			deleteButton := actionButtons.Objects[1].(*widget.Button)

			titleLabel.SetText(todo.Title)
			contentLabel.SetText(todo.Content)

			editButton.OnTapped = func() {
				showTodoForm(window, &todo, userID, updateTodoList)
			}

			deleteButton.OnTapped = func() {
				dialog.ShowConfirm("Delete Todo", "Are you sure you want to delete this todo?",
					func(ok bool) {
						if ok {
							utils.DeleteTodo(todo.ID, window)
							updateTodoList()

						}
					}, window)
			}
		},
	)

	// Pagination controls
	pagination := container.NewHBox()
	prevButton = widget.NewButton("Previous", func() {
		if currentPage > 1 {
			currentPage--
			updateTodoList()
		}
	})
	nextButton = widget.NewButton("Next", func() {
		if int(math.Ceil(float64(totalTodos)/float64(pageSize))) > currentPage {
			currentPage++
			updateTodoList()
		}
	})

	// Initialize page label
	pageLabel = widget.NewLabel(fmt.Sprintf("Page %d of %d", currentPage, int(math.Ceil(float64(totalTodos)/float64(pageSize)))))

	// Add buttons and label to the pagination container
	pagination.Add(prevButton)
	pagination.Add(pageLabel)
	pagination.Add(nextButton)

	// Center the pagination controls
	pagination = container.NewCenter(pagination)

	addTodoButton := widget.NewButton("Add Todo", func() {
		showTodoForm(window, nil, userID, updateTodoList)
	})

	// Bulk Upload button
	bulkUploadButton := widget.NewButton("Bulk Upload", func() {
		openFileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, window)
					return
				}
				if reader == nil {
					return
				}
				defer reader.Close()

				todos, parseErr := parseCSV(reader.URI().Path(), userID)
				if parseErr != nil {
					dialog.ShowError(parseErr, window)
					return
				}

				utils.BulkInsertTodos(todos, window)
				updateTodoList() // Refresh list after bulk upload
			}, window)
		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		openFileDialog.Show()
	})

	// Search functionality
	searchEntry = widget.NewEntry()
	searchEntry.SetPlaceHolder("Search Todos...")
	searchButton := widget.NewButton("Search Todos", func() {
		searchText := searchEntry.Text
		if searchText != "" {
			searchResults = utils.SearchTodos(searchText, userID, window)
			updateNoResultsLabel()
			currentPage = 1 // Reset to first page of search results
			updateTodoList()
		} else {
			// If search is cleared, reset the pagination and todo list
			searchResults = nil
			currentPage = 1
			updateTodoList()
		}
	})

	// Define functions for exporting data
	exportToCSV := widget.NewButton("export to csv", func() {
		todos := utils.GetTodosByUserID(userID, window)
		file, err := os.Create("todos.csv")
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Write header
		writer.Write([]string{"ID", "Title", "Content", "Done"})

		// Write todo data
		for _, todo := range todos {
			writer.Write([]string{
				todo.ID.Hex(),
				todo.Title,
				todo.Content,
				boolToString(todo.Done), // convert bool to string
			})
		}

		dialog.ShowInformation("Export Successful", "Todos have been exported to todos.csv", window)
	})

	exportToJSON := widget.NewButton("export to json", func() {
		todos := utils.GetTodosByUserID(userID, window) // Adjust as needed for paginated data
		file, err := os.Create("todos.json")
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(todos)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Export Successful", "Todos have been exported to todos.json", window)
	})

	// the search entry and bulk upload button
	searchContainer := container.New(layout.NewGridLayout(3), searchEntry, searchButton, bulkUploadButton)

	// No results label
	noResultsLabel = widget.NewLabel("No results found")
	noResultsLabel.Hide() // Hide by default

	// Load the initial set of todos
	updateTodoList()

	// grid for the add todo and export todos button
	exportButtonContainer := container.New(layout.NewGridLayout(4), addTodoButton, exportToCSV, exportToJSON)

	// Define the container for the list with pagination controls
	listContainer := container.NewBorder(titleRow, nil, nil, nil, todoList, noResultsLabel)

	listWrapper := container.NewBorder(exportButtonContainer, pagination, nil, nil, listContainer)

	// Return the final container with all elements
	return container.NewBorder(header, nil, nil, nil, container.NewBorder(searchContainer, nil, nil, nil, listWrapper))
}

// Function to display the todo form for adding or editing a todo
func showTodoForm(window fyne.Window, existing *models.Todo, UserID primitive.ObjectID, onSubmit func()) {
	var todo models.Todo
	isEdit := existing != nil
	if isEdit {
		todo = *existing
	}

	// Initialize form fields
	title := widget.NewEntry()
	title.SetPlaceHolder("Title")
	title.SetText(todo.Title)

	content := widget.NewEntry()
	content.SetPlaceHolder("Content")
	content.SetText(todo.Content)

	done := widget.NewCheck("Done", func(value bool) {
		todo.Done = value
	})
	done.SetChecked(todo.Done)

	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Title", Widget: title},
			{Text: "Content", Widget: content},
			{Text: "Done", Widget: done},
		},
		OnSubmit: func() {
			todo.Title = title.Text
			todo.Content = content.Text
			todo.Done = done.Checked

			if isEdit {
				utils.UpdateTodo(todo, window)
				// Create a new notification
				userID := utils.CurrentUserID
				newNotification := models.Notification{
					UserID:  userID,
					Message: "Todo edited successfully:" + todo.Title,
					IsRead:  false,
				}

				utils.AddNotification(newNotification, window)

				// Update the notification count
				updateNotificationCount(window)

			} else {
				todo.ID = primitive.NewObjectID()
				todo.UserID = UserID // Ensure UserID is set for new todos
				utils.AddTodo(todo, window)
				// Create a new notification
				userID := utils.CurrentUserID
				newNotification := models.Notification{
					UserID:  userID,
					Message: "Todo added successfully:" + todo.Title,
					IsRead:  false,
				}

				utils.AddNotification(newNotification, window)

				// Update the notification count
				updateNotificationCount(window)

			}

			if onSubmit != nil {
				onSubmit()
			}

		},
	}

	// Create a container for the form
	formContainer := container.NewVBox(form)
	formContainer.Resize(fyne.NewSize(400, 250))

	// Show the form dialog
	dialog.ShowForm("Todo Form", "Save", "Cancel", form.Items, func(ok bool) {
		if ok {
			form.OnSubmit() // Call OnSubmit if "Save" is clicked
		}
	}, window)
}

// Function to parse CSV and return a slice of todos
func parseCSV(filePath string, userID primitive.ObjectID) ([]models.Todo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		if len(record) < 3 {
			continue // Skip rows with insufficient columns
		}

		done, _ := strconv.ParseBool(record[2])
		/*if err != nil {
			return nil, fmt.Errorf("invalid done value in CSV: %v", err)
		}*/

		todo := models.Todo{
			ID:      primitive.NewObjectID(), // Generate a new unique ObjectID for each Todo
			Title:   record[0],
			Content: record[1],
			Done:    done,
			UserID:  userID,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Helper function to convert bool to string
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
