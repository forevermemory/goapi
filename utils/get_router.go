package utils

import (
	"bytes"
	"strings"
)

func GenerateRouterString(apis []*ApiConfig) string {

	buf := bytes.Buffer{}

	// 1. write package names
	buf.WriteString("package api\n")
	buf.WriteString("\n")

	var extraPackages = map[string]int{}

	buf2 := bytes.Buffer{}

	// 4. write api handles
	for _, api := range apis {
		for _, route := range api.Routes {
			// "method": "GET",
			// "path": "/user",
			// "handler": "UserController.List",
			// "description":"查询用户列表",

			buf2.WriteString("\t")
			buf2.WriteString("r." + route.Method + "(\"" + route.Path + "\", wrap(" + api.ModelName + ".Entity." + strings.Split(route.Handler, ".")[1] + "))")
			buf2.WriteString("\n")

			extraPackages[`"goapi/api/`+api.ModelName+`"`] = 1

		}
		buf2.WriteString("\n")
	}

	// 2. write packages
	buf.WriteString("import (\n")
	buf.WriteString("\t\"net/http\"")
	buf.WriteString("\n")
	buf.WriteString("\t\"github.com/gin-gonic/gin\"")
	buf.WriteString("\n")

	// 2.2 write extra package
	for k := range extraPackages {
		buf.WriteString("\t")
		buf.WriteString(k)
		buf.WriteString("\n")
	}

	buf.WriteString(")\n")
	buf.WriteString("\n")

	// 3. write func
	buf.WriteString(`
func wrap(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}
`)

	buf.WriteString(`
func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

`)

	buf.WriteString(buf2.String())
	// 5.
	buf.WriteString(`
		
	return r
}
`)

	return buf.String()
}
