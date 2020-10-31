package firestoredb

import (
	"context"
	"strconv"
	"todo/internal/models"

	"cloud.google.com/go/firestore"
	"github.com/fatih/structs"
	"google.golang.org/api/iterator"
)

const collection string = "todos"

// TodosRepository repository
type TodosRepository struct {
	ctx    context.Context
	client *firestore.Client
}

// New creats a client and context to interact with the firestore
func New(projectID string) (*TodosRepository, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	// Close client when done with
	// defer client.Close()
	return &TodosRepository{ctx, client}, nil
}

// Unmarshal takes a todo struct and returns the map[string]interface{} result
func Unmarshal(t *models.Todo) map[string]interface{} {
	return structs.Map(t)
}

// GetTodoByID Return a Todo by ID
func (tr *TodosRepository) GetTodoByID(todoID int) (*models.Todo, error) {
	data, err := tr.client.Collection(collection).Doc(strconv.Itoa(todoID)).Get(tr.ctx)
	if err != nil {
		return nil, err
	}
	var t models.Todo
	data.DataTo(&t)
	return &t, nil
}

// GetAllTodos returns all todos
func (tr *TodosRepository) GetAllTodos() (models.Todos, error) {
	iter := tr.client.Collection(collection).Documents(tr.ctx)
	var result models.Todos
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var t models.Todo
		doc.DataTo(&t)
		result = append(result, &t)
	}
	return result, nil
}

//CreateTodo creates a todo document
func (tr *TodosRepository) CreateTodo(todo *models.Todo) error {
	data := Unmarshal(todo)
	_, err := tr.client.Collection(collection).Doc(todo.IDString()).Set(tr.ctx, data)
	if err != nil {
		return err
	}
	return nil
}

//UpdateTodoByID updates an existing document or creates it if it doesn't exist
func (tr *TodosRepository) UpdateTodoByID(todo *models.Todo) error {
	data := Unmarshal(todo)
	_, err := tr.client.Collection(collection).Doc(todo.IDString()).Set(tr.ctx, data, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodoByID deletes a document by ID
func (tr *TodosRepository) DeleteTodoByID(todoID int) error {
	_, err := tr.client.Collection(collection).Doc(strconv.Itoa(todoID)).Delete(tr.ctx)
	if err != nil {
		return err
	}
	return nil
}
