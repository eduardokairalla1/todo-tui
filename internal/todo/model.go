/**
 * Defines the Model struct.
 */
package todo

// --- IMPORTS ---
import (
	"log"

	"github.com/eduardokairalla1/todo-tui/internal/storage"

	tea "github.com/charmbracelet/bubbletea"
)

// --- TYPES ---
type Model struct {
	Store *storage.Store
	Tasks []storage.Task

	Cursor int

	Adding bool
	Input  string
}

type tasksLoadedMsg struct {
	tasks []storage.Task
}

// --- CONSTRUCTOR ---
func NewModel(store *storage.Store) Model {
	return Model{
		Store:  store,
		Tasks:  []storage.Task{},
		Cursor: 0,
		Adding: false,
		Input:  "",
	}
}

// --- METHODS ---

/**
 * Initializes the model by loading tasks from the database.
 *
 * It returns a tea.Cmd that performs the loading of tasks and updates the model
 * with the loaded tasks.
 */
func (m Model) Init() tea.Cmd {

	// define a command to load tasks from the database
	loadTasks := func() tea.Msg {

		// load tasks from the database
		tasks, err := m.Store.ListTasks()

		// error in loading tasks: log error and return an empty slice of tasks
		if err != nil {
			log.Printf("Error loading tasks: %v", err)
			return tasksLoadedMsg{tasks: []storage.Task{}}
		}

		// successfully loaded tasks: return them in a tasksLoadedMsg
		return tasksLoadedMsg{tasks: tasks}
	}

	// return a batch of commands
	return tea.Batch(tea.ClearScreen, loadTasks)
}

/**
 * Updates the model based on the received message.
 *
 * It takes a tea.Msg as input and returns the updated model and an
 * tea.Cmd to perform side effects.
 */
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// get the type of the message and handle it accordingly
	switch msg := msg.(type) {

	// tasks have been loaded from the database: update the model with the
	// loaded tasks
	case tasksLoadedMsg:
		m.Tasks = msg.tasks
		return m, nil

	// key message received: handle user input and interactions
	case tea.KeyMsg:

		// quit keys: run the quit handler to exit the application
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			model, cmd := handleQuitKey(m)
			return model, cmd
		}

		// if adding new task mode is active: handle input for adding a new task
		if m.Adding {
			model, cmd := handleAddingInput(m, msg)
			return model, cmd
		}

		// initialize variables for the updated model and command
		var model Model
		var cmd tea.Cmd

		// handle other keys for navigating and managing tasks
		switch msg.String() {

		// add key: handle adding a new task
		case "a":
			model, cmd = handleAddKey(m)

		// delete key: handle deleting the selected task
		case "d":
			model, cmd = handleDeleteKey(m)

		// toggle keys: handle toggling the completion status of the
		// selected task
		case "enter", " ":
			model, cmd = handleToggleKey(m)

		// cursor movement keys: handle moving the cursor up and down the
		// task list
		case "up", "k":
			model, cmd = handleCursorUp(m)
		case "down", "j":
			model, cmd = handleCursorDown(m)

		// other keys: no action, return the current model and no command
		default:
			return m, nil
		}

		// return the updated model and command
		return model, cmd
	}

	// other message types: no action, return the current model and no command
	return m, nil
}
