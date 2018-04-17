package services

import "github.com/jinzhu/gorm"

type BaseService struct {
	DB *gorm.DB
}

func NewBaseService() *BaseService {
	return &BaseService{
		DB: db,
	}
}
