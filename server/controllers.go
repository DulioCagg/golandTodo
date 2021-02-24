package server

import (
	"github.com/DulioCagg/interview/domain"
	"github.com/DulioCagg/interview/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTodo list of all todos
func (s *server) GetTodos(c *gin.Context) {
	var todo []domain.Todo

	err := s.GetAllTodos(&todo)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if len(todo) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No 'to do' currently in database. Try adding one!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo a new todo entry
func (s *server) CreateATodo(c *gin.Context) {
	var todo domain.Todo

	c.BindJSON(&todo)

	// DB call to create a todo
	// DB.DB.CreateTodo(todo).Error
	err := s.CreateTodo(&todo)
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetTodo a particular Todo with the id
func (s *server) GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo domain.Todo

	err := s.GetTodo(&todo, id)
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateTodo an existing Todo
func (s *server) UpdateATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo domain.Todo

	err := s.GetTodo(&todo, id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if todo.Title == "" {
		err = errors.BadRequest("No record found with that id")
		c.JSON(err.Status, err)
	}

	c.BindJSON(&todo)
	err = s.UpdateTodo(&todo)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo a todo entry based on the id
func (s *server) DeleteATodo(c *gin.Context) {
	var todo domain.Todo

	id := c.Params.ByName("id")
	err := s.DeleteTodo(&todo, id)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record of ID: " + id + " deleted"})

}
