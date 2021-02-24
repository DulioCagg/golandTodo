package server

import (
	"github.com/DulioCagg/interview/domain"
	"github.com/DulioCagg/interview/errors"

	_ "github.com/go-sql-driver/mysql"
)

// Fetches all the todos from the database at once
func (s *server) GetAllTodos(todo *[]domain.Todo) *errors.RestErr {
	if err := s.DB.Find(todo).Error; err != nil {
		return errors.NotFound(err.Error())
	}
	return nil
}

// Fetches
func (s *server) CreateTodo(todo *domain.Todo) *errors.RestErr {
	if err := todo.Validate(); err != nil {
		return err
	}
	todo.Status = true
	if err := s.DB.Create(todo).Error; err != nil {
		return errors.NotFound(err.Error())
	}
	return nil
}

func (s *server) GetTodo(todo *domain.Todo, id string) *errors.RestErr {
	if err := s.DB.Where("id = ?", id).Find(todo).Error; err != nil {
		return errors.NotFound(err.Error())
	}
	return nil
}

func (s *server) UpdateTodo(todo *domain.Todo) *errors.RestErr {
	if err := todo.Validate(); err != nil {
		return err
	}
	s.DB.Save(todo)
	return nil
}
func (s *server) DeleteTodo(todo *domain.Todo, id string) *errors.RestErr {
	s.DB.Where("id = ?", id).Delete(todo)
	return nil
}
