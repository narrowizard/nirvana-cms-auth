package services

import (
	"strings"

	"github.com/narrowizard/nirvana-cms-auth/meta"
	"github.com/narrowizard/nirvana-cms-auth/models"

	"github.com/caicloud/nirvana/log"
)

var commonURL map[string]bool

func init() {
	commonURL = make(map[string]bool, 0)
	commonURL["/user/user/menulist"] = true
}

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
	// scan到slice中, 未查询到数据不会报错.
	// scan到struct中, 未查询到数据会报record not found的错误.
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

// UpdatePassword 修改密码
func (this *UserService) UpdatePassword(uid int, oldpwd, newpwd string) error {
	var u []models.User
	var err = this.DB.Table("users").Where("id=? and password=md5(concat(?,salt)) and status<100", uid, oldpwd).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return meta.TableQueryError.Error("users")
	}
	if len(u) == 0 {
		return meta.PasswordNotMatchError.Error()
	}
	if u[0].Status == 2 {
		return meta.UserBlockError.Error()
	}
	u[0].Password = newpwd
	u[0].Encrypt()
	err = this.DB.Table("users").Save(&u[0]).Error
	if err != nil {
		log.Error(err)
		return meta.TableUpdateError.Error("users")
	}
	return nil
}

// CheckURL 校验url权限
func (this *UserService) CheckURL(uid int, url string) error {
	// common url
	if commonURL[strings.ToLower(url)] {
		return nil
	}
	var c struct {
		C int
	}
	var err = this.DB.Raw(`select count(*) c from user_menus um 
											join menus m on um.menu_id = m.id 
											where um.user_id=? and m.url=? and um.status=1 and m.status=1`, uid, url).Scan(&c).Error
	if err != nil {
		log.Error(err)
		return meta.TableQueryError.Error("user_menus")
	}
	if c.C == 0 {
		return meta.UnAuthorizedError.Error()
	}
	return nil
}
