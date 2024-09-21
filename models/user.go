package models

import (
	"github.com/jinzhu/gorm"
)

// User model  
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

// GetAllUsers fetches all users, but now it accepts the DB connection as a parameter
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

// CreateUser creates a new user, and accepts the DB connection as a parameter
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
