package repositories

import "todo/internal/models"

// Todo is the interface that a todo repository should conform to.
type Todo interface {
	GetTodoByID(todoID string) (*models.Todo, error)
	GetAllTodos() (models.Todos, error)
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	DeleteTodoByID(todoID string) error
}
