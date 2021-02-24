package server

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter creates and establishes the app routes using the default gin Engine
func (s *server) setupRouter() {
	// Default returns a gin Engine instance with the logger and recovery middleware
	// already attached
	s.Router = gin.Default()

	// Group is used to access all the routes on /v1
	v1 := s.Router.Group("/v1")

	v1.GET("todo", s.GetTodos)
	v1.POST("todo", s.CreateATodo)
	v1.GET("todo/:id", s.GetATodo)
	v1.PUT("todo/:id", s.UpdateATodo)
	v1.DELETE("todo/:id", s.DeleteATodo)
}
