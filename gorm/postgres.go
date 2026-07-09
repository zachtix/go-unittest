package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost" // or the Docker service name if running in another container
	port     = 5432        // default PostgreSQL port
	user     = "admin"     // as defined in docker-compose.yml
	password = "admin"     // as defined in docker-compose.yml
	dbname   = "app"       // as defined in docker-compose.yml
)

// InitializePostgres opens a PostgreSQL connection and migrates the schema.
func InitializePostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	return db
}
