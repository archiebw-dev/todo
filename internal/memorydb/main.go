package memorydb

import (
	"errors"
	"sort"
	"todo/internal/models"
)

// TodosRepository is an in memory map of todos
type TodosRepository struct {
	todos map[int]models.Todo
}

// New creates an in memory map of Todo's
func New() (t *TodosRepository) {
	t = new(TodosRepository)
	t.todos = map[int]models.Todo{
		1: {ID: 1, Description: "clean the kitchen"},
		2: {ID: 2, Description: "feed Goblin"},
		3: {ID: 3, Description: "put the bins out"},
		4: {ID: 4, Description: "watch GO tutorial"},
		5: {ID: 5, Description: "go climbing"},
		6: {ID: 6, Description: "do a lunch workout"},
		7: {ID: 7, Description: "clean the kitchen"},
	}
	return
}

// GetTodoByID returns the Todo from the DB
func (t *TodosRepository) GetTodoByID(id int) (*models.Todo, error) {
	todo, present := t.todos[id]
	if !present {
		return nil, errors.New("item not found")
	}
	return &todo, nil
}

// GetAllTodos returns an array of Todo's from the DB
func (t *TodosRepository) GetAllTodos() (models.Todos, error) {
	var array models.Todos
	m := t.todos
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		val := m[k]
		array = append(array, &val)
	}
	return array, nil
}

// CreateTodo attempts to add a todo into the in memory db
func (t *TodosRepository) CreateTodo(todo *models.Todo) error {
	if _, present := t.todos[todo.ID]; present == true {
		return errors.New("cannot create item as it already exists")
	}
	return nil
}

// UpdateTodoByID adds/updates a todo into the in memory db
func (t *TodosRepository) UpdateTodoByID(todo *models.Todo) error {
	t.todos[todo.ID] = *todo
	return nil
}

// DeleteTodoByID removes a Todo from the db
func (t *TodosRepository) DeleteTodoByID(id int) error {
	delete(t.todos, id)
	return nil
}

// String returns a string to pretty print the collection of Todos.
func (t *TodosRepository) String() string {
	var result string
	for _, v := range t.todos {
		result += v.String()
	}
	return result
}
