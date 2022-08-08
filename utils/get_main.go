package utils

import "bytes"

func GenerateMainGoString() string {

	buf := bytes.Buffer{}

	buf.WriteString(`
package main

import (
	"fmt"
	"goapi/api"
	"goapi/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := api.InitRouter()

	log.Println("http server is runing on:", config.GlobalConfig.HttpConfig.Port)
	r.Run(fmt.Sprintf(":%v", config.GlobalConfig.HttpConfig.Port))
}
`)

	return buf.String()
}
