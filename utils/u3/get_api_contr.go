package u3

import (
	"goapi/utils"
	"os"
	"path"
	"strings"
)

func GenerateNewApiController() error {
	prepareData()

	// write to file
	modelStr := getControllerMethods(apiConfig.ModelName, utils.ToGoUpper(apiConfig.ModelName), goModName)
	fp, err := os.Create(path.Join(".", "api", apiConfig.ModelName, "controller.go"))
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err = fp.WriteString(modelStr); err != nil {
		return err
	}

	return nil
}

func getControllerMethods(apiName string, structName string, goModName string) string {

	var s = `package ` + apiName + `

import (
	"github.com/gin-gonic/gin"
	"` + goModName + `/config"
	"strconv"

)

var Entity = &GO_STRUCT_NAMEController{}

type GO_STRUCT_NAMEController struct {}

func (a *GO_STRUCT_NAMEController) Add(c *gin.Context) interface{} {
	var req = GO_STRUCT_NAME{}
	err := c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := GO_STRUCT_NAME{}
	data, err := b.Add(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}

func (a *GO_STRUCT_NAMEController) Update(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	var req = GO_STRUCT_NAME{}
	err = c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}
	req.ID = intv

	t := GO_STRUCT_NAME{}
	data, err := t.Update(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}

func (a *GO_STRUCT_NAMEController) Delete(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithData("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	t := GO_STRUCT_NAME{}
	err = t.Delete(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK
}

func (a *GO_STRUCT_NAMEController) List(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := GO_STRUCT_NAME{}
	datas, err := b.List(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}

func (a *GO_STRUCT_NAMEController) Count(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := GO_STRUCT_NAME{}
	datas, err := b.Count(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}

func (a *GO_STRUCT_NAMEController) GetByID(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := GO_STRUCT_NAME{}
	data, err := b.GetItemByID(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}`

	s2 := strings.ReplaceAll(s, "GO_STRUCT_NAME", structName)

	return s2

}
