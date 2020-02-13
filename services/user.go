package services

import (
	"strings"

	"github.com/narrowizard/nirvana-cms-auth/meta"
	"github.com/narrowizard/nirvana-cms-auth/models"
	"github.com/narrowizard/nirvana-cms-auth/utils"

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
func (this *UserService) Login(account, password, ip string) (uint, error) {
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
	// 登录成功 创建登录日志
	var ul models.UserLogin
	ul.UserID = u[0].ID
	ul.IP = ip
	this.DB.Create(&ul)
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
	var err = this.DB.Raw(`select count(*) c from role_menus rm 
											join users u on rm.role_id = u.role_id
											join menus m on rm.menu_id = m.id 
											where u.id=? and m.url=? and rm.status=1 and m.status=1 and u.status=1`, uid, url).Scan(&c).Error
	if err != nil {
		log.Error(err)
		return meta.TableQueryError.Error("user_menus")
	}
	if c.C > 0 {
		return nil
	}
	// 附加校验包含参数的路由
	var menus []models.Menu
	err = this.DB.Raw(`select m.* from role_menus rm
									join users u on rm.role_id = u.role_id
									join menus m on rm.menu_id = m.id
									where u.id=? and rm.status=1 and m.status=1 and u.status=1 and m.url like ?`, uid, `%{%`).Scan(&menus).Error
	if err != nil {
		log.Error(err)
		return meta.TableQueryError.Error("user_menus")
	}
	for _, v := range menus {
		if utils.MatchURL(v.URL, url) {
			return nil
		}
	}
	return meta.UnAuthorizedError.Error()
}

// URLCheckLog url校验日志
func (this *UserService) URLCheckLog(uid int, url, ip string) error {
	var ucl models.URLCheckLog
	ucl.IP = ip
	ucl.URL = url
	ucl.UserID = uint(uid)
	return this.DB.Create(&ucl).Error
}
