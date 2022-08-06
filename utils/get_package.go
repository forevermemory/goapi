package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
)

// GetStructMethods 获取结构体的所有方法
func GetStructMethods(structName string, fpath string, regStr string) map[string]string {
	b, _ := ioutil.ReadFile(fpath)
	a := string(b)

	// fmt.Println(a)
	// func[\s]+.*?UserController\)\s(.*?)\(.*?{
	// golang 的m模式 可以匹配多行
	reg, _ := regexp.Compile(regStr)
	// res := reg.FindAllStringIndex(a, -1)
	matchs := reg.FindAllStringSubmatchIndex(a, -1)
	// fmt.Println(matchs)
	var res = map[string]string{}

	// var first = res[0] // [2 58] [145 204] [230 285]]
	// 共匹配到几组: 4 [[5 61 30 34] [348 407 373 380] [432 487 457 460] [512 570 537 543]]
	// fmt.Println(a[2:145])
	// fmt.Println(a[145:230])
	// fmt.Println(a[230:])
	var length = len(matchs)
	var funcName string
	for i := 0; i < length; i++ {
		funcName = a[matchs[i][2]:matchs[i][3]]
		if i+1 != length {
			res[funcName] = a[matchs[i][0]:matchs[i+1][0]]
			continue
		}
		res[funcName] = a[matchs[i][0]:]
	}
	var result2 = map[string]string{}
	for k, v := range res {
		result2[k] = strings.ReplaceAll(v, "\"", `"`)
	}
	return result2
}

// // GetControllerFuncs 读取controller.go里面自定义的和重写的函数
// func GetControllerFuncs(controllerName, fpath string) map[string]string {

// }

// GetApiConfigs 读取api/item/config.json
func GetApiConfigs(project string) ([]*ApiConfig, error) {
	var apis = make([]*ApiConfig, 0)
	dirs, _ := ioutil.ReadDir(path.Join(project, "api"))
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		b, err := ioutil.ReadFile(path.Join(project, "api", dir.Name(), "config.json"))
		if err != nil {
			log.Println("err:", err)
			continue
		}

		var cfg = ApiConfig{}
		if err = json.Unmarshal(b, &cfg); err != nil {
			return nil, err
		}
		apis = append(apis, &cfg)
	}

	return apis, nil
}

// 读取一个文件 返回引入的包集合 可能没有包
func GetGoFilePackages(fpath string) map[string]int {
	b, _ := ioutil.ReadFile(fpath)
	ss := strings.Split(string(b), "\n")
	var result = map[string]int{}
	var multiPackage bool

	for _, line := range ss {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasSuffix(line, ")") && multiPackage {
			break
		}
		if multiPackage {
			result[line] = 1
			continue
		}

		// 只引入了一个包 import "fmt"
		if strings.HasPrefix(line, "import \"") {
			result[strings.Replace(line, "import ", "", -1)] = 1
			break
		}

		// 引入多个包 import ()
		if strings.HasPrefix(line, "import (") {
			multiPackage = true
			continue
		}

	}
	var result2 = map[string]int{}
	for k := range result {
		result2[strings.ReplaceAll(k, "\"", `"`)] = 1
	}
	return result2
}
