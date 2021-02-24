package domain

import (
	"github.com/DulioCagg/interview/errors"
	"strings"
)

type Todo struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// Validate checks that the Title and Description are valid strings
func (todo Todo) Validate() *errors.RestErr {
	todo.Title = strings.TrimSpace(todo.Title)
	todo.Description = strings.TrimSpace(todo.Description)

	if todo.Title == "" {
		return errors.BadRequest("Can't have empty title")
	}
	if todo.Description == "" {
		return errors.BadRequest("Can't have empty descirption")
	}

	return nil
}
