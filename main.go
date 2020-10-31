package main

import (
	"net/http"
	"strconv"
	"todo/internal/memorydb"
	"todo/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var e *echo.Echo
var todoDB *memorydb.Todos

func main() {
	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Info("Setting up routes...")
	setupRouting(e)
	e.Logger.Info("Initialising DB...")
	todoDB = memorydb.New()
	e.Logger.Fatal(e.Start(":8000"))
}

func setupRouting(e *echo.Echo) {
	e.GET("/todo/", getAllTodo)
	e.GET("/todo/:id", getTodo)
	e.POST("/todo", saveTodo)
	e.PUT("/todo/:id", updateTodo)
	e.DELETE("/todo/:id", deleteTodo)
}

func getAllTodo(c echo.Context) error {
	e.Logger.Info("GET - todo")
	t, err := todoDB.ReadAll()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, t)
}

func getTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	e.Logger.Infof("GET - todo/%d", id)
	t, err := todoDB.Read(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, t)
	}
	return c.JSON(http.StatusOK, t)
}

func saveTodo(c echo.Context) error {
	t := new(models.Todo)
	e.Logger.Infof("POST - /todo %d", t.ID)
	if err := c.Bind(t); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := todoDB.Create(t); err != nil {

	}
	return c.JSON(http.StatusCreated, t)
}

func updateTodo(c echo.Context) error {
	t := new(models.Todo)
	e.Logger.Infof("POST - /todo/%d", t.ID)
	if err := c.Bind(t); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := todoDB.Update(t); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusAccepted, t)
}

func deleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	e.Logger.Infof("DELETE - todo/%d", id)
	err := todoDB.Delete(id)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
