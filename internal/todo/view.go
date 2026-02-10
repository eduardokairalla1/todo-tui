/**
 * View rendering logic.
 */

package todo

// --- IMPORTS ---
import (
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
