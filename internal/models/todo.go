package models

import "fmt"

// Todo struct
type Todo struct {
	ID          int
	Description string
}

func (t Todo) String() string {
	return fmt.Sprintf(
		"ID: %d\n"+
			"Description: %q\n",
		t.ID, t.Description)
}
