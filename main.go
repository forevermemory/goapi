package main

import (
	"flag"
	"fmt"
	"goapi/utils/u2"
	"goapi/utils/u3"
	"os"
)

var argFmt = `./pragram -c init -db sqlite -name project1 \n`

var command, projectName, dbType string

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(argFmt)
		return
	}

	flag.StringVar(&command, "c", "init", "命令模式")
	flag.StringVar(&projectName, "name", "tmpgoproject.com", "项目名称")
	flag.StringVar(&dbType, "db", "sqlite", "数据库类型")

	flag.Parse()

	fmt.Printf("command = %v \n", command)
	fmt.Printf("projectName = %v \n", projectName)
	fmt.Printf("dbType = %v \n", dbType)

	switch command {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("./pragram init [project-name]")
			return
		}
		initProject()
		addExe()

	case "add":
		if projectName == "" {
			// 使用示例
			return
		}
		addApi()
	}

	return
}

func addExe() {
	u3.AddExeToProject(projectName)
}

func addApi() {
	u3.GenerateNewApiModel()
	u3.GenerateNewApiController()
	u3.AddApiToRouter()
}

func initProject() {

	// 初始化目录
	u2.Madir(projectName, "api")
	u2.Madir(projectName, "api/user")
	u2.Madir(projectName, "api/role")
	u2.Madir(projectName, "api/permission")
	u2.Madir(projectName, "api/role_permission")
	u2.Madir(projectName, "config")
	u2.Madir(projectName, "utils")

	// config
	u2.WriteConfig(projectName, dbType)
	u2.WriteMainGo(projectName)

	// 添加默认文件
	u2.WriteConfigYaml(projectName, dbType)
	u2.WriteDotGitignore(projectName)
	u2.WriteGoMod(projectName, dbType)
	u2.WriteReadme(projectName)
	u2.WriteDefaultModelConfig(projectName)

	// 写入utils
	u2.WriteUtilsCors(projectName)

	// 添加api
	u2.WriteApiRouter(projectName)

	u2.WriteApiUserModel(projectName)
	u2.WriteApiUserController(projectName)
	u2.WriteApiUserConfigJson(projectName)

	u2.WriteApiRoleModel(projectName)
	u2.WriteApiRoleConfigJson(projectName)
	u2.WriteApiRoleController(projectName)

	u2.WriteApiRolePermissionController(projectName)
	u2.WriteApiRolePermissionModel(projectName)
	u2.WriteApiRolePermissionConfigJson(projectName)

	u2.WriteApiPermissionController(projectName)
	u2.WriteApiPermissionModel(projectName)
	u2.WriteApiPermissionConfigJson(projectName)

	// 初始化数据库

}
