package server

import (
	"fmt"
	"github.com/DulioCagg/interview/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

type server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// Creates a new server with an active DB connection and a gin.Engine as a router
func New() *server {
	var s server

	var err error

	for i := 0; i < 12; i++ {
		// Creates a connection to the database
		s.DB, err = gorm.Open("mysql", dbURL(buildDBConfig()))

		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				fmt.Printf("Could not connect to database. Try #%d\n", i)
				time.Sleep(3 * time.Second)
				continue
			}
			if i == 10 {
				log.Fatal("status error: ", err)
			}
		}
	}

	fmt.Println("Connected to the database!")

	// run the migration : todo struct
	// AutoMigrate automatically migrates the schema, to keep the schema updated
	// Only creates tables and missing columns, wont modify existing columns
	s.DB.AutoMigrate(&domain.Todo{})
	// Setup routes
	s.setupRouter()

	return &s
}
