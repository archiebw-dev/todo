package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"todo/internal/firestoredb"
	"todo/internal/memorydb"
	"todo/internal/models"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var e *echo.Echo
var c *context.Context
var f *firestore.Client
var todoDB *memorydb.Todos

func main() {
	setupEcho()
	setupDB()
	setupFirestore()
	firestoredb.Add(*c, f)
	e.Logger.Info("GAC: " + os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	e.Logger.Fatal(e.Start(":8000"))
}

func setupEcho() {
	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Info("Setting up routes...")
	setupRouting(e)
	e.Logger.Info("Initialising DB...")
}

func setupDB() {
	todoDB = memorydb.New()
}

func setupFirestore() {
	ctx := context.Background()
	c = &ctx
	f = firestoredb.CreateClient(ctx, "archiebw-todo")
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
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, t)
}

func getTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	e.Logger.Infof("GET - todo/%d", id)
	t, err := todoDB.Read(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, t)
}

func saveTodo(c echo.Context) error {
	t := new(models.Todo)
	e.Logger.Infof("POST - /todo %d", t.ID)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := todoDB.Create(t); err != nil {

	}
	return c.JSON(http.StatusCreated, t)
}

func updateTodo(c echo.Context) error {
	t := new(models.Todo)
	e.Logger.Infof("POST - /todo/%d", t.ID)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := todoDB.Update(t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusAccepted, t)
}

func deleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	e.Logger.Infof("DELETE - todo/%d", id)
	err := todoDB.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
