package server

import (
	"fmt"
	"os"
)

// dbConfig represents the db configuration
type dbConfig struct {
	Host     string
	port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig returns a new DB configuration with the default settings
func buildDBConfig() *dbConfig {
	dbConfig := dbConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		port:     3306,
		User:     os.Getenv("MYSQL_USER"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
	return &dbConfig
}

// DBURL creates the string to connect to the database based on the fields from the
// configuration
func dbURL(config *dbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.port,
		config.DBName,
	)
}