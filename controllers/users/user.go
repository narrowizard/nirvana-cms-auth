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
	headers["Set-Cookie"] = meta.CreateCookie(services.ConfigInfo().SessionName, ss.SessionId(), 0).String()
	return headers, nil
}

func Authorize(ctx context.Context, request string) (int, error) {
	var uid, succ = userID(ctx)
	if !succ {
		return 0, meta.UnLoginError.Error()
	}
	var us = services.NewUserService()
	return uid, us.CheckURL(uid, request)
	// return uid, nil
}

func IsLogin(ctx context.Context, ssid string) (bool, error) {
	var _, succ = userID(ctx)
	return succ, nil
}

func Logout(ctx context.Context) error {
	var httpContext = service.HTTPContextFrom(ctx)
	var c, err = httpContext.Request().Cookie(services.ConfigInfo().SessionName)
	if err != nil {
		log.Error(err)
		return nil
	}
	var ssc = services.SessionContainer()
	var ss, ok = ssc.Session(c.Value)
	if !ok {
		return nil
	}
	ss.Die()
	return nil
}

// userID returns user id from session info
func userID(ctx context.Context) (int, bool) {
	var httpContext = service.HTTPContextFrom(ctx)
	var c, err = httpContext.Request().Cookie(services.ConfigInfo().SessionName)
	if err != nil {
		log.Error(err)
		return 0, false
	}
	var ssc = services.SessionContainer()
	var ss, ok = ssc.Session(c.Value)
	if !ok {
		return 0, false
	}
	uid, ok := ss.Int("UserID")
	return uid, ok
}
