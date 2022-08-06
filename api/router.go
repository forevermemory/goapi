package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func wrap(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}
func InitRouter() *gin.Engine {
	r := gin.Default()

	registerSystemApi(r)

	return r
}

func registerSystemApi(r *gin.Engine) {
	{
		// 用户
		// _user := r.Group("/user")
		// _user.GET("", wrap(user.Entity.List))
		// _user.GET("/:id", wrap(user.Entity.GetByID))
		// _user.POST("", wrap(user.Entity.Add))
		// _user.PUT("/:id", wrap(user.Entity.Update))
		// _user.DELETE("/:id", wrap(user.Entity.Delete))
	}

}
