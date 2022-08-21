package u2

import "path"

func WriteApiRouter(projectname string) error {
	var content = `package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"` + projectname + `/utils"
	"` + projectname + `/api/role"
	"` + projectname + `/api/user"
	"` + projectname + `/config"
)


func wrap(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.Cors())

	/////////// init table ///////////
	initTable()
	/////////// init table ///////////

	/////////// API START///////////
	handleApi(r)
	/////////// API END///////////

	return r
}

func initTable() {
	config.DATABASE.AutoMigrate(&role.Role{})
	config.DATABASE.AutoMigrate(&user.User{})
}

func handleApi(r *gin.Engine) {
	apis := r.Group("/api")

	apis.GET("/role", wrap(role.Entity.List))
	apis.GET("/role/count", wrap(role.Entity.Count))
	apis.GET("/role/:id", wrap(role.Entity.GetByID))
	apis.POST("/role", wrap(role.Entity.Add))
	apis.PUT("/role/:id", wrap(role.Entity.Update))
	apis.DELETE("/role/:id", wrap(role.Entity.Delete))

	apis.GET("/user", wrap(user.Entity.List))
	apis.GET("/user/count", wrap(user.Entity.Count))
	apis.GET("/user/:id", wrap(user.Entity.GetByID))
	apis.POST("/user", wrap(user.Entity.Add))
	apis.PUT("/user/:id", wrap(user.Entity.Update))
	apis.DELETE("/user/:id", wrap(user.Entity.Delete))
}

`
	return writeStr2File(path.Join(".", projectname, "api", "router.go"), content)

}
