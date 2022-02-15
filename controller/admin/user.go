package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/gink/internal/util"
	"github.com/yeyudekuangxiang/gink/service"
)

var DefaultUserController = UserController{}

type UserController struct {
}

func (UserController) GetUserInfo(c *gin.Context) (gin.H, error) {
	var form GetUserInfoForm
	if err := util.BindForm(c, &form); err != nil {
		return nil, err
	}
	user, err := service.DefaultUserService.GetUserById(form.ID)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"user": user,
	}, nil
}
