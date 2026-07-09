package gorm

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := User{Fullname: fullname, Age: age, Email: email}

	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	result := db.Create(&user)
	return result.Error
}
