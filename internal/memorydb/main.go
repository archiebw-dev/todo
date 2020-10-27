package memorydb

import (
	"todo/internal/models"
)

// TodoDB is an in memory database
type TodoDB struct {
	todos map[int]models.Todo
}

// Initialise creates an in memory map of Todo's
func (t *TodoDB) Initialise() {
	if t.todos == nil {
		t.todos = map[int]models.Todo{
			1: {ID: 1, Description: "clean the kitchen"},
			2: {ID: 2, Description: "feed Goblin"},
			3: {ID: 3, Description: "put the bins out"},
			4: {ID: 4, Description: "watch GO tutorial"},
			5: {ID: 5, Description: "go climbing"},
			6: {ID: 6, Description: "do a lunch workout"},
			7: {ID: 7, Description: "clean the kitchen"},
		}
	}
}

// Create attempts to add a todo into the in memory db
func (t *TodoDB) Create(todo models.Todo) (err bool) {
	if _, present := t.todos[todo.ID]; present == true {
		return true
	}
	t.todos[todo.ID] = todo
	return false
}

// Read returns the Todo from the DB
func (t *TodoDB) Read(id int) (todo models.Todo, present bool) {
	todo, present = t.todos[id]
	return
}

// Update adds/updates a todo into the in memory db
func (t *TodoDB) Update(todo models.Todo) {
	t.todos[todo.ID] = todo
}

// Delete removes a Todo from the db
func (t *TodoDB) Delete(id int) {
	delete(t.todos, id)
}

// String returns a string to pretty print the collection of Todos.
func (t *TodoDB) String() string {
	var result string
	for _, v := range t.todos {
		result += v.String()
	}
	return result
}
