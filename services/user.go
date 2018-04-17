package services

import (
	"nirvana-cms-auth/meta"
	"nirvana-cms-auth/models"
	"time"

	"github.com/caicloud/nirvana/log"
	"github.com/kaogps/tools/utils"
)

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	return &UserService{
		BaseService: *NewBaseService(),
	}
}

func (this *UserService) Login(account, password string) (string, error) {
	var u []models.User
	var err = this.DB.Table("users").Where("account=? and password=md5(concat(?,salt))", account, password).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return "", meta.TableQueryError.Error("users")
	}
	if len(u) == 0 {
		return "", meta.PasswordNotMatchError.Error()
	}
	// 是否清除上次的token
	var m models.UserLogin
	m.UserID = u[0].ID
	m.ExpireTime = time.Now().Add(time.Second * time.Duration(loginExpire))
	m.Status = 1
	m.Ticket = utils.RandStringBytesMaskImprSrc(32)
	err = this.DB.Model(&models.UserLogin{}).Create(&m).Error
	if err != nil {
		log.Error(err)
		return "", meta.TableInsertError.Error("user_logins")
	}
	return m.Ticket, nil
}
