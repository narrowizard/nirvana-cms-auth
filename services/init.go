package services

import (
	"nirvana-cms-auth/models"

	"github.com/kdada/tinygo/session"
	"github.com/narrowizard/tinygoext/sessionext"

	"github.com/jinzhu/gorm"
	"github.com/kdada/tinygo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var sessionContainer session.SessionContainer

var configInfo = models.Config{
	LoginExpire: 1800,
	SessionName: "nirvana_cms_ssid",
}

func init() {
	var cfg, err = config.NewConfig("ini", "./config/db.cfg")
	checkError(err)
	connString, err := cfg.GlobalSection().String("ConnectionString")
	checkError(err)
	db, err = gorm.Open("mysql", connString)
	checkError(err)
	createDatabase()
	// 创建session处理器
	var redisConnString = `{"Host":"10.0.0.11:6379","MaxIdle":10,"MaxActive":20,"IdleTimeout":60,"Wait":false,"DB":2,"Password":""}`
	session.Register("redis", sessionext.NewRediSessionContainer)
	sessionContainer, err = session.NewSessionContainer("redis", configInfo.LoginExpire, redisConnString)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createDatabase() {
}

// ConfigInfo return app config info
func ConfigInfo() models.Config {
	return configInfo
}

// SessionContainer return session container
func SessionContainer() session.SessionContainer {
	return sessionContainer
}
