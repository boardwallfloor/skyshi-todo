package main

import "time"

func getAllTodos() ([]Todo, error) {
	todo := []Todo{{
		ID:              1,
		ActivityGroupID: 1,
		Title:           "item 5.4",
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
		UpdatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
	}}
	return todo, nil
}

func getOneTodo(id int) (Todo, error) {
	todo := Todo{
		ID:              1,
		ActivityGroupID: 1,
		Title:           "item 5.4",
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
		UpdatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
	}
	return todo, nil
}

func createTodo(title, activity_id, is_active string) (Todo, error) {
	todo := Todo{
		ID:              1,
		ActivityGroupID: 1,
		Title:           "item 5.4",
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
		UpdatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
	}
	return todo, nil
}

func updateTodo(id int, title, activity_id, is_active string) (Todo, error) {
	todo := Todo{
		ID:              1,
		ActivityGroupID: 1,
		Title:           "item 5.4",
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
		UpdatedAt:       time.Date(2022, 11, 28, 7, 33, 19, 0, time.UTC),
	}
	return todo, nil
}

func deleteTodo(id int) (Todo, error) {
	return Todo{}, nil
}
