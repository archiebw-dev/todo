package repositories

import "todo/internal/models"

// Todo is the interface that a todo repository should conform to.
type Todo interface {
	GetTodoByID(todoID int) (*models.Todo, error)
	GetAllTodos() (models.Todos, error)
	CreateTodo(todo *models.Todo) error
	UpdateTodoByID(todo *models.Todo) error
	DeleteTodoByID(todoID int) error
}
