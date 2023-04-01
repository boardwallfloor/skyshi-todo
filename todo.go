package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Todo struct {
	ID              int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Todo Functions
func getAllTodoItems(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get All Todo Items")
}

func getTodoItem(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get One Todo Item")
}

func createTodoItem(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create Todo Item")
}

func updateTodoItem(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update Todo Item")
}

func deleteTodoItem(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete Todo Item")
}
