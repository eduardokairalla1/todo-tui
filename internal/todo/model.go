/**
 * Defines the Model struct.
 */
package todo

// --- IMPORTS ---
import (
	"github.com/eduardokairalla1/todo-tui/internal/storage"
)

// --- TYPES ---
type Model struct {
	Store *storage.Store
	Tasks []storage.Task

	Cursor int

	Adding bool
	Input  string
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
