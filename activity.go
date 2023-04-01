package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Activity struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func getAllActivitiesHandler(c echo.Context) error {
	res, err := getAllActivities()
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func getActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := getOneActivity(id)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func createActivityHandler(c echo.Context) error {
	title := c.FormValue("title")
	email := c.FormValue("email")

	res, err := createActivity(title, email)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func updateActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	title := c.FormValue("title")
	email := c.FormValue("email")

	res, err := updateActivity(id, title, email)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(res, "Success", "Success")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}

func deleteActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = deleteActivity(id)
	if err != nil {
		log.Fatal(err)
	}

	resJson, err := wrapResp(nil, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(http.StatusOK, resJson)
}
