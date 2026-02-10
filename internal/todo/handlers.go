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

/**
 * Handles the delete key input to delete the selected task.
 *
 * It takes the current model as input, deletes the selected task from the
 * database and the task list, and returns the updated model with no command.
 */
func handleDeleteKey(m Model) (Model, tea.Cmd) {

	// no tasks: nothing to delete, return the current model
	if len(m.Tasks) == 0 {
		return m, nil
	}

	// get the index of the selected task and delete it from the database
	i := m.Cursor
	err := m.Store.DeleteTask(m.Tasks[i].Id)

	// error in deleting task: log error and do not update the task list
	if err != nil {
		log.Printf("Error while deleting task: %v", err)
		return m, nil
	}

	// no error: remove the task from the task list and adjust the cursor
	m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)

	// cursor is now out of bounds: move it up
	if m.Cursor >= len(m.Tasks) && m.Cursor > 0 {
		m.Cursor--
	}

	// return the updated model and no command
	return m, nil
}

/**
 * Handles the toggle key input to toggle the completion status
 * of the selected task.
 *
 * It takes the current model as input, toggles the completion status of the
 * selected task in the database and updates the task list, and returns the
 * updated model with no command.
 */
func handleToggleKey(m Model) (Model, tea.Cmd) {

	// no tasks: nothing to toggle, return the current model
	if len(m.Tasks) == 0 {
		return m, nil
	}

	// get the selected task, toggle its completion status, and update it in
	// the database
	t := &m.Tasks[m.Cursor]
	t.Completed = !t.Completed
	err := m.Store.UpdateTask(t.Id, t.Completed)

	// error in updating task: log error and revert the completion status
	if err != nil {
		log.Printf("Error while updating task: %v", err)
		t.Completed = !t.Completed
	}

	// return the updated model and no command
	return m, nil
}
