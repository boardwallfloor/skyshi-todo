package main

import (
	"strings"
	"time"
)

func (app *App) getAllActivities() ([]Activity, error) {
	query := "SELECT * FROM activities"
	rows, err := app.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []Activity{}
	for rows.Next() {
		var (
			id        int
			title     string
			email     string
			createdAt []uint8
			updatedAt []uint8
		)

		err := rows.Scan(&id, &title, &email, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		activity := Activity{
			Activity_id: id,
			Title:       title,
			Email:       email,
			CreatedAt:   parseTime(createdAt),
			UpdatedAt:   parseTime(updatedAt),
		}
		activities = append(activities, activity)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (app *App) getOneActivity(id int) (Activity, error) {
	var activity Activity
	query := "SELECT * FROM activities WHERE activity_id = ?"
	var (
		created_at []uint8
		updated_at []uint8
	)
	err := app.db.QueryRow(query, id).Scan(&activity.Activity_id, &activity.Title, &activity.Email, &created_at, &updated_at)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Activity{}, ErrIdNotFound
		}
		return Activity{}, err
	}

	activity.CreatedAt = parseTime(created_at)
	activity.UpdatedAt = parseTime(updated_at)

	return activity, nil

}

func (app *App) createActivity(title, email string) (Activity, error) {
	activity := Activity{
		Title:     title,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	res, err := app.db.Exec("INSERT INTO activities (title, email) VALUES (?, ?)", activity.Title, activity.Email)
	if err != nil {
		return Activity{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Activity{}, err
	}
	activity.Activity_id = int(id)

	return activity, nil
}

func (app *App) updateActivity(id int, title, email string) (Activity, error) {
	activity := Activity{
		Title:     title,
		Email:     email,
		UpdatedAt: time.Now().UTC(),
	}

	query := "UPDATE activities SET"
	var values []interface{}

	if activity.Title != "" {
		query += " title = ?,"
		values = append(values, activity.Title)
	}

	if activity.Email != "" {
		query += " email = ?,"
		values = append(values, activity.Email)
	}

	if !activity.UpdatedAt.IsZero() {
		query += " updated_at = ?,"
		values = append(values, activity.UpdatedAt)
	}

	query = query[:len(query)-1]

	query += " WHERE activity_id = ?"

	values = append(values, id)

	_, err := app.db.Exec(query, values...)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Activity{}, ErrIdNotFound
		}
		return Activity{}, err
	}

	var (
		created_at []uint8
		updated_at []uint8
	)
	err = app.db.QueryRow("SELECT * FROM activities where activity_id  = ?", id).Scan(
		&activity.Activity_id,
		&activity.Title,
		&activity.Email,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return Activity{}, ErrIdNotFound
		}
		return Activity{}, err
	}

	activity.CreatedAt = parseTime(created_at)
	activity.UpdatedAt = parseTime(updated_at)

	return activity, nil
}

func (app *App) deleteActivity(id int) (Activity, error) {
	res, err := app.db.Exec("DELETE FROM activities WHERE activity_id = ?", id)
	if err != nil {

		return Activity{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {

		return Activity{}, err
	}

	if rowsAffected == 0 {
		return Activity{}, ErrIdNotFound
	}

	return Activity{}, nil
}
