package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
