package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserLogin struct {
	gorm.Model
	UserID     uint
	Ticket     string
	ExpireTime time.Time
	Status     int // 1-正常 101-无效
}
