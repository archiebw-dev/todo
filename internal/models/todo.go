package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Todo struct
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Todos is an array of *Todo
type Todos []*Todo

// IDString returns the ID as string
func (t *Todo) IDString() string {
	return strconv.Itoa(t.ID)
}

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
