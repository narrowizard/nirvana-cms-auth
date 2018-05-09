package meta

import (
	"net/http"
	"time"

	"github.com/caicloud/nirvana/errors"
)

var (
	TableQueryError       = errors.InternalServerError.Build("Database:TableQueryError", "table ${tableName} query error.")
	TableInsertError      = errors.InternalServerError.Build("Database:TableInsertError", "table ${tableName} insert error.")
	TableUpdateError      = errors.InternalServerError.Build("Database:TableUpdateError", "table ${tableName} update error.")
	UnexpectedParamError  = errors.BadRequest.Build("Params:UnexpectedParamError", "unexpected param: ${paramName}.")
	PasswordNotMatchError = errors.Forbidden.Build("Login:PasswordNotMatchError", "password not match account.")
	UserForbiddenError    = errors.Forbidden.Build("Login:UserForbiddenError", "user ${account} is forbidden to login.")
	UserBlockError        = errors.Forbidden.Build("User:UserBlockError", "you are blocked.")
	UnLoginError          = errors.Forbidden.Build("Login:UnLoginError", "need login.")
	UnAuthorizedError     = errors.Forbidden.Build("Auth:UnAuthorizedError", "you are not allowed to this api.")
	SessionCreateError    = errors.InternalServerError.Build("Session:SessionCreateError", "session create error.")
)

// createCookie 创建cookie
func CreateCookie(name string, id string, expire int) *http.Cookie {
	var cookieValue = new(http.Cookie)
	cookieValue.Name = name
	cookieValue.Value = id
	cookieValue.Path = "/"
	cookieValue.HttpOnly = true
	if expire > 0 {
		cookieValue.MaxAge = expire
		cookieValue.Expires = time.Now().Add(time.Second * time.Duration(expire))
	}
	return cookieValue
}
