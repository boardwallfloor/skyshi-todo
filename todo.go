package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func getAllTodosHandler(c echo.Context) error {
	res, err := getAllTodos()
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func getTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := getOneTodo(id)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func createTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	activity_id := c.FormValue("activity_id")
	is_active := c.FormValue("is_active")

	res, err := createTodo(title, activity_id, is_active)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func updateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	title := c.FormValue("title")
	activity_id := c.FormValue("activity_id")
	is_active := c.FormValue("is_active")

	res, err := updateTodo(id, title, activity_id, is_active)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func deleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = deleteTodo(id)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}
