package models

import "github.com/jinzhu/gorm"

type UserLogin struct {
	gorm.Model
	UserID uint
	IP     string
}
