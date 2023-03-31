package main

import (
	"net/http"

	"github.com/labstack/echo"
)

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
