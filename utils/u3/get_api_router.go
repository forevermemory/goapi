package u3

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

func AddApiToRouter() error {
	prepareData()

	var routerPath = path.Join(".", "api", "router.go")

	b, err := ioutil.ReadFile(routerPath)
	if err != nil {
		return err
	}
	tmp := string(b)

	//////////////// add package //////////////////
	reg, _ := regexp.Compile(`import\s+\(`)
	res := reg.FindAllStringIndex(tmp, 1)

	var s1 string
	if len(res) > 0 {
		s1 += tmp[0:res[0][1]]
		s1 += fmt.Sprintf("\n\t\"%v/api/%v\"\n", goModName, apiName)
		s1 += tmp[res[0][1]:]
	}

	///////////////// add api handle /////////////////
	reg2, _ := regexp.Compile(`(?m)^func\s+handleApi\(r\s+\*gin\.Engine\)\s+\{`)
	res2 := reg2.FindStringIndex(s1) // 返回[start end]
	var i2 = 1
	var count2 = 1

	for {
		b := s1[res2[1]+i2]
		if b == '{' {
			count2 += 1
		}
		if b == '}' {
			count2 -= 1
		}
		i2 += 1
		if count2 == 0 {
			break
		}
	}

	var s2 string
	s2 += s1[:res2[1]+i2-1]
	s2 += "\n"
	s2 += getHandleStr()
	s2 += "\n"
	s2 += "\n"
	s2 += "}"
	s2 += s1[res2[1]+i2:]

	///////////////// initTable ////////////////////
	reg3, _ := regexp.Compile(`(?m)^func\s+initTable\(\)\s+\{`)
	res3 := reg3.FindStringIndex(s2) // 返回[start end]
	var i3 = 1
	var count3 = 1

	for {
		b := s2[res3[1]+i3]
		if b == '{' {
			count3 += 1
		}
		if b == '}' {
			count3 -= 1
		}
		i3 += 1
		if count3 == 0 {
			break
		}
	}

	var s3 string
	s3 += s2[:res3[1]+i3-1]
	s3 += getInitTableStr()
	s3 += "\n"
	s3 += "\n"
	s3 += "}"
	s3 += s2[res3[1]+i3:]

	ioutil.WriteFile(routerPath, []byte(s3), fs.ModePerm)
	return nil

}

func getInitTableStr() string {
	var s = `	config.DATABASE.AutoMigrate(&` + apiName + `.` + apiStructName + `{})`
	return s
}

func getHandleStr() string {
	var s = `	apis.GET("/GO_APINAME", wrap(GO_APINAME.Entity.List))
	apis.GET("/GO_APINAME/count", wrap(GO_APINAME.Entity.Count))
	apis.GET("/GO_APINAME/:id", wrap(GO_APINAME.Entity.GetByID))
	apis.POST("/GO_APINAME", wrap(GO_APINAME.Entity.Add))
	apis.PUT("/GO_APINAME/:id", wrap(GO_APINAME.Entity.Update))
	apis.DELETE("/GO_APINAME/:id", wrap(GO_APINAME.Entity.Delete))`

	return strings.ReplaceAll(s, "GO_APINAME", apiName)
}
