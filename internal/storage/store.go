/**
 * Initialize SQLite database
 */

package storage

// --- IMPORTS ---
import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// --- TYPES ---
type Store struct {
	db *sql.DB
}

// --- CODE ---

/**
 * Creates a new Store instance by opening or creating a SQLite database file
 * at the specified path.
 *
 * The function returns a pointer to the Store instance and an error
 * if any occurs during initialization.
 */
func NewStore(path string) (*Store, error) {

	// open or create the SQLite database file
	db, err := sql.Open("sqlite", path)

	// errors while opening/creating the database: return it
	if err != nil {
		return nil, err
	}

	// create the tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed INTEGER NOT NULL DEFAULT 0,
			created_at TEXT NOT NULL DEFAULT (datetime('now'))
		);
	`)

	// errors while creating the table: close the database and return the error
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	// return the Store instance
	return &Store{db: db}, nil
}

/**
 * Closes the database connection.
 *
 * It returns an error if any occurs during closing the database.
 */
func (s *Store) Close() error { return s.db.Close() }
