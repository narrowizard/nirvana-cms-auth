package services

import (
	"nirvana-cms-auth/meta"
	"nirvana-cms-auth/models"
	"strings"

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
