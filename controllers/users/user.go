package users

import (
	"context"
	"strings"

	"github.com/narrowizard/nirvana-cms-auth/meta"
	"github.com/narrowizard/nirvana-cms-auth/services"

	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/service"
)

func Login(ctx context.Context, account, password, ip string) (map[string]string, error) {
	if ip == "" {
		var httpCtx = service.HTTPContextFrom(ctx)
		ip = httpCtx.Request().RemoteAddr
	} else {
		var ips = strings.Split(ip, ",")
		ip = ips[0]
	}
	var us = services.NewUserService()
	var uid, err = us.Login(account, password, ip)
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

func Authorize(ctx context.Context, request, ip string) (int, error) {
	var uid, succ = userID(ctx)
	if !succ {
		return 0, meta.UnLoginError.Error()
	}
	var us = services.NewUserService()
	var err = us.CheckURL(uid, request)
	if err != nil {
		return 0, err
	}
	if ip == "" {
		var httpCtx = service.HTTPContextFrom(ctx)
		ip = httpCtx.Request().RemoteAddr
	} else {
		var ips = strings.Split(ip, ",")
		ip = ips[0]
	}
	// 增加日志
	go us.URLCheckLog(uid, request, ip)
	return uid, nil
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

func UpdatePassword(ctx context.Context, oldpwd, newpwd string) error {
	var uid, succ = userID(ctx)
	if !succ {
		return meta.UnLoginError.Error()
	}
	var us = services.NewUserService()
	var err = us.UpdatePassword(uid, oldpwd, newpwd)
	if err != nil {
		return err
	}
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
