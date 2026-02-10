/**
 * View rendering logic.
 */

package todo

// --- IMPORTS ---
import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// --- GLOBALS ---
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED")).
			MarginBottom(1)

	taskStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED"))

	completedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Strikethrough(true)

	checkDone = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Render("[x]")

	checkPending = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Render("[ ]")

	cursorChar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED")).
			Bold(true).
			Render(">")

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED")).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			MarginTop(1)

	emptyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Italic(true)
)

// --- CODE ---

/**
 * Renders the view of the application based on the current state of the model.
 *
 * It returns a string representation of the view to be displayed in the
 * terminal.
 */
func (m Model) View() string {

	// start with the title of the application
	s := titleStyle.Render("Todo") + "\n"

	// adding new task mode is active: show the input field and help text
	if m.Adding {
		s += "\n"
		s += inputStyle.Render("New task:") + " " + m.Input + "\n"
		s += helpStyle.Render("enter confirm  |  esc cancel") + "\n"
		return s
	}

	// no tasks: show the empty state message
	if len(m.Tasks) == 0 {
		s += emptyStyle.Render("No tasks yet. Press 'a' to add one.") + "\n"

		// tasks exist: show the list of tasks with their status and cursor
	} else {

		// iterate over the tasks and render each one with its status and cursor
		for i, task := range m.Tasks {

			// determine the checkmark and title style based on the task's
			// completion status
			check := checkPending
			title := task.Title

			// task is completed: use the done checkmark and apply the
			if task.Completed {
				check = checkDone
				title = completedStyle.Render(title)
			}

			// format the line with the checkmark and title
			line := fmt.Sprintf("%s %s", check, title)

			// task is selected: render with the cursor and selected style
			if m.Cursor == i {
				s += fmt.Sprintf(
					" %s %s\n",
					cursorChar,
					selectedStyle.Render(line),
				)

				// task is not selected: render with the normal task style
			} else {
				s += taskStyle.Render(fmt.Sprintf("   %s", line)) + "\n"
			}
		}
	}

	// show the help text for the main view
	s += helpStyle.Render(
		"a add  |  d delete  |  enter/space toggle  |  q quit") + "\n"

	// return the complete view string
	return s
}
