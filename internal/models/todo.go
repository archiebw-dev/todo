package models

import (
	"encoding/json"
	"fmt"
)

// Todo struct
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (t Todo) String() string {
	return fmt.Sprintf(
		"ID: %d\n"+
			"Description: %q\n",
		t.ID, t.Description)
}

// JSON returns a Todo as JSON
func (t *Todo) JSON() ([]byte, error) {
	return json.Marshal(t)
}
