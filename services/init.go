package services

import (
	"nirvana-cms-auth/models"

	"github.com/jinzhu/gorm"
	"github.com/kdada/tinygo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

var loginExpire = 1800

func init() {
	var cfg, err = config.NewConfig("ini", "./config/db.cfg")
	checkError(err)
	connString, err := cfg.GlobalSection().String("ConnectionString")
	checkError(err)
	db, err = gorm.Open("mysql", connString)
	checkError(err)
	createDatabase()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createDatabase() {
	db.AutoMigrate(&models.UserLogin{})
}
