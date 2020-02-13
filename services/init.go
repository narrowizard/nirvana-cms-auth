package services

import (
	"github.com/narrowizard/nirvana-cms-auth/models"

	"github.com/kdada/tinygo/session"
	"github.com/narrowizard/tinygoext/sessionext"

	"github.com/jinzhu/gorm"
	"github.com/kdada/tinygo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var sessionContainer session.SessionContainer

var configInfo = models.Config{}

func init() {
	var cfg, err = config.NewConfig("ini", "./config/config.cfg")
	checkError(err)
	connString, err := cfg.GlobalSection().String("ConnectionString")
	checkError(err)
	db, err = gorm.Open("mysql", connString)
	checkError(err)
	redisConnString, err := cfg.GlobalSection().String("RedisConnString")
	checkError(err)
	configInfo.LoginExpire, err = cfg.GlobalSection().Int("LoginExpire")
	checkError(err)
	configInfo.SessionName, err = cfg.GlobalSection().String("SessionName")
	checkError(err)
	port, err := cfg.GlobalSection().Int("Port")
	checkError(err)
	configInfo.Port = uint16(port)
	// 创建session处理器
	session.Register("redis", sessionext.NewRediSessionContainer)
	sessionContainer, err = session.NewSessionContainer("redis", configInfo.LoginExpire, redisConnString)
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

// ConfigInfo return app config info
func ConfigInfo() models.Config {
	return configInfo
}

// SessionContainer return session container
func SessionContainer() session.SessionContainer {
	return sessionContainer
}
