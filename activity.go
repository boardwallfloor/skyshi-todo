package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Activity struct {
	Activity_id int       `json:"id"`
	Title       string    `json:"title" form:"title"`
	Email       string    `json:"email" form:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (app *App) getAllActivitiesHandler(c echo.Context) error {
	res, err := app.getAllActivities()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func (app *App) getActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}

	res, err := app.getOneActivity(id)
	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
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

func (app *App) createActivityHandler(c echo.Context) error {

	formData := new(Activity)
	err := c.Bind(formData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if formData.Title == "" {
		resJson, err := wrapResp(nil, "Bad Request", "title cannot be null")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSONBlob(http.StatusBadRequest, resJson)

	}
	res, err := app.createActivity(formData.Title, formData.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusCreated, resJson)
}

func (app *App) updateActivityHandler(c echo.Context) error {

	formData := new(Activity)
	err := c.Bind(formData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}
	res, err := app.updateActivity(id, formData.Title, formData.Email)

	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
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

func (app *App) deleteActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad ID")
	}

	_, err = app.deleteActivity(id)
	if err != nil {
		if errors.Is(err, ErrIdNotFound) {
			resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
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
