package service

import (
	"errors"
	"goapi/utils"
	"log"
	"os"
	"path"
	"strings"
)

func GenerateApiHandles() {
	runApiHandle.run()
}

var project = "/Users/liuqt/develop/go/projects/basestrapi/"

var runApiHandle = runApiHandles{
	project:   project,
	basecache: path.Join(project, ".cache"),
}

type runApiHandles struct {
	project   string
	basecache string
	apis      []*utils.ApiConfig
}

func mkdir(s string) error {
	err := os.Mkdir(s, 0777)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			// 已存在文件夹不用关心
			return nil
		}
		log.Println("initWorkdir err:", err)
		return err
	}
	return nil
}

func (c *runApiHandles) writeStringToFile(fpath, content string) error {

	fp, err := os.Create(fpath)
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

// copy 拷贝当前项目的所有文件到.cache下面
func (c *runApiHandles) copy() {
	utils.CopyDir(c.project, c.basecache)
}

func (c *runApiHandles) initWorkdir() {

	mkdir(c.basecache)
	// mkdir(path.Join(c.basecache, "api"))
	// mkdir(path.Join(c.basecache, "config"))

}

func (c *runApiHandles) run() {
	// utils.GetStructFromString("", "")
	// return
	c.initWorkdir()
	c.copy() // TODO

	// return

	var err error
	var project = "/Users/liuqt/develop/go/projects/basestrapi/"
	apis, err := utils.GetApiConfigs(project)
	if err != nil {
		log.Println(err)
		return
	}

	for _, api := range apis {
		log.Println(api.ControllerName)
		// 生成controller.go
		c.generateController(api)

		// 生成model.go
		c.generateModel(api)

		// 生成 routes
	}

	// 生成api/router.go
}

func (c *runApiHandles) generateModel(api *utils.ApiConfig) error {
	// 1. 获取包名
	modelPath := path.Join(utils.Workdir, "api", api.ModelName, "model.go")
	var structModelName = utils.ToFirstUpper(api.ModelName)

	// 2. 生成struct 定义
	var ok bool
	structString := utils.GetStructFromAPi(api)

	// 2.2 替换当前的model.go结构体
	utils.ReplaceCurrentModelStruct(api, structString)

	// 3. 获取结构体的函数 (生命周期和增删改查或者用户自定义)
	reg := `(?m)^func.+?` + structModelName + `\)\s+(.*?)\(.*?\{`
	rewriteModelHandles := utils.GetStructMethods(api.ControllerName, modelPath, reg)
	// log.Println(rewriteModelHandles)

	for k := range utils.InternelModelHandles {
		if _, ok = rewriteModelHandles[k]; !ok {
			// 模版替换
			s2, _ := utils.ParseModelHandleTemplate(structModelName, k)
			rewriteModelHandles[k] = s2
		}
	}
	var _handles = make([]string, 0)
	for _, v := range rewriteModelHandles {
		_handles = append(_handles, v)
	}

	// 4. 写入文件
	param := utils.ControllerAndModelHandleTemplate{
		StructString:   structString,
		ControllerName: api.ControllerName,
		ModelName:      api.ModelName,
		Packages:       api.ModelPackages,
		Handles:        _handles,
		StructName:     structModelName,
	}
	str3 := utils.GetModelGoString(&param)
	// mkdir(path.Join(c.basecache, "api", api.ModelName))
	c.writeStringToFile(path.Join(c.basecache, "api", api.ModelName, "model.go"), str3)

	return nil

}

func (c *runApiHandles) generateController(api *utils.ApiConfig) error {
	controllerPath := path.Join(utils.Workdir, "api", api.ModelName, "controller.go")

	// 1. 获取包名

	// return nil
	// 处理路由处理函数
	// 2. 获取handles
	var apiHandles = map[string]string{}
	reg := `(?m)^func[\s]+.*?` + api.ControllerName + `\)\s+(.*?)\(.*?{`
	rewriteApiHandles := utils.GetStructMethods(api.ControllerName, controllerPath, reg)

	var ok bool
	var err error
	var rewrite string
	for _, route := range api.Routes {

		name := strings.Split(route.Handler, ".")[1]
		if rewrite, ok = rewriteApiHandles[name]; ok {
			apiHandles[name] = rewrite
		} else {
			rewrite, err = utils.ParseControllerHandleTemplate(api.ControllerName, api.ModelName, name)
			if err != nil {
				log.Println(err)
				return err
			}
			apiHandles[name] = rewrite
		}
	}
	api.ControllerHandles = apiHandles

	// 3. 生成文件了
	str3 := utils.GetControllerGoString(api)

	// mkdir(path.Join(c.basecache, "api", api.ModelName))
	c.writeStringToFile(path.Join(c.basecache, "api", api.ModelName, "controller.go"), str3)

	return nil
}
