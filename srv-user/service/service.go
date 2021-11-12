package service

import (
	"leaf/srv-user/common"
	"leaf/srv-user/models"
	"leaf/srv-user/proto/out/user"
	"net/http"
)

func UserLogin(req *user.LoginRequest) (resp user.LoginResponse) {
	db := common.GetDB()

	// TODO 测试用验证码
	if req.GetCode() != "1234" {
		resp.Code = http.StatusBadRequest
		resp.Message = "验证码错误"
		return
	}

	var userModel models.User
	err := db.Where("phone = ?", req.GetPhone()).FirstOrCreate(&userModel).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
	}

	resp.Code = http.StatusOK
	resp.Message = common.MessageOk

	return
}
