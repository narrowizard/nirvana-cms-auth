package users

import (
	"context"
	"nirvana-cms-auth/meta"
	"nirvana-cms-auth/services"

	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/service"
)

func Login(ctx context.Context, account, password string) (map[string]string, error) {
	var us = services.NewUserService()
	var uid, err = us.Login(account, password)
	if err != nil {
		return nil, err
	}
	var ssc = services.SessionContainer()
	// login always create new session
	var ss, ok = ssc.CreateSession()
	if !ok {
		return nil, meta.SessionCreateError.Error()
	}
	ss.SetInt("UserID", int(uid))
	var headers = make(map[string]string)
	headers["Set-Cookie"] = meta.CreateCookie(services.ConfigInfo().SessionName, ss.SessionId(), services.ConfigInfo().LoginExpire).String()
	return headers, nil
}

func Authorize(ctx context.Context, token string, request string) (int, error) {

	return 0, nil
}

func IsLogin(ctx context.Context, ssid string) (bool, error) {
	var httpContext = service.HTTPContextFrom(ctx)
	var c, err = httpContext.Request().Cookie(services.ConfigInfo().SessionName)
	if err != nil {
		log.Error(err)
		return false, nil
	}
	var ssc = services.SessionContainer()
	var ss, ok = ssc.Session(c.Value)
	if !ok {
		return false, nil
	}
	uid, ok := ss.Int("UserID")
	if !ok || uid <= 0 {
		return false, nil
	}
	return true, nil
}
