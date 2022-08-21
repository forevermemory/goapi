package u2

import "path"

func WriteApiRolePermissionConfigJson(projectname string) error {
	return nil
}

func WriteApiRolePermissionController(projectname string) error {
	return nil
}

func WriteApiRolePermissionModel(projectname string) error {
	return nil
}

func WriteApiPermissionConfigJson(projectname string) error {
	return nil
}

func WriteApiPermissionModel(projectname string) error {
	return nil
}

func WriteApiPermissionController(projectname string) error {
	return nil
}
func WriteApiRoleController(projectname string) error {
	var content = `package role

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"` + projectname + `/config"

)

var Entity = &roleController{}

type roleController struct {}



func (a *roleController)  GetByID(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := Role{}
	data, err := b.GetItemByID(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}
	
	

func (a *roleController) Add(c *gin.Context) interface{} {
	var req = Role{}
	err := c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := Role{}
	data, err := b.Add(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}

	

func (a *roleController) Update(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	var req = Role{}
	err = c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}
	req.ID = intv

	t := Role{}
	data, err := t.Update(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}
	
	

func (a *roleController) Delete(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithData("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	t := Role{}
	err = t.Delete(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK
}
	
	

func (a *roleController) List(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := Role{}
	datas, err := b.List(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}

	

func (a *roleController) Count(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := Role{}
	datas, err := b.Count(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}



`

	return writeStr2File(path.Join(".", projectname, "api", "role", "controller.go"), content)
}

func WriteApiRoleConfigJson(projectname string) error {
	var content = `{
	"kind":"collectionType",
	"controllerName":"roleController",
	"modelName":"role",
	"modelCommnet":"角色",
	"description":"This represents the role Model",
	"attributes":[
		{
			"name":"name",
			"type":"string",
			"description":"角色名称"
		},
		{
			"name":"description",
			"type":"string",
			"description":"描述"
		}
	],
	"routes":[
		{
			"method": "GET",
			"path": "/role",
			"handler": "RoleController.List",
			"description":"查询角色列表",
			"policies": []
		},
		{
			"method": "GET",
			"path": "/role/count",
			"handler": "RoleController.Count",
			"description":"查询角色数量",
			"policies": []
		},
		{
			"method": "GET",
			"path": "/role/:id",
			"handler": "RoleController.GetByID",
			"description":"根据id查询角色",
			"policies": []
		},
		{
			"method": "POST",
			"path": "/role",
			"handler": "RoleController.Add",
			"description":"新增角色",
			"policies": []
		},
		{
			"method": "PUT",
			"path": "/role/:id",
			"handler": "RoleController.Update",
			"description":"更新角色",
			"policies": []
		},
		{
			"method": "DELETE",
			"path": "/role/:id",
			"handler": "RoleController.Delete",
			"description":"删除角色",
			"policies": []
		}
		
	]
}`
	return writeStr2File(path.Join(".", projectname, "api", "role", "config.json"), content)

}
func WriteApiRoleModel(projectname string) error {
	var content = `package role

import (
	"time"

	"` + projectname + `/config"
	"gorm.io/gorm"

)

type Role struct {
	ID	int	` + "`json:\"id\" gorm:\"column:id;primary_key;auto_increment;\"`" + `

	Name	string	` + "`json:\"name\" gorm:\"column:name\"`" + `
	Description	string	` + "`json:\"description\" gorm:\"column:description\"`" + `

	CreatedAt	time.Time	` + "`json:\"created_at\" gorm:\"column:created_at\"`" + `
	UpdatedAt	time.Time	` + "`json:\"updated_at\" gorm:\"column:updated_at\"`" + `
	DeletedAt	*time.Time	` + "`json:\"-\" gorm:\"column:deleted_at\"`" + `
}

func (o *Role) BeforeUpdate(tx *gorm.DB) error { return nil }


	
func (*Role) GetItemByID(id int, tx ...*gorm.DB) (*Role, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var rec = Role{}
	db = db.Model(&Role{})
	db = db.Where("deleted_at is null")
	err := db.Where("id = ?", id).First(&rec).Error
	return &rec, err
}

func (o *Role) AfterCreate(tx *gorm.DB) error { return nil }

func (o *Role) AfterDelete(tx *gorm.DB) error { return nil }

func (o *Role) TableName() string {return "role"}


func (*Role) Add(o *Role, tx ...*gorm.DB) (*Role, error) {
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


func (*Role) Delete(id int, tx ...*gorm.DB) error {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&Role{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	// err := db.Delete(&Role{}, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (o *Role) BeforeDelete(tx *gorm.DB) error { return nil }


	
func (*Role) Count(cond *config.Condition, tx ...*gorm.DB) (int64, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var count int64
	db = db.Model(&Role{})
	db = db.Where("deleted_at is null")
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field+" > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field+" >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field+" < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field+" <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field+" = ? ", v.Value)
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

func (o *Role) BeforeCreate(tx *gorm.DB) error { return nil }

func (o *Role) AfterUpdate(tx *gorm.DB) error { return nil }


	
func (*Role) Update(o *Role, tx ...*gorm.DB) (*Role, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&Role{}).Where("id = ?", o.ID).Updates(o).Error
	// err := db.Model(&Role{}).Where("id = ?", o.ID).Save(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Role) AfterSave(tx *gorm.DB) error { return nil }

func (o *Role) AfterFind(tx *gorm.DB) error   { return nil }


	
func (*Role) List(cond *config.Condition, tx ...*gorm.DB) ([]*Role, error) {
	var datas = make([]*Role, 0)
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	db = db.Model(&Role{})
	db = db.Where("deleted_at is null")
	// 条件
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field+" > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field+" >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field+" < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field+" <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field+" = ? ", v.Value)
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


`

	return writeStr2File(path.Join(".", projectname, "api", "role", "model.go"), content)
}

func WriteApiUserModel(projectname string) error {
	var content = `package user

import (
	"time"

	"` + projectname + `/api/role"
	"` + projectname + `/config"
	"gorm.io/gorm"

)

type User struct {
	ID	int	` + "`json:\"id\" gorm:\"column:id;primary_key;auto_increment;\"`" + `

	Username	string	` + "`json:\"username\" gorm:\"column:username\"`" + `
	Password	string	` + "`json:\"password\" gorm:\"column:password\"`" + `
	Email	string	` + "`json:\"email\" gorm:\"column:email\"`" + `
	Blocked	int	` + "`json:\"blocked\" gorm:\"column:blocked\"`" + `
	LoginAt	time.Time	` + "`json:\"login_at\" gorm:\"column:login_at\"`" + `
	RoleId	int	` + "`json:\"role_id\" gorm:\"column:role_id\"`" + `
	Role	role.Role	` + "`gorm:\"foreignKey:RoleId\"`" + `
	CreatedAt	time.Time	` + "`json:\"created_at\" gorm:\"column:created_at\"`" + `
	UpdatedAt	time.Time	` + "`json:\"updated_at\" gorm:\"column:updated_at\"`" + `
	DeletedAt	*time.Time	` + "`json:\"-\" gorm:\"column:deleted_at\"`" + `
}


func (*User) Delete(id int, tx ...*gorm.DB) error {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&User{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	// err := db.Delete(&User{}, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (o *User) BeforeCreate(tx *gorm.DB) error { return nil }

func (o *User) AfterDelete(tx *gorm.DB) error { return nil }


	
func (*User) List(cond *config.Condition, tx ...*gorm.DB) ([]*User, error) {
	var datas = make([]*User, 0)
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	db = db.Model(&User{})
	db = db.Where("deleted_at is null")
	// 条件
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field+" > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field+" >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field+" < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field+" <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field+" = ? ", v.Value)
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


func (*User) Add(o *User, tx ...*gorm.DB) (*User, error) {
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

func (o *User) AfterFind(tx *gorm.DB) error   { return nil }


	
func (*User) GetItemByID(id int, tx ...*gorm.DB) (*User, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var rec = User{}
	db = db.Model(&User{})
	db = db.Where("deleted_at is null")
	err := db.Where("id = ?", id).First(&rec).Error
	return &rec, err
}

func (o *User) AfterCreate(tx *gorm.DB) error { return nil }

func (o *User) TableName() string {return "user"}

func (o *User) AfterUpdate(tx *gorm.DB) error { return nil }

func (o *User) BeforeDelete(tx *gorm.DB) error { return nil }

func (o *User) BeforeUpdate(tx *gorm.DB) error { return nil }


	
func (*User) Count(cond *config.Condition, tx ...*gorm.DB) (int64, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	var count int64
	db = db.Model(&User{})
	db = db.Where("deleted_at is null")
	for _, v := range cond.Wheres {
		switch v.Method {
		case config.GT:
			db = db.Where(v.Field+" > ? ", v.Value)
		case config.GTE:
			db = db.Where(v.Field+" >= ? ", v.Value)
		case config.LT:
			db = db.Where(v.Field+" < ? ", v.Value)
		case config.LTE:
			db = db.Where(v.Field+" <= ? ", v.Value)
		case config.EQ:
			db = db.Where(v.Field+" = ? ", v.Value)
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

func (o *User) AfterSave(tx *gorm.DB) error { return nil }


	
func (*User) Update(o *User, tx ...*gorm.DB) (*User, error) {
	db := config.DATABASE
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Model(&User{}).Where("id = ?", o.ID).Updates(o).Error
	// err := db.Model(&User{}).Where("id = ?", o.ID).Save(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}
	
	
	
`
	return writeStr2File(path.Join(".", projectname, "api", "user", "model.go"), content)

}

func WriteApiUserConfigJson(projectname string) error {
	var content = `{
	"kind":"collectionType",
	"controllerName":"UserController",
	"modelName":"user",
	"modelCommnet":"用户",
	"description":"This represents the user Model",
	"attributes":[
		{
			"name":"username",
			"type":"string",
			"description":"描述"
		},
		{
			"name":"password",
			"type":"string",
			"description":"密码"
		},
		{
			"name":"email",
			"type":"string",
			"description":"邮箱"
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
	],
	"routes":[
		{
			"method": "GET",
			"path": "/user",
			"handler": "UserController.List",
			"description":"查询用户列表",
			"policies": []
		},
		{
			"method": "GET",
			"path": "/user/count",
			"handler": "UserController.Count",
			"description":"查询用户数量",
			"policies": []
		},
		{
			"method": "GET",
			"path": "/user/:id",
			"handler": "UserController.GetByID",
			"description":"根据id查询用户",
			"policies": []
		},
		{
			"method": "POST",
			"path": "/user",
			"handler": "UserController.Add",
			"description":"新增用户",
			"policies": []
		},
		{
			"method": "PUT",
			"path": "/user/:id",
			"handler": "UserController.Update",
			"description":"更新用户",
			"policies": []
		},
		{
			"method": "DELETE",
			"path": "/user/:id",
			"handler": "UserController.Delete",
			"description":"删除用户",
			"policies": []
		}
		
	]
}`
	return writeStr2File(path.Join(".", projectname, "api", "user", "config.json"), content)
}
func WriteApiUserController(projectname string) error {
	var content = `package user

import (
	"github.com/gin-gonic/gin"
	"` + projectname + `/config"
	"strconv"

)

var Entity = &UserController{}

type UserController struct {}



func (a *UserController) Update(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	var req = User{}
	err = c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}
	req.ID = intv

	t := User{}
	data, err := t.Update(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}
	
	

func (a *UserController) Delete(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithData("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	t := User{}
	err = t.Delete(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK
}
	
	

func (a *UserController) List(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := User{}
	datas, err := b.List(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}

	

func (a *UserController) Count(c *gin.Context) interface{} {
	var querys = c.Request.URL.Query()
	cond := config.ParseQuerys(querys)

	b := User{}
	datas, err := b.Count(cond)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(datas)
}



func (a *UserController)  GetByID(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		return config.ERROR.WithCode(config.ErrParamNotFull).WithMessage("参数不完整")
	}
	intv, err := strconv.Atoi(id)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := User{}
	data, err := b.GetItemByID(intv)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}
	
	

func (a *UserController) Add(c *gin.Context) interface{} {
	var req = User{}
	err := c.ShouldBind(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	b := User{}
	data, err := b.Add(&req)
	if err != nil {
		return config.ERROR.WithMessage(err.Error())
	}

	return config.OK.WithData(data)
}

		
	
	
`
	return writeStr2File(path.Join(".", projectname, "api", "user", "controller.go"), content)

}
