package main

import "github.com/labstack/echo"

func (app *App) getRoutes(e *echo.Echo) {
	// Activity Routes
	e.GET("/activity-groups", app.getAllActivitiesHandler)
	e.GET("/activity-groups/:id", app.getActivityHandler)
	e.POST("/activity-groups", app.createActivityHandler)
	e.PATCH("/activity-groups/:id", app.updateActivityHandler)
	e.DELETE("/activity-groups/:id", app.deleteActivityHandler)

	// Todo Routes
	e.GET("/todo-items", app.getAllTodosHandler)
	e.GET("/todo-items/:id", app.getTodoHandler)
	e.POST("/todo-items", app.createTodoHandler)
	e.PATCH("/todo-items/:id", app.updateTodoHandler)
	e.DELETE("/todo-items/:id", app.deleteTodoHandler)
}
