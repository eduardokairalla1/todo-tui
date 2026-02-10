/**
 * Task table operations
 */

package storage

/**
 * Retrieves all tasks from the database,
 * ordered by creation date descending.
 *
 * It returns a slice of Task structs
 * and an error if any occurs during the database query.
 */
func (s *Store) ListTasks() ([]Task, error) {

	// query all tasks ordered by creation date descending
	rows, err := s.db.Query(
		"SELECT id, title, completed FROM tasks ORDER BY created_at DESC",
	)

	// error in querying tasks: return the error
	if err != nil {
		return nil, err
	}

	// ensure rows are closed after reading
	defer rows.Close()

	// read all tasks from the result set
	var tasks []Task

	// iterate over the rows
	for rows.Next() {

		// scan the row into a Task struct
		var t Task

		// temporary variable to hold the integer representation of 'completed'
		var completed int

		// scan the id, title, and completed status from the row
		if err := rows.Scan(&t.Id, &t.Title, &completed); err != nil {
			return nil, err
		}

		// convert the integer 'completed' to a boolean
		t.Completed = completed != 0

		// append the task to the slice
		tasks = append(tasks, t)
	}

	// check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// return the slice of tasks and nil error
	return tasks, nil
}

/**
 * Creates a new task with the given title in the database.
 *
 * It returns the created Task struct
 * and an error if any occurs during the database operation.
 */
func (s *Store) CreateTask(title string) (Task, error) {

	// insert a new task with the given title
	result, err := s.db.Exec("INSERT INTO tasks (title) VALUES (?)", title)

	// error in inserting task: return the error
	if err != nil {
		return Task{}, err
	}

	// get the ID of the newly inserted task
	id, err := result.LastInsertId()

	// error in getting last insert ID: return the error
	if err != nil {
		return Task{}, err
	}

	// return the created task with the new ID, title, and status (false)
	return Task{Id: id, Title: title, Completed: false}, nil
}
