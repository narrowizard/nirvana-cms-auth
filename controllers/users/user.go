package users

import (
	"context"
	"nirvana-cms-auth/services"
)

func Login(ctx context.Context, account, password string) (map[string]string, error) {
	var service = services.NewUserService()
	var token, err = service.Login(account, password)
	var headers = make(map[string]string, 0)
	headers["Set-Cookie"] = "nirvana_cms_auth_ssid=" + token + "; Path=/; HttpOnly"
	return headers, err
}
