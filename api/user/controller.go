package user

import (
	"commonapi"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 用户自定义一些接口
func (u *UserController) Register(c *gin.Context) interface{} {
	return commonapi.OK
}
