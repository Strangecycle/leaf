package handler

import (
	context "context"
	"leaf/srv-user/proto/out/user"
	"leaf/srv-user/service"
)

type User struct{}

func (u User) UserLogin(ctx context.Context, request *user.LoginRequest, response *user.LoginResponse) error {
	*response = service.UserLogin(request)
	return nil
}
