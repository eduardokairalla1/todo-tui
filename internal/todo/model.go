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
