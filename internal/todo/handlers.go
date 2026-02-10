/**
 * Handler functions for processing user input and interactions.
 */

package todo

// --- IMPORTS ---
import (
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
