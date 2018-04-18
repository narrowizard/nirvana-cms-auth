package services

import (
	"nirvana-cms-auth/meta"
	"nirvana-cms-auth/models"

	"github.com/caicloud/nirvana/log"
)

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	return &UserService{
		BaseService: *NewBaseService(),
	}
}

// Login returns user id
func (this *UserService) Login(account, password string) (uint, error) {
	var u []models.User
	var err = this.DB.Table("users").Where("account=? and password=md5(concat(?,salt)) and status<100", account, password).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return 0, meta.TableQueryError.Error("users")
	}
	if len(u) == 0 {
		return 0, meta.PasswordNotMatchError.Error()
	}
	if u[0].Status == 2 {
		return 0, meta.UserForbiddenError.Error(account)
	}
	return u[0].ID, nil
}
