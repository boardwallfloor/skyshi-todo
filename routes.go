package main

import "github.com/labstack/echo"

func getRoutes(e *echo.Echo) {
	// Activity Routes
	e.GET("/activity-groups", getAllActivities)
	e.GET("/activity-groups/:id", getActivity)
	e.POST("/activity-groups", createActivity)
	e.PATCH("/activity-groups/:id", updateActivity)
	e.DELETE("/activity-groups/:id", deleteActivity)

	// Todo Routes
	e.GET("/todo-items", getAllTodoItems)
	e.GET("/todo-items/:id", getTodoItem)
	e.POST("/todo-items", createTodoItem)
	e.PATCH("/todo-items/:id", updateTodoItem)
	e.DELETE("/todo-items/:id", deleteTodoItem)
}
