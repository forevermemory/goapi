package utils

import (
	"bytes"
	"html/template"
)

const (
	Count_Controller_Template string = `
func (a *{{ .ControllerName }}) Count{{ .ModelName }}(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := {{ .ModelName }}{}
	datas, err := b.Count(cond)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK.WithData(datas)
}

`

	List_Controller_Template string = `
func (a *{{ .ControllerName }}) List{{ .ModelName }}(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := {{ .ModelName }}{}
	datas, err := b.List(cond)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK.WithData(datas)
}

	`
	GetByID_Controller_Template string = `
func (a *{{ .ControllerName }})  Get{{ .ModelName }}ByID(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return goapi.ERROR.WithCode(ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	b := {{ .ModelName }}{}
	data, err := b.GetItemByID(intv)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK.WithData(data)
}
	
	`
	Add_Controller_Template string = `
func (a *{{ .ControllerName }}) Add{{ .ModelName }}(c *gin.Context) interface{} {
	var req = {{ .ModelName }}{}
	err := c.ShouldBind(&req)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	b := {{ .ModelName }}{}
	data, err := b.Add(&req)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK.WithData(data)
}

	`
	Update_Controller_Template string = `
func (a *{{ .ControllerName }}) Update{{ .ModelName }}(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return goapi.ERROR.WithCode(ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	var req = {{ .ModelName }}{}
	err = c.ShouldBind(&req)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}
	req.ID = intv

	t := {{ .ModelName }}{}
	data, err := t.Update(&req)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK.WithData(data)
}
	
	`
	Delete_Controller_Template string = `
func (a *{{ .ControllerName }}) Delete{{ .ModelName }}(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return goapi.ERROR.WithCode(ErrParamNotFull).WithData("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	t := {{ .ModelName }}{}
	err = t.Delete(intv)
	if err != nil {
		return goapi.ERROR.WithMessage(err)
	}

	return goapi.OK
}
	
	`

	Controller_go_template string = `
	package {{ .ModelName }}

	import (
		{{ .Packages }}
	)
	
	var Entity = &{{ .ControllerName }}{}
	
	type {{ .ControllerName }} struct {
	}
	
	{{ .Handles }}
	`
)

var InternelControllerHandles = map[string]string{
	"List":    List_Controller_Template,
	"Count":   Count_Controller_Template,
	"GetByID": GetByID_Controller_Template,
	"Add":     Add_Controller_Template,
	"Update":  Update_Controller_Template,
	"Delete":  Delete_Controller_Template,
}

type ControllerAndModelHandleTemplate struct {
	StructString   string
	ControllerName string
	Packages       map[string]int
	Handles        []string
	ModelName      string
	StructName     string
}

func ParseControllerHandleTemplate(controllerName string, modelName string, apiname string) (string, error) {
	var err error
	var tpl *template.Template

	tplname := InternelControllerHandles[apiname]
	tpl, err = template.New(apiname).Parse(tplname)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, &ControllerAndModelHandleTemplate{ControllerName: controllerName, ModelName: modelName})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetControllerGoString(param *ApiConfig) string {

	buf := &bytes.Buffer{}
	buf.WriteString(`package `)
	buf.WriteString(param.ModelName)
	buf.WriteString("\n")
	buf.WriteString("\n")

	// packages
	buf.WriteString("import (\n")
	for k := range param.ControllerPackages {
		buf.WriteString("\t")
		buf.WriteString(k)
		buf.WriteString("\n")
	}

	buf.WriteString("\n")
	buf.WriteString(")\n")
	buf.WriteString("\n")
	buf.WriteString("var Entity = &")
	buf.WriteString(param.ControllerName)
	buf.WriteString("{}\n\n")
	buf.WriteString("type ")
	buf.WriteString(param.ControllerName)
	buf.WriteString(" struct {}\n")
	buf.WriteString("\n")
	buf.WriteString("\n")

	for _, v := range param.ControllerHandles {
		buf.WriteString(v)
		buf.WriteString("\n")
	}
	buf.WriteString("\n")

	return buf.String()
}
