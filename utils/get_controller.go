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
		return ERROR.WithMessage(err)
	}

	return OK.WithData(datas)
}

`

	List_Controller_Template string = `
func (a *{{ .ControllerName }}) List{{ .ModelName }}(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := {{ .ModelName }}{}
	datas, err := b.List(cond)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	return OK.WithData(datas)
}

	`
	GetByID_Controller_Template string = `
func (a *{{ .ControllerName }})  Get{{ .ModelName }}ByID(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return ERROR.WithCode(ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	b := {{ .ModelName }}{}
	data, err := b.GetItemByID(intv)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	return OK.WithData(data)
}
	
	`
	Add_Controller_Template string = `
func (a *{{ .ControllerName }}) Add{{ .ModelName }}(c *gin.Context) interface{} {
	var req = {{ .ModelName }}{}
	err := c.ShouldBind(&req)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	b := {{ .ModelName }}{}
	data, err := b.Add(&req)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	return OK.WithData(data)
}

	`
	Update_Controller_Template string = `
func (a *{{ .ControllerName }}) Update{{ .ModelName }}(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return ERROR.WithCode(ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	var req = {{ .ModelName }}{}
	err = c.ShouldBind(&req)
	if err != nil {
		return ERROR.WithMessage(err)
	}
	req.ID = intv

	t := {{ .ModelName }}{}
	data, err := t.Update(&req)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	return OK.WithData(data)
}
	
	`
	Delete_Controller_Template string = `
func (a *{{ .ControllerName }}) Delete{{ .ModelName }}(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return ERROR.WithCode(ErrParamNotFull).WithData("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	t := {{ .ModelName }}{}
	err = t.Delete(intv)
	if err != nil {
		return ERROR.WithMessage(err)
	}

	return OK
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

type ControllerAndModelHandleTemplate struct {
	StructString   string
	ControllerName string
	Packages       []string
	Handles        []string
	ModelName      string
	StructName     string
}

func ParseControllerHandleTemplate(controllerName string, modelName string, apiname string) (string, error) {
	var err error
	var tpl *template.Template
	// log.Println("apiname:::", apiname)
	switch apiname {
	case "List":
		tpl, err = template.New(apiname).Parse(List_Controller_Template)
	case "GetByID":
		tpl, err = template.New(apiname).Parse(GetByID_Controller_Template)
	case "Add":
		tpl, err = template.New(apiname).Parse(Add_Controller_Template)
	case "Update":
		tpl, err = template.New(apiname).Parse(Update_Controller_Template)
	case "Delete":
		tpl, err = template.New(apiname).Parse(Delete_Controller_Template)
	case "Count":
		tpl, err = template.New(apiname).Parse(Delete_Controller_Template)
	}
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

func GetControllerGoString(param *ControllerAndModelHandleTemplate) string {

	buf := &bytes.Buffer{}
	buf.WriteString(`package `)
	buf.WriteString(param.ModelName)
	buf.WriteString("\n")
	buf.WriteString("\n")
	buf.WriteString("import (\n")
	for _, v := range param.Packages {
		buf.WriteString("\t")
		buf.WriteString(v)
		buf.WriteString("\n")
	}
	// 必须要写入一个config
	buf.WriteString("\t\"commonapi/config\"\n")

	buf.WriteString("\n")
	buf.WriteString(")\n")
	buf.WriteString("\n")
	buf.WriteString("var Entity = &")
	buf.WriteString(param.ControllerName)
	buf.WriteString("{}\n\n")
	buf.WriteString("type ")
	buf.WriteString(param.ControllerName)
	buf.WriteString(" struct {\n")
	buf.WriteString("}")
	buf.WriteString("\n")
	buf.WriteString("\n")

	for _, v := range param.Handles {
		buf.WriteString(v)
		buf.WriteString("\n")
	}
	buf.WriteString("\n")

	return buf.String()
}
