package u3

import (
	"bytes"
	"goapi/utils"
	"os"
	"path"
	"strings"
)

func GenerateNewApiModel() error {
	prepareData()

	// parse
	var pkgs = map[string]int{
		// `"a.liuqt.com/api/role"`: 1,
		`"gorm.io/gorm"`:             1,
		`"time"`:                     1,
		`"` + goModName + `/config"`: 1,
	}

	apiMethods := getModelMethods(apiStructName)

	structStr, extraPackages := utils.GetStructFromCreateNewApi(apiStructName, goModName, apiConfig.Attributes)
	if len(extraPackages) > 0 {
		for k := range extraPackages {
			pkgs[k] = 1
		}
	}
	pkgsStr := getPackageStr(apiName, pkgs)

	// write to file
	modelStr := pkgsStr + structStr + apiMethods
	fp, err := os.Create(path.Join(".", "api", apiName, "model.go"))
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err = fp.WriteString(modelStr); err != nil {
		return err
	}

	return nil
}

func getPackageStr(modelName string, pkgs map[string]int) string {

	buf := &bytes.Buffer{}
	buf.WriteString(`package `)
	buf.WriteString(modelName)
	buf.WriteString("\n")
	buf.WriteString("\n")

	// 包
	buf.WriteString("import (\n")
	for k := range pkgs {
		buf.WriteString("\t")
		buf.WriteString(k)
		buf.WriteString("\n")
	}

	buf.WriteString("\n")
	buf.WriteString(")\n")
	buf.WriteString("\n")

	return buf.String()
}

func getModelMethods(structName string) string {
	var s = `func (*GO_STRUCT_NAME) List(cond *config.Condition, tx ...*gorm.DB) ([]*GO_STRUCT_NAME, error) {
	var datas = make([]*GO_STRUCT_NAME, 0)
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	db = db.Model(&GO_STRUCT_NAME{})
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
		db = db.Order(v.Field + " " + v.Value)
	}

	// 分页
	db = db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)

	// 查询
	err := db.Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}


	
func (*GO_STRUCT_NAME) Count(cond *config.Condition, tx ...*gorm.DB) (int64, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var count int64
	db = db.Model(&GO_STRUCT_NAME{})
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
}

func (*GO_STRUCT_NAME) Add(o *GO_STRUCT_NAME, tx ...*gorm.DB) (*GO_STRUCT_NAME, error) {
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
}
	
func (*GO_STRUCT_NAME) Update(o *GO_STRUCT_NAME, tx ...*gorm.DB) (*GO_STRUCT_NAME, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&GO_STRUCT_NAME{}).Where("id = ?", o.ID).Updates(o).Error
	// err := db.Model(&GO_STRUCT_NAME{}).Where("id = ?", o.ID).Save(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}


func (*GO_STRUCT_NAME) Delete(id int, tx ...*gorm.DB) error {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&GO_STRUCT_NAME{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	// err := db.Delete(&GO_STRUCT_NAME{}, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (*GO_STRUCT_NAME) GetItemByID(id int, tx ...*gorm.DB) (*GO_STRUCT_NAME, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var rec = GO_STRUCT_NAME{}
	db = db.Model(&GO_STRUCT_NAME{})
	db = db.Where("deleted_at is null")
	err := db.Where("id = ?", id).First(&rec).Error
	return &rec, err
}

func (o *GO_STRUCT_NAME) TableName() string {return "GO_API_NAME"}

func (o *GO_STRUCT_NAME) BeforeUpdate(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) BeforeDelete(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) AfterFind(tx *gorm.DB) error   { return nil }

func (o *GO_STRUCT_NAME) BeforeCreate(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) AfterSave(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) AfterDelete(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) AfterCreate(tx *gorm.DB) error { return nil }

func (o *GO_STRUCT_NAME) AfterUpdate(tx *gorm.DB) error { return nil }

`

	s2 := strings.ReplaceAll(s, "GO_STRUCT_NAME", structName)
	s3 := strings.ReplaceAll(s2, "GO_API_NAME", apiName)
	return s3
}
