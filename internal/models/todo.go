package models

import (
	"encoding/json"
	"fmt"
)

// Todo struct
type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Toggled     bool   `json:"toggled"`
}

// Todos is an array of *Todo
type Todos []*Todo

// String return formatted string representation
func (t *Todo) String() string {
	return fmt.Sprintf(
		"ID: %d\n"+
			"Description: %q\n",
		t.ID, t.Description)
}

// JSON returns a Todo as JSON
func (t *Todo) JSON() ([]byte, error) {
	return json.Marshal(t)
}
