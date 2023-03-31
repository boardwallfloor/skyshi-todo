package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Activity Functions
func getAllActivities(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get All Activities")
}

func getActivity(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get One Activity")
}

func createActivity(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create Activity")
}

func updateActivity(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update Activity")
}

func deleteActivity(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete Activity")
}
