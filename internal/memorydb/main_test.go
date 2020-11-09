package memorydb_test

import (
	"testing"
	"todo/internal/memorydb"
	"todo/internal/models"
)

func TestGetTodo(t *testing.T) {
	db := memorydb.New()
	expected := models.Todo{ID: "1", Description: "clean the kitchen", Toggled: true}
	got, _ := db.GetTodoByID(expected.ID)
	if expected.ID != got.ID {
		t.Errorf("Expected ID: %s, got ID: %s", expected.ID, got.ID)
	}
	if expected.Description != got.Description {
		t.Errorf("Expected DESC: %s, got Desc: %s", expected.Description, got.Description)
	}
	if expected.Toggled != got.Toggled {
		t.Errorf("Expected Toggled: %t, got ID: %t", expected.Toggled, got.Toggled)
	}
}

func TestCreateTodo(t *testing.T) {
	db := memorydb.New()
	expected := models.Todo{ID: "8", Description: "new todo", Toggled: false}
	db.CreateTodo(&expected)
	got, _ := db.GetTodoByID(expected.ID)
	if expected.ID != got.ID {
		t.Errorf("Expected ID: %s, got ID: %s", expected.ID, got.ID)
	}
	if expected.Description != got.Description {
		t.Errorf("Expected DESC: %s, got Desc: %s", expected.Description, got.Description)
	}
	if expected.Toggled != got.Toggled {
		t.Errorf("Expected Toggled: %t, got ID: %t", expected.Toggled, got.Toggled)
	}
}

func TestUpdateTodo(t *testing.T) {
	db := memorydb.New()
	expected := models.Todo{ID: "5", Description: "updated todo", Toggled: false}
	db.UpdateTodo(&expected)
	got, _ := db.GetTodoByID(expected.ID)
	if expected.ID != got.ID {
		t.Errorf("Expected ID: %s, got ID: %s", expected.ID, got.ID)
	}
	if expected.Description != got.Description {
		t.Errorf("Expected DESC: %s, got Desc: %s", expected.Description, got.Description)
	}
	if expected.Toggled != got.Toggled {
		t.Errorf("Expected Toggled: %t, got ID: %t", expected.Toggled, got.Toggled)
	}
}

func TestDeleteTodoByID(t *testing.T) {
	db := memorydb.New()
	db.DeleteTodoByID("2")
	got, _ := db.GetTodoByID("2")
	var expected *models.Todo
	if got != expected {
		t.Fail()
	}
}
