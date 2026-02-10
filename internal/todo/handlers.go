/**
 * Handler functions for processing user input and interactions.
 */

package todo

// --- IMPORTS ---
import (
	"log"

	"github.com/eduardokairalla1/todo-tui/internal/storage"

	tea "github.com/charmbracelet/bubbletea"
)

// --- CODE ---

/**
 * Handles the quit key input.
 *
 * It takes the current model as input and returns the same model
 * along with a tea.Cmd that signals the application to quit.
 */
func handleQuitKey(m Model) (Model, tea.Cmd) {
	return m, tea.Quit
}

/**
 * Handles the add key input to start adding a new task.
 *
 * It takes the current model as input, sets the Adding flag to true,
 * and returns the updated model with no command.
 */
func handleAddKey(m Model) (Model, tea.Cmd) {
	m.Adding = true
	m.Input = ""
	return m, nil
}

/**
 * Handles the input for adding a new task.
 *
 * It takes the current model and the key message as input, processes the key
 * input to update the Input field of the model, and returns the updated model
 * along with any command if necessary (e.g., when the user presses enter to
 * create a new task).
 */
func handleAddingInput(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {

	// process the key input based on the key pressed
	switch msg.String() {

	// escape key: cancel adding a new task, reset the input and exit
	// adding mode
	case "esc":
		m.Input = ""
		m.Adding = false

	// enter key: create a new task with the current input, add it to the
	// task list, and exit adding mode
	case "enter":

		// input is not empty: create a new task
		if m.Input != "" {

			// create a new task in the database
			task, err := m.Store.CreateTask(m.Input)

			// error in creating task: log error and do not update the task list
			if err != nil {
				log.Printf("Error while creating task: %v", err)

				// no error: add the new task and reset the cursor
			} else {
				m.Tasks = append([]storage.Task{task}, m.Tasks...)
				m.Cursor = 0
			}
		}

		// exit adding mode and reset the input
		m.Adding = false
		m.Input = ""

	// backspace key: remove the last character from the input if it's not empty
	case "backspace":
		if len(m.Input) > 0 {
			m.Input = m.Input[:len(m.Input)-1]
		}

	// space key: add a space character to the input
	case " ":
		m.Input += " "

	// other keys: add the character to the input
	default:
		if msg.Type == tea.KeyRunes {
			m.Input += msg.String()
		}
	}

	// return the updated model and no command
	return m, nil
}
