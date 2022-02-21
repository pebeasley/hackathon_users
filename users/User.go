package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"-"`
	Email     string `json:"email"`
}
