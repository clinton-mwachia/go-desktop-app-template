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

func TodosView(window fyne.Window, userID primitive.ObjectID) fyne.CanvasObject {
	var todoList *widget.List
	var todos []models.Todo

	updateTodoList := func() {
		todos = loadTodosByUserID(window, userID)
		todoList.Refresh()
	}

	titleRow := container.NewHBox(
		widget.NewLabelWithStyle("Title", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Content", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Actions", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	todoList = widget.NewList(
		func() int { return len(todos) },
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

	addTodoButton := widget.NewButton("Add Todo", func() {
		showTodoForm(window, nil, userID, updateTodoList)
	})
	// Define the container for the list
	listContainer := container.NewBorder(titleRow, nil, nil, nil, todoList)

	// Use a container to make the list responsive
	listWrapper := container.NewBorder(addTodoButton, nil, nil, nil, listContainer)

	updateTodoList()

	return listWrapper
}

// Function to load todos by UserID
func loadTodosByUserID(window fyne.Window, userID primitive.ObjectID) []models.Todo {
	return utils.GetTodosByUserID(userID, window)
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
			} else {
				todo.ID = primitive.NewObjectID()
				todo.UserID = UserID // Ensure UserID is set for new todos
				utils.AddTodo(todo, window)
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
