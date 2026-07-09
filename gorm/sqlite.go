package gorm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// InitializeSQLite opens a SQLite connection (file-based) and migrates the schema.
func InitializeSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	return db
}
