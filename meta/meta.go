package meta

import (
	"github.com/caicloud/nirvana/errors"
)

var (
	TableQueryError       = errors.InternalServerError.Build("Database:TableQueryError", "table ${tableName} query error.")
	TableInsertError      = errors.InternalServerError.Build("Database:TableInsertError", "table ${tableName} insert error.")
	TableUpdateError      = errors.InternalServerError.Build("Database:TableUpdateError", "table ${tableName} update error.")
	UnexpectedParamError  = errors.BadRequest.Build("Params:UnexpectedParamError", "unexpected param: ${paramName}.")
	PasswordNotMatchError = errors.Forbidden.Build("Login:PasswordNotMatchError", "password not match account.")
)
