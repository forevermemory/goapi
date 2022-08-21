package utils

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

const (
	List_Model_Template string = `
	
func (*{{ .StructName }}) List(cond *config.Condition, tx ...*gorm.DB) ([]*{{ .StructName }}, error) {
	var datas = make([]*{{ .StructName }}, 0)
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	db = db.Model(&{{ .StructName }}{})
	db = db.Where("deleted_at is null")
	// 条件
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field + " > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field + " >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field + " < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field + " <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field + " = ? ", v.Value)
		case config.CONTAINS:
			db = db.Where(v.Field + " like '%" + v.Value + "%' ")
		}
	}

	// 排序
	for _, v := range cond.Sorts {
		db = db.Order(fmt.Sprintf(" %s %s ", v.Field, v.Value))
	}

	// 分页
	db = db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)

	// 查询
	err := db.Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}`
	GetByID_Model_Template string = `
	
func (*{{ .StructName }}) GetItemByID(id int, tx ...*gorm.DB) (*{{ .StructName }}, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var rec = {{ .StructName }}{}
	db = db.Model(&{{ .StructName }}{})
	db = db.Where("deleted_at is null")
	err := db.Where("id = ?", id).First(&rec).Error
	return &rec, err
}`
	Count_Model_Template string = `
	
func (*{{ .StructName }}) Count(cond *config.Condition, tx ...*gorm.DB) (int64, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var count int64
	db = db.Model(&{{ .StructName }}{})
	db = db.Where("deleted_at is null")
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field + " > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field + " >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field + " < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field + " <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field + " = ? ", v.Value)
		case config.CONTAINS:
			db = db.Where(v.Field + " like '%" + v.Value + "%' ")
		}
	}
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return 0, nil
}`
	Add_Model_Template string = `
func (*{{ .StructName }}) Add(o *{{ .StructName }}, tx ...*gorm.DB) (*{{ .StructName }}, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	o.CreatedAt = time.Now()
	o.UpdatedAt = o.CreatedAt
	err := db.Create(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}`
	Update_Model_Template string = `
	
func (*{{ .StructName }}) Update(o *{{ .StructName }}, tx ...*gorm.DB) (*{{ .StructName }}, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&{{ .StructName }}{}).Where("id = ?", o.ID).Updates(o).Error
	// err := db.Model(&{{ .StructName }}{}).Where("id = ?", o.ID).Save(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}`
	Delete_Model_Template string = `
func (*{{ .StructName }}) Delete(id int, tx ...*gorm.DB) error {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&{{ .StructName }}{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	// err := db.Delete(&{{ .StructName }}{}, id).Error

	if err != nil {
		return err
	}
	return nil
}`

	// 生命周期
	TableName_Model_Template    string = `func (o *{{ .StructName }}) TableName() string {return "{{ .ModelName }}"}`
	BeforeCreate_Model_Template string = `func (o *{{ .StructName }}) BeforeCreate(tx *gorm.DB) error { return nil }`
	AfterCreate_Model_Template  string = `func (o *{{ .StructName }}) AfterCreate(tx *gorm.DB) error { return nil }`
	AfterSave_Model_Template    string = `func (o *{{ .StructName }}) AfterSave(tx *gorm.DB) error { return nil }`

	BeforeUpdate_Model_Template string = `func (o *{{ .StructName }}) BeforeUpdate(tx *gorm.DB) error { return nil }`
	AfterUpdate_Model_Template  string = `func (o *{{ .StructName }}) AfterUpdate(tx *gorm.DB) error { return nil }`
	BeforeDelete_Model_Template string = `func (o *{{ .StructName }}) BeforeDelete(tx *gorm.DB) error { return nil }`
	AfterDelete_Model_Template  string = `func (o *{{ .StructName }}) AfterDelete(tx *gorm.DB) error { return nil }`

	AfterFind_Model_Template string = `func (o *{{ .StructName }}) AfterFind(tx *gorm.DB) error   { return nil }`
)

var InternelModelHandles = map[string]string{
	"List":    List_Model_Template,
	"Count":   Count_Model_Template,
	"GetByID": GetByID_Model_Template,
	"Add":     Add_Model_Template,
	"Update":  Update_Model_Template,
	"Delete":  Delete_Model_Template,

	"TableName":    TableName_Model_Template,
	"BeforeCreate": BeforeCreate_Model_Template,
	"AfterCreate":  AfterCreate_Model_Template,
	"AfterSave":    AfterSave_Model_Template,

	"BeforeUpdate": BeforeUpdate_Model_Template,
	"AfterUpdate":  AfterUpdate_Model_Template,
	"BeforeDelete": BeforeDelete_Model_Template,
	"AfterDelete":  AfterDelete_Model_Template,

	"AfterFind": AfterFind_Model_Template,
}

// GetModelGoString 生成.cache 下面的api/xxx/model.go的字符串
func GetModelGoString(param *ControllerAndModelHandleTemplate) string {

	buf := &bytes.Buffer{}
	buf.WriteString(`package `)
	buf.WriteString(param.ModelName)
	buf.WriteString("\n")
	buf.WriteString("\n")

	// 包
	buf.WriteString("import (\n")
	for k := range param.Packages {
		buf.WriteString("\t")
		buf.WriteString(k)
		buf.WriteString("\n")
	}

	buf.WriteString("\n")
	buf.WriteString(")\n")
	buf.WriteString("\n")

	// 结构体
	buf.WriteString(param.StructString)

	// 方法
	for _, v := range param.Handles {
		buf.WriteString(v)
		buf.WriteString("\n")
		buf.WriteString("\n")
	}
	buf.WriteString("\n")

	return buf.String()
}

// ParseModelHandleTemplate 替换model.go 里面结构体方法的模版变量
func ParseModelHandleTemplate(modelName string, apiname string) (string, error) {
	var err error
	var tpl *template.Template

	tmpname := InternelModelHandles[apiname]
	tpl, err = template.New(apiname).Parse(tmpname)

	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, &ControllerAndModelHandleTemplate{StructName: ToGoUpper(modelName), ModelName: modelName})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ReplaceCurrentModelStruct 替换主文件的api/xxx/model.go的struct定义
func ReplaceCurrentModelStruct(api *ApiConfig, structString string) {

	modelPath := path.Join(Workdir, "api", api.ModelName, "model.go")
	b, _ := ioutil.ReadFile(modelPath)

	oldStr := string(b)

	// 获取到struct的索引开始和结束位置
	pos := GetStructPositionFromString(oldStr, ToGoUpper(api.ModelName))

	// 更新import的包
	// pos2 := GetImportPositionFromString(oldStr) TODO

	newString := oldStr[0:pos[0]] + structString + oldStr[pos[1]:]

	fp, err := os.Create(modelPath)
	if err != nil {
		return
	}
	defer fp.Close()
	fp.WriteString(newString)

}

// GetStructFromAPi 根据传入的config.json 生成 type xxx struct{} 字符串
func GetStructFromAPi(api *ApiConfig) (string, map[string]int) {
	var extraPackages = map[string]int{}
	buf := bytes.Buffer{}
	buf.WriteString("type ")
	buf.WriteString(ToGoUpper(api.ModelName))
	buf.WriteString(" struct {")
	buf.WriteString("\n")

	buf.WriteString("\tID\tint\t`json:\"id\" gorm:\"column:id;primary_key;auto_increment;\"`")
	buf.WriteString("\n")
	buf.WriteString("\n")

	var gotype string
	for _, m := range api.Attributes {
		buf.WriteString("\t")
		buf.WriteString(ToGoUpper(m.Name))
		buf.WriteString("\t")

		switch m.Type {
		case Boolean_TYPE:
			gotype = "int"
		case Number_TYPE:
			gotype = "int"
		case String_TYPE:
			gotype = "string"
		case Datetime_TYPE:
			gotype = "time.Time"
		case Reference_Belong_TYPE:
			gotype = "int"
			goto BELONG_TYPE

		}

		buf.WriteString(gotype)
		buf.WriteString("\t")
		buf.WriteString("`json:\"" + m.Name + "\" gorm:\"column:" + m.Name + "\"`")
		buf.WriteString("\n")
		continue

	BELONG_TYPE:
		buf.WriteString(gotype)
		buf.WriteString("\t")
		buf.WriteString("`json:\"" + m.Name + "\" gorm:\"column:" + m.Name + "\"`")
		buf.WriteString("\n")
		// Company      Company `gorm:"foreignKey:CompanyRefer"`

		belongToModel := ToGoUpper(m.Model)
		belong := "\t" + belongToModel + "\t" + m.Model + "." + belongToModel + "\t`gorm:\"foreignKey:" + ToGoUpper(m.Name) + "\"`"
		buf.WriteString(belong)

		extraPackages[`"goapi/api/`+m.Model+`"`] = 1

	}

	buf.WriteString(addTimeField())

	buf.WriteString("}")
	buf.WriteString("\n")
	buf.WriteString("\n")

	return buf.String(), extraPackages

}

// GetStructFromCreateNewApi 根据传入的config.json 生成 type xxx struct{} 字符串
func GetStructFromCreateNewApi(structName string, gomodName string, attributes []*ApiField) (string, map[string]int) {
	var extraPackages = map[string]int{}
	buf := bytes.Buffer{}
	buf.WriteString("type ")
	buf.WriteString(structName)
	buf.WriteString(" struct {")
	buf.WriteString("\n")

	buf.WriteString("\tID\tint\t`json:\"id\" gorm:\"column:id;primary_key;auto_increment;\"`")
	buf.WriteString("\n")
	buf.WriteString("\n")

	var gotype string
	for _, m := range attributes {
		buf.WriteString("\t")
		buf.WriteString(ToGoUpper(m.Name))
		buf.WriteString("\t")

		switch m.Type {
		case Boolean_TYPE:
			gotype = "int"
		case Number_TYPE:
			gotype = "int"
		case String_TYPE:
			gotype = "string"
		case Datetime_TYPE:
			gotype = "time.Time"
		case Reference_Belong_TYPE:
			gotype = "int"
			goto BELONG_TYPE

		}

		buf.WriteString(gotype)
		buf.WriteString("\t")
		buf.WriteString("`json:\"" + m.Name + "\" gorm:\"column:" + m.Name + "\"`")
		buf.WriteString("\n")
		continue

	BELONG_TYPE:
		buf.WriteString(gotype)
		buf.WriteString("\t")
		buf.WriteString("`json:\"" + m.Name + "\" gorm:\"column:" + m.Name + "\"`")
		buf.WriteString("\n")
		// Company      Company `gorm:"foreignKey:CompanyRefer"`

		belongToModel := ToGoUpper(m.Model)
		belong := "\t" + belongToModel + "\t" + m.Model + "." + belongToModel + "\t`gorm:\"foreignKey:" + ToGoUpper(m.Name) + "\"`"
		buf.WriteString(belong)

		extraPackages[`"`+gomodName+`/api/`+m.Model+`"`] = 1

	}

	buf.WriteString(addTimeField())

	buf.WriteString("}")
	buf.WriteString("\n")
	buf.WriteString("\n")

	return buf.String(), extraPackages

}

func addTimeField() string {
	buf := bytes.Buffer{}
	buf.WriteString("\n")
	buf.WriteString("\tCreatedAt\ttime.Time\t`json:\"created_at\" gorm:\"column:created_at\"`")
	buf.WriteString("\n")
	buf.WriteString("\tUpdatedAt\ttime.Time\t`json:\"updated_at\" gorm:\"column:updated_at\"`")
	buf.WriteString("\n")
	buf.WriteString("\tDeletedAt\t*time.Time\t`json:\"-\" gorm:\"column:deleted_at\"`")
	buf.WriteString("\n")

	return buf.String()
}
