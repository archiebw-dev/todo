package main

import (
	"net/http"
	"strconv"
	"todo/internal/memorydb"
	"todo/internal/models"

	"github.com/labstack/echo/v4"
)

var todoDB *memorydb.TodoDB

func main() {
	e := echo.New()
	setupRouting(e)
	initialiseDB()
	e.Logger.Fatal(e.Start(":8000"))
}

func setupRouting(e *echo.Echo) {
	e.GET("/todo/:id", getTodo)
	e.POST("/todo", saveTodo)
	e.PUT("/todo/:id", updateTodo)
	e.DELETE("/todo/:id", deleteTodo)
}

func getTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	t, ok := todoDB.Read(id)
	if !ok {
		return c.JSON(http.StatusNotFound, t)
	}
	return c.JSON(http.StatusCreated, t)
}

func saveTodo(c echo.Context) error {
	t := new(models.Todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}

func updateTodo(c echo.Context) error {
	// User ID from path `todo/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteTodo(c echo.Context) error {
	// User ID from path `todo/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func initialiseDB() {
	todoDB = new(memorydb.TodoDB)
	todoDB.Initialise()
}
