package models

import (
	"github.com/jinzhu/gorm"
)

type URLCheckLog struct {
	gorm.Model
	UserID uint
	URL    string
	IP     string
}
