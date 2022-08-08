package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"goapi/api/role"
	"goapi/api/user"
)


func wrap(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/role", wrap(role.Entity.List))
	r.GET("/role/count", wrap(role.Entity.Count))
	r.GET("/role/:id", wrap(role.Entity.GetByID))
	r.POST("/role", wrap(role.Entity.Add))
	r.PUT("/role/:id", wrap(role.Entity.Update))
	r.DELETE("/role/:id", wrap(role.Entity.Delete))

	r.GET("/user", wrap(user.Entity.List))
	r.GET("/user/count", wrap(user.Entity.Count))
	r.GET("/user/:id", wrap(user.Entity.GetByID))
	r.POST("/user", wrap(user.Entity.Add))
	r.PUT("/user/:id", wrap(user.Entity.Update))
	r.DELETE("/user/:id", wrap(user.Entity.Delete))


		
	return r
}
