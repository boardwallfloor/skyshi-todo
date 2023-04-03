package main

import (
	"strings"
	"time"
)

func (app *App) getAllTodos(qp string) ([]Todo, error) {
	query := "SELECT * FROM todos"

	if qp != "" {
		query += " WHERE activity_group_id = " + qp
	}
	rows, err := app.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var (
			id                int
			title             string
			activity_group_id int
			is_active         bool
			priority          string
			createdAt         []uint8
			updatedAt         []uint8
		)

		err := rows.Scan(&id, &activity_group_id, &title, &is_active, &priority, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		todo := Todo{
			Todo_id:         id,
			Title:           title,
			ActivityGroupID: &activity_group_id,
			IsActive:        is_active,
			Priority:        priority,
			CreatedAt:       parseTime(createdAt),
			UpdatedAt:       parseTime(updatedAt),
		}
		todos = append(todos, todo)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (app *App) getOneTodo(id int) (Todo, error) {

	var todo Todo
	query := "SELECT * FROM todos WHERE todo_id = ?"
	var (
		created_at []uint8
		updated_at []uint8
	)
	err := app.db.QueryRow(query, id).Scan(&todo.Todo_id, &todo.ActivityGroupID, &todo.Title, &todo.IsActive, &todo.Priority, &created_at, &updated_at)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Todo{}, ErrIdNotFound
		}
		return Todo{}, err
	}

	todo.CreatedAt = parseTime(created_at)
	todo.UpdatedAt = parseTime(updated_at)

	return todo, nil

}

func (app *App) createTodo(activity_id *int, isActive bool, title, priority string) (Todo, error) {

	if priority == "" {
		priority = "very-high"
	}
	if !isActive {
		isActive = true
	}
	todo := Todo{
		ActivityGroupID: activity_id,
		Title:           title,
		IsActive:        isActive,
		Priority:        priority,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	res, err := app.db.Exec("INSERT INTO todos (activity_group_id, title,is_active,priority) VALUES (?, ?,?,?)", todo.ActivityGroupID, todo.Title, todo.IsActive, todo.Priority)
	if err != nil {
		return Todo{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	todo.Todo_id = int(id)

	return todo, nil
}

func (app *App) updateTodo(id int, activity_id *int, isActive bool, title, priority string) (Todo, error) {

	todo := &Todo{

		Title:     title,
		IsActive:  isActive,
		Priority:  priority,
		UpdatedAt: time.Now().UTC(),
	}

	if activity_id != nil {
		todo.ActivityGroupID = activity_id
	}

	query := "UPDATE todos SET"
	var values []interface{}

	if todo.ActivityGroupID != nil {
		query += " activity_group_id = ?,"
		values = append(values, todo.ActivityGroupID)
	}

	if todo.Title != "" {
		query += " title = ?,"
		values = append(values, todo.Title)
	}

	if !todo.IsActive {
		query += " is_active = ?,"
		values = append(values, todo.IsActive)
	}

	if todo.Title != "" {
		query += " updated_at = ?,"
		values = append(values, todo.UpdatedAt)
	}
	if !todo.UpdatedAt.IsZero() {
		query += " updated_at = ?,"
		values = append(values, todo.UpdatedAt)
	}

	query = query[:len(query)-1]

	query += " WHERE todo_id = ?"

	values = append(values, id)

	_, err := app.db.Exec(query, values...)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Todo{}, ErrIdNotFound
		}
		return Todo{}, err
	}

	var (
		created_at []uint8
		updated_at []uint8
	)
	err = app.db.QueryRow("SELECT * FROM todos WHERE todo_id = ?", id).Scan(
		&todo.Todo_id, &todo.ActivityGroupID, &todo.Title, &todo.IsActive, &todo.Priority, &created_at, &updated_at)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Todo{}, ErrIdNotFound
		}
		return Todo{}, err
	}

	todo.CreatedAt = parseTime(created_at)
	todo.UpdatedAt = parseTime(updated_at)

	return *todo, nil
}

func (app *App) deleteTodo(id int) (Todo, error) {
	res, err := app.db.Exec("DELETE FROM todos WHERE todo_id = ?", id)
	if err != nil {

		return Todo{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return Todo{}, err
	}

	if rowsAffected == 0 {
		return Todo{}, ErrIdNotFound
	}

	return Todo{}, nil
}
