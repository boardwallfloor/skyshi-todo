package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Todo struct {
	Todo_id         int       `json:"id"`
	ActivityGroupID *int      `json:"activity_group_id" form:"activity_group_id"`
	Title           string    `json:"title" form:"title"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	Priority        string    `json:"priority" form:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (app *App) getAllTodosHandler(c echo.Context) error {
	var qp string
	if c.QueryParam("activity_group_id") != "" {
		qp = c.QueryParam("activity_group_id")
	}

	res, err := app.getAllTodos(qp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func (app *App) getTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}

	res, err := app.getOneTodo(id)
	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSONBlob(http.StatusNotFound, resJson)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func (app *App) createTodoHandler(c echo.Context) error {
	formData := new(Todo)
	if err := c.Bind(formData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if formData.Title == "" {
		resJson, err := wrapResp(nil, "Bad Request", "title cannot be null")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSONBlob(http.StatusBadRequest, resJson)
	}

	if formData.ActivityGroupID == nil {
		resJson, err := wrapResp(nil, "Bad Request", "activity_group_id cannot be null")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSONBlob(http.StatusBadRequest, resJson)
	}

	res, err := app.createTodo(formData.ActivityGroupID, formData.IsActive, formData.Title, formData.Priority)
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusCreated, resJson)
}

func (app *App) updateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}

	formData := new(Todo)
	err = c.Bind(formData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := app.updateTodo(id, formData.ActivityGroupID, formData.IsActive, formData.Title, formData.Priority)
	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSONBlob(http.StatusNotFound, resJson)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func (app *App) deleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}

	_, err = app.deleteTodo(id)
	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSONBlob(http.StatusNotFound, resJson)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(nil, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, resJson)
}
