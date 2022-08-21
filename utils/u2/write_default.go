package u2

import (
	"log"
	"os"
	"path"
)

func WriteMainGo(projectname string) error {
	var content = `package main

import (
	"fmt"
	"` + projectname + `/api"
	"` + projectname + `/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := api.InitRouter()

	log.Println("http server is runing on:", config.GlobalConfig.HttpConfig.Port)
	r.Run(fmt.Sprintf(":%v", config.GlobalConfig.HttpConfig.Port))
}
`

	return writeStr2File(path.Join(".", projectname, "main.go"), content)
}

func WriteDefaultModelConfig(projectname string) error {
	var content = `{
	"modelName":"修改这里",
	"attributes":[
		{
			"name":"username",
			"type":"string",
			"description":"描述"
		},
		{
			"name":"blocked",
			"type":"boolean",
			"description":"是否锁定"
		},
		{
			"name":"login_at",
			"type":"datetime",
			"description":"上次登录时间"
		}
		,
		{
			"name":"role_id",
			"model":"role",
			"type":"reference-belong",
			"description":"角色id"
		}
	]
}`

	return writeStr2File(path.Join(".", projectname, "config_model.json"), content)
}

func WriteDotGitignore(projectname string) error {
	var content = `.DS_Store
.cache`

	return writeStr2File(path.Join(".", projectname, ".gitignore"), content)
}

func WriteReadme(projectname string) error {
	var content = `### this project is open-source CMS by golang
#### project init 
	- go mod tidy
	- go build main.go

### add api
	-   ./program -c add
		-   program read config_model.json

`

	return writeStr2File(path.Join(".", projectname, "README.md"), content)
}

func WriteGoMod(projectname, dbType string) error {
	var content = ""
	switch dbType {
	case "sqlserver":
		content += `module ` + projectname + `
		
go 1.16
		
require (
	github.com/gin-gonic/gin v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.8
)
`
	case "postgres":
		content += `module ` + projectname + `
		
go 1.16
		
require (
	github.com/gin-gonic/gin v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.3.9
	gorm.io/gorm v1.23.8
)
`
	case "sqlite":
		content += `module ` + projectname + `
		
go 1.16
		
require (
	github.com/gin-gonic/gin v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)
`
	case "mysql":
		content += `module ` + projectname + `

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)
`

	default:
		content += `module ` + projectname + `

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)
`

	}

	return writeStr2File(path.Join(".", projectname, "go.mod"), content)
}

func WriteConfigYaml(projectname string, dbType string) error {
	var content string
	switch dbType {
	case "sqlserver":
		content = `#### change config when need 
sqlserver_config:
    user: root
    password: "123456"
    host: localhost
    port: 3306
    database: taobao
http_config:
    port: 80`
	case "postgres":
		content = `#### change config when need 
postgres_config:
    user: root
    password: "123456"
    host: localhost
    port: 3306
    database: taobao
    sslmode: disable
    time_zone: Asia/Shanghai
http_config:
    port: 80`
	case "sqlite":
		content = `#### change config when need 
sqlite3_config:
	file_path: "api.db"
http_config:
	port: 80`
	case "mysql":
		content = `#### change config when need 
mysql_config:
    user: root
    password: "123456"
    host: localhost
    port: 3306
    database: taobao
http_config:
    port: 80`
	default:
		content = `#### change config when need 
sqlite3_config:
    file_path: "api.db"
http_config:
    port: 80`
	}

	return writeStr2File(path.Join(".", projectname, "config.yaml"), content)
}

func writeStr2File(filepath, content string) error {
	fp, err := os.Create(filepath)
	if err != nil {
		log.Println("createFile err:", err)
		return err
	}
	defer fp.Close()

	if _, err = fp.WriteString(content); err != nil {
		log.Println("createFile err:", err)
		return err
	}
	return nil
}
