package main

import "github.com/labstack/echo"

func getRoutes(e *echo.Echo) {
	// Activity Routes
	e.GET("/activity-groups", getAllActivitiesHandler)
	e.GET("/activity-groups/:id", getActivityHandler)
	e.POST("/activity-groups", createActivityHandler)
	e.PATCH("/activity-groups/:id", updateActivityHandler)
	e.DELETE("/activity-groups/:id", deleteActivityHandler)

	// Todo Routes
	e.GET("/todo-items", getAllTodosHandler)
	e.GET("/todo-items/:id", getTodoHandler)
	e.POST("/todo-items", createTodoHandler)
	e.PATCH("/todo-items/:id", updateTodoHandler)
	e.DELETE("/todo-items/:id", deleteTodoHandler)
}
