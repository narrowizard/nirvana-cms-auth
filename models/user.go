package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account  string
	Password string
	Salt     string
	Status   int
}
