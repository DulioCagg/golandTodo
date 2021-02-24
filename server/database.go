package server

import "fmt"

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
		Host:     "my_go_app_db",
		port:     3306,
		User:     "docker",
		DBName:   "todos-docker",
		Password: "password",
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