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
	e.GET("/todo-items", getAllTodoItems)
	e.GET("/todo-items/:id", getTodoItem)
	e.POST("/todo-items", createTodoItem)
	e.PATCH("/todo-items/:id", updateTodoItem)
	e.DELETE("/todo-items/:id", deleteTodoItem)
}
