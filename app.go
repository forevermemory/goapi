package commonapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalConfig *Config = func() *Config {
	cfg, err := InitConfig("")
	if err != nil {
		log.Println("初始化配置文件失败:", err)
	}
	log.Println("初始化配置文件成功...")

	return cfg
}()

func InitConfig(f string) (*Config, error) {
	if f == "" {
		f = "./config.yaml"
	}
	in, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	err = yaml.Unmarshal(in, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type Config struct {
	MysqlConfig *MysqlConfig `yaml:"mysql_config"`
	HttpConfig  *HttpConfig  `yaml:"http_config"`
}

type HttpConfig struct {
	Port int `yaml:"port"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

var MYSQL *gorm.DB

func ConnectToMysql() {

	var cfg = GlobalConfig.MysqlConfig
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		log.Println("连接mysql数据库失败:", err)
		return
	}

	log.Println("连接mysql数据库成功...")
	MYSQL = db

}

// ?page=1&pageSize=10&_sort[0]=id:DESC&_sort[1]=age:DESC&_where[0][status_lte]=1&_where[1][alarm_time_lte]=2022-07-01T04:00:00.000Z&_where[2][content_contains]=1

type Condition struct {
	Page     int
	PageSize int
	Sorts    []*Item
	Wheres   []*Item
}

type QueryMethod string

const (
	GT        QueryMethod = "gt"
	GTE       QueryMethod = "gte"
	EQ        QueryMethod = "eq"
	LT        QueryMethod = "lt"
	LTE       QueryMethod = "lte"
	CONTAINS  QueryMethod = "contains"
	CONTAINSS QueryMethod = "containss"
)

type Item struct {
	Field string
	// > >= = < <= contains containss
	// gt gte eq lt lte contains containss
	Method QueryMethod
	Value  string
}

func ParseQuerys(querys url.Values) *Condition {
	var cond = Condition{
		Page:     1,
		PageSize: 10,
		Sorts:    make([]*Item, 0),
		Wheres:   make([]*Item, 0),
	}

	if _page := querys.Get("page"); _page != "" {
		ipage, err := strconv.Atoi(_page)
		if err == nil {
			cond.Page = ipage
		}
	}
	if _pageSize := querys.Get("pageSize"); _pageSize != "" {
		ipageSize, err := strconv.Atoi(_pageSize)
		if err == nil {
			cond.PageSize = ipageSize
		}
	}

	condReg, _ := regexp.Compile(`_where\[\d+\]\[(.*?)\]`)
	for k, v := range querys {
		if strings.HasPrefix(k, "_sort") && len(v) > 0 {
			tmp := strings.Split(v[0], ":")
			cond.Sorts = append(cond.Sorts, &Item{Field: tmp[0], Value: tmp[1]})
		}
		if strings.HasPrefix(k, "_where") && len(v) > 0 {
			match := condReg.FindStringSubmatch(k)
			tmps := strings.Split(match[1], "_")
			_method := QueryMethod(tmps[len(tmps)-1])
			_field := strings.Join(tmps[0:len(tmps)-1], "_")
			cond.Wheres = append(cond.Wheres, &Item{Field: _field, Method: _method, Value: v[0]})
		}
	}

	return &cond
}

// Response code 0 成功 code -1 失败
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	ErrParamNotFull int = 10000 // 参数不完整
)

var (
	ERROR = &Response{Code: -1}
	OK    = &Response{Code: 0}
)

func (r *Response) clone() *Response {
	n := new(Response)
	n.Code = r.Code
	n.Data = r.Data
	return n
}

func (r *Response) WithMessage(msg string) *Response {
	n := r.clone()
	n.Msg = msg
	return n
}

func (r *Response) WithData(data interface{}) *Response {
	n := r.clone()
	n.Data = data
	return n
}

func (r *Response) WithCode(code int) *Response {
	n := r.clone()
	n.Code = code
	return n
}
