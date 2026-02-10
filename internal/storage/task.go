/**
 * Defines the Task struct representing a task in the to-do list application.
 */

package storage

// --- TYPES ---
type Task struct {
	Id        int64
	Title     string
	Completed bool
}
