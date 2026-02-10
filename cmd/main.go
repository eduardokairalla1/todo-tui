/**
 * Application entry point
 */

package main

// --- IMPORTS ---
import (
	"fmt"
	"os"

	"github.com/eduardokairalla1/todo-tui/internal/storage"
	"github.com/eduardokairalla1/todo-tui/internal/todo"

	tea "github.com/charmbracelet/bubbletea"
)

// --- CODE ---
func main() {

	dbPath := "todo.db"

	// initialize storage
	store, err := storage.NewStore(dbPath)

	// error in initializing storage: log and exit
	if err != nil {
		fmt.Println("Error initializing database:", err)
		os.Exit(1)
	}

	// schedule storage to be closed when main function exits
	defer store.Close()

	// initialize application model
	model := todo.NewModel(store)

	// start the TUI application
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Todo:", err)
		os.Exit(1)
	}
}
